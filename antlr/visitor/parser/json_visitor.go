// Code generated from JSON.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // JSON

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by JSONParser.
type JSONVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JSONParser#json.
	VisitJson(ctx *JsonContext) interface{}

	// Visit a parse tree produced by JSONParser#obj.
	VisitObj(ctx *ObjContext) interface{}

	// Visit a parse tree produced by JSONParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by JSONParser#arr.
	VisitArr(ctx *ArrContext) interface{}

	// Visit a parse tree produced by JSONParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
