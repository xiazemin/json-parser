//line calc.y:8
package main

import __yyfmt__ "fmt"

//line calc.y:8
import (
	"fmt"
)

var idValueMap = map[string]int{}

//line calc.y:17
type calcSymType struct {
	yys   int
	value int
	id    string
}

const NUMBER = 57346
const ID = 57347
const ADD = 57348
const SUB = 57349
const MUL = 57350
const DIV = 57351
const ABS = 57352
const LPAREN = 57353
const RPAREN = 57354
const ASSIGN = 57355
const EOL = 57356

var calcToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"ID",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"ABS",
	"LPAREN",
	"RPAREN",
	"ASSIGN",
	"EOL",
}
var calcStatenames = [...]string{}

const calcEofCode = 1
const calcErrCode = 2
const calcInitialStackSize = 16

//line calc.y:63

//line yacctab:1
var calcExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const calcPrivate = 57344

const calcLast = 39

var calcAct = [...]int{

	5, 10, 11, 10, 11, 12, 10, 11, 15, 24,
	1, 9, 23, 0, 21, 22, 6, 16, 6, 3,
	2, 0, 7, 8, 7, 8, 4, 13, 14, 17,
	0, 0, 0, 20, 0, 0, 0, 18, 19,
}
var calcPact = [...]int{

	-1000, 14, -3, -8, 19, -1000, -1000, 12, 12, -1000,
	12, 12, 12, 12, 12, -1000, -1000, 0, 19, 19,
	-5, -1000, -1000, -1000, -1000,
}
var calcPgo = [...]int{

	0, 20, 26, 0, 10,
}
var calcR1 = [...]int{

	0, 4, 4, 4, 1, 1, 1, 2, 2, 2,
	3, 3, 3, 3,
}
var calcR2 = [...]int{

	0, 0, 3, 5, 1, 3, 3, 1, 3, 3,
	1, 1, 2, 3,
}
var calcChk = [...]int{

	-1000, -4, -1, 5, -2, -3, 4, 10, 11, 14,
	6, 7, 13, 8, 9, -3, 5, -1, -2, -2,
	-1, -3, -3, 12, 14,
}
var calcDef = [...]int{

	1, -2, 0, 11, 4, 7, 10, 0, 0, 2,
	0, 0, 0, 0, 0, 12, 11, 0, 5, 6,
	0, 8, 9, 13, 3,
}
var calcTok1 = [...]int{

	1,
}
var calcTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14,
}
var calcTok3 = [...]int{
	0,
}

var calcErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	calcDebug        = 0
	calcErrorVerbose = false
)

type calcLexer interface {
	Lex(lval *calcSymType) int
	Error(s string)
}

type calcParser interface {
	Parse(calcLexer) int
	Lookahead() int
}

type calcParserImpl struct {
	lval  calcSymType
	stack [calcInitialStackSize]calcSymType
	char  int
}

func (p *calcParserImpl) Lookahead() int {
	return p.char
}

func calcNewParser() calcParser {
	return &calcParserImpl{}
}

const calcFlag = -1000

