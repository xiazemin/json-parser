// Code generated from MarkdownLexer.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 49, 260, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 3, 2, 3, 
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 
	3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 
	18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 
	3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 
	27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 
	3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 6, 37, 173, 
	10, 37, 13, 37, 14, 37, 174, 3, 37, 3, 37, 5, 37, 179, 10, 37, 3, 37, 3, 
	37, 3, 38, 6, 38, 184, 10, 38, 13, 38, 14, 38, 185, 3, 38, 3, 38, 5, 38, 
	190, 10, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 
	39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 
	3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 215, 10, 39, 3, 39, 3, 39, 3, 40, 3, 
	40, 3, 41, 3, 41, 3, 41, 3, 41, 7, 41, 225, 10, 41, 12, 41, 14, 41, 228, 
	11, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 7, 42, 235, 10, 42, 12, 42, 
	14, 42, 238, 11, 42, 3, 42, 3, 42, 3, 42, 3, 43, 5, 43, 244, 10, 43, 3, 
	43, 3, 43, 3, 43, 5, 43, 249, 10, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 
	3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 2, 2, 49, 3, 3, 5, 4, 7, 5, 9, 6, 11, 
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 
	31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 
	49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 
	67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 
	85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 3, 2, 15, 4, 2, 70, 70, 
	102, 102, 4, 2, 75, 75, 107, 107, 4, 2, 88, 88, 120, 120, 4, 2, 85, 85, 
	117, 117, 4, 2, 82, 82, 114, 114, 4, 2, 67, 67, 99, 99, 4, 2, 80, 80, 112, 
	112, 4, 2, 74, 74, 106, 106, 4, 2, 84, 84, 116, 116, 6, 2, 67, 72, 90, 
	90, 99, 104, 122, 122, 4, 2, 67, 92, 99, 124, 3, 2, 50, 59, 9, 2, 36, 36, 
	41, 41, 46, 46, 65, 65, 94, 94, 98, 98, 128, 128, 2, 274, 2, 3, 3, 2, 2, 
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 
	2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 
	3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 
	43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 
	2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 
	2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 
	2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 
	2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 
	3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 
	89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 
	3, 97, 3, 2, 2, 2, 5, 99, 3, 2, 2, 2, 7, 101, 3, 2, 2, 2, 9, 103, 3, 2, 
	2, 2, 11, 105, 3, 2, 2, 2, 13, 107, 3, 2, 2, 2, 15, 109, 3, 2, 2, 2, 17, 
	111, 3, 2, 2, 2, 19, 113, 3, 2, 2, 2, 21, 115, 3, 2, 2, 2, 23, 117, 3, 
	2, 2, 2, 25, 119, 3, 2, 2, 2, 27, 121, 3, 2, 2, 2, 29, 123, 3, 2, 2, 2, 
	31, 125, 3, 2, 2, 2, 33, 127, 3, 2, 2, 2, 35, 129, 3, 2, 2, 2, 37, 135, 
	3, 2, 2, 2, 39, 137, 3, 2, 2, 2, 41, 139, 3, 2, 2, 2, 43, 141, 3, 2, 2, 
	2, 45, 143, 3, 2, 2, 2, 47, 145, 3, 2, 2, 2, 49, 147, 3, 2, 2, 2, 51, 149, 
	3, 2, 2, 2, 53, 151, 3, 2, 2, 2, 55, 153, 3, 2, 2, 2, 57, 155, 3, 2, 2, 
	2, 59, 157, 3, 2, 2, 2, 61, 159, 3, 2, 2, 2, 63, 161, 3, 2, 2, 2, 65, 163, 
	3, 2, 2, 2, 67, 165, 3, 2, 2, 2, 69, 167, 3, 2, 2, 2, 71, 169, 3, 2, 2, 
	2, 73, 172, 3, 2, 2, 2, 75, 183, 3, 2, 2, 2, 77, 214, 3, 2, 2, 2, 79, 218, 
	3, 2, 2, 2, 81, 220, 3, 2, 2, 2, 83, 236, 3, 2, 2, 2, 85, 248, 3, 2, 2, 
	2, 87, 250, 3, 2, 2, 2, 89, 252, 3, 2, 2, 2, 91, 254, 3, 2, 2, 2, 93, 256, 
	3, 2, 2, 2, 95, 258, 3, 2, 2, 2, 97, 98, 7, 34, 2, 2, 98, 4, 3, 2, 2, 2, 
	99, 100, 7, 11, 2, 2, 100, 6, 3, 2, 2, 2, 101, 102, 7, 44, 2, 2, 102, 8, 
	3, 2, 2, 2, 103, 104, 7, 47, 2, 2, 104, 10, 3, 2, 2, 2, 105, 106, 7, 97, 
	2, 2, 106, 12, 3, 2, 2, 2, 107, 108, 7, 64, 2, 2, 108, 14, 3, 2, 2, 2, 
	109, 110, 7, 42, 2, 2, 110, 16, 3, 2, 2, 2, 111, 112, 7, 43, 2, 2, 112, 
	18, 3, 2, 2, 2, 113, 114, 7, 93, 2, 2, 114, 20, 3, 2, 2, 2, 115, 116, 7, 
	95, 2, 2, 116, 22, 3, 2, 2, 2, 117, 118, 7, 36, 2, 2, 118, 24, 3, 2, 2, 
	2, 119, 120, 7, 41, 2, 2, 120, 26, 3, 2, 2, 2, 121, 122, 7, 60, 2, 2, 122, 
	28, 3, 2, 2, 2, 123, 124, 7, 61, 2, 2, 124, 30, 3, 2, 2, 2, 125, 126, 7, 
	66, 2, 2, 126, 32, 3, 2, 2, 2, 127, 128, 7, 35, 2, 2, 128, 34, 3, 2, 2, 
	2, 129, 130, 5, 93, 47, 2, 130, 131, 5, 33, 17, 2, 131, 132, 5, 9, 5, 2, 
	132, 133, 5, 9, 5, 2, 133, 134, 6, 18, 2, 2, 134, 36, 3, 2, 2, 2, 135, 
	136, 7, 49, 2, 2, 136, 38, 3, 2, 2, 2, 137, 138, 7, 48, 2, 2, 138, 40, 
	3, 2, 2, 2, 139, 140, 7, 63, 2, 2, 140, 42, 3, 2, 2, 2, 141, 142, 7, 40, 
	2, 2, 142, 44, 3, 2, 2, 2, 143, 144, 7, 94, 2, 2, 144, 46, 3, 2, 2, 2, 
	145, 146, 7, 98, 2, 2, 146, 48, 3, 2, 2, 2, 147, 148, 7, 45, 2, 2, 148, 
	50, 3, 2, 2, 2, 149, 150, 7, 125, 2, 2, 150, 52, 3, 2, 2, 2, 151, 152, 
	7, 127, 2, 2, 152, 54, 3, 2, 2, 2, 153, 154, 9, 2, 2, 2, 154, 56, 3, 2, 
	2, 2, 155, 156, 9, 3, 2, 2, 156, 58, 3, 2, 2, 2, 157, 158, 9, 4, 2, 2, 
	158, 60, 3, 2, 2, 2, 159, 160, 9, 5, 2, 2, 160, 62, 3, 2, 2, 2, 161, 162, 
	9, 6, 2, 2, 162, 64, 3, 2, 2, 2, 163, 164, 9, 7, 2, 2, 164, 66, 3, 2, 2, 
	2, 165, 166, 9, 8, 2, 2, 166, 68, 3, 2, 2, 2, 167, 168, 9, 9, 2, 2, 168, 
	70, 3, 2, 2, 2, 169, 170, 9, 10, 2, 2, 170, 72, 3, 2, 2, 2, 171, 173, 5, 
	41, 21, 2, 172, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 172, 3, 2, 
	2, 2, 174, 175, 3, 2, 2, 2, 175, 178, 3, 2, 2, 2, 176, 179, 5, 85, 43, 
	2, 177, 179, 5, 81, 41, 2, 178, 176, 3, 2, 2, 2, 178, 177, 3, 2, 2, 2, 
	179, 180, 3, 2, 2, 2, 180, 181, 6, 37, 3, 2, 181, 74, 3, 2, 2, 2, 182, 
	184, 5, 9, 5, 2, 183, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 183, 
	3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 189, 3, 2, 2, 2, 187, 190, 5, 85, 
	43, 2, 188, 190, 5, 81, 41, 2, 189, 187, 3, 2, 2, 2, 189, 188, 3, 2, 2, 
	2, 190, 191, 3, 2, 2, 2, 191, 192, 6, 38, 4, 2, 192, 76, 3, 2, 2, 2, 193, 
	194, 7, 37, 2, 2, 194, 195, 7, 37, 2, 2, 195, 196, 7, 37, 2, 2, 196, 197, 
	7, 37, 2, 2, 197, 198, 7, 37, 2, 2, 198, 215, 7, 37, 2, 2, 199, 200, 7, 
	37, 2, 2, 200, 201, 7, 37, 2, 2, 201, 202, 7, 37, 2, 2, 202, 203, 7, 37, 
	2, 2, 203, 215, 7, 37, 2, 2, 204, 205, 7, 37, 2, 2, 205, 206, 7, 37, 2, 
	2, 206, 207, 7, 37, 2, 2, 207, 215, 7, 37, 2, 2, 208, 209, 7, 37, 2, 2, 
	209, 210, 7, 37, 2, 2, 210, 215, 7, 37, 2, 2, 211, 212, 7, 37, 2, 2, 212, 
	215, 7, 37, 2, 2, 213, 215, 7, 37, 2, 2, 214, 193, 3, 2, 2, 2, 214, 199, 
	3, 2, 2, 2, 214, 204, 3, 2, 2, 2, 214, 208, 3, 2, 2, 2, 214, 211, 3, 2, 
	2, 2, 214, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 217, 6, 39, 5, 2, 
	217, 78, 3, 2, 2, 2, 218, 219, 7, 37, 2, 2, 219, 80, 3, 2, 2, 2, 220, 221, 
	5, 3, 2, 2, 221, 226, 5, 3, 2, 2, 222, 225, 5, 3, 2, 2, 223, 225, 5, 5, 
	3, 2, 224, 222, 3, 2, 2, 2, 224, 223, 3, 2, 2, 2, 225, 228, 3, 2, 2, 2, 
	226, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 229, 3, 2, 2, 2, 228, 
	226, 3, 2, 2, 2, 229, 230, 5, 85, 43, 2, 230, 231, 6, 41, 6, 2, 231, 82, 
	3, 2, 2, 2, 232, 235, 5, 3, 2, 2, 233, 235, 5, 5, 3, 2, 234, 232, 3, 2, 
	2, 2, 234, 233, 3, 2, 2, 2, 235, 238, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 
	236, 237, 3, 2, 2, 2, 237, 239, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 239, 
	240, 5, 85, 43, 2, 240, 241, 6, 42, 7, 2, 241, 84, 3, 2, 2, 2, 242, 244, 
	7, 15, 2, 2, 243, 242, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 245, 3, 2, 
	2, 2, 245, 249, 7, 12, 2, 2, 246, 247, 7, 15, 2, 2, 247, 249, 8, 43, 2, 
	2, 248, 243, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 249, 86, 3, 2, 2, 2, 250, 
	251, 9, 11, 2, 2, 251, 88, 3, 2, 2, 2, 252, 253, 9, 12, 2, 2, 253, 90, 
	3, 2, 2, 2, 254, 255, 9, 13, 2, 2, 255, 92, 3, 2, 2, 2, 256, 257, 7, 62, 
	2, 2, 257, 94, 3, 2, 2, 2, 258, 259, 9, 14, 2, 2, 259, 96, 3, 2, 2, 2, 
	14, 2, 174, 178, 185, 189, 214, 224, 226, 234, 236, 243, 248, 3, 3, 43, 
	2,
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
	"", "' '", "'\t'", "'*'", "'-'", "'_'", "'>'", "'('", "')'", "'['", "']'", 
	"'\"'", "'''", "':'", "';'", "'@'", "'!'", "", "'/'", "'.'", "'='", "'&'", 
	"'\\'", "'`'", "'+'", "'{'", "'}'", "", "", "", "", "", "", "", "", "", 
	"", "", "", "'#'", "", "", "", "", "", "", "'<'",
}

