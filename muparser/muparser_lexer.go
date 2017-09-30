// Generated from MuParser.g4 by ANTLR 4.7.

package muparser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 37, 283,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 158, 10, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 172,
	10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 28, 3, 28, 3, 29, 6, 29, 231, 10, 29, 13, 29, 14, 29, 232, 3,
	30, 6, 30, 236, 10, 30, 13, 30, 14, 30, 237, 3, 30, 3, 30, 7, 30, 242,
	10, 30, 12, 30, 14, 30, 245, 11, 30, 3, 30, 3, 30, 6, 30, 249, 10, 30,
	13, 30, 14, 30, 250, 5, 30, 253, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 7, 35, 275, 10, 35, 12, 35, 14, 35,
	278, 11, 35, 3, 36, 3, 36, 3, 36, 3, 36, 2, 2, 37, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15,
	29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24,
	47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33,
	65, 34, 67, 35, 69, 36, 71, 37, 3, 2, 6, 3, 2, 50, 59, 5, 2, 67, 92, 97,
	97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 311, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3,
	2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47,
	3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2,
	55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2,
	2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2,
	2, 2, 71, 3, 2, 2, 2, 3, 73, 3, 2, 2, 2, 5, 75, 3, 2, 2, 2, 7, 157, 3,
	2, 2, 2, 9, 171, 3, 2, 2, 2, 11, 173, 3, 2, 2, 2, 13, 175, 3, 2, 2, 2,
	15, 178, 3, 2, 2, 2, 17, 181, 3, 2, 2, 2, 19, 184, 3, 2, 2, 2, 21, 187,
	3, 2, 2, 2, 23, 190, 3, 2, 2, 2, 25, 193, 3, 2, 2, 2, 27, 196, 3, 2, 2,
	2, 29, 199, 3, 2, 2, 2, 31, 202, 3, 2, 2, 2, 33, 205, 3, 2, 2, 2, 35, 207,
	3, 2, 2, 2, 37, 209, 3, 2, 2, 2, 39, 211, 3, 2, 2, 2, 41, 213, 3, 2, 2,
	2, 43, 215, 3, 2, 2, 2, 45, 217, 3, 2, 2, 2, 47, 219, 3, 2, 2, 2, 49, 221,
	3, 2, 2, 2, 51, 223, 3, 2, 2, 2, 53, 225, 3, 2, 2, 2, 55, 227, 3, 2, 2,
	2, 57, 230, 3, 2, 2, 2, 59, 252, 3, 2, 2, 2, 61, 254, 3, 2, 2, 2, 63, 259,
	3, 2, 2, 2, 65, 265, 3, 2, 2, 2, 67, 268, 3, 2, 2, 2, 69, 272, 3, 2, 2,
	2, 71, 279, 3, 2, 2, 2, 73, 74, 7, 12, 2, 2, 74, 4, 3, 2, 2, 2, 75, 76,
	7, 46, 2, 2, 76, 6, 3, 2, 2, 2, 77, 78, 7, 117, 2, 2, 78, 79, 7, 107, 2,
	2, 79, 158, 7, 112, 2, 2, 80, 81, 7, 101, 2, 2, 81, 82, 7, 113, 2, 2, 82,
	158, 7, 117, 2, 2, 83, 84, 7, 118, 2, 2, 84, 85, 7, 99, 2, 2, 85, 158,
	7, 112, 2, 2, 86, 87, 7, 99, 2, 2, 87, 88, 7, 117, 2, 2, 88, 89, 7, 107,
	2, 2, 89, 158, 7, 112, 2, 2, 90, 91, 7, 99, 2, 2, 91, 92, 7, 101, 2, 2,
	92, 93, 7, 113, 2, 2, 93, 158, 7, 117, 2, 2, 94, 95, 7, 99, 2, 2, 95, 96,
	7, 118, 2, 2, 96, 97, 7, 99, 2, 2, 97, 158, 7, 112, 2, 2, 98, 99, 7, 117,
	2, 2, 99, 100, 7, 107, 2, 2, 100, 101, 7, 112, 2, 2, 101, 158, 7, 106,
	2, 2, 102, 103, 7, 101, 2, 2, 103, 104, 7, 113, 2, 2, 104, 105, 7, 117,
	2, 2, 105, 158, 7, 106, 2, 2, 106, 107, 7, 118, 2, 2, 107, 108, 7, 99,
	2, 2, 108, 109, 7, 112, 2, 2, 109, 158, 7, 106, 2, 2, 110, 111, 7, 99,
	2, 2, 111, 112, 7, 117, 2, 2, 112, 113, 7, 107, 2, 2, 113, 114, 7, 112,
	2, 2, 114, 158, 7, 106, 2, 2, 115, 116, 7, 99, 2, 2, 116, 117, 7, 101,
	2, 2, 117, 118, 7, 113, 2, 2, 118, 119, 7, 117, 2, 2, 119, 158, 7, 106,
	2, 2, 120, 121, 7, 99, 2, 2, 121, 122, 7, 118, 2, 2, 122, 123, 7, 99, 2,
	2, 123, 124, 7, 112, 2, 2, 124, 158, 7, 106, 2, 2, 125, 126, 7, 110, 2,
	2, 126, 127, 7, 113, 2, 2, 127, 128, 7, 105, 2, 2, 128, 158, 7, 52, 2,
	2, 129, 130, 7, 110, 2, 2, 130, 131, 7, 113, 2, 2, 131, 132, 7, 105, 2,
	2, 132, 133, 7, 51, 2, 2, 133, 158, 7, 50, 2, 2, 134, 135, 7, 110, 2, 2,
	135, 136, 7, 113, 2, 2, 136, 158, 7, 105, 2, 2, 137, 138, 7, 110, 2, 2,
	138, 158, 7, 112, 2, 2, 139, 140, 7, 103, 2, 2, 140, 141, 7, 122, 2, 2,
	141, 158, 7, 114, 2, 2, 142, 143, 7, 117, 2, 2, 143, 144, 7, 115, 2, 2,
	144, 145, 7, 116, 2, 2, 145, 158, 7, 118, 2, 2, 146, 147, 7, 117, 2, 2,
	147, 148, 7, 107, 2, 2, 148, 149, 7, 105, 2, 2, 149, 158, 7, 112, 2, 2,
	150, 151, 7, 116, 2, 2, 151, 152, 7, 107, 2, 2, 152, 153, 7, 112, 2, 2,
	153, 158, 7, 118, 2, 2, 154, 155, 7, 99, 2, 2, 155, 156, 7, 100, 2, 2,
	156, 158, 7, 117, 2, 2, 157, 77, 3, 2, 2, 2, 157, 80, 3, 2, 2, 2, 157,
	83, 3, 2, 2, 2, 157, 86, 3, 2, 2, 2, 157, 90, 3, 2, 2, 2, 157, 94, 3, 2,
	2, 2, 157, 98, 3, 2, 2, 2, 157, 102, 3, 2, 2, 2, 157, 106, 3, 2, 2, 2,
	157, 110, 3, 2, 2, 2, 157, 115, 3, 2, 2, 2, 157, 120, 3, 2, 2, 2, 157,
	125, 3, 2, 2, 2, 157, 129, 3, 2, 2, 2, 157, 134, 3, 2, 2, 2, 157, 137,
	3, 2, 2, 2, 157, 139, 3, 2, 2, 2, 157, 142, 3, 2, 2, 2, 157, 146, 3, 2,
	2, 2, 157, 150, 3, 2, 2, 2, 157, 154, 3, 2, 2, 2, 158, 8, 3, 2, 2, 2, 159,
	160, 7, 111, 2, 2, 160, 161, 7, 107, 2, 2, 161, 172, 7, 112, 2, 2, 162,
	163, 7, 111, 2, 2, 163, 164, 7, 99, 2, 2, 164, 172, 7, 122, 2, 2, 165,
	166, 7, 117, 2, 2, 166, 167, 7, 119, 2, 2, 167, 172, 7, 111, 2, 2, 168,
	169, 7, 99, 2, 2, 169, 170, 7, 120, 2, 2, 170, 172, 7, 105, 2, 2, 171,
	159, 3, 2, 2, 2, 171, 162, 3, 2, 2, 2, 171, 165, 3, 2, 2, 2, 171, 168,
	3, 2, 2, 2, 172, 10, 3, 2, 2, 2, 173, 174, 7, 63, 2, 2, 174, 12, 3, 2,
	2, 2, 175, 176, 7, 45, 2, 2, 176, 177, 7, 63, 2, 2, 177, 14, 3, 2, 2, 2,
	178, 179, 7, 47, 2, 2, 179, 180, 7, 63, 2, 2, 180, 16, 3, 2, 2, 2, 181,
	182, 7, 44, 2, 2, 182, 183, 7, 63, 2, 2, 183, 18, 3, 2, 2, 2, 184, 185,
	7, 49, 2, 2, 185, 186, 7, 63, 2, 2, 186, 20, 3, 2, 2, 2, 187, 188, 7, 40,
	2, 2, 188, 189, 7, 40, 2, 2, 189, 22, 3, 2, 2, 2, 190, 191, 7, 126, 2,
	2, 191, 192, 7, 126, 2, 2, 192, 24, 3, 2, 2, 2, 193, 194, 7, 62, 2, 2,
	194, 195, 7, 63, 2, 2, 195, 26, 3, 2, 2, 2, 196, 197, 7, 64, 2, 2, 197,
	198, 7, 63, 2, 2, 198, 28, 3, 2, 2, 2, 199, 200, 7, 35, 2, 2, 200, 201,
	7, 63, 2, 2, 201, 30, 3, 2, 2, 2, 202, 203, 7, 63, 2, 2, 203, 204, 7, 63,
	2, 2, 204, 32, 3, 2, 2, 2, 205, 206, 7, 62, 2, 2, 206, 34, 3, 2, 2, 2,
	207, 208, 7, 64, 2, 2, 208, 36, 3, 2, 2, 2, 209, 210, 7, 45, 2, 2, 210,
	38, 3, 2, 2, 2, 211, 212, 7, 47, 2, 2, 212, 40, 3, 2, 2, 2, 213, 214, 7,
	44, 2, 2, 214, 42, 3, 2, 2, 2, 215, 216, 7, 49, 2, 2, 216, 44, 3, 2, 2,
	2, 217, 218, 7, 96, 2, 2, 218, 46, 3, 2, 2, 2, 219, 220, 7, 35, 2, 2, 220,
	48, 3, 2, 2, 2, 221, 222, 7, 65, 2, 2, 222, 50, 3, 2, 2, 2, 223, 224, 7,
	60, 2, 2, 224, 52, 3, 2, 2, 2, 225, 226, 7, 42, 2, 2, 226, 54, 3, 2, 2,
	2, 227, 228, 7, 43, 2, 2, 228, 56, 3, 2, 2, 2, 229, 231, 9, 2, 2, 2, 230,
	229, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 230, 3, 2, 2, 2, 232, 233,
	3, 2, 2, 2, 233, 58, 3, 2, 2, 2, 234, 236, 9, 2, 2, 2, 235, 234, 3, 2,
	2, 2, 236, 237, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2,
	238, 239, 3, 2, 2, 2, 239, 243, 7, 48, 2, 2, 240, 242, 9, 2, 2, 2, 241,
	240, 3, 2, 2, 2, 242, 245, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 243, 244,
	3, 2, 2, 2, 244, 253, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 246, 248, 7, 48,
	2, 2, 247, 249, 9, 2, 2, 2, 248, 247, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2,
	250, 248, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 253, 3, 2, 2, 2, 252,
	235, 3, 2, 2, 2, 252, 246, 3, 2, 2, 2, 253, 60, 3, 2, 2, 2, 254, 255, 7,
	118, 2, 2, 255, 256, 7, 116, 2, 2, 256, 257, 7, 119, 2, 2, 257, 258, 7,
	103, 2, 2, 258, 62, 3, 2, 2, 2, 259, 260, 7, 104, 2, 2, 260, 261, 7, 99,
	2, 2, 261, 262, 7, 110, 2, 2, 262, 263, 7, 117, 2, 2, 263, 264, 7, 103,
	2, 2, 264, 64, 3, 2, 2, 2, 265, 266, 7, 97, 2, 2, 266, 267, 7, 103, 2,
	2, 267, 66, 3, 2, 2, 2, 268, 269, 7, 97, 2, 2, 269, 270, 7, 114, 2, 2,
	270, 271, 7, 107, 2, 2, 271, 68, 3, 2, 2, 2, 272, 276, 9, 3, 2, 2, 273,
	275, 9, 4, 2, 2, 274, 273, 3, 2, 2, 2, 275, 278, 3, 2, 2, 2, 276, 274,
	3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 70, 3, 2, 2, 2, 278, 276, 3, 2,
	2, 2, 279, 280, 9, 5, 2, 2, 280, 281, 3, 2, 2, 2, 281, 282, 8, 36, 2, 2,
	282, 72, 3, 2, 2, 2, 11, 2, 157, 171, 232, 237, 243, 250, 252, 276, 3,
	8, 2, 2,
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
	"", "'\n'", "','", "", "", "'='", "'+='", "'-='", "'*='", "'/='", "'&&'",
	"'||'", "'<='", "'>='", "'!='", "'=='", "'<'", "'>'", "'+'", "'-'", "'*'",
	"'/'", "'^'", "'!'", "'?'", "':'", "'('", "')'", "", "", "'true'", "'false'",
	"'_e'", "'_pi'",
}

