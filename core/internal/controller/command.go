package controller

type Command struct {
	Subject    string      `json:"subject"`
	Subcommand interface{} `json:"subcommand"`
}
