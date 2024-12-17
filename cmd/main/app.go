package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"restapitry/cmd/internal/user"
	"time"
)

func main() {
	log.Println("create router ")
	// Создаем новый роутер
	router := httprouter.New()

	log.Println("register user handler ")
	handler := user.NewHandler()
	handler.Register(router)
	start(router)
}

func start(router *httprouter.Router) {
	log.Println("start application ")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	// Создаем сервер с настройками таймаутов
	server := &http.Server{
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}

	log.Println("server is listening port 0.0.0.0:1234 ")
	// Запускаем сервер и логируем ошибку, если она произойдет
	log.Fatalln(server.Serve(listener))
}
