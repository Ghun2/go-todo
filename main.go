package main

import (
	"go-todo/app"
	"log"
	"net/http"
)

func main() {
	m := app.MakeHandler("./todo.db")
	defer m.Close()

	log.Fatal(http.ListenAndServe(":3000", m))
}
