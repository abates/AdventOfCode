package util

import (
	"bytes"
	"fmt"
)

type StringWriter struct {
	buffer bytes.Buffer
}

func (s *StringWriter) Write(p string) (int, error) {
	return (&s.buffer).Write([]byte(p))
}

func (s *StringWriter) Writef(format string, args ...interface{}) {
	s.Write(fmt.Sprintf(format, args...))
}

func (s *StringWriter) String() string {
	return (&s.buffer).String()
}
