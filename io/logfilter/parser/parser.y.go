// Code generated by goyacc -o parser.y.go -p cond parser.y. DO NOT EDIT.

//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2

import (
	"fmt"
	"math"
	"time"
)

//line parser.y:11
type condSymType struct {
	yys   int
	node  Node
	nodes []Node

	item Item

	strings   []string
	float     float64
	duration  time.Duration
	timestamp time.Time
}

const EQ = 57346
const COLON = 57347
const SEMICOLON = 57348
const COMMA = 57349
const COMMENT = 57350
const DURATION = 57351
const EOF = 57352
const ERROR = 57353
const ID = 57354
const LEFT_BRACE = 57355
const LEFT_BRACKET = 57356
const LEFT_PAREN = 57357
const NUMBER = 57358
const RIGHT_BRACE = 57359
const RIGHT_BRACKET = 57360
const RIGHT_PAREN = 57361
const SPACE = 57362
const STRING = 57363
const QUOTED_STRING = 57364
const NAMESPACE = 57365
const DOT = 57366
const operatorsStart = 57367
const ADD = 57368
const DIV = 57369
const GTE = 57370
const GT = 57371
const LT = 57372
const LTE = 57373
const MOD = 57374
const MUL = 57375
const NEQ = 57376
const POW = 57377
const SUB = 57378
const operatorsEnd = 57379
const keywordsStart = 57380
const AS = 57381
const ASC = 57382
const AUTO = 57383
const BY = 57384
const DESC = 57385
const TRUE = 57386
const FALSE = 57387
const FILTER = 57388
const IDENTIFIER = 57389
const IN = 57390
const AND = 57391
const LINK = 57392
const LIMIT = 57393
const SLIMIT = 57394
const OR = 57395
const NIL = 57396
const NULL = 57397
const OFFSET = 57398
const SOFFSET = 57399
const ORDER = 57400
const RE = 57401
const INT = 57402
const FLOAT = 57403
const POINT = 57404
const TIMEZONE = 57405
const WITH = 57406
const keywordsEnd = 57407
const startSymbolsStart = 57408
const START_STMTS = 57409
const START_BINARY_EXPRESSION = 57410
const START_FUNC_EXPRESSION = 57411
const startSymbolsEnd = 57412

var condToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"EQ",
	"COLON",
	"SEMICOLON",
	"COMMA",
	"COMMENT",
	"DURATION",
	"EOF",
	"ERROR",
	"ID",
	"LEFT_BRACE",
	"LEFT_BRACKET",
	"LEFT_PAREN",
	"NUMBER",
	"RIGHT_BRACE",
	"RIGHT_BRACKET",
	"RIGHT_PAREN",
	"SPACE",
	"STRING",
	"QUOTED_STRING",
	"NAMESPACE",
	"DOT",
	"operatorsStart",
	"ADD",
	"DIV",
	"GTE",
	"GT",
	"LT",
	"LTE",
	"MOD",
	"MUL",
	"NEQ",
	"POW",
	"SUB",
	"operatorsEnd",
	"keywordsStart",
	"AS",
	"ASC",
	"AUTO",
	"BY",
	"DESC",
	"TRUE",
	"FALSE",
	"FILTER",
	"IDENTIFIER",
	"IN",
	"AND",
	"LINK",
	"LIMIT",
	"SLIMIT",
	"OR",
	"NIL",
	"NULL",
	"OFFSET",
	"SOFFSET",
	"ORDER",
	"RE",
	"INT",
	"FLOAT",
	"POINT",
	"TIMEZONE",
	"WITH",
	"keywordsEnd",
	"startSymbolsStart",
	"START_STMTS",
	"START_BINARY_EXPRESSION",
	"START_FUNC_EXPRESSION",
	"startSymbolsEnd",
}
var condStatenames = [...]string{}

const condEofCode = 1
const condErrCode = 2
const condInitialStackSize = 16

//line parser.y:974

//line yacctab:1
var condExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 9,
	1, 2,
	10, 2,
	-2, 16,
	-1, 17,
	15, 144,
	-2, 18,
	-1, 18,
	15, 145,
	-2, 19,
	-1, 45,
	7, 88,
	17, 88,
	-2, 16,
	-1, 46,
	7, 89,
	17, 89,
	-2, 14,
	-1, 102,
	15, 144,
	-2, 18,
}

const condPrivate = 57344

const condLast = 320

