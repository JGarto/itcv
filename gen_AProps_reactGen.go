// Code generated by reactGen. DO NOT EDIT.

package react

// APropsDef defines the properties for the <a> element
type AProps struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTMLDef
	Href                    string
	ID                      string
	Key                     string

	OnChange
	OnClick

	Role   string
	Style  *CSS
	Target string
}

func (a *AProps) assign(v *_AProps) {

	v.ClassName = a.ClassName

	v.DangerouslySetInnerHTML = a.DangerouslySetInnerHTML

	v.Href = a.Href

	if a.ID != "" {
		v.ID = a.ID
	}

	if a.Key != "" {
		v.Key = a.Key
	}

	if a.OnChange != nil {
		v.o.Set("onChange", a.OnChange.OnChange)
	}

	if a.OnClick != nil {
		v.o.Set("onClick", a.OnClick.OnClick)
	}

	v.Role = a.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = a.Style.hack()

	v.Target = a.Target

}
