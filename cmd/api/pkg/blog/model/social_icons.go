package model

type SocialIcons struct {
	Icons  []SocialIcon
	Size   Size
	Layout Layout
	Align  Align
}

type SocialIcon struct {
	PlatformName string
	PlatformLogo string
	ProfileLink  string
}
