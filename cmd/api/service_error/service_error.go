package err

type ServiceError struct {
	Message    string          `json:"message"`
	Code       string          `json:"code"`
	Properties []PropertyError `json:"properties,omitempty"`
}

type PropertyError struct {
	Property string `json:"property"`
	Message  string `json:"message"`
}

func New(msg string, code string) ServiceError {
	return ServiceError{Message: msg, Code: code}
}

func NewProps(props []PropertyError) ServiceError {
	return ServiceError{
		Message:    MsgProps,
		Code:       CodeProps,
		Properties: props}
}

func NewProp(prop string, msg string, e error) ServiceError {
	propErr := PropertyError{Property: prop, Message: msg}

	if e != nil {
		if se, ok := e.(ServiceError); ok {
			se.Properties = append(se.Properties, propErr)
			return se
		}
	}

	return ServiceError{
		Message:    MsgProps,
		Code:       CodeProps,
		Properties: []PropertyError{propErr}}
}

func (e ServiceError) Error() string {
	return e.Message
}
