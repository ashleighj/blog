package model

type Size string

const (
	SizeXS Size = "XSMALL"
	SizeS  Size = "SMALL"
	SizeM  Size = "MEDIUM"
	SizeL  Size = "LARGE"
	SizeXL Size = "XLARGE"
)

type Layout string

const (
	LayoutVert = "LAYOUT_VERT"
	LayoutHor  = "LAYOUT_HOR"
)

type Align string

const (
	AlignTop    Align = "ALIGN_TOP"
	AlignBottom Align = "ALIGN_BOTTOM"
	AlignLeft   Align = "ALIGN_LEFT"
	AlighRight  Align = "ALIGN_RIGHT"
	AlignCenter Align = "ALIGN_CENTER"
)

type TextType string

const (
	TxtTitle    TextType = "TXT_TITLE"
	TxtHeading1 TextType = "TXT_HEAD_1"
	TxtHeading2 TextType = "TXT_HEAD_2"
	TxtHeading3 TextType = "TXT_HEAD_3"
	TxtHeading4 TextType = "TXT_HEAD_4"
	TxtNormal   TextType = "TXT_NORM"
)
