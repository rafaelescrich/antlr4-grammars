// Generated from DiceNotationLexer.g4 by ANTLR 4.7.

package dice

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 37, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 3, 2, 3, 2, 5, 2, 18, 10, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 6, 6, 27, 10, 6, 13, 6, 14, 6, 28, 3, 7, 6, 7, 32, 10, 7, 13, 7, 14,
	7, 33, 3, 7, 3, 7, 2, 2, 8, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 3, 2,
	4, 4, 2, 70, 70, 102, 102, 4, 2, 11, 12, 15, 15, 2, 39, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 3, 17, 3, 2, 2, 2, 5, 19, 3, 2, 2, 2, 7, 21, 3, 2,
	2, 2, 9, 23, 3, 2, 2, 2, 11, 26, 3, 2, 2, 2, 13, 31, 3, 2, 2, 2, 15, 18,
	5, 5, 3, 2, 16, 18, 5, 7, 4, 2, 17, 15, 3, 2, 2, 2, 17, 16, 3, 2, 2, 2,
	18, 4, 3, 2, 2, 2, 19, 20, 7, 45, 2, 2, 20, 6, 3, 2, 2, 2, 21, 22, 7, 47,
	2, 2, 22, 8, 3, 2, 2, 2, 23, 24, 9, 2, 2, 2, 24, 10, 3, 2, 2, 2, 25, 27,
	4, 50, 59, 2, 26, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 26, 3, 2, 2,
	2, 28, 29, 3, 2, 2, 2, 29, 12, 3, 2, 2, 2, 30, 32, 9, 3, 2, 2, 31, 30,
	3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2,
	34, 35, 3, 2, 2, 2, 35, 36, 8, 7, 2, 2, 36, 14, 3, 2, 2, 2, 6, 2, 17, 28,
	33, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "OPERATOR", "ADD", "SUB", "DSEPARATOR", "DIGIT", "WS",
}

var lexerRuleNames = []string{
	"OPERATOR", "ADD", "SUB", "DSEPARATOR", "DIGIT", "WS",
}

type DiceNotationLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewDiceNotationLexer(input antlr.CharStream) *DiceNotationLexer {

	l := new(DiceNotationLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "DiceNotationLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DiceNotationLexer tokens.
const (
	DiceNotationLexerOPERATOR   = 1
	DiceNotationLexerADD        = 2
	DiceNotationLexerSUB        = 3
	DiceNotationLexerDSEPARATOR = 4
	DiceNotationLexerDIGIT      = 5
	DiceNotationLexerWS         = 6
)
