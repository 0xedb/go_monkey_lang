package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/0xedb/go_monkey_lang/token"

	"github.com/0xedb/go_monkey_lang/lexer"
)

const PROMPT = "🔥🔥🔥 "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf("%s", PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
