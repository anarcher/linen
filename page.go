package linen

type Page struct {
}

func NewPage(site *Site) (Page, error) {
	page := Page{}
	return page, nil
}
