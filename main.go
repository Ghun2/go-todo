package main

import (
	"go-todo/app"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	m := app.MakeHandler("./todo.db")
	defer m.Close()

	log.Fatal(http.ListenAndServe(":" + port, m))
}
