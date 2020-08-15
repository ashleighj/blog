package content

import "time"

type Article struct {
	Title            string            `json:"title"`
	Created          time.Time         `json:"created"`
	LastMod          time.Time         `json:"last_mod"`
	PubDate          time.Time         `json:"pub_date"`
	IsPublished      bool              `json:"is_published"`
	ReadTime         time.Time         `json:"read_time"`
	CreatedBy        Contributor       `json:"created_by"`
	Contributors     []Contributor     `json:"contributors"`
	Tags             []Tag             `json:"tags"`
	EditCache        string            `json:"edit_cache"`
	HTML             string            `json:"html"`
	PreviousVersions []PreviousVersion `json:"previous_versions"`
	Images           []Image           `json:"images"`
	Comments         []Comment         `json:"comments"`
	Upvote           []Upvote          `json:"upvote"`
	SEO              SEO               `json:"seo"`
}

type PreviousVersion struct {
	VersionNo int    `json:"version"`
	HTML      string `json:"html"`
}

type Tag struct {
	Tag string `json:"tag"`
}

type Comment struct {
	Contributor Contributor `json:"contributor"`
	Date        time.Time   `json:"date"`
	Text        string      `json:"test"`
	Replies     []Comment   `json:"replies"`
	Upvotes     []Upvote    `json:"upvotes"`
}

type Upvote struct {
	Contributor Contributor `json:"contributor"`
	Date        time.Time   `json:"date"`
}

type SEO struct {
}
