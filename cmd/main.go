package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dvd-denis/IT-Platform"
	"github.com/dvd-denis/IT-Platform/packages/handler"
	"github.com/dvd-denis/IT-Platform/packages/repository"
	"github.com/dvd-denis/IT-Platform/packages/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func main() {
	// ? logrus.SetFormatter(new(logrus.JSONFormatter)) json формат

	customFormatter := new(prefixed.TextFormatter)
	customFormatter.DisableColors = false
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	customFormatter.ForceFormatting = true

	logrus.SetFormatter(customFormatter)
	// gin.SetMode(gin.ReleaseMode)

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializning configs: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	f, err := os.OpenFile(viper.GetString("logfile"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logrus.Fatalf("Error initializning log file")
	}
	defer f.Close()

	logrus.SetOutput(io.MultiWriter(f, os.Stderr)) // ? Запись логов

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
	go func() {
		if err := srv.Run(viper.GetString("port"), viper.GetString("portTLS"), handlers.InitRouters()); err != nil {
			if err != http.ErrServerClosed {
				logrus.Fatalf("error occured while running http server: %s", err.Error())
			}
		}
	}()

	logrus.Info("IT-Platform Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("IT-Platform Shutting Down")

	errs := srv.Shutdown(context.Background())

	for _, err := range errs {
		if err != nil {
			logrus.Errorf("error occured on server shutting down: %s", err.Error())
		}
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}
