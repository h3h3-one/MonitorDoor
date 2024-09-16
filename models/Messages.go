package models

type Messages struct {
	X     int8
	Y     int8
	Color int8
	Text  string
}

func NewMessages(x, y, color int8, text string) *Messages {
	return &Messages{
		X:     x,
		Y:     y,
		Color: color,
		Text:  text,
	}
}
