package errors

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func Trace(err error) error {
	if err == nil {
		return nil
	}
	return NewEx(2, err, nil)
}

func NewEx(depth int, err error, fields Fields) error {
	stackFrame := Get(depth)
	se, ok := err.(*stackError)
	if !ok {
		if fields == nil {
			fields = make(Fields)
		}
		return &stackError{
			err:    err,
			fields: fields,
			stacks: []string{stackFrame},
		}
	}
	se.stacks = append(se.stacks, stackFrame)
	if se.fields == nil {
		se.fields = make(Fields)
	}
	for k, v := range fields {
		se.fields[k] = v
	}
	return se
}

func TraceWithFields(err error, fields Fields) error {
	return TraceWithFieldsEx(err, fields, 1)
}

func TraceWithFieldsEx(err error, fields Fields, depth int) error {
	if err == nil {
		return nil
	}
	return NewEx(depth+2, err, fields)
}

func TraceWithFieldEx(err error, key string, val interface{}, depth int) error {
	if err == nil {
		return nil
	}
	fs := make(Fields)
	fs[key] = val
	return NewEx(depth+2, err, fs)
}

func TraceWithField(err error, key string, val interface{}) error {
	return TraceWithFieldEx(err, key, val, 1)

}

func Cause(err error) error {
	if stackErr, ok := err.(*stackError); ok {
		return stackErr.Cause()
	}
	return err
}

func Is(src, dst error) bool {
	return Cause(src) == Cause(dst)
}

func Get(depth int) string {
	pc, _, line, ok := runtime.Caller(depth + 1)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v:%v", filepath.Base(runtime.FuncForPC(pc).Name()), line)
}