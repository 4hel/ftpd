package command

type CommandType int

const (
	Close CommandType = 1 << iota
	User
)

type Command struct {
	Kind CommandType
}