func calcTokname(c int) string {
	if c >= 1 && c-1 < len(calcToknames) {
		if calcToknames[c-1] != "" {
			return calcToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func calcStatname(s int) string {
	if s >= 0 && s < len(calcStatenames) {
		if calcStatenames[s] != "" {
			return calcStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func calcErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !calcErrorVerbose {
		return "syntax error"
	}

	for _, e := range calcErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + calcTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := calcPact[state]
	for tok := TOKSTART; tok-1 < len(calcToknames); tok++ {
		if n := base + tok; n >= 0 && n < calcLast && calcChk[calcAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if calcDef[state] == -2 {
		i := 0
		for calcExca[i] != -1 || calcExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; calcExca[i] >= 0; i += 2 {
			tok := calcExca[i]
			if tok < TOKSTART || calcExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if calcExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += calcTokname(tok)
	}
	return res
}

func calclex1(lex calcLexer, lval *calcSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = calcTok1[0]
		goto out
	}
	if char < len(calcTok1) {
		token = calcTok1[char]
		goto out
	}
	if char >= calcPrivate {
		if char < calcPrivate+len(calcTok2) {
			token = calcTok2[char-calcPrivate]
			goto out
		}
	}
	for i := 0; i < len(calcTok3); i += 2 {
		token = calcTok3[i+0]
		if token == char {
			token = calcTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = calcTok2[1] /* unknown char */
	}
	if calcDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", calcTokname(token), uint(char))
	}
	return char, token
}

func calcParse(calclex calcLexer) int {
	return calcNewParser().Parse(calclex)
}

func (calcrcvr *calcParserImpl) Parse(calclex calcLexer) int {
	var calcn int
	var calcVAL calcSymType
	var calcDollar []calcSymType
	_ = calcDollar // silence set and not used
	calcS := calcrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	calcstate := 0
	calcrcvr.char = -1
	calctoken := -1 // calcrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		calcstate = -1
		calcrcvr.char = -1
		calctoken = -1
	}()
	calcp := -1
	goto calcstack

ret0:
	return 0

ret1:
	return 1

calcstack:
	/* put a state and value onto the stack */
	if calcDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", calcTokname(calctoken), calcStatname(calcstate))
	}

	calcp++
	if calcp >= len(calcS) {
		nyys := make([]calcSymType, len(calcS)*2)
		copy(nyys, calcS)
		calcS = nyys
	}
	calcS[calcp] = calcVAL
	calcS[calcp].yys = calcstate

calcnewstate:
	calcn = calcPact[calcstate]
	if calcn <= calcFlag {
		goto calcdefault /* simple state */
	}
	if calcrcvr.char < 0 {
		calcrcvr.char, calctoken = calclex1(calclex, &calcrcvr.lval)
	}
	calcn += calctoken
	if calcn < 0 || calcn >= calcLast {
		goto calcdefault
	}
	calcn = calcAct[calcn]
	if calcChk[calcn] == calctoken { /* valid shift */
		calcrcvr.char = -1
		calctoken = -1
		calcVAL = calcrcvr.lval
		calcstate = calcn
		if Errflag > 0 {
			Errflag--
		}
		goto calcstack
	}

calcdefault:
	/* default state action */
	calcn = calcDef[calcstate]
	if calcn == -2 {
		if calcrcvr.char < 0 {
			calcrcvr.char, calctoken = calclex1(calclex, &calcrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if calcExca[xi+0] == -1 && calcExca[xi+1] == calcstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			calcn = calcExca[xi+0]
			if calcn < 0 || calcn == calctoken {
				break
			}
		}
		calcn = calcExca[xi+1]
		if calcn < 0 {
			goto ret0
		}
	}
	if calcn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			calclex.Error(calcErrorMessage(calcstate, calctoken))
			Nerrs++
			if calcDebug >= 1 {
				__yyfmt__.Printf("%s", calcStatname(calcstate))
				__yyfmt__.Printf(" saw %s\n", calcTokname(calctoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for calcp >= 0 {
				calcn = calcPact[calcS[calcp].yys] + calcErrCode
				if calcn >= 0 && calcn < calcLast {
					calcstate = calcAct[calcn] /* simulate a shift of "error" */
					if calcChk[calcstate] == calcErrCode {
						goto calcstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if calcDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", calcS[calcp].yys)
				}
				calcp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if calcDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", calcTokname(calctoken))
			}
			if calctoken == calcEofCode {
				goto ret1
			}
			calcrcvr.char = -1
			calctoken = -1
			goto calcnewstate /* try again in the same state */
		}
	}

	/* reduction by production calcn */
	if calcDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", calcn, calcStatname(calcstate))
	}

	calcnt := calcn
	calcpt := calcp
	_ = calcpt // guard against "declared and not used"

	calcp -= calcR2[calcn]
	// calcp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if calcp+1 >= len(calcS) {
		nyys := make([]calcSymType, len(calcS)*2)
		copy(nyys, calcS)
		calcS = nyys
	}
	calcVAL = calcS[calcp+1]

	/* consult goto table to find next state */
	calcn = calcR1[calcn]
	calcg := calcPgo[calcn]
	calcj := calcg + calcS[calcp].yys + 1

	if calcj >= calcLast {
		calcstate = calcAct[calcg]
	} else {
		calcstate = calcAct[calcj]
		if calcChk[calcstate] != -calcn {
			calcstate = calcAct[calcg]
		}
	}
	// dummy call; replaced with literal code
	switch calcnt {

	case 2:
		calcDollar = calcS[calcpt-3 : calcpt+1]
		//line calc.y:33
		{
			idValueMap["_"] = calcDollar[2].value
			fmt.Printf("= %v\n", calcDollar[2].value)
		}
	case 3:
		calcDollar = calcS[calcpt-5 : calcpt+1]
		//line calc.y:37
		{
			idValueMap["_"] = calcDollar[4].value
			idValueMap[calcDollar[2].id] = calcDollar[4].value
			fmt.Printf("= %v\n", calcDollar[4].value)
		}
	case 4:
		calcDollar = calcS[calcpt-1 : calcpt+1]
		//line calc.y:45
		{
			calcVAL.value = calcDollar[1].value
		}
	case 5:
		calcDollar = calcS[calcpt-3 : calcpt+1]
		//line calc.y:46
		{
			calcVAL.value = calcDollar[1].value + calcDollar[3].value
		}
	case 6:
		calcDollar = calcS[calcpt-3 : calcpt+1]
		//line calc.y:47
		{
			calcVAL.value = calcDollar[1].value - calcDollar[3].value
		}
	case 7:
		calcDollar = calcS[calcpt-1 : calcpt+1]
		//line calc.y:51
		{
			calcVAL.value = calcDollar[1].value
		}
	case 8:
		calcDollar = calcS[calcpt-3 : calcpt+1]
		//line calc.y:52
		{
			calcVAL.value = calcDollar[1].value * calcDollar[3].value
		}
	case 9:
		calcDollar = calcS[calcpt-3 : calcpt+1]
		//line calc.y:53
		{
			calcVAL.value = calcDollar[1].value / calcDollar[3].value
		}
	case 10:
		calcDollar = calcS[calcpt-1 : calcpt+1]
		//line calc.y:57
		{
			calcVAL.value = calcDollar[1].value
		}
	case 11:
		calcDollar = calcS[calcpt-1 : calcpt+1]
		//line calc.y:58
		{
			calcVAL.value = idValueMap[calcDollar[1].id]
		}
	case 12:
		calcDollar = calcS[calcpt-2 : calcpt+1]
		//line calc.y:59
		{
			if calcDollar[2].value >= 0 {
				calcVAL.value = calcDollar[2].value
			} else {
				calcVAL.value = -calcDollar[2].value
			}
		}
	case 13:
		calcDollar = calcS[calcpt-3 : calcpt+1]
		//line calc.y:60
		{
			calcVAL.value = calcDollar[2].value
		}
	}
	goto calcstack /* stack new state and value */
}
