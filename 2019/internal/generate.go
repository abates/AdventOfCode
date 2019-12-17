package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
)

type values struct {
	Name     string
	Day      string
	DayNum   string
	Receiver string
	TypeName string
}

func readTemplate(name, filename string) *template.Template {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read %q: %v\n", filename, err)
		os.Exit(-1)
	}

	tmpl, err := template.New(name).Delims("${", "}").Parse(string(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse %q: %v\n", filename, err)
		os.Exit(-1)
	}
	return tmpl
}

func write(tmpl *template.Template, filename string, v *values) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	//f, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create %q: %v\n", filename, err)
		os.Exit(-1)
	}
	defer f.Close()

	buf := bytes.NewBuffer(make([]byte, 0))
	err = tmpl.Execute(buf, v)
	b, err := format.Source(buf.Bytes())
	if err != nil {
		f.Write(buf.Bytes()) // This is here to debug bad format
		fmt.Fprintf(os.Stderr, "Failed to format %q: %v\n", filename, err)
	}

	f.Write(b)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <day number>\n", os.Args[0])
		os.Exit(-1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Day number must be an integer\n")
		fmt.Fprintf(os.Stderr, "Usage: %s <day number>\n", os.Args[0])
		os.Exit(-1)
	}

	v := &values{
		Name:     fmt.Sprintf("Day %02d", day),
		Day:      fmt.Sprintf("%02d", day),
		DayNum:   fmt.Sprintf("%d", day),
		Receiver: fmt.Sprintf("d%d", day),
		TypeName: fmt.Sprintf("D%d", day),
	}

	write(readTemplate("day", "internal/day.tmpl"), fmt.Sprintf("day%02d.go", day), v)
	write(readTemplate("test", "internal/day_test.tmpl"), fmt.Sprintf("day%02d_test.go", day), v)
}
