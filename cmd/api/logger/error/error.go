package error

type Error struct {
	Message  string
	Code     string
	Original error
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) Log() {
	
}
