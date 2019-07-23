package routes

import (
	"github.com/pubgo/vapper/jsvapper"
	"github.com/pubgo/vfrontend/views"
)

func init() {
	app := jsvapper.Default()
	app.Route("/", views.NewPage())
	app.Route("/app", views.NewPage())
	app.Route("/app/{id}", views.NewPage())
	app.NotFound(views.NewNotFound())
}
