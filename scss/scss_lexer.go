// Generated from ScssLexer.g4 by ANTLR 4.7.

package scss

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 63, 559,
	8, 1, 8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6,
	9, 6, 4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4,
	12, 9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17,
	9, 17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9,
	22, 4, 23, 9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27,
	4, 28, 9, 28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4,
	33, 9, 33, 4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38,
	9, 38, 4, 39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9,
	43, 4, 44, 9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48,
	4, 49, 9, 49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4,
	54, 9, 54, 4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59,
	9, 59, 4, 60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9,
	64, 4, 65, 9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69,
	4, 70, 9, 70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 193, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5,
	5, 199, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3,
	24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28,
	3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35,
	3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3,
	41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3,
	44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3,
	47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49,
	3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3,
	50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 7, 51, 376, 10, 51,
	12, 51, 14, 51, 379, 11, 51, 3, 51, 3, 51, 3, 51, 7, 51, 384, 10, 51, 12,
	51, 14, 51, 387, 11, 51, 5, 51, 389, 10, 51, 3, 51, 3, 51, 3, 52, 3, 52,
	7, 52, 395, 10, 52, 12, 52, 14, 52, 398, 11, 52, 3, 52, 3, 52, 3, 52, 7,
	52, 403, 10, 52, 12, 52, 14, 52, 406, 11, 52, 3, 52, 5, 52, 409, 10, 52,
	3, 53, 3, 53, 3, 54, 3, 54, 7, 54, 415, 10, 54, 12, 54, 14, 54, 418, 11,
	54, 3, 54, 5, 54, 421, 10, 54, 3, 54, 6, 54, 424, 10, 54, 13, 54, 14, 54,
	425, 3, 54, 7, 54, 429, 10, 54, 12, 54, 14, 54, 432, 11, 54, 3, 54, 5,
	54, 435, 10, 54, 3, 54, 6, 54, 438, 10, 54, 13, 54, 14, 54, 439, 5, 54,
	442, 10, 54, 3, 55, 3, 55, 6, 55, 446, 10, 55, 13, 55, 14, 55, 447, 3,
	56, 3, 56, 3, 56, 6, 56, 453, 10, 56, 13, 56, 14, 56, 454, 3, 56, 3, 56,
	3, 57, 3, 57, 3, 57, 3, 57, 7, 57, 463, 10, 57, 12, 57, 14, 57, 466, 11,
	57, 3, 57, 3, 57, 3, 57, 5, 57, 471, 10, 57, 5, 57, 473, 10, 57, 3, 57,
	3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 7, 58, 481, 10, 58, 12, 58, 14, 58,
	484, 11, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3,
	59, 3, 60, 3, 60, 6, 60, 497, 10, 60, 13, 60, 14, 60, 498, 5, 60, 501,
	10, 60, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 62,
	3, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3,
	65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68,
	3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 70, 3,
	70, 3, 70, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3, 72,
	3, 72, 3, 73, 3, 73, 3, 73, 3, 73, 3, 73, 3, 482, 2, 74, 5, 3, 7, 4, 9,
	5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
	29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
	47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
	65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
	83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50,
	101, 51, 103, 52, 105, 2, 107, 53, 109, 54, 111, 55, 113, 56, 115, 57,
	117, 58, 119, 59, 121, 60, 123, 2, 125, 61, 127, 2, 129, 62, 131, 2, 133,
	63, 135, 2, 137, 2, 139, 2, 141, 2, 143, 2, 145, 2, 147, 2, 5, 2, 3, 4,
	10, 6, 2, 67, 92, 97, 97, 99, 124, 258, 0, 8, 2, 47, 47, 50, 59, 67, 92,
	97, 97, 99, 124, 258, 0, 5, 2, 12, 12, 15, 15, 36, 36, 5, 2, 12, 12, 15,
	15, 41, 41, 5, 2, 50, 59, 67, 72, 99, 104, 5, 2, 11, 12, 15, 15, 34, 34,
	4, 2, 12, 12, 15, 15, 6, 2, 12, 12, 15, 15, 43, 43, 61, 61, 2, 593, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2,
	2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2,
	2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3,
	2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51,
	3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2,
	59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2,
	2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2,
	2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2,
	2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3,
	2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97,
	3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2,
	2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3,
	2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 3, 119, 3, 2, 2, 2, 3,
	121, 3, 2, 2, 2, 4, 123, 3, 2, 2, 2, 4, 125, 3, 2, 2, 2, 4, 127, 3, 2,
	2, 2, 4, 129, 3, 2, 2, 2, 4, 131, 3, 2, 2, 2, 4, 133, 3, 2, 2, 2, 4, 135,
	3, 2, 2, 2, 4, 137, 3, 2, 2, 2, 4, 139, 3, 2, 2, 2, 4, 141, 3, 2, 2, 2,
	4, 143, 3, 2, 2, 2, 4, 145, 3, 2, 2, 2, 4, 147, 3, 2, 2, 2, 5, 149, 3,
	2, 2, 2, 7, 154, 3, 2, 2, 2, 9, 192, 3, 2, 2, 2, 11, 198, 3, 2, 2, 2, 13,
	200, 3, 2, 2, 2, 15, 204, 3, 2, 2, 2, 17, 209, 3, 2, 2, 2, 19, 211, 3,
	2, 2, 2, 21, 213, 3, 2, 2, 2, 23, 215, 3, 2, 2, 2, 25, 217, 3, 2, 2, 2,
	27, 219, 3, 2, 2, 2, 29, 221, 3, 2, 2, 2, 31, 223, 3, 2, 2, 2, 33, 225,
	3, 2, 2, 2, 35, 227, 3, 2, 2, 2, 37, 229, 3, 2, 2, 2, 39, 231, 3, 2, 2,
	2, 41, 233, 3, 2, 2, 2, 43, 235, 3, 2, 2, 2, 45, 237, 3, 2, 2, 2, 47, 239,
	3, 2, 2, 2, 49, 241, 3, 2, 2, 2, 51, 243, 3, 2, 2, 2, 53, 246, 3, 2, 2,
	2, 55, 248, 3, 2, 2, 2, 57, 250, 3, 2, 2, 2, 59, 252, 3, 2, 2, 2, 61, 254,
	3, 2, 2, 2, 63, 256, 3, 2, 2, 2, 65, 264, 3, 2, 2, 2, 67, 267, 3, 2, 2,
	2, 69, 270, 3, 2, 2, 2, 71, 272, 3, 2, 2, 2, 73, 275, 3, 2, 2, 2, 75, 278,
	3, 2, 2, 2, 77, 285, 3, 2, 2, 2, 79, 295, 3, 2, 2, 2, 81, 301, 3, 2, 2,
	2, 83, 304, 3, 2, 2, 2, 85, 308, 3, 2, 2, 2, 87, 313, 3, 2, 2, 2, 89, 320,
	3, 2, 2, 2, 91, 326, 3, 2, 2, 2, 93, 335, 3, 2, 2, 2, 95, 343, 3, 2, 2,
	2, 97, 351, 3, 2, 2, 2, 99, 356, 3, 2, 2, 2, 101, 364, 3, 2, 2, 2, 103,
	388, 3, 2, 2, 2, 105, 408, 3, 2, 2, 2, 107, 410, 3, 2, 2, 2, 109, 441,
	3, 2, 2, 2, 111, 443, 3, 2, 2, 2, 113, 452, 3, 2, 2, 2, 115, 458, 3, 2,
	2, 2, 117, 476, 3, 2, 2, 2, 119, 490, 3, 2, 2, 2, 121, 500, 3, 2, 2, 2,
	123, 502, 3, 2, 2, 2, 125, 507, 3, 2, 2, 2, 127, 512, 3, 2, 2, 2, 129,
	516, 3, 2, 2, 2, 131, 518, 3, 2, 2, 2, 133, 522, 3, 2, 2, 2, 135, 524,
	3, 2, 2, 2, 137, 529, 3, 2, 2, 2, 139, 534, 3, 2, 2, 2, 141, 539, 3, 2,
	2, 2, 143, 544, 3, 2, 2, 2, 145, 549, 3, 2, 2, 2, 147, 554, 3, 2, 2, 2,
	149, 150, 7, 112, 2, 2, 150, 151, 7, 119, 2, 2, 151, 152, 7, 110, 2, 2,
	152, 153, 7, 110, 2, 2, 153, 6, 3, 2, 2, 2, 154, 155, 7, 107, 2, 2, 155,
	156, 7, 112, 2, 2, 156, 8, 3, 2, 2, 2, 157, 193, 7, 39, 2, 2, 158, 159,
	7, 114, 2, 2, 159, 193, 7, 122, 2, 2, 160, 161, 7, 101, 2, 2, 161, 193,
	7, 111, 2, 2, 162, 163, 7, 111, 2, 2, 163, 193, 7, 111, 2, 2, 164, 165,
	7, 107, 2, 2, 165, 193, 7, 112, 2, 2, 166, 167, 7, 114, 2, 2, 167, 193,
	7, 118, 2, 2, 168, 169, 7, 114, 2, 2, 169, 193, 7, 101, 2, 2, 170, 171,
	7, 103, 2, 2, 171, 193, 7, 111, 2, 2, 172, 173, 7, 103, 2, 2, 173, 193,
	7, 122, 2, 2, 174, 175, 7, 102, 2, 2, 175, 176, 7, 103, 2, 2, 176, 193,
	7, 105, 2, 2, 177, 178, 7, 116, 2, 2, 178, 179, 7, 99, 2, 2, 179, 193,
	7, 102, 2, 2, 180, 181, 7, 105, 2, 2, 181, 182, 7, 116, 2, 2, 182, 183,
	7, 99, 2, 2, 183, 193, 7, 102, 2, 2, 184, 185, 7, 111, 2, 2, 185, 193,
	7, 117, 2, 2, 186, 193, 7, 117, 2, 2, 187, 188, 7, 106, 2, 2, 188, 193,
	7, 124, 2, 2, 189, 190, 7, 109, 2, 2, 190, 191, 7, 106, 2, 2, 191, 193,
	7, 124, 2, 2, 192, 157, 3, 2, 2, 2, 192, 158, 3, 2, 2, 2, 192, 160, 3,
	2, 2, 2, 192, 162, 3, 2, 2, 2, 192, 164, 3, 2, 2, 2, 192, 166, 3, 2, 2,
	2, 192, 168, 3, 2, 2, 2, 192, 170, 3, 2, 2, 2, 192, 172, 3, 2, 2, 2, 192,
	174, 3, 2, 2, 2, 192, 177, 3, 2, 2, 2, 192, 180, 3, 2, 2, 2, 192, 184,
	3, 2, 2, 2, 192, 186, 3, 2, 2, 2, 192, 187, 3, 2, 2, 2, 192, 189, 3, 2,
	2, 2, 193, 10, 3, 2, 2, 2, 194, 195, 7, 40, 2, 2, 195, 199, 7, 40, 2, 2,
	196, 197, 7, 126, 2, 2, 197, 199, 7, 126, 2, 2, 198, 194, 3, 2, 2, 2, 198,
	196, 3, 2, 2, 2, 199, 12, 3, 2, 2, 2, 200, 201, 7, 48, 2, 2, 201, 202,
	7, 48, 2, 2, 202, 203, 7, 48, 2, 2, 203, 14, 3, 2, 2, 2, 204, 205, 5, 49,
	24, 2, 205, 206, 5, 21, 10, 2, 206, 207, 3, 2, 2, 2, 207, 208, 8, 7, 2,
	2, 208, 16, 3, 2, 2, 2, 209, 210, 7, 42, 2, 2, 210, 18, 3, 2, 2, 2, 211,
	212, 7, 43, 2, 2, 212, 20, 3, 2, 2, 2, 213, 214, 7, 125, 2, 2, 214, 22,
	3, 2, 2, 2, 215, 216, 7, 127, 2, 2, 216, 24, 3, 2, 2, 2, 217, 218, 7, 93,
	2, 2, 218, 26, 3, 2, 2, 2, 219, 220, 7, 95, 2, 2, 220, 28, 3, 2, 2, 2,
	221, 222, 7, 64, 2, 2, 222, 30, 3, 2, 2, 2, 223, 224, 7, 128, 2, 2, 224,
	32, 3, 2, 2, 2, 225, 226, 7, 62, 2, 2, 226, 34, 3, 2, 2, 2, 227, 228, 7,
	60, 2, 2, 228, 36, 3, 2, 2, 2, 229, 230, 7, 61, 2, 2, 230, 38, 3, 2, 2,
	2, 231, 232, 7, 46, 2, 2, 232, 40, 3, 2, 2, 2, 233, 234, 7, 48, 2, 2, 234,
	42, 3, 2, 2, 2, 235, 236, 7, 38, 2, 2, 236, 44, 3, 2, 2, 2, 237, 238, 7,
	66, 2, 2, 238, 46, 3, 2, 2, 2, 239, 240, 7, 40, 2, 2, 240, 48, 3, 2, 2,
	2, 241, 242, 7, 37, 2, 2, 242, 50, 3, 2, 2, 2, 243, 244, 7, 60, 2, 2, 244,
	245, 7, 60, 2, 2, 245, 52, 3, 2, 2, 2, 246, 247, 7, 45, 2, 2, 247, 54,
	3, 2, 2, 2, 248, 249, 7, 44, 2, 2, 249, 56, 3, 2, 2, 2, 250, 251, 7, 49,
	2, 2, 251, 58, 3, 2, 2, 2, 252, 253, 7, 47, 2, 2, 253, 60, 3, 2, 2, 2,
	254, 255, 7, 39, 2, 2, 255, 62, 3, 2, 2, 2, 256, 257, 7, 119, 2, 2, 257,
	258, 7, 116, 2, 2, 258, 259, 7, 110, 2, 2, 259, 260, 3, 2, 2, 2, 260, 261,
	5, 17, 8, 2, 261, 262, 3, 2, 2, 2, 262, 263, 8, 31, 3, 2, 263, 64, 3, 2,
	2, 2, 264, 265, 7, 63, 2, 2, 265, 266, 7, 63, 2, 2, 266, 66, 3, 2, 2, 2,
	267, 268, 7, 35, 2, 2, 268, 269, 7, 63, 2, 2, 269, 68, 3, 2, 2, 2, 270,
	271, 7, 63, 2, 2, 271, 70, 3, 2, 2, 2, 272, 273, 7, 126, 2, 2, 273, 274,
	7, 63, 2, 2, 274, 72, 3, 2, 2, 2, 275, 276, 7, 128, 2, 2, 276, 277, 7,
	63, 2, 2, 277, 74, 3, 2, 2, 2, 278, 279, 7, 66, 2, 2, 279, 280, 7, 111,
	2, 2, 280, 281, 7, 107, 2, 2, 281, 282, 7, 122, 2, 2, 282, 283, 7, 107,
	2, 2, 283, 284, 7, 112, 2, 2, 284, 76, 3, 2, 2, 2, 285, 286, 7, 66, 2,
	2, 286, 287, 7, 104, 2, 2, 287, 288, 7, 119, 2, 2, 288, 289, 7, 112, 2,
	2, 289, 290, 7, 101, 2, 2, 290, 291, 7, 118, 2, 2, 291, 292, 7, 107, 2,
	2, 292, 293, 7, 113, 2, 2, 293, 294, 7, 112, 2, 2, 294, 78, 3, 2, 2, 2,
	295, 296, 7, 66, 2, 2, 296, 297, 7, 103, 2, 2, 297, 298, 7, 110, 2, 2,
	298, 299, 7, 117, 2, 2, 299, 300, 7, 103, 2, 2, 300, 80, 3, 2, 2, 2, 301,
	302, 7, 107, 2, 2, 302, 303, 7, 104, 2, 2, 303, 82, 3, 2, 2, 2, 304, 305,
	7, 66, 2, 2, 305, 306, 7, 107, 2, 2, 306, 307, 7, 104, 2, 2, 307, 84, 3,
	2, 2, 2, 308, 309, 7, 66, 2, 2, 309, 310, 7, 104, 2, 2, 310, 311, 7, 113,
	2, 2, 311, 312, 7, 116, 2, 2, 312, 86, 3, 2, 2, 2, 313, 314, 7, 66, 2,
	2, 314, 315, 7, 121, 2, 2, 315, 316, 7, 106, 2, 2, 316, 317, 7, 107, 2,
	2, 317, 318, 7, 110, 2, 2, 318, 319, 7, 103, 2, 2, 319, 88, 3, 2, 2, 2,
	320, 321, 7, 66, 2, 2, 321, 322, 7, 103, 2, 2, 322, 323, 7, 99, 2, 2, 323,
	324, 7, 101, 2, 2, 324, 325, 7, 106, 2, 2, 325, 90, 3, 2, 2, 2, 326, 327,
	7, 66, 2, 2, 327, 328, 7, 107, 2, 2, 328, 329, 7, 112, 2, 2, 329, 330,
	7, 101, 2, 2, 330, 331, 7, 110, 2, 2, 331, 332, 7, 119, 2, 2, 332, 333,
	7, 102, 2, 2, 333, 334, 7, 103, 2, 2, 334, 92, 3, 2, 2, 2, 335, 336, 7,
	66, 2, 2, 336, 337, 7, 107, 2, 2, 337, 338, 7, 111, 2, 2, 338, 339, 7,
	114, 2, 2, 339, 340, 7, 113, 2, 2, 340, 341, 7, 116, 2, 2, 341, 342, 7,
	118, 2, 2, 342, 94, 3, 2, 2, 2, 343, 344, 7, 66, 2, 2, 344, 345, 7, 116,
	2, 2, 345, 346, 7, 103, 2, 2, 346, 347, 7, 118, 2, 2, 347, 348, 7, 119,
	2, 2, 348, 349, 7, 116, 2, 2, 349, 350, 7, 112, 2, 2, 350, 96, 3, 2, 2,
	2, 351, 352, 7, 104, 2, 2, 352, 353, 7, 116, 2, 2, 353, 354, 7, 113, 2,
	2, 354, 355, 7, 111, 2, 2, 355, 98, 3, 2, 2, 2, 356, 357, 7, 118, 2, 2,
	357, 358, 7, 106, 2, 2, 358, 359, 7, 116, 2, 2, 359, 360, 7, 113, 2, 2,
	360, 361, 7, 119, 2, 2, 361, 362, 7, 105, 2, 2, 362, 363, 7, 106, 2, 2,
	363, 100, 3, 2, 2, 2, 364, 365, 7, 35, 2, 2, 365, 366, 7, 102, 2, 2, 366,
	367, 7, 103, 2, 2, 367, 368, 7, 104, 2, 2, 368, 369, 7, 99, 2, 2, 369,
	370, 7, 119, 2, 2, 370, 371, 7, 110, 2, 2, 371, 372, 7, 118, 2, 2, 372,
	102, 3, 2, 2, 2, 373, 377, 9, 2, 2, 2, 374, 376, 9, 3, 2, 2, 375, 374,
	3, 2, 2, 2, 376, 379, 3, 2, 2, 2, 377, 375, 3, 2, 2, 2, 377, 378, 3, 2,
	2, 2, 378, 389, 3, 2, 2, 2, 379, 377, 3, 2, 2, 2, 380, 381, 7, 47, 2, 2,
	381, 385, 9, 2, 2, 2, 382, 384, 9, 3, 2, 2, 383, 382, 3, 2, 2, 2, 384,
	387, 3, 2, 2, 2, 385, 383, 3, 2, 2, 2, 385, 386, 3, 2, 2, 2, 386, 389,
	3, 2, 2, 2, 387, 385, 3, 2, 2, 2, 388, 373, 3, 2, 2, 2, 388, 380, 3, 2,
	2, 2, 389, 390, 3, 2, 2, 2, 390, 391, 8, 51, 2, 2, 391, 104, 3, 2, 2, 2,
	392, 396, 7, 36, 2, 2, 393, 395, 10, 4, 2, 2, 394, 393, 3, 2, 2, 2, 395,
	398, 3, 2, 2, 2, 396, 394, 3, 2, 2, 2, 396, 397, 3, 2, 2, 2, 397, 399,
	3, 2, 2, 2, 398, 396, 3, 2, 2, 2, 399, 409, 7, 36, 2, 2, 400, 404, 7, 41,
	2, 2, 401, 403, 10, 5, 2, 2, 402, 401, 3, 2, 2, 2, 403, 406, 3, 2, 2, 2,
	404, 402, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 407, 3, 2, 2, 2, 406,
	404, 3, 2, 2, 2, 407, 409, 7, 41, 2, 2, 408, 392, 3, 2, 2, 2, 408, 400,
	3, 2, 2, 2, 409, 106, 3, 2, 2, 2, 410, 411, 5, 105, 52, 2, 411, 108, 3,
	2, 2, 2, 412, 420, 7, 47, 2, 2, 413, 415, 4, 50, 59, 2, 414, 413, 3, 2,
	2, 2, 415, 418, 3, 2, 2, 2, 416, 414, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2,
	417, 419, 3, 2, 2, 2, 418, 416, 3, 2, 2, 2, 419, 421, 7, 48, 2, 2, 420,
	416, 3, 2, 2, 2, 420, 421, 3, 2, 2, 2, 421, 423, 3, 2, 2, 2, 422, 424,
	4, 50, 59, 2, 423, 422, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 423, 3,
	2, 2, 2, 425, 426, 3, 2, 2, 2, 426, 442, 3, 2, 2, 2, 427, 429, 4, 50, 59,
	2, 428, 427, 3, 2, 2, 2, 429, 432, 3, 2, 2, 2, 430, 428, 3, 2, 2, 2, 430,
	431, 3, 2, 2, 2, 431, 433, 3, 2, 2, 2, 432, 430, 3, 2, 2, 2, 433, 435,
	7, 48, 2, 2, 434, 430, 3, 2, 2, 2, 434, 435, 3, 2, 2, 2, 435, 437, 3, 2,
	2, 2, 436, 438, 4, 50, 59, 2, 437, 436, 3, 2, 2, 2, 438, 439, 3, 2, 2,
	2, 439, 437, 3, 2, 2, 2, 439, 440, 3, 2, 2, 2, 440, 442, 3, 2, 2, 2, 441,
	412, 3, 2, 2, 2, 441, 434, 3, 2, 2, 2, 442, 110, 3, 2, 2, 2, 443, 445,
	7, 37, 2, 2, 444, 446, 9, 6, 2, 2, 445, 444, 3, 2, 2, 2, 446, 447, 3, 2,
	2, 2, 447, 445, 3, 2, 2, 2, 447, 448, 3, 2, 2, 2, 448, 112, 3, 2, 2, 2,
	449, 453, 9, 7, 2, 2, 450, 451, 7, 15, 2, 2, 451, 453, 7, 12, 2, 2, 452,
	449, 3, 2, 2, 2, 452, 450, 3, 2, 2, 2, 453, 454, 3, 2, 2, 2, 454, 452,
	3, 2, 2, 2, 454, 455, 3, 2, 2, 2, 455, 456, 3, 2, 2, 2, 456, 457, 8, 56,
	4, 2, 457, 114, 3, 2, 2, 2, 458, 459, 7, 49, 2, 2, 459, 460, 7, 49, 2,
	2, 460, 464, 3, 2, 2, 2, 461, 463, 10, 8, 2, 2, 462, 461, 3, 2, 2, 2, 463,
	466, 3, 2, 2, 2, 464, 462, 3, 2, 2, 2, 464, 465, 3, 2, 2, 2, 465, 472,
	3, 2, 2, 2, 466, 464, 3, 2, 2, 2, 467, 473, 7, 12, 2, 2, 468, 470, 7, 15,
	2, 2, 469, 471, 7, 12, 2, 2, 470, 469, 3, 2, 2, 2, 470, 471, 3, 2, 2, 2,
	471, 473, 3, 2, 2, 2, 472, 467, 3, 2, 2, 2, 472, 468, 3, 2, 2, 2, 473,
	474, 3, 2, 2, 2, 474, 475, 8, 57, 4, 2, 475, 116, 3, 2, 2, 2, 476, 477,
	7, 49, 2, 2, 477, 478, 7, 44, 2, 2, 478, 482, 3, 2, 2, 2, 479, 481, 11,
	2, 2, 2, 480, 479, 3, 2, 2, 2, 481, 484, 3, 2, 2, 2, 482, 483, 3, 2, 2,
	2, 482, 480, 3, 2, 2, 2, 483, 485, 3, 2, 2, 2, 484, 482, 3, 2, 2, 2, 485,
	486, 7, 44, 2, 2, 486, 487, 7, 49, 2, 2, 487, 488, 3, 2, 2, 2, 488, 489,
	8, 58, 4, 2, 489, 118, 3, 2, 2, 2, 490, 491, 5, 19, 9, 2, 491, 492, 3,
	2, 2, 2, 492, 493, 8, 59, 5, 2, 493, 120, 3, 2, 2, 2, 494, 501, 5, 105,
	52, 2, 495, 497, 10, 9, 2, 2, 496, 495, 3, 2, 2, 2, 497, 498, 3, 2, 2,
	2, 498, 496, 3, 2, 2, 2, 498, 499, 3, 2, 2, 2, 499, 501, 3, 2, 2, 2, 500,
	494, 3, 2, 2, 2, 500, 496, 3, 2, 2, 2, 501, 122, 3, 2, 2, 2, 502, 503,
	5, 21, 10, 2, 503, 504, 3, 2, 2, 2, 504, 505, 8, 61, 5, 2, 505, 506, 8,
	61, 6, 2, 506, 124, 3, 2, 2, 2, 507, 508, 5, 113, 56, 2, 508, 509, 3, 2,
	2, 2, 509, 510, 8, 62, 5, 2, 510, 511, 8, 62, 4, 2, 511, 126, 3, 2, 2,
	2, 512, 513, 5, 43, 21, 2, 513, 514, 3, 2, 2, 2, 514, 515, 8, 63, 7, 2,
	515, 128, 3, 2, 2, 2, 516, 517, 5, 15, 7, 2, 517, 130, 3, 2, 2, 2, 518,
	519, 5, 23, 11, 2, 519, 520, 3, 2, 2, 2, 520, 521, 8, 65, 8, 2, 521, 132,
	3, 2, 2, 2, 522, 523, 5, 103, 51, 2, 523, 134, 3, 2, 2, 2, 524, 525, 5,
	13, 6, 2, 525, 526, 3, 2, 2, 2, 526, 527, 8, 67, 5, 2, 527, 528, 8, 67,
	9, 2, 528, 136, 3, 2, 2, 2, 529, 530, 5, 41, 20, 2, 530, 531, 3, 2, 2,
	2, 531, 532, 8, 68, 5, 2, 532, 533, 8, 68, 10, 2, 533, 138, 3, 2, 2, 2,
	534, 535, 5, 17, 8, 2, 535, 536, 3, 2, 2, 2, 536, 537, 8, 69, 5, 2, 537,
	538, 8, 69, 11, 2, 538, 140, 3, 2, 2, 2, 539, 540, 5, 19, 9, 2, 540, 541,
	3, 2, 2, 2, 541, 542, 8, 70, 5, 2, 542, 543, 8, 70, 12, 2, 543, 142, 3,
	2, 2, 2, 544, 545, 5, 35, 17, 2, 545, 546, 3, 2, 2, 2, 546, 547, 8, 71,
	5, 2, 547, 548, 8, 71, 13, 2, 548, 144, 3, 2, 2, 2, 549, 550, 5, 39, 19,
	2, 550, 551, 3, 2, 2, 2, 551, 552, 8, 72, 5, 2, 552, 553, 8, 72, 14, 2,
	553, 146, 3, 2, 2, 2, 554, 555, 5, 37, 18, 2, 555, 556, 3, 2, 2, 2, 556,
	557, 8, 73, 5, 2, 557, 558, 8, 73, 15, 2, 558, 148, 3, 2, 2, 2, 29, 2,
	3, 4, 192, 198, 377, 385, 388, 396, 404, 408, 416, 420, 425, 430, 434,
	439, 441, 447, 452, 454, 464, 470, 472, 482, 498, 500, 16, 7, 4, 2, 7,
	3, 2, 8, 2, 2, 6, 2, 2, 9, 11, 2, 9, 22, 2, 9, 12, 2, 9, 7, 2, 9, 21, 2,
	9, 9, 2, 9, 10, 2, 9, 18, 2, 9, 20, 2, 9, 19, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "URL_STARTED", "IDENTIFY",
}

