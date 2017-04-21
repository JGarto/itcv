// Code generated by reactGen. DO NOT EDIT.

package react

// SpanProps defines the properties for the <p> element
type SpanProps struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTMLDef
	ID                      string
	Key                     string
	OnChange                func(e *SyntheticEvent)
	OnClick                 func(e *SyntheticMouseEvent)
	Role                    string
	Style                   *CSS
}

func (s *SpanProps) assign(v *_SpanProps) {

	v.ClassName = s.ClassName

	v.DangerouslySetInnerHTML = s.DangerouslySetInnerHTML

	if s.ID != "" {
		v.ID = s.ID
	}

	if s.Key != "" {
		v.Key = s.Key
	}

	v.OnChange = s.OnChange

	v.OnClick = s.OnClick

	v.Role = s.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = s.Style.hack()

}
