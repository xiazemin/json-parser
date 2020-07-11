$java -jar ../calc/antlr-4.8-complete.jar -Dlanguage=Go -o parser IDL.g4

#查看没有提交记录diff的方法，clone master
#copy 进但强项目，覆盖，看git diff

http://m.wfuyu.com/php/24877.html
$go mod init github.com/xiazemin/json-parser/antlr/idl
$go mod edit -replace github.com/xiazemin/json-parser/antlr/idl/parser=./parser
$go run main.go