var lexerLiteralNames = []string{
	"", "'null'", "'in'", "", "", "'...'", "", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "'>'", "'~'", "'<'", "':'", "';'", "','", "'.'", "'$'", "'@'",
	"'&'", "'#'", "'::'", "'+'", "'*'", "'/'", "'-'", "'%'", "", "'=='", "'!='",
	"'='", "'|='", "'~='", "'@mixin'", "'@function'", "'@else'", "'if'", "'@if'",
	"'@for'", "'@while'", "'@each'", "'@include'", "'@import'", "'@return'",
	"'from'", "'through'", "'!default'",
}

var lexerSymbolicNames = []string{
	"", "NULL", "IN", "Unit", "COMBINE_COMPARE", "Ellipsis", "InterpolationStart",
	"LPAREN", "RPAREN", "BlockStart", "BlockEnd", "LBRACK", "RBRACK", "GT",
	"TIL", "LT", "COLON", "SEMI", "COMMA", "DOT", "DOLLAR", "AT", "AND", "HASH",
	"COLONCOLON", "PLUS", "TIMES", "DIV", "MINUS", "PERC", "UrlStart", "EQEQ",
	"NOTEQ", "EQ", "PIPE_EQ", "TILD_EQ", "MIXIN", "FUNCTION", "AT_ELSE", "IF",
	"AT_IF", "AT_FOR", "AT_WHILE", "AT_EACH", "INCLUDE", "IMPORT", "RETURN",
	"FROM", "THROUGH", "POUND_DEFAULT", "Identifier", "StringLiteral", "Number",
	"Color", "WS", "SL_COMMENT", "COMMENT", "UrlEnd", "Url", "SPACE", "InterpolationStartAfter",
	"IdentifierAfter",
}

