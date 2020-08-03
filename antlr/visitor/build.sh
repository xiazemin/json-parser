java -jar ../calc/antlr-4.8-complete.jar -visitor -Dlanguage=Go -o parser JSON.g4

-no-listener

栈保存返回值

$tree
.
|____build.sh
|____JSON.g4
|____parser
| |____JSON.interp
| |____JSON.tokens
| |____json_base_listener.go
| |____json_base_visitor.go
| |____json_lexer.go
| |____json_listener.go
| |____json_parser.go
| |____json_visitor.go
| |____JSONLexer.interp
| |____JSONLexer.tokens

https://github.com/nikunjy/antlr-calc-go/blob/master/main.go

$go mod init github.com/xiazemin/json-parser/antlr/visitor
go: creating new go.mod: module github.com/xiazemin/json-parser/antlr/visitor

$go mod edit -replace github.com/xiazemin/json-parser/antlr/visitor/parser=./parser

export GOPROXY=https://goproxy.cn

go run ./

https://liangshuang.name/2017/08/20/antlr/

https://zhuanlan.zhihu.com/p/47179842