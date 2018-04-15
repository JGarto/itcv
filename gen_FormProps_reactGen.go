// Code generated by reactGen. DO NOT EDIT.

package react

// FormProps defines the properties for the <form> element
type FormProps struct {
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

func (f *FormProps) assign(_v *_FormProps) {

	if f.AriaSet != nil {
		for dk, dv := range f.AriaSet {
			_v.o.Set("aria-"+dk, dv)
		}
	}

	_v.ClassName = f.ClassName

	_v.DangerouslySetInnerHTML = f.DangerouslySetInnerHTML

	if f.DataSet != nil {
		for dk, dv := range f.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if f.ID != "" {
		_v.ID = f.ID
	}

	if f.Key != "" {
		_v.Key = f.Key
	}

	if f.OnChange != nil {
		_v.o.Set("onChange", f.OnChange.OnChange)
	}

	if f.OnClick != nil {
		_v.o.Set("onClick", f.OnClick.OnClick)
	}

	if f.Ref != nil {
		_v.o.Set("ref", f.Ref.Ref)
	}

	_v.Role = f.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = f.Style.hack()

}
