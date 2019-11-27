package components

import (
	"github.com/pubgo/vapper/jsvapper"
	"github.com/pubgo/vfrontend/actions"
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"github.com/tulir/gopher-ace"
)

type Editor struct {
	vecty.Core

	Text string `vecty:"prop"`

	editor   ace.Editor
	id, lang string
	readonly bool
	app      *jsvapper.Vapper
}

func NewEditor(id, lang, text string, readonly bool) *Editor {
	v := &Editor{
		lang:     lang,
		id:       id,
		Text:     text,
		readonly: readonly,
	}
	return v
}

func (v *Editor) Init(app *jsvapper.Vapper) {
	v.app = app
}

func (v *Editor) Mount() {
	v.editor = ace.Edit(v.id)
	v.editor.SetOptions(map[string]interface{}{
		"mode": "ace/mode/" + v.lang,
	})
	if v.Text != "" {
		v.editor.SetValue(v.Text)
		v.editor.ClearSelection()
		v.editor.MoveCursorTo(0, 0)
	}
	var changes int
	v.editor.OnChange(func(ev *js.Object) {
		changes++
		before := changes

		go func() {
			<-time.After(time.Millisecond * 250)
			if before == changes {
				v.app.Dispatch(&actions.UserChangedTextAction{
					Text: v.editor.GetValue(),
				})
			}
		}()
	})
}

func (v *Editor) Render() vecty.ComponentOrHTML {
	if !v.readonly && v.editor.Object != nil && v.Text != v.editor.GetValue() {
		// only update the editor if the text is changed
		v.editor.SetValue(v.Text)
		v.editor.ClearSelection()
		v.editor.MoveCursorTo(0, 0)
	}

	return elem.Div(
		vecty.Markup(
			prop.ID(v.id),
			vecty.Class("editor"),
			event.Change(func(i *vecty.Event) {
				v.Text = i.Value.String()
				vecty.Rerender(v)
			}),
		),
		vecty.Text(v.Text),
	)
}
