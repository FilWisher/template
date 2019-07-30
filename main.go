package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var stdin bool

func main() {
	flag.BoolVar(&stdin, "stdin", false, "Read template from standard in")
	flag.Parse()

	cmd := os.Args[0]
	args := flag.Args()

	if !stdin {

		if len(args) == 0 {
			fmt.Fprintf(os.Stderr, "usage: %s <file> [args]\n", cmd)
			os.Exit(1)
		}

		file, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", cmd, err)
			os.Exit(1)
		}

		buf, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", cmd, err)
			os.Exit(1)
		}

		tmpl, err := template.New("").Parse(string(buf))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", cmd, err)
			os.Exit(1)
		}

		ctx, err := parseargs(args[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", cmd, err)
			os.Exit(1)
		}

		err = tmpl.Execute(os.Stdout, ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", cmd, err)
			os.Exit(1)
		}

		return
	}

	ctx, err := parseargs(args[0:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", cmd, err)
		os.Exit(1)
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		tmpl, err := template.New("").Parse(sc.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", cmd, err)
			os.Exit(1)
		}

		err = tmpl.Execute(os.Stdout, ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", cmd, err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout)
	}

}

func parseargs(args []string) (map[string]string, error) {
	m := make(map[string]string)
	for _, arg := range args {
		split := strings.Split(arg, "=")
		if len(split) <= 1 {
			return nil, fmt.Errorf("parseargs: badly formatted argument %s", arg)
		}
		m[split[0]] = strings.Join(split[1:], "=")
	}
	return m, nil
}
