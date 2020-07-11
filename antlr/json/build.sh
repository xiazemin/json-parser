$java -jar ../calc/antlr-4.8-complete.jar -o parser JSON.g4
$java -jar ../calc/antlr-4.8-complete.jar -Dlanguage=Go -o parser JSON.g4
$go mod init github.com/xiazemin/json-parser/antlr/json
go: creating new go.mod: module github.com/xiazemin/json-parser/antlr/json
$go mod edit -replace github.com/xiazemin/json-parser/antlr/json/parser=./parser
go run example1.go
