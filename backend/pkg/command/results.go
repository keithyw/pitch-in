package command

type CommandResults struct {
	Data map[string]any
	Errors string
	IsSuccess bool
}