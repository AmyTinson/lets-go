package main

import (
	"fmt"
	"os"
)

type List struct {
	Title string
	Body []byte
}

func (p *List) save() error {
	filename :=  p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadList(title string) (*List, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
			return nil, err
	}
	return &List{Title: title, Body: body}, nil
}

func main() {
	p1 := &List{Title: "TestList", Body: []byte("This is a sample List.")}
	p1.save()
	p2, _ := loadList("TestList")
	fmt.Println(string(p2.Body))
}