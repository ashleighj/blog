package content

type PageHeader struct {
	NavMenu Menu
	Logo    string
}

type PageFooter struct {
	Menus             []Menu
	Logo              string
	Copyright         string
	Social            SocialIcons
	LegalLink         string
	PrivacyPolicyLink string
}
