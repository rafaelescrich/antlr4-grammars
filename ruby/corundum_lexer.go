// Generated from Corundum.g4 by ANTLR 4.7.

package ruby

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 62, 409,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 7, 3, 132, 10, 3, 12, 3, 14, 3, 135, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	7, 3, 141, 10, 3, 12, 3, 14, 3, 144, 11, 3, 3, 3, 5, 3, 147, 10, 3, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 6, 5, 6, 154, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33,
	3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3,
	37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40,
	3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3,
	45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 5, 47,
	308, 10, 47, 3, 48, 3, 48, 3, 48, 3, 48, 5, 48, 314, 10, 48, 3, 49, 3,
	49, 3, 49, 3, 49, 5, 49, 320, 10, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52,
	3, 52, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 7, 55, 336,
	10, 55, 12, 55, 14, 55, 339, 11, 55, 3, 55, 5, 55, 342, 10, 55, 3, 55,
	3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3,
	56, 7, 56, 356, 10, 56, 12, 56, 14, 56, 359, 11, 56, 3, 56, 3, 56, 3, 56,
	3, 56, 3, 56, 3, 56, 5, 56, 367, 10, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3,
	57, 6, 57, 374, 10, 57, 13, 57, 14, 57, 375, 3, 57, 3, 57, 3, 58, 6, 58,
	381, 10, 58, 13, 58, 14, 58, 382, 3, 59, 7, 59, 386, 10, 59, 12, 59, 14,
	59, 389, 11, 59, 3, 59, 3, 59, 6, 59, 393, 10, 59, 13, 59, 14, 59, 394,
	3, 60, 3, 60, 7, 60, 399, 10, 60, 12, 60, 14, 60, 402, 11, 60, 3, 61, 3,
	61, 3, 61, 3, 62, 3, 62, 3, 62, 5, 133, 142, 357, 2, 63, 3, 2, 5, 3, 7,
	4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
	14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
	23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63,
	32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81,
	41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99,
	50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 56, 113, 57, 115,
	58, 117, 59, 119, 60, 121, 61, 123, 62, 3, 2, 8, 4, 2, 12, 12, 15, 15,
	4, 2, 11, 11, 34, 34, 3, 2, 50, 59, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2,
	50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 65, 65, 2, 425, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2,
	2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3,
	2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37,
	3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2,
	45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2,
	2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2,
	2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2,
	2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3,
	2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83,
	3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2,
	91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2,
	2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2,
	2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113,
	3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2,
	2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 3, 125, 3, 2, 2, 2, 5, 146, 3,
	2, 2, 2, 7, 148, 3, 2, 2, 2, 9, 150, 3, 2, 2, 2, 11, 153, 3, 2, 2, 2, 13,
	157, 3, 2, 2, 2, 15, 165, 3, 2, 2, 2, 17, 169, 3, 2, 2, 2, 19, 173, 3,
	2, 2, 2, 21, 180, 3, 2, 2, 2, 23, 184, 3, 2, 2, 2, 25, 187, 3, 2, 2, 2,
	27, 192, 3, 2, 2, 2, 29, 198, 3, 2, 2, 2, 31, 205, 3, 2, 2, 2, 33, 211,
	3, 2, 2, 2, 35, 217, 3, 2, 2, 2, 37, 223, 3, 2, 2, 2, 39, 227, 3, 2, 2,
	2, 41, 232, 3, 2, 2, 2, 43, 238, 3, 2, 2, 2, 45, 240, 3, 2, 2, 2, 47, 242,
	3, 2, 2, 2, 49, 244, 3, 2, 2, 2, 51, 246, 3, 2, 2, 2, 53, 248, 3, 2, 2,
	2, 55, 251, 3, 2, 2, 2, 57, 254, 3, 2, 2, 2, 59, 257, 3, 2, 2, 2, 61, 259,
	3, 2, 2, 2, 63, 261, 3, 2, 2, 2, 65, 264, 3, 2, 2, 2, 67, 267, 3, 2, 2,
	2, 69, 269, 3, 2, 2, 2, 71, 272, 3, 2, 2, 2, 73, 275, 3, 2, 2, 2, 75, 278,
	3, 2, 2, 2, 77, 281, 3, 2, 2, 2, 79, 284, 3, 2, 2, 2, 81, 288, 3, 2, 2,
	2, 83, 290, 3, 2, 2, 2, 85, 292, 3, 2, 2, 2, 87, 294, 3, 2, 2, 2, 89, 296,
	3, 2, 2, 2, 91, 299, 3, 2, 2, 2, 93, 307, 3, 2, 2, 2, 95, 313, 3, 2, 2,
	2, 97, 319, 3, 2, 2, 2, 99, 321, 3, 2, 2, 2, 101, 323, 3, 2, 2, 2, 103,
	325, 3, 2, 2, 2, 105, 327, 3, 2, 2, 2, 107, 329, 3, 2, 2, 2, 109, 333,
	3, 2, 2, 2, 111, 347, 3, 2, 2, 2, 113, 373, 3, 2, 2, 2, 115, 380, 3, 2,
	2, 2, 117, 387, 3, 2, 2, 2, 119, 396, 3, 2, 2, 2, 121, 403, 3, 2, 2, 2,
	123, 406, 3, 2, 2, 2, 125, 126, 7, 94, 2, 2, 126, 127, 7, 36, 2, 2, 127,
	4, 3, 2, 2, 2, 128, 133, 7, 36, 2, 2, 129, 132, 5, 3, 2, 2, 130, 132, 10,
	2, 2, 2, 131, 129, 3, 2, 2, 2, 131, 130, 3, 2, 2, 2, 132, 135, 3, 2, 2,
	2, 133, 134, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 134, 136, 3, 2, 2, 2, 135,
	133, 3, 2, 2, 2, 136, 147, 7, 36, 2, 2, 137, 142, 7, 41, 2, 2, 138, 141,
	5, 3, 2, 2, 139, 141, 10, 2, 2, 2, 140, 138, 3, 2, 2, 2, 140, 139, 3, 2,
	2, 2, 141, 144, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2,
	143, 145, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 147, 7, 41, 2, 2, 146,
	128, 3, 2, 2, 2, 146, 137, 3, 2, 2, 2, 147, 6, 3, 2, 2, 2, 148, 149, 7,
	46, 2, 2, 149, 8, 3, 2, 2, 2, 150, 151, 7, 61, 2, 2, 151, 10, 3, 2, 2,
	2, 152, 154, 7, 15, 2, 2, 153, 152, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154,
	155, 3, 2, 2, 2, 155, 156, 7, 12, 2, 2, 156, 12, 3, 2, 2, 2, 157, 158,
	7, 116, 2, 2, 158, 159, 7, 103, 2, 2, 159, 160, 7, 115, 2, 2, 160, 161,
	7, 119, 2, 2, 161, 162, 7, 107, 2, 2, 162, 163, 7, 116, 2, 2, 163, 164,
	7, 103, 2, 2, 164, 14, 3, 2, 2, 2, 165, 166, 7, 103, 2, 2, 166, 167, 7,
	112, 2, 2, 167, 168, 7, 102, 2, 2, 168, 16, 3, 2, 2, 2, 169, 170, 7, 102,
	2, 2, 170, 171, 7, 103, 2, 2, 171, 172, 7, 104, 2, 2, 172, 18, 3, 2, 2,
	2, 173, 174, 7, 116, 2, 2, 174, 175, 7, 103, 2, 2, 175, 176, 7, 118, 2,
	2, 176, 177, 7, 119, 2, 2, 177, 178, 7, 116, 2, 2, 178, 179, 7, 112, 2,
	2, 179, 20, 3, 2, 2, 2, 180, 181, 7, 114, 2, 2, 181, 182, 7, 107, 2, 2,
	182, 183, 7, 116, 2, 2, 183, 22, 3, 2, 2, 2, 184, 185, 7, 107, 2, 2, 185,
	186, 7, 104, 2, 2, 186, 24, 3, 2, 2, 2, 187, 188, 7, 103, 2, 2, 188, 189,
	7, 110, 2, 2, 189, 190, 7, 117, 2, 2, 190, 191, 7, 103, 2, 2, 191, 26,
	3, 2, 2, 2, 192, 193, 7, 103, 2, 2, 193, 194, 7, 110, 2, 2, 194, 195, 7,
	117, 2, 2, 195, 196, 7, 107, 2, 2, 196, 197, 7, 104, 2, 2, 197, 28, 3,
	2, 2, 2, 198, 199, 7, 119, 2, 2, 199, 200, 7, 112, 2, 2, 200, 201, 7, 110,
	2, 2, 201, 202, 7, 103, 2, 2, 202, 203, 7, 117, 2, 2, 203, 204, 7, 117,
	2, 2, 204, 30, 3, 2, 2, 2, 205, 206, 7, 121, 2, 2, 206, 207, 7, 106, 2,
	2, 207, 208, 7, 107, 2, 2, 208, 209, 7, 110, 2, 2, 209, 210, 7, 103, 2,
	2, 210, 32, 3, 2, 2, 2, 211, 212, 7, 116, 2, 2, 212, 213, 7, 103, 2, 2,
	213, 214, 7, 118, 2, 2, 214, 215, 7, 116, 2, 2, 215, 216, 7, 123, 2, 2,
	216, 34, 3, 2, 2, 2, 217, 218, 7, 100, 2, 2, 218, 219, 7, 116, 2, 2, 219,
	220, 7, 103, 2, 2, 220, 221, 7, 99, 2, 2, 221, 222, 7, 109, 2, 2, 222,
	36, 3, 2, 2, 2, 223, 224, 7, 104, 2, 2, 224, 225, 7, 113, 2, 2, 225, 226,
	7, 116, 2, 2, 226, 38, 3, 2, 2, 2, 227, 228, 7, 118, 2, 2, 228, 229, 7,
	116, 2, 2, 229, 230, 7, 119, 2, 2, 230, 231, 7, 103, 2, 2, 231, 40, 3,
	2, 2, 2, 232, 233, 7, 104, 2, 2, 233, 234, 7, 99, 2, 2, 234, 235, 7, 110,
	2, 2, 235, 236, 7, 117, 2, 2, 236, 237, 7, 103, 2, 2, 237, 42, 3, 2, 2,
	2, 238, 239, 7, 45, 2, 2, 239, 44, 3, 2, 2, 2, 240, 241, 7, 47, 2, 2, 241,
	46, 3, 2, 2, 2, 242, 243, 7, 44, 2, 2, 243, 48, 3, 2, 2, 2, 244, 245, 7,
	49, 2, 2, 245, 50, 3, 2, 2, 2, 246, 247, 7, 39, 2, 2, 247, 52, 3, 2, 2,
	2, 248, 249, 7, 44, 2, 2, 249, 250, 7, 44, 2, 2, 250, 54, 3, 2, 2, 2, 251,
	252, 7, 63, 2, 2, 252, 253, 7, 63, 2, 2, 253, 56, 3, 2, 2, 2, 254, 255,
	7, 35, 2, 2, 255, 256, 7, 63, 2, 2, 256, 58, 3, 2, 2, 2, 257, 258, 7, 64,
	2, 2, 258, 60, 3, 2, 2, 2, 259, 260, 7, 62, 2, 2, 260, 62, 3, 2, 2, 2,
	261, 262, 7, 62, 2, 2, 262, 263, 7, 63, 2, 2, 263, 64, 3, 2, 2, 2, 264,
	265, 7, 64, 2, 2, 265, 266, 7, 63, 2, 2, 266, 66, 3, 2, 2, 2, 267, 268,
	7, 63, 2, 2, 268, 68, 3, 2, 2, 2, 269, 270, 7, 45, 2, 2, 270, 271, 7, 63,
	2, 2, 271, 70, 3, 2, 2, 2, 272, 273, 7, 47, 2, 2, 273, 274, 7, 63, 2, 2,
	274, 72, 3, 2, 2, 2, 275, 276, 7, 44, 2, 2, 276, 277, 7, 63, 2, 2, 277,
	74, 3, 2, 2, 2, 278, 279, 7, 49, 2, 2, 279, 280, 7, 63, 2, 2, 280, 76,
	3, 2, 2, 2, 281, 282, 7, 39, 2, 2, 282, 283, 7, 63, 2, 2, 283, 78, 3, 2,
	2, 2, 284, 285, 7, 44, 2, 2, 285, 286, 7, 44, 2, 2, 286, 287, 7, 63, 2,
	2, 287, 80, 3, 2, 2, 2, 288, 289, 7, 40, 2, 2, 289, 82, 3, 2, 2, 2, 290,
	291, 7, 126, 2, 2, 291, 84, 3, 2, 2, 2, 292, 293, 7, 96, 2, 2, 293, 86,
	3, 2, 2, 2, 294, 295, 7, 128, 2, 2, 295, 88, 3, 2, 2, 2, 296, 297, 7, 62,
	2, 2, 297, 298, 7, 62, 2, 2, 298, 90, 3, 2, 2, 2, 299, 300, 7, 64, 2, 2,
	300, 301, 7, 64, 2, 2, 301, 92, 3, 2, 2, 2, 302, 303, 7, 99, 2, 2, 303,
	304, 7, 112, 2, 2, 304, 308, 7, 102, 2, 2, 305, 306, 7, 40, 2, 2, 306,
	308, 7, 40, 2, 2, 307, 302, 3, 2, 2, 2, 307, 305, 3, 2, 2, 2, 308, 94,
	3, 2, 2, 2, 309, 310, 7, 113, 2, 2, 310, 314, 7, 116, 2, 2, 311, 312, 7,
	126, 2, 2, 312, 314, 7, 126, 2, 2, 313, 309, 3, 2, 2, 2, 313, 311, 3, 2,
	2, 2, 314, 96, 3, 2, 2, 2, 315, 316, 7, 112, 2, 2, 316, 317, 7, 113, 2,
	2, 317, 320, 7, 118, 2, 2, 318, 320, 7, 35, 2, 2, 319, 315, 3, 2, 2, 2,
	319, 318, 3, 2, 2, 2, 320, 98, 3, 2, 2, 2, 321, 322, 7, 42, 2, 2, 322,
	100, 3, 2, 2, 2, 323, 324, 7, 43, 2, 2, 324, 102, 3, 2, 2, 2, 325, 326,
	7, 93, 2, 2, 326, 104, 3, 2, 2, 2, 327, 328, 7, 95, 2, 2, 328, 106, 3,
	2, 2, 2, 329, 330, 7, 112, 2, 2, 330, 331, 7, 107, 2, 2, 331, 332, 7, 110,
	2, 2, 332, 108, 3, 2, 2, 2, 333, 337, 7, 37, 2, 2, 334, 336, 10, 2, 2,
	2, 335, 334, 3, 2, 2, 2, 336, 339, 3, 2, 2, 2, 337, 335, 3, 2, 2, 2, 337,
	338, 3, 2, 2, 2, 338, 341, 3, 2, 2, 2, 339, 337, 3, 2, 2, 2, 340, 342,
	7, 15, 2, 2, 341, 340, 3, 2, 2, 2, 341, 342, 3, 2, 2, 2, 342, 343, 3, 2,
	2, 2, 343, 344, 7, 12, 2, 2, 344, 345, 3, 2, 2, 2, 345, 346, 8, 55, 2,
	2, 346, 110, 3, 2, 2, 2, 347, 348, 7, 63, 2, 2, 348, 349, 7, 100, 2, 2,
	349, 350, 7, 103, 2, 2, 350, 351, 7, 105, 2, 2, 351, 352, 7, 107, 2, 2,
	352, 353, 7, 112, 2, 2, 353, 357, 3, 2, 2, 2, 354, 356, 11, 2, 2, 2, 355,
	354, 3, 2, 2, 2, 356, 359, 3, 2, 2, 2, 357, 358, 3, 2, 2, 2, 357, 355,
	3, 2, 2, 2, 358, 360, 3, 2, 2, 2, 359, 357, 3, 2, 2, 2, 360, 361, 7, 63,
	2, 2, 361, 362, 7, 103, 2, 2, 362, 363, 7, 112, 2, 2, 363, 364, 7, 102,
	2, 2, 364, 366, 3, 2, 2, 2, 365, 367, 7, 15, 2, 2, 366, 365, 3, 2, 2, 2,
	366, 367, 3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368, 369, 7, 12, 2, 2, 369,
	370, 3, 2, 2, 2, 370, 371, 8, 56, 2, 2, 371, 112, 3, 2, 2, 2, 372, 374,
	9, 3, 2, 2, 373, 372, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375, 373, 3, 2,
	2, 2, 375, 376, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 378, 8, 57, 2, 2,
	378, 114, 3, 2, 2, 2, 379, 381, 9, 4, 2, 2, 380, 379, 3, 2, 2, 2, 381,
	382, 3, 2, 2, 2, 382, 380, 3, 2, 2, 2, 382, 383, 3, 2, 2, 2, 383, 116,
	3, 2, 2, 2, 384, 386, 9, 4, 2, 2, 385, 384, 3, 2, 2, 2, 386, 389, 3, 2,
	2, 2, 387, 385, 3, 2, 2, 2, 387, 388, 3, 2, 2, 2, 388, 390, 3, 2, 2, 2,
	389, 387, 3, 2, 2, 2, 390, 392, 7, 48, 2, 2, 391, 393, 9, 4, 2, 2, 392,
	391, 3, 2, 2, 2, 393, 394, 3, 2, 2, 2, 394, 392, 3, 2, 2, 2, 394, 395,
	3, 2, 2, 2, 395, 118, 3, 2, 2, 2, 396, 400, 9, 5, 2, 2, 397, 399, 9, 6,
	2, 2, 398, 397, 3, 2, 2, 2, 399, 402, 3, 2, 2, 2, 400, 398, 3, 2, 2, 2,
	400, 401, 3, 2, 2, 2, 401, 120, 3, 2, 2, 2, 402, 400, 3, 2, 2, 2, 403,
	404, 7, 38, 2, 2, 404, 405, 5, 119, 60, 2, 405, 122, 3, 2, 2, 2, 406, 407,
	5, 119, 60, 2, 407, 408, 9, 7, 2, 2, 408, 124, 3, 2, 2, 2, 21, 2, 131,
	133, 140, 142, 146, 153, 307, 313, 319, 337, 341, 357, 366, 375, 382, 387,
	394, 400, 3, 8, 2, 2,
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
	"", "", "','", "';'", "", "'require'", "'end'", "'def'", "'return'", "'pir'",
	"'if'", "'else'", "'elsif'", "'unless'", "'while'", "'retry'", "'break'",
	"'for'", "'true'", "'false'", "'+'", "'-'", "'*'", "'/'", "'%'", "'**'",
	"'=='", "'!='", "'>'", "'<'", "'<='", "'>='", "'='", "'+='", "'-='", "'*='",
	"'/='", "'%='", "'**='", "'&'", "'|'", "'^'", "'~'", "'<<'", "'>>'", "",
	"", "", "'('", "')'", "'['", "']'", "'nil'",
}

