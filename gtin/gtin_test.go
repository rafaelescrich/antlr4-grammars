// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package gtin_test contains tests for the gtin grammar.
// The tests should be run with the -timeout flag, to ensure the parser doesn't
// get stuck.
//
// Do not edit this file, it is generated by make.go
//
package gtin_test

import (
	"bramp.net/antlr4/gtin"
	"bramp.net/antlr4/internal"

	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
	"testing"
)

const MAX_TOKENS = 1000000

var examples = []string{
	"grammars-v4/gtin/examples/bookland1.txt",
	"grammars-v4/gtin/examples/bookland2.txt",
	"grammars-v4/gtin/examples/ean8.txt",
	"grammars-v4/gtin/examples/gtin13.txt",
	"grammars-v4/gtin/examples/gtin14_1.txt",
	"grammars-v4/gtin/examples/ismn.txt",
	"grammars-v4/gtin/examples/issn.txt",
	"grammars-v4/gtin/examples/sup2.txt",
	"grammars-v4/gtin/examples/sup5.txt",
	"grammars-v4/gtin/examples/upc_a_1.txt",
	"grammars-v4/gtin/examples/upc_a_1_hyphen.txt",
	"grammars-v4/gtin/examples/upc_e_1.txt",
}

type exampleListener struct {
	*gtin.BasegtinListener
}

func (l *exampleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func Example() {
	// Setup the input
	is := antlr.NewInputStream("...some text to parse...")

	// Create the Lexer
	lexer := gtin.NewgtinLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := gtin.NewgtinParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Finally walk the tree
	tree := p.Gtin()
	antlr.ParseTreeWalkerDefault.Walk(&exampleListener{}, tree)
}

func newCharStream(filename string) (antlr.CharStream, error) {
	var input antlr.CharStream
	input, err := antlr.NewFileStream(filepath.Join("../", filename))
	if err != nil {
		return nil, err
	}

	return input, nil
}

func TestGtinLexer(t *testing.T) {
	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := gtin.NewgtinLexer(input)

		// Try and read all tokens
		i := 0
		for ; i < MAX_TOKENS; i++ {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
		}

		// If we read too many tokens, then perhaps there is a problem with the lexer.
		if i >= MAX_TOKENS {
			t.Errorf("NewgtinLexer(%q) read %d tokens without finding EOF", file, i)
		}
	}
}

func TestGtinParser(t *testing.T) {
	// TODO(bramp): Run this test with and without p.BuildParseTrees

	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := gtin.NewgtinLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := gtin.NewgtinParser(stream)
		p.BuildParseTrees = true
		p.AddErrorListener(internal.NewTestingErrorListener(t, file))

		// Finally test
		p.Gtin()

		// TODO(bramp): If there is a "file.tree", then compare the output
	}
}
