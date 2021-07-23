package main

import (
	"os"

	"github.com/dvd-denis/poem-app"
	"github.com/dvd-denis/poem-app/packages/handler"
	"github.com/dvd-denis/poem-app/packages/repository"
	"github.com/dvd-denis/poem-app/packages/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {
	// ! logrus.SetFormatter(new(logrus.JSONFormatter)) для сохранения логов

	// ! gin.SetMode(gin.ReleaseMode) сделать при релизе

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializning configs: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		UserName: viper.GetString("db.username"),
		DBname:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initilize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(*repos)
	handlers := handler.NewHandler(service)

	srv := new(poem.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}
