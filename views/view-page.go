package views

import (
	"fmt"
	"github.com/dave/splitter"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/pubgo/vapper/jsvapper"
	"github.com/pubgo/vfrontend/actions"
	"github.com/pubgo/vfrontend/compontents"
	"github.com/pubgo/vfrontend/stores"
)

type Page struct {
	vecty.Core
	app   *jsvapper.Vapper
	split *splitter.Split

	editor  *stores.EditorStore
	_review int
}

func NewPage() *Page {
	v := &Page{}
	return v
}

func (t *Page) Mount() {
	t.app.Watch(t, func(done chan struct{}) {
		defer close(done)
		vecty.Rerender(t)
	})

	fmt.Println("ReadyStateComplete")
	t.split = splitter.New("split")
	t.split.Init(
		js.S{"#left", "#right"},
		js.M{"sizes": []float64{50, 50}},
	)
}

func (t *Page) Unmount() {
	t.app.Delete(t)
}

func (t *Page) Handle(ctx *jsvapper.Context) {
	vecty.RenderBody(t)
}

func (t *Page) Init(app *jsvapper.Vapper, editor *stores.EditorStore) {
	t.app = app
	t.editor = editor
}

// ReadyStateComplete call when ReadyState is Complete
func (t *Page) ReadyStateComplete() {

}

func (t *Page) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("container-fluid", "p-0", "split", "split-horizontal"),
			),
			t.renderLeft(),
			t.renderRight(),
		),
	)
}

func (t *Page) renderLeft() *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			prop.ID("left"),
			vecty.Class("split"),
		),
		compontents.NewEditor("html-editor", "html", t.editor.Html(), true, func(value string) {
			t.app.Dispatch(&actions.UserChangedTextAction{
				Text: value,
			})
		}),
	)
}

func (t *Page) renderRight() *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			prop.ID("right"),
			vecty.Class("split"),
		),
		compontents.NewEditor("code-editor", "golang", t.editor.Code(), false, nil),
	)
}
