package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var tmpl = `
package main

var (
    replName     = "{{.Name}}"
    prompt       = "{{.Name}} > "
    exitMessage  = "exiting {{.Name}}... bye!"
)
`

type config struct {
	Name string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: genconfig <name>")
		os.Exit(1)
	}
	name := strings.TrimSpace(os.Args[1])

	cfg := config{Name: name}

	f, err := os.Create("config.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	t := template.Must(template.New("").Parse(tmpl))
	if err := t.Execute(f, cfg); err != nil {
		panic(err)
	}
}
