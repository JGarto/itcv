// Code generated by reactGen. DO NOT EDIT.

package main

import "myitcv.io/react"

type FooBarElem struct {
	react.Element
}

func (f FooBarDef) ShouldComponentUpdateIntf(nextProps, prevState, nextState interface{}) bool {
	res := false

	{
		res = f.Props() != nextProps.(FooBarProps) || res
	}
	v := prevState.(FooBarState)
	res = !v.EqualsIntf(nextState) || res
	return res
}

func buildFooBar(cd react.ComponentDef) react.Component {
	return FooBarDef{ComponentDef: cd}
}

func buildFooBarElem(props FooBarProps, children ...react.Element) *FooBarElem {
	return &FooBarElem{
		Element: react.CreateElement(buildFooBar, props),
	}
}

// SetState is an auto-generated proxy proxy to update the state for the
// FooBar component.  SetState does not immediately mutate f.State()
// but creates a pending state transition.
func (f FooBarDef) SetState(state FooBarState) {
	f.ComponentDef.SetState(state)
}

// State is an auto-generated proxy to return the current state in use for the
// render of the FooBar component
func (f FooBarDef) State() FooBarState {
	return f.ComponentDef.State().(FooBarState)
}

// IsState is an auto-generated definition so that FooBarState implements
// the myitcv.io/react.State interface.
func (f FooBarState) IsState() {}

var _ react.State = FooBarState{}

// GetInitialStateIntf is an auto-generated proxy to GetInitialState
func (f FooBarDef) GetInitialStateIntf() react.State {
	return FooBarState{}
}

func (f FooBarState) EqualsIntf(val interface{}) bool {
	return f == val.(FooBarState)
}

// Props is an auto-generated proxy to the current props of FooBar
func (f FooBarDef) Props() FooBarProps {
	uprops := f.ComponentDef.Props()
	return uprops.(FooBarProps)
}

func (f FooBarProps) EqualsIntf(val interface{}) bool {
	return f == val.(FooBarProps)
}

var _ react.Equals = FooBarProps{}