var lexerRuleNames = []string{
	"NULL", "IN", "Unit", "COMBINE_COMPARE", "Ellipsis", "InterpolationStart",
	"LPAREN", "RPAREN", "BlockStart", "BlockEnd", "LBRACK", "RBRACK", "GT",
	"TIL", "LT", "COLON", "SEMI", "COMMA", "DOT", "DOLLAR", "AT", "AND", "HASH",
	"COLONCOLON", "PLUS", "TIMES", "DIV", "MINUS", "PERC", "UrlStart", "EQEQ",
	"NOTEQ", "EQ", "PIPE_EQ", "TILD_EQ", "MIXIN", "FUNCTION", "AT_ELSE", "IF",
	"AT_IF", "AT_FOR", "AT_WHILE", "AT_EACH", "INCLUDE", "IMPORT", "RETURN",
	"FROM", "THROUGH", "POUND_DEFAULT", "Identifier", "STRING", "StringLiteral",
	"Number", "Color", "WS", "SL_COMMENT", "COMMENT", "UrlEnd", "Url", "BlockStart_ID",
	"SPACE", "DOLLAR_ID", "InterpolationStartAfter", "InterpolationEnd_ID",
	"IdentifierAfter", "Ellipsis_ID", "DOT_ID", "LPAREN_ID", "RPAREN_ID", "COLON_ID",
	"COMMA_ID", "SEMI_ID",
}

