// Code generated from hello.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // hello

import "github.com/antlr/antlr4/runtime/Go/antlr"

// helloListener is a complete listener for a parse tree produced by helloParser.
type helloListener interface {
	antlr.ParseTreeListener

	// EnterHello is called when entering the hello production.
	EnterHello(c *HelloContext)

	// ExitHello is called when exiting the hello production.
	ExitHello(c *HelloContext)
}
