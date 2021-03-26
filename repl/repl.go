package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/MihaiBlebea/go-interpreter/lexer"
	"github.com/MihaiBlebea/go-interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	history := make([]string, 0)

	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		history = append(history, line)

		if line == "clear" {
			clearConsole()
			continue
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
