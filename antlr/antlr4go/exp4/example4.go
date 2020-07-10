package main
import (
"bramp.net/antlr4/json" // The parser
"github.com/antlr/antlr4/runtime/Go/antlr"
)

type exampleListener struct {
// https://godoc.org/bramp.net/antlr4/json#BaseJSONListener
*json.BaseJSONListener
}

func main() {
// Setup the input
is := antlr.NewInputStream(`
		{
			"example": "json",
			"with": ["an", "array"]
		}`)


// Create the JSON Lexer
lexer := json.NewJSONLexer(is)
stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

// Create the JSON Parser
p := json.NewJSONParser(stream)

// Finally walk the tree
antlr.ParseTreeWalkerDefault.Walk(&exampleListener{}, p.Json())
}