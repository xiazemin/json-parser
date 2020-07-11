$java -jar ../calc/antlr-4.8-complete.jar -Dlanguage=Go -o parser XMLLexer.g4 XMLParser.g4
$go mod init github.com/xiazemin/json-parser/antlr/xml
go: creating new go.mod: module github.com/xiazemin/json-parser/antlr/xml
$go mod edit -replace github.com/xiazemin/json-parser/antlr/xml/parser=./parser
$go run example1.go