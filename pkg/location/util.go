package location

import "time"

type Message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
	Time     string `json:"time"`
}

func FormatMessage(username string, msg string) Message {
	return Message{
		Username: username,
		Text:     msg,
		Time:     time.Now().Format(time.RFC3339),
	}
}
