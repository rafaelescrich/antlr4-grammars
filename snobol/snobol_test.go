// Package snobol_test contains tests for the snobol grammar.
// The tests should be run with the -timeout flag, to ensure the parser doesn't
// get stuck.
//
// Do not edit this file, it is generated by maketest.go
//
package snobol_test

import (
	"bramp.net/antlr4test-go/snobol"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
	"testing"
)

const MAX_TOKENS = 1000000

var examples = []string{
	"grammars-v4/snobol/examples/example1.sno",
	"grammars-v4/snobol/examples/example2.sno",
	"grammars-v4/snobol/examples/hello.sno",
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

func TestsnobolLexer(t *testing.T) {
	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := snobol.NewsnobolLexer(input)

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
			t.Errorf("NewsnobolLexer(%q) read %d tokens without finding EOF", file, i)
		}
	}
}

func TestsnobolParser(t *testing.T) {
	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := snobol.NewsnobolLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := snobol.NewsnobolParser(stream)
		p.BuildParseTrees = true
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true)) // TODO Change this
		p.AddErrorListener(antlr.NewConsoleErrorListener())

		// Finally test
		p.Prog()

		// TODO Check for errors
	}
}