type ScssLexer struct {
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

func NewScssLexer(input antlr.CharStream) *ScssLexer {

	l := new(ScssLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ScssLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ScssLexer tokens.
const (
	ScssLexerNULL                    = 1
	ScssLexerIN                      = 2
	ScssLexerUnit                    = 3
	ScssLexerCOMBINE_COMPARE         = 4
	ScssLexerEllipsis                = 5
	ScssLexerInterpolationStart      = 6
	ScssLexerLPAREN                  = 7
	ScssLexerRPAREN                  = 8
	ScssLexerBlockStart              = 9
	ScssLexerBlockEnd                = 10
	ScssLexerLBRACK                  = 11
	ScssLexerRBRACK                  = 12
	ScssLexerGT                      = 13
	ScssLexerTIL                     = 14
	ScssLexerLT                      = 15
	ScssLexerCOLON                   = 16
	ScssLexerSEMI                    = 17
	ScssLexerCOMMA                   = 18
	ScssLexerDOT                     = 19
	ScssLexerDOLLAR                  = 20
	ScssLexerAT                      = 21
	ScssLexerAND                     = 22
	ScssLexerHASH                    = 23
	ScssLexerCOLONCOLON              = 24
	ScssLexerPLUS                    = 25
	ScssLexerTIMES                   = 26
	ScssLexerDIV                     = 27
	ScssLexerMINUS                   = 28
	ScssLexerPERC                    = 29
	ScssLexerUrlStart                = 30
	ScssLexerEQEQ                    = 31
	ScssLexerNOTEQ                   = 32
	ScssLexerEQ                      = 33
	ScssLexerPIPE_EQ                 = 34
	ScssLexerTILD_EQ                 = 35
	ScssLexerMIXIN                   = 36
	ScssLexerFUNCTION                = 37
	ScssLexerAT_ELSE                 = 38
	ScssLexerIF                      = 39
	ScssLexerAT_IF                   = 40
	ScssLexerAT_FOR                  = 41
	ScssLexerAT_WHILE                = 42
	ScssLexerAT_EACH                 = 43
	ScssLexerINCLUDE                 = 44
	ScssLexerIMPORT                  = 45
	ScssLexerRETURN                  = 46
	ScssLexerFROM                    = 47
	ScssLexerTHROUGH                 = 48
	ScssLexerPOUND_DEFAULT           = 49
	ScssLexerIdentifier              = 50
	ScssLexerStringLiteral           = 51
	ScssLexerNumber                  = 52
	ScssLexerColor                   = 53
	ScssLexerWS                      = 54
	ScssLexerSL_COMMENT              = 55
	ScssLexerCOMMENT                 = 56
	ScssLexerUrlEnd                  = 57
	ScssLexerUrl                     = 58
	ScssLexerSPACE                   = 59
	ScssLexerInterpolationStartAfter = 60
	ScssLexerIdentifierAfter         = 61
)

// ScssLexer modes.
const (
	ScssLexerURL_STARTED = iota + 1
	ScssLexerIDENTIFY
)