var lexerSymbolicNames = []string{
	"", "", "", "FUNCTION", "FUNCTIONMULTI", "ASSIGN", "ASSIGNADD", "ASSIGNSUB",
	"ASSIGNMUL", "ASSIGNDIV", "AND", "OR", "LTEQ", "GTEQ", "NEQ", "EQ", "LT",
	"GT", "ADD", "SUB", "MUL", "DIV", "POW", "NOT", "QUESTION", "COLON", "OPAR",
	"CPAR", "INT", "FLOAT", "TRUE", "FALSE", "E", "PI", "ID", "SPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "FUNCTION", "FUNCTIONMULTI", "ASSIGN", "ASSIGNADD", "ASSIGNSUB",
	"ASSIGNMUL", "ASSIGNDIV", "AND", "OR", "LTEQ", "GTEQ", "NEQ", "EQ", "LT",
	"GT", "ADD", "SUB", "MUL", "DIV", "POW", "NOT", "QUESTION", "COLON", "OPAR",
	"CPAR", "INT", "FLOAT", "TRUE", "FALSE", "E", "PI", "ID", "SPACE",
}

type MuParserLexer struct {
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

func NewMuParserLexer(input antlr.CharStream) *MuParserLexer {

	l := new(MuParserLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "MuParser.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MuParserLexer tokens.
const (
	MuParserLexerT__0          = 1
	MuParserLexerT__1          = 2
	MuParserLexerFUNCTION      = 3
	MuParserLexerFUNCTIONMULTI = 4
	MuParserLexerASSIGN        = 5
	MuParserLexerASSIGNADD     = 6
	MuParserLexerASSIGNSUB     = 7
	MuParserLexerASSIGNMUL     = 8
	MuParserLexerASSIGNDIV     = 9
	MuParserLexerAND           = 10
	MuParserLexerOR            = 11
	MuParserLexerLTEQ          = 12
	MuParserLexerGTEQ          = 13
	MuParserLexerNEQ           = 14
	MuParserLexerEQ            = 15
	MuParserLexerLT            = 16
	MuParserLexerGT            = 17
	MuParserLexerADD           = 18
	MuParserLexerSUB           = 19
	MuParserLexerMUL           = 20
	MuParserLexerDIV           = 21
	MuParserLexerPOW           = 22
	MuParserLexerNOT           = 23
	MuParserLexerQUESTION      = 24
	MuParserLexerCOLON         = 25
	MuParserLexerOPAR          = 26
	MuParserLexerCPAR          = 27
	MuParserLexerINT           = 28
	MuParserLexerFLOAT         = 29
	MuParserLexerTRUE          = 30
	MuParserLexerFALSE         = 31
	MuParserLexerE             = 32
	MuParserLexerPI            = 33
	MuParserLexerID            = 34
	MuParserLexerSPACE         = 35
)