var lexerSymbolicNames = []string{
	"", "LITERAL", "COMMA", "SEMICOLON", "CRLF", "REQUIRE", "END", "DEF", "RETURN",
	"PIR", "IF", "ELSE", "ELSIF", "UNLESS", "WHILE", "RETRY", "BREAK", "FOR",
	"TRUE", "FALSE", "PLUS", "MINUS", "MUL", "DIV", "MOD", "EXP", "EQUAL",
	"NOT_EQUAL", "GREATER", "LESS", "LESS_EQUAL", "GREATER_EQUAL", "ASSIGN",
	"PLUS_ASSIGN", "MINUS_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN", "MOD_ASSIGN",
	"EXP_ASSIGN", "BIT_AND", "BIT_OR", "BIT_XOR", "BIT_NOT", "BIT_SHL", "BIT_SHR",
	"AND", "OR", "NOT", "LEFT_RBRACKET", "RIGHT_RBRACKET", "LEFT_SBRACKET",
	"RIGHT_SBRACKET", "NIL", "SL_COMMENT", "ML_COMMENT", "WS", "INT", "FLOAT",
	"ID", "ID_GLOBAL", "ID_FUNCTION",
}

var lexerRuleNames = []string{
	"ESCAPED_QUOTE", "LITERAL", "COMMA", "SEMICOLON", "CRLF", "REQUIRE", "END",
	"DEF", "RETURN", "PIR", "IF", "ELSE", "ELSIF", "UNLESS", "WHILE", "RETRY",
	"BREAK", "FOR", "TRUE", "FALSE", "PLUS", "MINUS", "MUL", "DIV", "MOD",
	"EXP", "EQUAL", "NOT_EQUAL", "GREATER", "LESS", "LESS_EQUAL", "GREATER_EQUAL",
	"ASSIGN", "PLUS_ASSIGN", "MINUS_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN", "MOD_ASSIGN",
	"EXP_ASSIGN", "BIT_AND", "BIT_OR", "BIT_XOR", "BIT_NOT", "BIT_SHL", "BIT_SHR",
	"AND", "OR", "NOT", "LEFT_RBRACKET", "RIGHT_RBRACKET", "LEFT_SBRACKET",
	"RIGHT_SBRACKET", "NIL", "SL_COMMENT", "ML_COMMENT", "WS", "INT", "FLOAT",
	"ID", "ID_GLOBAL", "ID_FUNCTION",
}

