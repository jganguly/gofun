package main

/* Struct for Req for both input from client and output from server */
type Ses struct {
	Key string // uppercase for json.Marshall to work
	Val string
}
