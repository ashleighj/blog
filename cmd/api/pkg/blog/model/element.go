package model

type Element struct {
	Name         string
	Text         string
	Color        string
	Size         Size
	ChildElement *Element
}
