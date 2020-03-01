package model

/*Command struct*/
type Command struct {
	Call   string `json:"call"`
	Result string `json:"result"`
}

/*Validate Command struct*/
func (command *Command) Validate() (string, bool) {

	if command.Call == "" {
		return "command call is missing", false
	}
	return "", true
}
