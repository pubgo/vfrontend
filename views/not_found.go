package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/pubgo/vapper/jsvapper"
)

func NewNotFound() *NotFound {
	return &NotFound{}
}

type NotFound struct {
	vecty.Core
}

func (t *NotFound) ReadyStateComplete() {

}
func (t *NotFound) Handle(ctx *jsvapper.Context) {
	vecty.RenderBody(t)
}

func (t *NotFound) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(prop.ID("home-view")),
		elem.Div(
			vecty.Markup(prop.ID("home-top")),
			elem.Heading1(
				vecty.Text("page not found 🤦🏻‍♂️"),
			),
		),
	)
}
