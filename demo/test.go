package main

import (
	errors2 "errors"
	"fmt"
	"github.com/qingwenjie/errors"
)

func main() {
	err := err1()
	e := errors.TraceWithField(err, "f1", "aaaa")
	fmt.Println(e)
}

func err1() error {
	c := testErr{}
	return c.error1()
}

type testErr struct{}

func (s *testErr) error1() error {
	err := s.error2()
	return errors.Trace(err)
}

func (s *testErr) error2() error {
	err := errors2.New("test error2")
	return errors.Trace(err)
}
