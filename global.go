package tools

import "log"

const (
	DEBUG   = 1
	RELEASE = 0
)

var Mode = RELEASE
var Logger *log.Logger
