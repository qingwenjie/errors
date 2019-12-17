package errors

import (
	"bytes"
	"strings"
)

type stackError struct {
	fields Fields
	stacks []string
	err    error
}

func (s *stackError) WithField(key string, val interface{}) {
	s.fields[key] = val
}

func (s *stackError) WithFields(fields Fields) {
	s.fields.Merge(fields)
}

func (s *stackError) Fields() Fields {
	s.fields["stack"] = s.stack()
	return s.fields
}

func (s stackError) stack() string {
	if len(s.stacks) == 0 {
		return ""
	}
	prev := s.stacks[0]
	var b bytes.Buffer
	for index, stack := range s.stacks {
		if index != 0 &&
			stack == prev {
			continue
		}
		prev = stack
		b.WriteString(stack)
		b.WriteString(";")
	}
	return strings.TrimSuffix(b.String(), ";")
}

func (s *stackError) Error() string {
	return s.err.Error()
}

func (s *stackError) Cause() error {
	return s.err
}