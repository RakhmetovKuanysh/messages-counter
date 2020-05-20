package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	nethttp "net/http"
	"os"
	"otus/messages-counter/app/di"
	"otus/messages-counter/app/http/router"
	"otus/messages-counter/messages_api"
	"otus/messages-counter/provider"
)

// конфиг микросервиса
var config struct {
	Listen string
	Debug  bool
	Redis  provider.Options
}

// инициализация конфига
func Init(configPath string) {
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		logrus.Fatalln("Не удалось загрузить конфиг", err)
	}
}

// точка входа
func main() {
	configPath := flag.String("config", "config/testing/config.toml", "Путь до файла конфига toml")
	flag.Parse()

	if *configPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	Init(*configPath)

	appDI := initDI()

	nethttp.Handle("/", router.Router(appDI, config.Debug))

	logrus.Info("ms-messages-counter - сервис запущен")
	logrus.Fatal(nethttp.ListenAndServe(config.Listen, nil))
}

// инициализация приложения
func initDI() *di.DI {
	messagesAPI := messages_api.NewMessagesAPI()
	cacheProvider := new(provider.Cache).Connect(config.Redis)
	core := di.NewDI(cacheProvider, messagesAPI)

	return &core
}
