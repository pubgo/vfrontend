package main

import (
	"github.com/gopherjs/vecty"
	"github.com/pubgo/vapper/vapper"
	_ "github.com/pubgo/vfrontend/config"
	_ "github.com/pubgo/vfrontend/routes"
	"github.com/pubgo/vfrontend/views"
	"github.com/vincent-petithory/dataurl"
)

func main() {
	app := vapper.Default()
	app.ForceHashURL = false
	app.Start()
	vecty.AddStylesheet(dataurl.New([]byte(views.Styles), "text/css").String())
}
