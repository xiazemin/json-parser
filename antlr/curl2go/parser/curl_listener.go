// Code generated from curl.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // curl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// curlListener is a complete listener for a parse tree produced by curlParser.
type curlListener interface {
	antlr.ParseTreeListener

	// EnterCurl is called when entering the curl production.
	EnterCurl(c *CurlContext)

	// EnterShortflag is called when entering the shortflag production.
	EnterShortflag(c *ShortflagContext)

	// EnterLongflag is called when entering the longflag production.
	EnterLongflag(c *LongflagContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// ExitCurl is called when exiting the curl production.
	ExitCurl(c *CurlContext)

	// ExitShortflag is called when exiting the shortflag production.
	ExitShortflag(c *ShortflagContext)

	// ExitLongflag is called when exiting the longflag production.
	ExitLongflag(c *LongflagContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)
}
