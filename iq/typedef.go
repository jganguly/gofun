package iq

type structLoc struct {
	line int
	cpos int
}

type Ses struct {
	Key  string			// uppercase for json.Marshall to work
	Val	 string
}

