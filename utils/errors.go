package utils

type ErrorCode int

const (
	MissingToken      ErrorCode = iota
	DiscordConnection ErrorCode = iota
)

var errorMessages = map[ErrorCode]string{
	MissingToken:      "Help! Cant find token from file or env",
	DiscordConnection: "Could not connect to discord... :(",
}

func (e ErrorCode) ErrorMessage() string {
	return errorMessages[e]
}