type CorundumLexer struct {
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

func NewCorundumLexer(input antlr.CharStream) *CorundumLexer {

	l := new(CorundumLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Corundum.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CorundumLexer tokens.
const (
	CorundumLexerLITERAL        = 1
	CorundumLexerCOMMA          = 2
	CorundumLexerSEMICOLON      = 3
	CorundumLexerCRLF           = 4
	CorundumLexerREQUIRE        = 5
	CorundumLexerEND            = 6
	CorundumLexerDEF            = 7
	CorundumLexerRETURN         = 8
	CorundumLexerPIR            = 9
	CorundumLexerIF             = 10
	CorundumLexerELSE           = 11
	CorundumLexerELSIF          = 12
	CorundumLexerUNLESS         = 13
	CorundumLexerWHILE          = 14
	CorundumLexerRETRY          = 15
	CorundumLexerBREAK          = 16
	CorundumLexerFOR            = 17
	CorundumLexerTRUE           = 18
	CorundumLexerFALSE          = 19
	CorundumLexerPLUS           = 20
	CorundumLexerMINUS          = 21
	CorundumLexerMUL            = 22
	CorundumLexerDIV            = 23
	CorundumLexerMOD            = 24
	CorundumLexerEXP            = 25
	CorundumLexerEQUAL          = 26
	CorundumLexerNOT_EQUAL      = 27
	CorundumLexerGREATER        = 28
	CorundumLexerLESS           = 29
	CorundumLexerLESS_EQUAL     = 30
	CorundumLexerGREATER_EQUAL  = 31
	CorundumLexerASSIGN         = 32
	CorundumLexerPLUS_ASSIGN    = 33
	CorundumLexerMINUS_ASSIGN   = 34
	CorundumLexerMUL_ASSIGN     = 35
	CorundumLexerDIV_ASSIGN     = 36
	CorundumLexerMOD_ASSIGN     = 37
	CorundumLexerEXP_ASSIGN     = 38
	CorundumLexerBIT_AND        = 39
	CorundumLexerBIT_OR         = 40
	CorundumLexerBIT_XOR        = 41
	CorundumLexerBIT_NOT        = 42
	CorundumLexerBIT_SHL        = 43
	CorundumLexerBIT_SHR        = 44
	CorundumLexerAND            = 45
	CorundumLexerOR             = 46
	CorundumLexerNOT            = 47
	CorundumLexerLEFT_RBRACKET  = 48
	CorundumLexerRIGHT_RBRACKET = 49
	CorundumLexerLEFT_SBRACKET  = 50
	CorundumLexerRIGHT_SBRACKET = 51
	CorundumLexerNIL            = 52
	CorundumLexerSL_COMMENT     = 53
	CorundumLexerML_COMMENT     = 54
	CorundumLexerWS             = 55
	CorundumLexerINT            = 56
	CorundumLexerFLOAT          = 57
	CorundumLexerID             = 58
	CorundumLexerID_GLOBAL      = 59
	CorundumLexerID_FUNCTION    = 60
)
