package models

type ViewSearchResult struct {
	MemberType string
	Name       string
}

func (view *ViewSearchResult) setView(memberType string, name string) {
	view.MemberType = memberType
	view.Name = name
}
