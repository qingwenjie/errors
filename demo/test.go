package main

import (
	errors2 "errors"
	"github.com/qingwenjie/errors"
)

func main() {
	err := errors2.New("test error")
	errors.Trace(err)
}
