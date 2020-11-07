package model

import "time"

type ImageType string

const (
	ImgHero      ImageType = "IMG_HERO"
	ImgGallery   ImageType = "IMG_GALLERY"
	ImgArticle   ImageType = "IMG_ARTICLE"
	ImgThumbnail ImageType = "IMG_THUMBNAIL"
)

type Image struct {
	Title    string    `json:"title"`
	Uploaded time.Time `json:"uploaded"`
	Size     Size
	Type     ImageType
	Overlay  ImageOverlay
}

type ImageOverlay struct {
	Heading string
	Text    string
	Button  Element
}

type ImageGallery struct {
	Images     []Image
	Columns    int
	Rows       int
	SeeMoreRow *int
}
