package user

//go:generate myenumstr -type Status,Color

type Status int

const (
	Offline Status = iota
	Online
	Disable
	Deleted
)

type Color int

const (
	Write Color = iota
	Red
	Blue
)
