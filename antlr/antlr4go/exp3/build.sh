$go mod init github.com/xiazemin/json-parser/antlr/antlr4go
go: creating new go.mod: module github.com/xiazemin/json-parser/antlr/antlr4go
$go mod edit -replace github.com/xiazemin/json-parser/antlr/antlr4go/parser=../parser
$go run example3.go
go: finding github.com/antlr/antlr4 latest
8