var condAct = [...]int{

	17, 11, 107, 12, 27, 41, 101, 25, 30, 27,
	97, 61, 58, 32, 28, 117, 32, 68, 38, 28,
	14, 9, 18, 10, 44, 37, 45, 42, 39, 46,
	5, 65, 64, 63, 62, 64, 35, 36, 132, 29,
	131, 20, 115, 15, 29, 120, 33, 34, 40, 67,
	48, 24, 103, 104, 114, 55, 56, 121, 58, 32,
	94, 113, 112, 41, 41, 91, 92, 71, 133, 119,
	102, 74, 75, 76, 77, 78, 79, 80, 81, 82,
	83, 84, 85, 86, 87, 42, 42, 118, 121, 110,
	109, 45, 108, 100, 46, 2, 3, 4, 106, 124,
	70, 69, 110, 109, 116, 108, 89, 90, 93, 66,
	121, 111, 105, 88, 8, 6, 102, 1, 23, 127,
	128, 122, 110, 109, 111, 129, 123, 110, 109, 130,
	108, 16, 19, 27, 21, 126, 25, 30, 73, 100,
	22, 125, 32, 28, 111, 99, 13, 38, 72, 111,
	98, 96, 47, 48, 37, 7, 43, 39, 55, 56,
	27, 58, 59, 25, 30, 35, 36, 26, 29, 32,
	28, 31, 0, 0, 38, 33, 34, 0, 0, 0,
	24, 37, 0, 0, 39, 0, 60, 0, 0, 0,
	0, 0, 35, 36, 0, 29, 0, 0, 0, 0,
	0, 95, 33, 34, 0, 0, 0, 24, 47, 48,
	49, 50, 53, 54, 55, 56, 57, 58, 59, 27,
	0, 0, 0, 30, 0, 0, 0, 0, 32, 28,
	0, 51, 0, 38, 0, 52, 0, 0, 0, 0,
	37, 60, 0, 39, 0, 0, 0, 0, 0, 0,
	0, 35, 36, 60, 29, 0, 0, 0, 0, 0,
	0, 33, 34, 47, 48, 49, 50, 53, 54, 55,
	56, 57, 58, 59, 0, 47, 48, 49, 50, 53,
	54, 55, 56, 57, 58, 59, 51, 60, 0, 0,
	52, 0, 0, 0, 0, 0, 0, 0, 51, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 47,
	48, 49, 50, 53, 54, 55, 56, 57, 58, 59,
}
var condPact = [...]int{

	28, 105, 101, 148, -3, -1000, -1000, -1000, 148, -1000,
	237, -37, -1000, -1000, -1000, 10, 9, 8, 7, -1000,
	-1000, -1000, -1000, -1000, 94, 148, 86, -1000, -1000, 85,
	-1000, 51, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 8, 7, 131, -1000, -1000, -1000, 148, 148, 148,
	148, 148, 148, 148, 148, 148, 148, 148, 148, 148,
	148, 99, -3, -3, -3, -3, 38, 182, -1000, -8,
	-5, -1000, -1000, 148, 23, -23, 126, 126, 283, 249,
	126, 126, -23, -23, 126, -23, 23, 126, 207, -1000,
	-1000, -1000, -1000, 43, 42, -1000, 35, -1000, -1000, -1000,
	237, 207, 11, 72, 54, 26, -1000, 103, -1000, -1000,
	8, 7, -1000, -1000, -1000, -8, 81, 121, -3, -3,
	-1000, 207, -1000, -1000, -1000, 237, 207, 21, 19, -1000,
	50, -1000, -1000, -1000,
}
var condPgo = [...]int{

	0, 171, 171, 171, 167, 0, 167, 167, 167, 156,
	155, 155, 151, 151, 151, 151, 151, 151, 151, 151,
	151, 151, 151, 151, 151, 151, 151, 3, 2, 151,
	151, 22, 17, 23, 24, 151, 10, 43, 150, 20,
	146, 145, 145, 145, 145, 145, 145, 145, 1, 140,
	41, 134, 132, 131, 118, 118, 118, 118, 118, 117,
}
var condR1 = [...]int{

	0, 59, 59, 59, 59, 59, 13, 13, 13, 14,
	14, 14, 33, 33, 33, 33, 33, 33, 48, 48,
	31, 31, 1, 1, 50, 51, 51, 49, 49, 39,
	37, 53, 53, 12, 12, 12, 12, 28, 28, 28,
	27, 27, 27, 27, 27, 27, 54, 36, 36, 36,
	36, 38, 38, 17, 17, 18, 2, 2, 8, 8,
	15, 16, 19, 19, 6, 35, 35, 35, 35, 35,
	35, 35, 11, 11, 7, 7, 7, 7, 42, 42,
	42, 42, 10, 10, 9, 9, 9, 9, 34, 34,
	32, 32, 32, 32, 32, 32, 32, 32, 32, 32,
	32, 32, 32, 32, 32, 44, 44, 45, 45, 46,
	46, 46, 56, 56, 56, 47, 47, 47, 47, 47,
	47, 43, 43, 26, 26, 23, 23, 25, 25, 22,
	22, 24, 24, 20, 20, 21, 21, 30, 30, 30,
	29, 3, 3, 3, 4, 4, 52, 52, 58, 58,
	55, 55, 57, 57, 40, 40, 41, 41, 5, 5,
	5,
}
var condR2 = [...]int{

	0, 2, 2, 2, 2, 1, 1, 1, 1, 1,
	3, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 3, 1, 1, 1, 1, 1, 1, 1, 3,
	4, 3, 3, 3, 2, 1, 0, 3, 1, 0,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 3, 5, 5, 3, 5, 1, 1, 1, 3,
	9, 3, 1, 2, 2, 1, 1, 1, 3, 3,
	3, 2, 4, 0, 1, 3, 2, 0, 1, 1,
	3, 3, 3, 0, 1, 3, 2, 0, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 5, 6, 0, 1, 0, 2,
	1, 0, 2, 1, 0, 2, 2, 4, 5, 1,
	0, 1, 1, 4, 0, 2, 0, 2, 0, 2,
	0, 2, 0, 2, 0, 3, 0, 3, 1, 0,
	2, 1, 1, 0, 1, 1, 1, 2, 10, 5,
	1, 2, 1, 1, 4, 4, 4, 4, 1, 1,
	4,
}
var condChk = [...]int{

	-1000, -59, 67, 68, 69, 2, 10, -10, 13, -32,
	-33, -48, -27, -40, -39, -37, -53, -5, -31, -52,
	-50, -51, -49, -54, 59, 15, -4, 12, 22, 47,
	16, -1, 21, 54, 55, 44, 45, 33, 26, 36,
	-37, -5, -31, -9, -34, -32, -39, 26, 27, 28,
	29, 49, 53, 30, 31, 32, 33, 34, 35, 36,
	4, 48, 24, 24, 24, 24, 15, -33, -32, 15,
	15, 16, 17, 7, -33, -33, -33, -33, -33, -33,
	-33, -33, -33, -33, -33, -33, -33, -33, 14, -37,
	-37, -5, -5, -50, 22, 19, -12, -36, -38, -41,
	-33, 14, -5, 60, 61, -50, -34, -28, -27, -48,
	-5, -31, 19, 19, 19, 7, -28, 4, 15, 15,
	19, 7, 18, -36, 18, -33, 14, -5, -5, -27,
	-28, 19, 19, 18,
}
var condDef = [...]int{

	0, -2, 83, 0, 0, 5, 4, 1, 87, -2,
	0, 42, 12, 13, 14, 15, 17, -2, -2, 40,
	41, 43, 44, 45, 0, 0, 0, 158, 159, 0,
	146, 0, 24, 25, 26, 27, 28, 46, 22, 23,
	3, 144, 145, 0, 84, -2, -2, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 16, 36,
	0, 147, 82, 86, 90, 91, 92, 93, 94, 95,
	96, 97, 98, 99, 100, 101, 102, 103, 39, 31,
	32, 20, 21, 0, 0, 29, 0, 35, 47, 48,
	49, 39, -2, 0, 0, 0, 85, 0, 38, 42,
	18, 19, 154, 155, 30, 34, 0, 0, 0, 0,
	160, 0, 104, 33, 50, 51, 39, 0, 0, 37,
	0, 156, 157, 52,
}
var condTok1 = [...]int{

	1,
}
var condTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70,
}
var condTok3 = [...]int{
	0,
}

var condErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	condDebug        = 0
	condErrorVerbose = false
)

type condLexer interface {
	Lex(lval *condSymType) int
	Error(s string)
}

type condParser interface {
	Parse(condLexer) int
	Lookahead() int
}

type condParserImpl struct {
	lval  condSymType
	stack [condInitialStackSize]condSymType
	char  int
}

func (p *condParserImpl) Lookahead() int {
	return p.char
}

func condNewParser() condParser {
	return &condParserImpl{}
}

const condFlag = -1000

func condTokname(c int) string {
	if c >= 1 && c-1 < len(condToknames) {
		if condToknames[c-1] != "" {
			return condToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func condStatname(s int) string {
	if s >= 0 && s < len(condStatenames) {
		if condStatenames[s] != "" {
			return condStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func condErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !condErrorVerbose {
		return "syntax error"
	}

	for _, e := range condErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + condTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := condPact[state]
	for tok := TOKSTART; tok-1 < len(condToknames); tok++ {
		if n := base + tok; n >= 0 && n < condLast && condChk[condAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if condDef[state] == -2 {
		i := 0
		for condExca[i] != -1 || condExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; condExca[i] >= 0; i += 2 {
			tok := condExca[i]
			if tok < TOKSTART || condExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if condExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += condTokname(tok)
	}
	return res
}

func condlex1(lex condLexer, lval *condSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = condTok1[0]
		goto out
	}
	if char < len(condTok1) {
		token = condTok1[char]
		goto out
	}
	if char >= condPrivate {
		if char < condPrivate+len(condTok2) {
			token = condTok2[char-condPrivate]
			goto out
		}
	}
	for i := 0; i < len(condTok3); i += 2 {
		token = condTok3[i+0]
		if token == char {
			token = condTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = condTok2[1] /* unknown char */
	}
	if condDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", condTokname(token), uint(char))
	}
	return char, token
}

func condParse(condlex condLexer) int {
	return condNewParser().Parse(condlex)
}

func (condrcvr *condParserImpl) Parse(condlex condLexer) int {
	var condn int
	var condVAL condSymType
	var condDollar []condSymType
	_ = condDollar // silence set and not used
	condS := condrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	condstate := 0
	condrcvr.char = -1
	condtoken := -1 // condrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		condstate = -1
		condrcvr.char = -1
		condtoken = -1
	}()
	condp := -1
	goto condstack

ret0:
	return 0

ret1:
	return 1

condstack:
	/* put a state and value onto the stack */
	if condDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", condTokname(condtoken), condStatname(condstate))
	}

	condp++
	if condp >= len(condS) {
		nyys := make([]condSymType, len(condS)*2)
		copy(nyys, condS)
		condS = nyys
	}
	condS[condp] = condVAL
	condS[condp].yys = condstate

condnewstate:
	condn = condPact[condstate]
	if condn <= condFlag {
		goto conddefault /* simple state */
	}
	if condrcvr.char < 0 {
		condrcvr.char, condtoken = condlex1(condlex, &condrcvr.lval)
	}
	condn += condtoken
	if condn < 0 || condn >= condLast {
		goto conddefault
	}
	condn = condAct[condn]
	if condChk[condn] == condtoken { /* valid shift */
		condrcvr.char = -1
		condtoken = -1
		condVAL = condrcvr.lval
		condstate = condn
		if Errflag > 0 {
			Errflag--
		}
		goto condstack
	}

conddefault:
	/* default state action */
	condn = condDef[condstate]
	if condn == -2 {
		if condrcvr.char < 0 {
			condrcvr.char, condtoken = condlex1(condlex, &condrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if condExca[xi+0] == -1 && condExca[xi+1] == condstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			condn = condExca[xi+0]
			if condn < 0 || condn == condtoken {
				break
			}
		}
		condn = condExca[xi+1]
		if condn < 0 {
			goto ret0
		}
	}
	if condn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			condlex.Error(condErrorMessage(condstate, condtoken))
			Nerrs++
			if condDebug >= 1 {
				__yyfmt__.Printf("%s", condStatname(condstate))
				__yyfmt__.Printf(" saw %s\n", condTokname(condtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for condp >= 0 {
				condn = condPact[condS[condp].yys] + condErrCode
				if condn >= 0 && condn < condLast {
					condstate = condAct[condn] /* simulate a shift of "error" */
					if condChk[condstate] == condErrCode {
						goto condstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if condDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", condS[condp].yys)
				}
				condp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if condDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", condTokname(condtoken))
			}
			if condtoken == condEofCode {
				goto ret1
			}
			condrcvr.char = -1
			condtoken = -1
			goto condnewstate /* try again in the same state */
		}
	}

	/* reduction by production condn */
	if condDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", condn, condStatname(condstate))
	}

	condnt := condn
	condpt := condp
	_ = condpt // guard against "declared and not used"

	condp -= condR2[condn]
	// condp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if condp+1 >= len(condS) {
		nyys := make([]condSymType, len(condS)*2)
		copy(nyys, condS)
		condS = nyys
	}
	condVAL = condS[condp+1]

	/* consult goto table to find next state */
	condn = condR1[condn]
	condg := condPgo[condn]
	condj := condg + condS[condp].yys + 1

	if condj >= condLast {
		condstate = condAct[condg]
	} else {
		condstate = condAct[condj]
		if condChk[condstate] != -condn {
			condstate = condAct[condg]
		}
	}
	// dummy call; replaced with literal code
	switch condnt {

	case 1:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:136
		{
			yylex.(*parser).parseResult = condDollar[2].nodes
		}
	case 2:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:140
		{
			yylex.(*parser).parseResult = condDollar[2].node
		}
	case 3:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:144
		{
			yylex.(*parser).parseResult = condDollar[2].node
		}
	case 5:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:149
		{
			yylex.(*parser).unexpected("", "")
		}
	case 6:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:155
		{
			condVAL.node = condDollar[1].node
		}
	case 7:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:157
		{
			condVAL.node = condDollar[1].node
		}
	case 8:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:159
		{
			condVAL.node = condDollar[1].node
		}
	case 9:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:163
		{
			condVAL.node = Stmts{condDollar[1].node}
		}
	case 10:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:167
		{
			arr := condDollar[1].node.(Stmts)
			arr = append(arr, condDollar[3].node)
			condVAL.node = arr
		}
	case 11:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:173
		{
			condVAL.node = condDollar[1].node
		}
	case 18:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:183
		{
			condVAL.node = &Identifier{Name: condDollar[1].item.Val}
		}
	case 19:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:187
		{
			condVAL.node = condDollar[1].node
		}
	case 20:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:193
		{
			condVAL.node = &AttrExpr{
				Obj:  &Identifier{Name: condDollar[1].item.Val},
				Attr: &Identifier{Name: condDollar[3].item.Val},
			}
		}
	case 21:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:200
		{
			condVAL.node = &AttrExpr{
				Obj:  condDollar[1].node.(*AttrExpr),
				Attr: &Identifier{Name: condDollar[3].item.Val},
			}
		}
	case 24:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:213
		{
			condVAL.node = &StringLiteral{Val: yylex.(*parser).unquoteString(condDollar[1].item.Val)}
		}
	case 25:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:219
		{
			condVAL.node = &NilLiteral{}
		}
	case 26:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:223
		{
			condVAL.node = &NilLiteral{}
		}
	case 27:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:229
		{
			condVAL.node = &BoolLiteral{Val: true}
		}
	case 28:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:233
		{
			condVAL.node = &BoolLiteral{Val: false}
		}
	case 29:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:239
		{
			condVAL.node = &ParenExpr{Param: condDollar[2].node}
		}
	case 30:
		condDollar = condS[condpt-4 : condpt+1]
//line parser.y:245
		{
			condVAL.node = yylex.(*parser).newFunc(condDollar[1].item.Val, condDollar[3].nodes)
		}
	case 31:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:251
		{
			condVAL.node = &CascadeFunctions{Funcs: []*FuncExpr{condDollar[1].node.(*FuncExpr), condDollar[3].node.(*FuncExpr)}}
		}
	case 32:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:255
		{
			fc := condDollar[1].node.(*CascadeFunctions)
			fc.Funcs = append(fc.Funcs, condDollar[3].node.(*FuncExpr))
			condVAL.node = fc
		}
	case 33:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:263
		{
			condVAL.nodes = append(condVAL.nodes, condDollar[3].node)
		}
	case 35:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:268
		{
			condVAL.nodes = []Node{condDollar[1].node}
		}
	case 36:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:272
		{
			condVAL.nodes = nil
		}
	case 37:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:278
		{
			nl := condVAL.node.(NodeList)
			nl = append(nl, condDollar[3].node)
			condVAL.node = nl
		}
	case 38:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:284
		{
			condVAL.node = NodeList{condDollar[1].node}
		}
	case 39:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:288
		{
			condVAL.node = NodeList{}
		}
	case 46:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:302
		{
			condVAL.node = &Star{}
		}
	case 50:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:311
		{
			condVAL.node = getFuncArgList(condDollar[2].node.(NodeList))
		}
	case 51:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:317
		{
			condVAL.node = &FuncArg{ArgName: condDollar[1].item.Val, ArgVal: condDollar[3].node}
		}
	case 52:
		condDollar = condS[condpt-5 : condpt+1]
//line parser.y:321
		{
			condVAL.node = &FuncArg{
				ArgName: condDollar[1].item.Val,
				ArgVal:  getFuncArgList(condDollar[4].node.(NodeList)),
			}
		}
	case 53:
		condDollar = condS[condpt-5 : condpt+1]
//line parser.y:330
		{
			var cFuns *OuterFuncs
			chainFuncs, err := yylex.(*parser).newOuterFunc(cFuns, condDollar[1].node.(*FuncExpr))
			if err != nil {
				yylex.(*parser).addParseErr(nil, err)
			} else {
				switch chainFuncs.(type) {
				case *OuterFuncs:
					condVAL.node = chainFuncs.(*OuterFuncs)
				case *Show:
					show := chainFuncs.(*Show)
					if condDollar[2].nodes != nil {
						show.WhereCondition = condDollar[2].nodes
					}
					if condDollar[3].node != nil {
						show.TimeRange = condDollar[3].node.(*TimeRange)
					}
					if condDollar[4].node != nil {
						show.Limit = condDollar[4].node.(*Limit)
					}
					if condDollar[5].node != nil {
						show.Offset = condDollar[5].node.(*Offset)
					}
					condVAL.node = show
				case *DeleteFunc:
					condVAL.node = chainFuncs.(*DeleteFunc)
				default:
					yylex.(*parser).addParseErr(nil, fmt.Errorf("outer func error"))
				}
			}
		}
	case 54:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:362
		{
			cFuns := condDollar[1].node.(*OuterFuncs)
			chainFuncs, err := yylex.(*parser).newOuterFunc(cFuns, condDollar[3].node.(*FuncExpr))
			if err != nil {
				yylex.(*parser).addParseErr(nil, err)
			}
			condVAL.node = chainFuncs.(*OuterFuncs)
		}
	case 55:
		condDollar = condS[condpt-5 : condpt+1]
//line parser.y:374
		{
			m := yylex.(*parser).newLambda(condDollar[1].node.(*DFQuery), condDollar[2].item, condDollar[5].nodes)
			for _, n := range condDollar[3].nodes {
				m.Right = append(m.Right, n.(*DFQuery))
			}
			condVAL.node = m
		}
	case 56:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:384
		{
			condVAL.item = Item{Typ: FILTER}
		}
	case 57:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:388
		{
			condVAL.item = Item{Typ: LINK}
		}
	case 58:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:394
		{
			condVAL.nodes = []Node{condDollar[1].node}
		}
	case 59:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:398
		{
			condVAL.nodes = append(condDollar[1].nodes, condDollar[3].node)
		}
	case 60:
		condDollar = condS[condpt-9 : condpt+1]
//line parser.y:404
		{
			m := condDollar[1].node.(*DFQuery)

			if condDollar[2].node != nil {
				m.TimeRange = condDollar[2].node.(*TimeRange)
			}

			if condDollar[3].node != nil {
				m.GroupBy = condDollar[3].node.(*GroupBy)
			}

			if condDollar[4].node != nil {
				m.OrderBy = condDollar[4].node.(*OrderBy)
			}

			if condDollar[5].node != nil {
				m.Limit = condDollar[5].node.(*Limit)
			}

			if condDollar[6].node != nil {
				m.Offset = condDollar[6].node.(*Offset)
			}

			if condDollar[7].node != nil {
				m.SLimit = condDollar[7].node.(*SLimit)
			}

			if condDollar[8].node != nil {
				m.SOffset = condDollar[8].node.(*SOffset)
			}

			if condDollar[9].node != nil {
				m.TimeZone = condDollar[9].node.(*TimeZone)
			}

			condVAL.node = m
		}
	case 61:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:444
		{
			m := condDollar[1].node.(*DFQuery)
			m.Targets = yylex.(*parser).newTargets(condDollar[2].nodes)
			m.WhereCondition = condDollar[3].nodes
			condVAL.node = m
		}
	case 62:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:453
		{
			condVAL.node = condDollar[1].node
		}
	case 63:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:457
		{
			q := condDollar[2].node.(*DFQuery)
			q.Namespace = condDollar[1].item.Val
			condVAL.node = q
		}
	case 64:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:465
		{
			condVAL.item = condDollar[1].item
		}
	case 65:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:471
		{
			q, err := yylex.(*parser).newQuery(condDollar[1].node)
			if err != nil {
				log.Errorf("newQuery: %s", err)
			}
			condVAL.node = q
		}
	case 66:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:479
		{
			// FIXME: only func:: support attr_expr in from-clause
			x := condDollar[1].node.(*AttrExpr)
			q, err := yylex.(*parser).newQuery(&StringLiteral{Val: fmt.Sprintf("%s__%s", x.Obj, x.Attr)})
			if err != nil {
				log.Errorf("newQuery: %s", err)
			}
			condVAL.node = q
		}
	case 67:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:489
		{
			q, err := yylex.(*parser).newQuery(condDollar[1].item)
			if err != nil {
				log.Errorf("newQuery: %s", err)
			}
			condVAL.node = q
		}
	case 68:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:497
		{
			condVAL.node = yylex.(*parser).newSubquery(condDollar[2].node.(*DFQuery))
		}
	case 69:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:501
		{
			q := condDollar[1].node.(*DFQuery)
			if err := q.appendFrom(condDollar[3].node); err != nil {
				log.Debugf("appendFrom: %s", err.Error())
			}
			condVAL.node = q
		}
	case 70:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:509
		{
			q := condDollar[1].node.(*DFQuery)
			if err := q.appendFrom(condDollar[3].item.Val); err != nil {
				log.Debugf("appendFrom: %s", err.Error())
			}
			condVAL.node = q
		}
	case 71:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:517
		{
			condVAL.node = condDollar[1].node
		}
	case 72:
		condDollar = condS[condpt-4 : condpt+1]
//line parser.y:523
		{
			condVAL.nodes = condDollar[3].nodes
		}
	case 73:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:527
		{
			condVAL.nodes = nil
		}
	case 74:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:531
		{
			condVAL.nodes = []Node{condDollar[1].node}
		}
	case 75:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:535
		{
			condVAL.nodes = append(condDollar[1].nodes, condDollar[3].node)
		}
	case 77:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:540
		{
			condVAL.nodes = nil
		}
	case 78:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:544
		{
			nl, err := yylex.(*parser).newTarget(condDollar[1].node, "")
			if err != nil {
				yylex.(*parser).addParseErr(nil, err)
			}
			condVAL.node = nl
		}
	case 79:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:552
		{
			nl, err := yylex.(*parser).newTarget(condDollar[1].node, "")
			if err != nil {
				yylex.(*parser).addParseErr(nil, err)
			}
			condVAL.node = nl
		}
	case 80:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:560
		{
			nl, err := yylex.(*parser).newTarget(condDollar[1].node, condDollar[3].item.Val)
			if err != nil {
				yylex.(*parser).addParseErr(nil, err)
			}
			condVAL.node = nl
		}
	case 81:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:568
		{
			nl, err := yylex.(*parser).newTarget(condDollar[1].node, condDollar[3].item.Val)
			if err != nil {
				yylex.(*parser).addParseErr(nil, err)
			}
			condVAL.node = nl
		}
	case 82:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:578
		{
			condVAL.nodes = yylex.(*parser).newWhereConditions(condDollar[2].nodes)
		}
	case 83:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:582
		{
			condVAL.nodes = yylex.(*parser).newWhereConditions(nil)
		}
	case 84:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:589
		{
			condVAL.nodes = []Node{condDollar[1].node}
		}
	case 85:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:593
		{
			condVAL.nodes = append(condVAL.nodes, condDollar[3].node)
		}
	case 87:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:598
		{
			condVAL.nodes = nil
		}
	case 90:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:605
		{
			condVAL.node = yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
		}
	case 91:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:609
		{
			condVAL.node = yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
		}
	case 92:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:613
		{
			bexpr := yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
			bexpr.ReturnBool = true
			condVAL.node = bexpr
		}
	case 93:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:619
		{
			bexpr := yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
			bexpr.ReturnBool = true
			condVAL.node = bexpr
		}
	case 94:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:625
		{
			bexpr := yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
			bexpr.ReturnBool = true
			condVAL.node = bexpr
		}
	case 95:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:631
		{
			bexpr := yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
			bexpr.ReturnBool = true
			condVAL.node = bexpr
		}
	case 96:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:637
		{
			bexpr := yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
			bexpr.ReturnBool = true
			condVAL.node = bexpr
		}
	case 97:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:643
		{
			bexpr := yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
			bexpr.ReturnBool = true
			condVAL.node = bexpr
		}
	case 98:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:649
		{
			bexpr := yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
			condVAL.node = bexpr
		}
	case 99:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:654
		{
			bexpr := yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
			condVAL.node = bexpr
		}
	case 100:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:659
		{
			bexpr := yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
			bexpr.ReturnBool = true
			condVAL.node = bexpr
		}
	case 101:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:665
		{
			bexpr := yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
			condVAL.node = bexpr
		}
	case 102:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:670
		{
			bexpr := yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
			condVAL.node = bexpr
		}
	case 103:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:675
		{
			bexpr := yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[3].node, condDollar[2].item)
			bexpr.ReturnBool = true
			condVAL.node = bexpr
		}
	case 104:
		condDollar = condS[condpt-5 : condpt+1]
//line parser.y:681
		{
			bexpr := yylex.(*parser).newBinExpr(condDollar[1].node, condDollar[4].node, condDollar[2].item)
			bexpr.ReturnBool = true
			condVAL.node = bexpr
		}
	case 105:
		condDollar = condS[condpt-6 : condpt+1]
//line parser.y:689
		{
			condVAL.node = yylex.(*parser).newTimeRangeOpt(condDollar[2].node, condDollar[3].node, condDollar[4].node, condDollar[5].duration)
		}
	case 106:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:693
		{
			condVAL.node = yylex.(*parser).newTimeRangeOpt(nil, nil, nil, time.Duration(0))
		}
	case 107:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:699
		{
			condVAL.node = condDollar[1].node
		}
	case 108:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:701
		{
			condVAL.node = nil
		}
	case 109:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:705
		{
			condVAL.node = condDollar[2].node
		}
	case 110:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:707
		{
			condVAL.node = nil
		}
	case 111:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:709
		{
			condVAL.node = nil
		}
	case 112:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:713
		{
			condVAL.duration = condDollar[2].duration
		}
	case 113:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:715
		{
			condVAL.duration = time.Duration(0)
		}
	case 114:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:717
		{
			condVAL.duration = time.Duration(0)
		}
	case 115:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:721
		{
			condVAL.node = &TimeResolution{Duration: condDollar[2].duration}
		}
	case 116:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:725
		{
			condVAL.node = yylex.(*parser).newTimeResolution(nil, true) // Deprecated
		}
	case 117:
		condDollar = condS[condpt-4 : condpt+1]
//line parser.y:729
		{
			condVAL.node = yylex.(*parser).newTimeResolution(nil, false)
		}
	case 118:
		condDollar = condS[condpt-5 : condpt+1]
//line parser.y:733
		{
			condVAL.node = yylex.(*parser).newTimeResolution(condDollar[4].node.(*NumberLiteral), false)
		}
	case 119:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:737
		{
			condVAL.node = nil
		}
	case 120:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:739
		{
			condVAL.node = nil
		}
	case 121:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:743
		{
			condVAL.node = yylex.(*parser).newTimeExpr(&TimeExpr{IsDuration: true, Duration: condDollar[1].duration})
		}
	case 122:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:747
		{
			condVAL.node = yylex.(*parser).newTimeExpr(&TimeExpr{Time: condDollar[1].timestamp})
		}
	case 123:
		condDollar = condS[condpt-4 : condpt+1]
//line parser.y:753
		{
			condVAL.node = yylex.(*parser).newTimeZone(condDollar[3].node.(*StringLiteral))
		}
	case 124:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:757
		{
			condVAL.node = nil
		}
	case 125:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:761
		{
			condVAL.node = yylex.(*parser).newOffset(condDollar[2].node.(*NumberLiteral))
		}
	case 126:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:765
		{
			condVAL.node = yylex.(*parser).newOffset(nil)
		}
	case 127:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:771
		{
			condVAL.node = yylex.(*parser).newSOffset(condDollar[2].node.(*NumberLiteral))
		}
	case 128:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:775
		{
			condVAL.node = yylex.(*parser).newSOffset(nil)
		}
	case 129:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:781
		{
			condVAL.node = yylex.(*parser).newLimit(condDollar[2].node.(*NumberLiteral))
		}
	case 130:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:785
		{
			condVAL.node = yylex.(*parser).newLimit(nil)
		}
	case 131:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:791
		{
			condVAL.node = yylex.(*parser).newSLimit(condDollar[2].node.(*NumberLiteral))
		}
	case 132:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:795
		{
			condVAL.node = yylex.(*parser).newSLimit(nil)
		}
	case 133:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:802
		{
			nl := condDollar[2].node.(NodeList)
			if len(nl) == 0 {
				yylex.(*parser).addParseErrf(condDollar[1].item.PositionRange(), "group by list empty")
			}

			condVAL.node = &GroupBy{List: nl}
		}
	case 134:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:811
		{
			condVAL.node = nil
		}
	case 135:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:816
		{
			condVAL.node = yylex.(*parser).newOrderBy(condDollar[3].node.(NodeList))
		}
	case 136:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:820
		{
			condVAL.node = yylex.(*parser).newOrderBy(nil)
		}
	case 137:
		condDollar = condS[condpt-3 : condpt+1]
//line parser.y:826
		{
			nl := condDollar[1].node.(NodeList)
			nl = append(nl, condDollar[3].node)
			condVAL.node = nl
		}
	case 138:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:832
		{
			condVAL.node = NodeList{condDollar[1].node}
		}
	case 139:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:836
		{
			condVAL.node = NodeList{}
		}
	case 140:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:842
		{
			condVAL.node = yylex.(*parser).newOrderByElem(condDollar[1].item.Val, condDollar[2].item)
		}
	case 143:
		condDollar = condS[condpt-0 : condpt+1]
//line parser.y:849
		{
			condVAL.item = Item{Typ: ASC}
		}
	case 144:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:854
		{
			condVAL.item = condDollar[1].item
		}
	case 145:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:858
		{
			condVAL.item = Item{Val: condDollar[1].node.(*AttrExpr).String()}
		}
	case 146:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:865
		{
			condVAL.node = yylex.(*parser).number(condDollar[1].item.Val)
		}
	case 147:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:869
		{
			num := yylex.(*parser).number(condDollar[2].item.Val)
			switch condDollar[1].item.Typ {
			case ADD: // pass
			case SUB:
				if num.IsInt {
					num.Int = -num.Int
				} else {
					num.Float = -num.Float
				}
			}
			condVAL.node = num
		}
	case 148:
		condDollar = condS[condpt-10 : condpt+1]
//line parser.y:885
		{
			timestr := fmt.Sprintf("%s-%02s-%02s %02s:%02s:%02s", condDollar[1].item.Val, condDollar[3].item.Val, condDollar[5].item.Val, condDollar[6].item.Val, condDollar[8].item.Val, condDollar[10].item.Val)
			t, err := time.ParseInLocation("2006-01-02 15:04:05", timestr, time.UTC)
			if err != nil {
				yylex.(*parser).addParseErrf(condDollar[1].item.PositionRange(), "invalid date string: %s", timestr)
			}

			condVAL.timestamp = t
		}
	case 149:
		condDollar = condS[condpt-5 : condpt+1]
//line parser.y:895
		{
			timestr := fmt.Sprintf("%s-%02s-%02s", condDollar[1].item.Val, condDollar[3].item.Val, condDollar[5].item.Val)
			t, err := time.ParseInLocation("2006-01-02", timestr, time.UTC)
			if err != nil {
				yylex.(*parser).addParseErrf(condDollar[1].item.PositionRange(), "invalid date string: %s", timestr)
			}
			condVAL.timestamp = t
		}
	case 150:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:906
		{
			du, err := yylex.(*parser).parseDuration(condDollar[1].item.Val)
			if err != nil {
				yylex.(*parser).addParseErr(condDollar[1].item.PositionRange(), err)
			} else {
				condVAL.duration = du
			}
		}
	case 151:
		condDollar = condS[condpt-2 : condpt+1]
//line parser.y:915
		{
			du, err := yylex.(*parser).parseDuration(condDollar[2].item.Val)
			if err != nil {
				yylex.(*parser).addParseErr(condDollar[2].item.PositionRange(), err)
			} else {
				switch condDollar[1].item.Typ {
				case ADD:
					condVAL.duration = du
				case SUB:
					condVAL.duration = -du
				}
			}
		}
	case 153:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:932
		{
			nl := condDollar[1].node.(*NumberLiteral)
			var t time.Time
			if nl.IsInt {
				t = time.Unix(nl.Int, 0)
			} else {
				i, f := math.Modf(nl.Float)
				t = time.Unix(int64(i), int64(f*float64(time.Second)))
			}
			condVAL.timestamp = t
		}
	case 154:
		condDollar = condS[condpt-4 : condpt+1]
//line parser.y:946
		{
			condVAL.node = &Regex{Regex: condDollar[3].node.(*StringLiteral).Val}
		}
	case 155:
		condDollar = condS[condpt-4 : condpt+1]
//line parser.y:950
		{
			condVAL.node = &Regex{Regex: yylex.(*parser).unquoteString(condDollar[3].item.Val)}
		}
	case 156:
		condDollar = condS[condpt-4 : condpt+1]
//line parser.y:956
		{
			condVAL.node = &StaticCast{IsInt: true, Val: &Identifier{Name: condDollar[3].item.Val}}
		}
	case 157:
		condDollar = condS[condpt-4 : condpt+1]
//line parser.y:960
		{
			condVAL.node = &StaticCast{IsFloat: true, Val: &Identifier{Name: condDollar[3].item.Val}}
		}
	case 159:
		condDollar = condS[condpt-1 : condpt+1]
//line parser.y:967
		{
			condVAL.item.Val = yylex.(*parser).unquoteString(condDollar[1].item.Val)
		}
	case 160:
		condDollar = condS[condpt-4 : condpt+1]
//line parser.y:971
		{
			condVAL.item.Val = condDollar[3].node.(*StringLiteral).Val
		}
	}
	goto condstack /* stack new state and value */
}
