package model

type CardGallery struct {
	Cards []Card
}

type Card struct {
	Title  string
	Blurb  string
	isHero bool
	Image  Image
}
