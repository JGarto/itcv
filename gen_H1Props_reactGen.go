// Code generated by reactGen. DO NOT EDIT.

package react

// H1Props defines the properties for the <h1> element
type H1Props struct {
	AriaSet
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	ID  string
	Key string

	OnChange
	OnClick

	Ref
	Role  string
	Style *CSS
}

func (h *H1Props) assign(_v *_H1Props) {

	if h.AriaSet != nil {
		for dk, dv := range h.AriaSet {
			_v.o.Set("aria-"+dk, dv)
		}
	}

	_v.ClassName = h.ClassName

	_v.DangerouslySetInnerHTML = h.DangerouslySetInnerHTML

	if h.DataSet != nil {
		for dk, dv := range h.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if h.ID != "" {
		_v.ID = h.ID
	}

	if h.Key != "" {
		_v.Key = h.Key
	}

	if h.OnChange != nil {
		_v.o.Set("onChange", h.OnChange.OnChange)
	}

	if h.OnClick != nil {
		_v.o.Set("onClick", h.OnClick.OnClick)
	}

	if h.Ref != nil {
		_v.o.Set("ref", h.Ref.Ref)
	}

	_v.Role = h.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = h.Style.hack()

}