var lexerSymbolicNames = []string{
	"", "SPACE", "TAB", "EMPH", "MINUS", "UNDERSCORE", "CLOSE_ANGLE_BRACKET", 
	"OPEN_PAREN", "CLOSE_PAREN", "OPEN_SB", "CLOSE_SB", "DOUBLE_QUOTE", "SINGLE_QUOTE", 
	"COLON", "SEMI_COLON", "AT", "EXCLAMATION_MARK", "HTML_COMMENT_OPEN", "SLASH", 
	"PERIOD", "EQUAL", "AMPERSAND", "BACKSLASH", "BACKTICK", "PLUS", "OPEN_CURLY", 
	"CLOSE_CURLY", "D", "I", "V", "S", "P", "A", "N", "H", "R", "SETEXT_BOTTOM_1", 
	"SETEXT_BOTTOM_2", "ATX_START", "SHARP", "LINE_BREAK", "BLANK_LINE", "NEWLINE", 
	"HEX_CHAR", "NORMAL_CHAR", "DIGIT", "OPEN_ANGLE_BRACKET", "SPECIAL_CHAR",
}

var lexerRuleNames = []string{
	"SPACE", "TAB", "EMPH", "MINUS", "UNDERSCORE", "CLOSE_ANGLE_BRACKET", "OPEN_PAREN", 
	"CLOSE_PAREN", "OPEN_SB", "CLOSE_SB", "DOUBLE_QUOTE", "SINGLE_QUOTE", "COLON", 
	"SEMI_COLON", "AT", "EXCLAMATION_MARK", "HTML_COMMENT_OPEN", "SLASH", "PERIOD", 
	"EQUAL", "AMPERSAND", "BACKSLASH", "BACKTICK", "PLUS", "OPEN_CURLY", "CLOSE_CURLY", 
	"D", "I", "V", "S", "P", "A", "N", "H", "R", "SETEXT_BOTTOM_1", "SETEXT_BOTTOM_2", 
	"ATX_START", "SHARP", "LINE_BREAK", "BLANK_LINE", "NEWLINE", "HEX_CHAR", 
	"NORMAL_CHAR", "DIGIT", "OPEN_ANGLE_BRACKET", "SPECIAL_CHAR",
}

type MarkdownLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewMarkdownLexer(input antlr.CharStream) *MarkdownLexer {

	l := new(MarkdownLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "MarkdownLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MarkdownLexer tokens.
const (
	MarkdownLexerSPACE = 1
	MarkdownLexerTAB = 2
	MarkdownLexerEMPH = 3
	MarkdownLexerMINUS = 4
	MarkdownLexerUNDERSCORE = 5
	MarkdownLexerCLOSE_ANGLE_BRACKET = 6
	MarkdownLexerOPEN_PAREN = 7
	MarkdownLexerCLOSE_PAREN = 8
	MarkdownLexerOPEN_SB = 9
	MarkdownLexerCLOSE_SB = 10
	MarkdownLexerDOUBLE_QUOTE = 11
	MarkdownLexerSINGLE_QUOTE = 12
	MarkdownLexerCOLON = 13
	MarkdownLexerSEMI_COLON = 14
	MarkdownLexerAT = 15
	MarkdownLexerEXCLAMATION_MARK = 16
	MarkdownLexerHTML_COMMENT_OPEN = 17
	MarkdownLexerSLASH = 18
	MarkdownLexerPERIOD = 19
	MarkdownLexerEQUAL = 20
	MarkdownLexerAMPERSAND = 21
	MarkdownLexerBACKSLASH = 22
	MarkdownLexerBACKTICK = 23
	MarkdownLexerPLUS = 24
	MarkdownLexerOPEN_CURLY = 25
	MarkdownLexerCLOSE_CURLY = 26
	MarkdownLexerD = 27
	MarkdownLexerI = 28
	MarkdownLexerV = 29
	MarkdownLexerS = 30
	MarkdownLexerP = 31
	MarkdownLexerA = 32
	MarkdownLexerN = 33
	MarkdownLexerH = 34
	MarkdownLexerR = 35
	MarkdownLexerSETEXT_BOTTOM_1 = 36
	MarkdownLexerSETEXT_BOTTOM_2 = 37
	MarkdownLexerATX_START = 38
	MarkdownLexerSHARP = 39
	MarkdownLexerLINE_BREAK = 40
	MarkdownLexerBLANK_LINE = 41
	MarkdownLexerNEWLINE = 42
	MarkdownLexerHEX_CHAR = 43
	MarkdownLexerNORMAL_CHAR = 44
	MarkdownLexerDIGIT = 45
	MarkdownLexerOPEN_ANGLE_BRACKET = 46
	MarkdownLexerSPECIAL_CHAR = 47
)


    int htmlElementCount;



func (l *MarkdownLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 41:
			l.NEWLINE_Action(localctx, actionIndex)


	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *MarkdownLexer) NEWLINE_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:
			setCharPositionInLine(0);


	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}


func (l *MarkdownLexer) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 16:
			return l.HTML_COMMENT_OPEN_Sempred(localctx, predIndex)


	case 35:
			return l.SETEXT_BOTTOM_1_Sempred(localctx, predIndex)


	case 36:
			return l.SETEXT_BOTTOM_2_Sempred(localctx, predIndex)


	case 37:
			return l.ATX_START_Sempred(localctx, predIndex)


	case 39:
			return l.LINE_BREAK_Sempred(localctx, predIndex)


	case 40:
			return l.BLANK_LINE_Sempred(localctx, predIndex)



	default:
		panic("No registered predicate for: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MarkdownLexer) HTML_COMMENT_OPEN_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return _tokenStartCharPositionInLine == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MarkdownLexer) SETEXT_BOTTOM_1_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
			return _tokenStartCharPositionInLine == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MarkdownLexer) SETEXT_BOTTOM_2_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
			return _tokenStartCharPositionInLine == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MarkdownLexer) ATX_START_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
			return _tokenStartCharPositionInLine == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MarkdownLexer) LINE_BREAK_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
			return _tokenStartCharPositionInLine > 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MarkdownLexer) BLANK_LINE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
			return _tokenStartCharPositionInLine == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

