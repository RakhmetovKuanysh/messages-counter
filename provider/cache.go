package provider

import (
	redigo "github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"time"
)

// Интерфейс кэша
type Cachier interface {
	Set(key string, val interface{}, lifetime int)
	Get(key string) (interface{}, error)
	GetInt(key string) (int, error)
	Delete(key string)
}

// Настройки
type Options struct {
	Addr    string
	Timeout int
	Db      int
	Pool    *redigo.Pool
}

// Структура кэша
type Cache struct {
	provider *redigo.Pool
}

var options Options

const (
	DefaultPoolMaxIdle     = 3
	DefaultPoolIdleTimeout = 240
)

// Инициализация настроек redis
func Init(opt Options) {
	options = opt

	if options.Pool == nil {
		options.Pool = &redigo.Pool{
			MaxIdle:     DefaultPoolMaxIdle,
			IdleTimeout: DefaultPoolIdleTimeout,
		}
	}

	options.Pool.IdleTimeout *= time.Second
	options.Pool.MaxConnLifetime *= time.Second
	options.Pool.Dial = func() (redigo.Conn, error) {
		return redigo.Dial("tcp", options.Addr)
	}
}

// Подключение к кэшу
func (c *Cache) Connect(options Options) Cachier {
	Init(options)

	c.provider = Pool()

	return c
}

// Получает пулл соединений
func Pool() *redigo.Pool {
	return options.Pool
}

// Сохранение в кэш
func (c *Cache) Set(key string, val interface{}, lifetime int) {
	conn := c.provider.Get()
	defer func() {
		_ = conn.Close()
	}()

	_, err := conn.Do("SETEX", key, lifetime, val)

	if err != nil {
		logrus.WithError(err).Errorf("Не удалось сохранить данные в кэш key: %s, val: %s", key, val)
	}
}

// Получение данных из кэша
func (c *Cache) Get(key string) (interface{}, error) {
	conn := c.provider.Get()
	defer func() {
		_ = conn.Close()
	}()

	return conn.Do("GET", key)
}

// Получения данных из кэша с типом int
func (c *Cache) GetInt(key string) (int, error) {
	return redigo.Int(c.Get(key))
}

// Удаление кэш по ключу
func (c *Cache) Delete(key string) {
	conn := c.provider.Get()
	defer func() {
		_ = conn.Close()
	}()

	_, err := conn.Do("DEL", key)

	if err != nil {
		logrus.WithField("key", key).
			WithError(err).
			Error("Не удалось удалить кэш")
	}
}
