package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tokenizer/parser"
)

const PROMPT = ">>"

func main() {
	in := os.Stdin
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		fields := strings.Fields(line)
		for _, f := range fields {
			p := parser.New(f)
			t := p.ParseNumber()
			fmt.Println(t.String())
		}
	}
}
