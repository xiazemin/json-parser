package main


import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xiazemin/json-parser/antlr/jsonToGo/parser"
	"runtime"
	"strings"
)

type Listener struct {
	*parser.BaseJSONListener
	gocodeMap map[antlr.Tree]string `json:"gocode_map"`
	gocode string
	gotestMap map[antlr.Tree]string
	gotest string
}

func NewJsonToGoListener() *Listener {
	return &Listener{
		BaseJSONListener:&parser.BaseJSONListener{},
		gocodeMap:           make(map[antlr.Tree]string),
		gotestMap:           make(map[antlr.Tree]string),
	}
}
// EnterJson is called when entering the json production.
func (l*Listener)EnterJson(c *parser.JsonContext){
    l.gocodeMap[c]=`
package main

type JsonStruct struct {

`
fmt.Println(fmt.Sprintf("EnterJson %#v",*c))
}

// EnterObj is called when entering the obj production.
func (l*Listener)EnterObj(c *parser.ObjContext){
	fmt.Println(fmt.Sprintf("EnterObj %#v",*c))
}

// EnterPair is called when entering the pair production.
func (l*Listener)EnterPair(c *parser.PairContext){
	fmt.Println(fmt.Sprintf("EnterPair %#v",*c))
}

// EnterArr is called when entering the arr production.
func (l*Listener)EnterArr(c *parser.ArrContext){
	fmt.Println(fmt.Sprintf("EnterArr %#v",*c))
}

// EnterValue is called when entering the value production.
func (l*Listener)EnterValue(c *parser.ValueContext){
	fmt.Println(fmt.Sprintf("EnterValue %#v",*c))
}

// ExitJson is called when exiting the json production.
func (l*Listener)ExitJson(c *parser.JsonContext){
	fmt.Println(fmt.Sprintf("ExitJson %#v",l.gocodeMap[c.GetChild(0)]),c.GetAltNumber(),c.GetChild(0))
    if s,ok:=l.gocodeMap[c];ok{
    	l.gocode+=s
    	l.gocode+=l.gocodeMap[c.GetChild(0)]
    	l.gocode+=`
	}
    `
		return
	}
}

// ExitObj is called when exiting the obj production.
func (l*Listener)ExitObj(c *parser.ObjContext){
	sb := strings.Builder{}
	sb.WriteString("{\n")
	for _, p := range c.AllPair() {
		fmt.Println(fmt.Sprintf("ExitObj %#v",p))
		sb.WriteString(l.gocodeMap[p])
		sb.WriteString("\n")
	}
	sb.WriteString("}\n")
	l.gocodeMap[c]=sb.String()
}

// ExitPair is called when exiting the pair production.
func (l*Listener)ExitPair(c *parser.PairContext){
	fmt.Println(fmt.Sprintf("ExitPair %#v  %#v",c.STRING().GetText(),c.Value().GetText()))
	k := stripQuotes(c.STRING().GetText())
	//v := c.Value().GetText()
    //l.gocodeMap[c]=fmt.Sprintf("%s string `json:\"%s\" //%s `",toCamel(k),k,v)
//todo 根据value的类型定 key的类型，可能是简单类型也可能是复杂类型
	funcName, file, line, ok := runtime.Caller(0)
	fmt.Println("go->: k,v:", toCamel(k),l.gocodeMap[c.Value()],"******",k,funcName, file, line, ok)
	l.gocodeMap[c]=fmt.Sprintf("%s %s\n",toCamel(k),l.gocodeMap[c.Value()])
}

// ExitArr is called when exiting the arr production.
func (l*Listener)ExitArr(c *parser.ArrContext){
	fmt.Println(fmt.Sprintf("ExitArr %#v",c.Value(0)))
	sb := strings.Builder{}
	sb.WriteString("[] ArrayType struct \n")
	sb.WriteString(l.gocodeMap[c.Value(0)])
	l.gocodeMap[c]=sb.String()
}

// ExitValue is called when exiting the value production.
func (l*Listener)ExitValue(c *parser.ValueContext){
	fmt.Println(fmt.Sprintf("ExitValue %#v ",c.GetText()))
     // l.gocodeMap[c]=fmt.Sprintf("%s string `json:\"%s\" //%s ` ",toCamel(c.GetText()),c.GetText(),c.GetText())
	//l.gocodeMap[c]=c.GetText()
     //todo 类型推断
	fmt.Println("value1:", c.GetText())
	if c.Obj()!=nil {
		fmt.Println("value2:", c.Obj().GetText())
		l.gocodeMap[c]=fmt.Sprintf("struct //%s",c.Obj().GetText())
	}
	fmt.Println("value:3", c.GetChild(0))
	fmt.Println("value:4", c.GetAltNumber())
	fmt.Println("value:5",c.GetStop().GetText() )
	fmt.Println("value:6", c.GetStart().GetText())
	fmt.Println("value:8", c.BaseParserRuleContext.GetText())
	//fmt.Println(fmt.Sprintf("value:%#v", c.EnterRule(l)))
	fmt.Println("value:9", c.GetParser().GetTokenStream().GetAllText())
	if c.NUMBER()!=nil {
		fmt.Println("value:10", c.NUMBER().GetText())
		l.gocodeMap[c]=fmt.Sprintf(" float64 //%s",c.NUMBER().GetText())
	}
	if c.STRING()!=nil {
		fmt.Println("value:11", c.STRING().GetText())
		l.gocodeMap[c]=fmt.Sprintf("string //%s",c.STRING().GetText())
	}
	//fmt.Println("value:", c.IsValueContext())
	fmt.Println("value:12", c.GetParent())
	fmt.Println("value:13", c.GetSourceInterval().String())
	fmt.Println("value:14", c.IsEmpty())
	if c.Arr()!=nil {
		fmt.Println("value:7", c.Arr().GetText(),c.Arr().GetChildCount())
		if c.Arr().GetChildCount()==2 && c.GetText()=="[]" {
			l.gocodeMap[c] = fmt.Sprintf("[] interface{}{}//%s", c.Arr().GetText())
		}else{
			l.gocodeMap[c] = fmt.Sprintf("[] %s //%s",l.gocodeMap[c.GetChild(0)] ,c.Arr().GetText())
		}
	}

	if c.Obj()!=nil{
		if c.Obj().GetChildCount()==2 && c.GetText()=="{}"{
			//todo 空格问题 ｛  ｝
			l.gocodeMap[c] = fmt.Sprintf("interface{}{}//%s", c.Obj().GetText())
		}else {
			fmt.Println("go->:", " struct {", c.Obj().GetText())
		}
	}

    if c.NUMBER()!=nil{
		fmt.Println("go->:", " float64 ",c.NUMBER().GetText())
	}

	if c.STRING()!=nil{
		fmt.Println("go->:", " string",c.STRING().GetText())
	}
}


func toCamel(s string)string{
	var upperStr string
	vv := []rune(s)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return "S"+s
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
func stripQuotes(s string) string {
	if s == "" || ! strings.Contains(s, "\"") {
		return s
	}
	return s[1 : len(s)-1]
}

func (l*Listener)GetGo()string{
	/*
	return `
      type MyStruct struct{

     }
     `

	 */
	return l.gocode
}