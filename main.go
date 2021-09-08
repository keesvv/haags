package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

type Dict = map[*regexp.Regexp]string

type DictRaw = map[string]string

func main() {
	dict := *getDict()
	r := bufio.NewReader(os.Stdin)

	read := func() ([]byte, error) {
		return r.ReadBytes('\n')
	}

	for b, err := read(); err != io.EOF; b, err = read() {
		line := string(b)

		for key, value := range dict {
			line = key.ReplaceAllLiteralString(line, value)
		}

		fmt.Print(line)
	}
}

func getDict() *Dict {
	raw := getRawDict()
	dict := Dict{}

	for k, v := range *raw {
		dict[regexp.MustCompile(k)] = v
	}

	return &dict
}
