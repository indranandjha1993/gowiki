package wiki

import "os"

type Page struct {
	Title string
	Body  []byte
}

func NewPage(title string, body []byte) error {
	p := &Page{Title: title, Body: body}
	err := p.save()
	return err
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func InsertDummyData() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
