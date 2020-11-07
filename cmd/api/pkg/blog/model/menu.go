package model

type MenuType string

const (
	MenuTypeVert MenuType = "MENU_VERT"
	MenuTypeHor  MenuType = "MENU_HOR"
)

type MenuItemType string

const (
	MenuItemMenu MenuItemType = "MENU_ITEM_MENU"
	MenuItemLink MenuItemType = "MENU_ITEM_LINK"
	MenuItemPage MenuItemType = "MENU_ITEM_PAGE"
)

type Menu struct {
	VAlign Align
	HAlign Align
	Type   MenuType
	Items  []MenuItem
}

type MenuItem struct {
	Name string
	Type MenuItemType
	// ChildMenu would contain a ChildMenu, if Type == MenuItemMenu
	ChildMenu Menu
	// Link would be populated with an external link if Type == MenuTypeLink
	Link string
	// Page would have a page ID if Type == MenuTypePage
	PageID int64
}
