package utils

type CreateOptions struct {
	debug   bool
	provider string
	masterCount int
	workerCount int
}

var createOptions = &CreateOptions{}