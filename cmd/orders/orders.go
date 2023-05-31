package main

import (
	"net/http"
	"orders/pkg/api"
	"orders/pkg/db"
	"log"
	//"sf/36_WebApp/orders/pkg/db"
)

func main() {
    // Инициализация БД в памяти.
    db := db.New()
    // Создание объекта API, использующего БД в памяти.
    api := api.New(db)
    // Запуск сетевой службы и HTTP-сервера
    // на всех локальных IP-адресах на порту 80.
    err:=http.ListenAndServe(":80", api.Router())
    if err != nil {
       log.Fatal(err)
    }
}