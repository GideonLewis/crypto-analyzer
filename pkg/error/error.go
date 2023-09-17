package error

type Error struct {
	Raw      error  `json:"raw"`
	Message  string `json:"message"`
	Classify string `json:"classify"`
}

func (e Error) ToString() string {
	if e.Message == "" {
		e.Message = e.Raw.Error()
	}
	return e.Message
}

func (e Error) Is(compare string) bool {
	return true
}

func (e Error) Format() Error {
	return Error{}
}

func (e Error) GetRaw() error {
	return e.Raw
}

func (e Error) GetMessage() string {
	return e.Message
}

func (e Error) GetClassify() string {
	return e.Classify
}

func NewError(raw error, message string, classify string) Error {
	return Error{
		Raw:      raw,
		Message:  message,
		Classify: classify,
	}
}
