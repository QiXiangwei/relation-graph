package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"relation-graph/controllers"
)

func main()  {
	var (
		app *iris.Application
		err error
	)

	app = iris.Default()

	for domain, handler := range controllers.Controllers {
		mvc.Configure(app.Party(domain), handler)
	}

	if err = app.Run(iris.Addr(":8088")); err != nil {
		fmt.Println("iris run error:", err.Error())
	}
}
