package repl

import (
	"bufio"
	"fmt"
	"io"
	"lana/lexer"
	"lana/parser"
)

const PROMPT = "ᏊᎾ ꈊᎾ Ꮚ >> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scannd := scanner.Scan()
		if !scannd {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, LANA_FACE)
	io.WriteString(out, "Baaaa! We ran into some lana business here🐑\n")
	io.WriteString(out, "parser errors!\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

const LANA_FACE = `
　 　　 /⌒⌒⌒ヽ
　  /~´　　　｀⌒ヽ
 　(　　 /⌒⌒⌒ヽ 　 )
　( ＠/ ・　 ・ ヽ＠ )
　(　し､　ꈊ　　Ｕ　　) < ᵇᵃᵃ ~
　　(　　ヽ——イ　　) 　
　 (　　　　　　　　) 　
　 　　“し——–Ｕ“
`
