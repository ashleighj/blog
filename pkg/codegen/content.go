package codegen

type Type string
type TypeMod string

const (
	TypeString  Type    = "string"
	TypeInt     Type    = "int"
	TypeInt32   Type    = "int32"
	TypeInt64   Type    = "int64"
	TypeFloat   Type    = "float"
	TypeFloat32 Type    = "float32"
	TypeFloat64 Type    = "float64"
	TypeModList TypeMod = "[]"
)

type NewContent struct {
	Name  string
	Attrs []Attribute
}

type Attribute struct {
	Name string
	Type string
}
