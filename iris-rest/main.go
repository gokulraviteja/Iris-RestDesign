package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/gorillamux"
)

func main() {
	app := iris.New()

	// DevLogger returns a new Logger which prints both ProdMode and DevMode messages to the default global logger printer.
	app.Adapt(iris.DevLogger())

	//returns a mux router plugged inside iris
	app.Adapt(gorillamux.New())

	router := app.Party("v1")

	router.Get("/iristest", irisTest)

	router.Post("/createbook", createBook)
	router.Get("/getbooks", getBooks)
	router.Get("/book/{name}", getBook)
	router.Delete("/book/{name}", deleteBook)
	router.Put("/book/{name}", updateBook)

	app.Listen(":8011")

}
