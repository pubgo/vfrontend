// This file was created with https://github.com/pubgo/factor
// using https://jsgo.io/dave/html2vecty
package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type Test struct {
	vecty.Core
}

func (p *Test) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Heading1(
			vecty.Text("html2vecty"),
		),
		elem.Paragraph(
			vecty.Text("Enter HTML here and the vecty syntax will appear opposite."),
		),
		elem.Heading2(
			vecty.Text("Class attributes"),
		),
		elem.Paragraph(
			vecty.Markup(
				vecty.Class("foo", "bar", "baz"),
			),
		),
		elem.Heading2(
			vecty.Text("Style attributes"),
		),
		elem.Paragraph(
			vecty.Markup(
				vecty.Style("border", "2px"),
				vecty.Style("color", "red!important"),
			),
		),
		elem.Heading2(
			vecty.Text("Special properties"),
		),
		elem.Input(
			vecty.Markup(
				prop.Type(prop.TypeCheckbox),
				prop.Checked(true),
				prop.Autofocus(true),
			),
		),
		elem.Anchor(
			vecty.Markup(
				prop.Href("href"),
				prop.ID("id"),
				vecty.Data("foo", "bar"),
			),
			vecty.Text("Props"),
		),
		elem.Heading2(
			vecty.Text("An example"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("modal"),
				vecty.Attribute("tabindex", "-1"),
				vecty.Attribute("role", "dialog"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("modal-dialog"),
					vecty.Attribute("role", "document"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("modal-content"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("modal-header"),
						),
						elem.Heading5(
							vecty.Markup(
								vecty.Class("modal-title"),
							),
							vecty.Text("Modal title"),
						),
						elem.Button(
							vecty.Markup(
								prop.Type(prop.TypeButton),
								vecty.Class("close"),
								vecty.Data("dismiss", "modal"),
								vecty.Attribute("aria-label", "Close"),
							),
							elem.Span(
								vecty.Markup(
									vecty.Attribute("aria-hidden", "true"),
								),
								vecty.Text("x"),
							),
						),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("modal-body"),
						),
						elem.Paragraph(
							vecty.Text("Modal body text goes here."),
						),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("modal-footer"),
						),
						elem.Button(
							vecty.Markup(
								prop.Type(prop.TypeButton),
								vecty.Class("btn", "btn-primary"),
							),
							vecty.Text("Save changes"),
						),
						elem.Button(
							vecty.Markup(
								prop.Type(prop.TypeButton),
								vecty.Class("btn", "btn-secondary"),
								vecty.Data("dismiss", "modal"),
							),
							vecty.Text("Close"),
						),
					),
				),
			),
		),
	)
}
