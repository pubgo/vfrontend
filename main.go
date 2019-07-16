package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/pubgo/vfrontend/routes"
	"github.com/pubgo/vfrontend/stores"
	"github.com/pubgo/vfrontend/views"
	. "github.com/siongui/godom"
	"github.com/vincent-petithory/dataurl"
)

func main() {
	if Document.Get("ReadyState") == js.Undefined {
		go run()
		return
	}

	Document.AddEventListener("DOMContentLoaded", func(event Event) {
		go run()
	})
}

func run() {
	vecty.AddStylesheet(dataurl.New([]byte(views.Styles), "text/css").String())

	/*
	vapper serve
	vapper build

	vapper.Start(func(app vapper){
		config.Init(app)
		routes.InitRouter(app)
	})
	*/


	app := &stores.App{}
	app.Init()
	p := views.NewPage(app)

	r := routes.New()
	r.ForceHashURL = false
	//r.ShouldInterceptLinks=true

	r.HandleFunc("/{name}", func(context *routes.Context) {
		fmt.Printf("Hello, %s\n", context.Params["name"])
		vecty.RenderBody(p)
	})

	r.HandleFunc("/", func(context *routes.Context) {
		fmt.Println(context, "context")
		vecty.RenderBody(p)
	})

	pt := Window.Get("location").Get("pathname").String()
	if r.CanNavigate(pt) {
		r.Navigate(pt)
	} else {
		r.Navigate("/")
	}
}
