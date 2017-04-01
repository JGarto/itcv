// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package immtodoapp

import (
	"fmt"

	r "github.com/myitcv/gopherjs/react"
	"honnef.co/go/js/dom"
)

//go:generate reactGen
//go:generate immutableGen

// TodoAppDef is the definition fo the TodoApp component
type TodoAppDef struct {
	r.ComponentDef
}

type _Imm_item struct {
	name string
}

type _Imm_itemS []*item

// TodoAppState is the state type for the TodoApp component
type TodoAppState struct {
	items    *itemS
	currItem string
}

// TodoApp creates instances of the TodoApp component
func TodoApp() *TodoAppDef {
	res := &TodoAppDef{}
	r.BlessElement(res, nil)
	return res
}

func (t *TodoAppDef) GetInitialState() TodoAppState {
	return TodoAppState{
		items: new(itemS),
	}
}

// Render renders the TodoApp component
func (t *TodoAppDef) Render() r.Element {
	var entries []*r.LiDef

	for _, v := range t.State().items.Range() {
		entry := r.Li(nil, r.S(v.name()))
		entries = append(entries, entry)
	}

	return r.Div(nil,
		r.H3(nil, r.S("TODO")),
		r.Ul(nil, entries...),
		r.Form(&r.FormProps{ClassName: "form-inline"},
			r.Div(
				&r.DivProps{ClassName: "form-group"},
				r.Label(&r.LabelProps{ClassName: "sr-only", For: "todoText"}, r.S("Todo Item")),
				r.Input(&r.InputProps{
					Type:        "text",
					ClassName:   "form-control",
					ID:          "todoText",
					Placeholder: "Todo Item",
					Value:       t.State().currItem,
					OnChange:    t.onCurrItemChange,
				}),
				r.Button(&r.ButtonProps{
					Type:      "submit",
					ClassName: "btn btn-default",
					OnClick:   t.onAddClicked,
				}, r.S(fmt.Sprintf("Add #%v", t.State().items.Len()+1))),
			),
		),
	)
}

func (t *TodoAppDef) onCurrItemChange(se *r.SyntheticEvent) {
	target := se.Target().(*dom.HTMLInputElement)

	ns := t.State()
	ns.currItem = target.Value

	t.SetState(ns)
}

func (t *TodoAppDef) onAddClicked(se *r.SyntheticMouseEvent) {
	ns := t.State()

	ns.items = ns.items.Append(new(item).setName(ns.currItem))
	ns.currItem = ""

	t.SetState(ns)

	se.PreventDefault()
}
