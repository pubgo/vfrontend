package routes

import (
	"github.com/pubgo/vapper/vapper"
	"github.com/pubgo/vfrontend/views"
)

func init() {
	vapper.Route("/", views.NewPage())
	vapper.Route("/app", views.NewPage())
	vapper.Route("/app/{id}", views.NewPage())
}
