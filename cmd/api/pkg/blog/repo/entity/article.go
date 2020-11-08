package entity

type Article struct {
	ID                   int64
	Title                string
	Created              string
	LastMod              string
	PubDate              string
	IsPublished          bool
	ReadTime             string
	PrimaryContributorID int64
	HTML                 string
}

type PreviousVersion struct {
	ArticleID int64
	VersionNo int
	HTML      string
}

type Tag struct {
	ID        int64
	ArticleID int64
	Tag       string
}

type Comment struct {
	ID              int64
	ArticleID       int64
	ContributorID   int64
	ParentCommentID int64
	Date            string
	Text            string
	Upvotes         []Upvote
}

type Upvote struct {
	ID            int64
	ContributorID int64
	TargetID      int64
	Date          string
}

// COMMENT or ARTICLE
type UpvoteTarget struct {
	ID     int64
	Target string
}

type SEO struct {
}
