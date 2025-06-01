package main

import (
	todo "alisherAsd"
	"alisherAsd/pkg/handler"
	"alisherAsd/pkg/repository"
	"alisherAsd/pkg/service"

	"log"

	"github.com/spf13/viper"
)

// Точка входа приложения
// Создаём экземпляр структуры Handler
// Запускаем сервер
func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error init config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	server := new(todo.Server)
	// Из server.go и handler.go запускаем сервер с обработчиками
	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("Error run http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}