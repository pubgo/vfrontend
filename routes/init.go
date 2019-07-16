package routes

import (
	"github.com/pubgo/vapper/vapper"
	"github.com/pubgo/vfrontend/views"
)

func init() {
	app := vapper.Default()
	app.Route("/", views.NewPage())
	app.Route("/app", views.NewPage())
	app.Route("/app/{id}", views.NewPage())
}
