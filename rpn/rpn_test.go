// Package rpn_test contains tests for the rpn grammar.
// The tests should be run with the -timeout flag, to ensure the parser doesn't
// get stuck.
//
// Do not edit this file, it is generated by maketest.go
//
package rpn_test

import (
	"bramp.net/antlr4test-go/rpn"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
	"testing"
)

const MAX_TOKENS = 1000000

var examples = []string{
	"grammars-v4/rpn/examples/cos.txt",
	"grammars-v4/rpn/examples/number1.txt",
	"grammars-v4/rpn/examples/number2.txt",
	"grammars-v4/rpn/examples/number3.txt",
	"grammars-v4/rpn/examples/number4.txt",
	"grammars-v4/rpn/examples/number5.txt",
	"grammars-v4/rpn/examples/pow1.txt",
	"grammars-v4/rpn/examples/precedence1.txt",
	"grammars-v4/rpn/examples/pythagoras.txt",
	"grammars-v4/rpn/examples/pythagoras2.txt",
	"grammars-v4/rpn/examples/simple.txt",
	"grammars-v4/rpn/examples/variable1.txt",
	"grammars-v4/rpn/examples/variable2.txt",
}

func newCharStream(filename string) (antlr.CharStream, error) {
	var input antlr.CharStream
	input, err := antlr.NewFileStream(filepath.Join("..", filename))
	if err != nil {
		return nil, err
	}

	return input, nil
}

// TODO Add an Example func

func TestrpnLexer(t *testing.T) {
	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := rpn.NewrpnLexer(input)

		// Try and read all tokens
		i := 0
		for ; i < MAX_TOKENS; i++ {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
		}

		// If we read too many tokens, then perhaps there is a problem with the lexer.
		if i == MAX_TOKENS {
			t.Errorf("NewrpnLexer(%q) read %d tokens without finding EOF", file, i)
		}
	}
}

func TestrpnParser(t *testing.T) {
	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := rpn.NewrpnLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := rpn.NewrpnParser(stream)
		p.BuildParseTrees = true
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true)) // TODO Change this
		p.AddErrorListener(antlr.NewConsoleErrorListener())

		// Finally test
		p.Expression()

		// TODO Check for errors
	}
}
