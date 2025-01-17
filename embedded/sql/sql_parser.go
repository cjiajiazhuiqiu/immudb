// Code generated by goyacc -l -o sql_parser.go sql_grammar.y. DO NOT EDIT.
package sql

import __yyfmt__ "fmt"

import "fmt"

func setResult(l yyLexer, stmts []SQLStmt) {
	l.(*lexer).result = stmts
}

type yySymType struct {
	yys        int
	stmts      []SQLStmt
	stmt       SQLStmt
	colsSpec   []*ColSpec
	colSpec    *ColSpec
	cols       []*ColSelector
	rows       []*RowSpec
	row        *RowSpec
	values     []ValueExp
	value      ValueExp
	id         string
	number     uint64
	str        string
	boolean    bool
	blob       []byte
	sqlType    SQLValueType
	aggFn      AggregateFn
	ids        []string
	col        *ColSelector
	sel        Selector
	sels       []Selector
	distinct   bool
	ds         DataSource
	tableRef   *tableRef
	joins      []*JoinSpec
	join       *JoinSpec
	joinType   JoinType
	exp        ValueExp
	binExp     ValueExp
	err        error
	ordcols    []*OrdCol
	opt_ord    bool
	logicOp    LogicOperator
	cmpOp      CmpOperator
	pparam     int
	update     *colUpdate
	updates    []*colUpdate
	onConflict *OnConflictDo
}

const CREATE = 57346
const USE = 57347
const DATABASE = 57348
const SNAPSHOT = 57349
const SINCE = 57350
const UP = 57351
const TO = 57352
const TABLE = 57353
const UNIQUE = 57354
const INDEX = 57355
const ON = 57356
const ALTER = 57357
const ADD = 57358
const COLUMN = 57359
const PRIMARY = 57360
const KEY = 57361
const BEGIN = 57362
const TRANSACTION = 57363
const COMMIT = 57364
const ROLLBACK = 57365
const INSERT = 57366
const UPSERT = 57367
const INTO = 57368
const VALUES = 57369
const DELETE = 57370
const UPDATE = 57371
const SET = 57372
const CONFLICT = 57373
const DO = 57374
const NOTHING = 57375
const SELECT = 57376
const DISTINCT = 57377
const FROM = 57378
const BEFORE = 57379
const TX = 57380
const JOIN = 57381
const HAVING = 57382
const WHERE = 57383
const GROUP = 57384
const BY = 57385
const LIMIT = 57386
const ORDER = 57387
const ASC = 57388
const DESC = 57389
const AS = 57390
const NOT = 57391
const LIKE = 57392
const IF = 57393
const EXISTS = 57394
const IN = 57395
const IS = 57396
const AUTO_INCREMENT = 57397
const NULL = 57398
const NPARAM = 57399
const CAST = 57400
const PPARAM = 57401
const JOINTYPE = 57402
const LOP = 57403
const CMPOP = 57404
const IDENTIFIER = 57405
const TYPE = 57406
const NUMBER = 57407
const VARCHAR = 57408
const BOOLEAN = 57409
const BLOB = 57410
const AGGREGATE_FUNC = 57411
const ERROR = 57412
const STMT_SEPARATOR = 57413

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"CREATE",
	"USE",
	"DATABASE",
	"SNAPSHOT",
	"SINCE",
	"UP",
	"TO",
	"TABLE",
	"UNIQUE",
	"INDEX",
	"ON",
	"ALTER",
	"ADD",
	"COLUMN",
	"PRIMARY",
	"KEY",
	"BEGIN",
	"TRANSACTION",
	"COMMIT",
	"ROLLBACK",
	"INSERT",
	"UPSERT",
	"INTO",
	"VALUES",
	"DELETE",
	"UPDATE",
	"SET",
	"CONFLICT",
	"DO",
	"NOTHING",
	"SELECT",
	"DISTINCT",
	"FROM",
	"BEFORE",
	"TX",
	"JOIN",
	"HAVING",
	"WHERE",
	"GROUP",
	"BY",
	"LIMIT",
	"ORDER",
	"ASC",
	"DESC",
	"AS",
	"NOT",
	"LIKE",
	"IF",
	"EXISTS",
	"IN",
	"IS",
	"AUTO_INCREMENT",
	"NULL",
	"NPARAM",
	"CAST",
	"PPARAM",
	"JOINTYPE",
	"LOP",
	"CMPOP",
	"IDENTIFIER",
	"TYPE",
	"NUMBER",
	"VARCHAR",
	"BOOLEAN",
	"BLOB",
	"AGGREGATE_FUNC",
	"ERROR",
	"','",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'.'",
	"STMT_SEPARATOR",
	"'('",
	"')'",
	"'['",
	"']'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 94,
	50, 123,
	53, 123,
	-2, 112,
	-1, 154,
	39, 90,
	-2, 85,
	-1, 187,
	39, 90,
	-2, 87,
}

const yyPrivate = 57344

const yyLast = 333

var yyAct = [...]int{
	226, 269, 54, 132, 91, 200, 203, 114, 225, 6,
	75, 186, 67, 199, 123, 61, 70, 88, 17, 238,
	242, 130, 130, 130, 196, 251, 130, 246, 245, 243,
	220, 197, 244, 96, 131, 241, 98, 209, 191, 204,
	110, 108, 106, 109, 183, 32, 159, 107, 158, 102,
	103, 104, 105, 55, 205, 201, 96, 97, 129, 98,
	116, 208, 101, 110, 108, 106, 109, 164, 148, 93,
	107, 141, 102, 103, 104, 105, 55, 146, 139, 140,
	97, 120, 111, 125, 90, 101, 141, 80, 78, 135,
	136, 138, 137, 139, 140, 19, 181, 144, 145, 66,
	65, 128, 147, 141, 135, 136, 138, 137, 219, 79,
	160, 149, 79, 49, 153, 223, 151, 268, 263, 154,
	168, 210, 242, 138, 137, 99, 156, 141, 157, 152,
	222, 155, 161, 141, 139, 140, 170, 171, 172, 173,
	174, 175, 141, 163, 68, 135, 136, 138, 137, 182,
	140, 135, 136, 138, 137, 184, 180, 56, 56, 53,
	135, 136, 138, 137, 55, 56, 190, 130, 119, 51,
	74, 55, 222, 127, 112, 85, 194, 230, 77, 207,
	193, 202, 198, 162, 56, 89, 192, 166, 71, 150,
	124, 126, 121, 76, 118, 82, 72, 57, 211, 212,
	117, 32, 214, 44, 41, 36, 113, 189, 218, 177,
	237, 206, 236, 81, 38, 217, 176, 229, 228, 141,
	143, 233, 234, 227, 178, 124, 58, 179, 239, 270,
	271, 255, 133, 262, 249, 232, 68, 248, 250, 213,
	84, 63, 62, 253, 73, 30, 37, 34, 17, 256,
	260, 252, 258, 240, 48, 167, 165, 29, 261, 115,
	264, 10, 11, 28, 20, 266, 267, 215, 86, 64,
	39, 272, 12, 2, 273, 259, 31, 7, 169, 8,
	9, 13, 14, 83, 59, 15, 16, 60, 45, 46,
	47, 17, 21, 35, 134, 40, 27, 22, 24, 23,
	43, 25, 26, 92, 18, 221, 69, 142, 216, 235,
	254, 265, 195, 231, 95, 94, 247, 188, 187, 185,
	42, 33, 52, 50, 100, 224, 257, 87, 122, 5,
	4, 3, 1,
}

var yyPact = [...]int{
	257, -1000, -1000, 18, -1000, -1000, -1000, 243, -1000, -1000,
	286, 295, 285, 237, 231, 209, 138, 212, -1000, 257,
	-1000, 142, 163, 163, 282, 141, 292, 140, 138, 138,
	138, 224, 37, 95, -1000, -1000, -1000, 134, 177, 270,
	163, -1000, 205, 203, 253, 22, 21, 195, 125, 133,
	208, -1000, 99, 130, -1000, 10, 36, 9, 161, 132,
	269, -1000, 202, 110, 251, 122, 122, 298, 7, 103,
	-1000, 144, -1000, -18, 102, -1000, -1000, 131, 94, 129,
	127, -1000, 5, 128, 108, -1000, 127, -21, 96, -1000,
	-45, 188, 281, 32, 171, -1000, 7, 7, -1, -1000,
	-1000, 7, -1000, -1000, -1000, -1000, -10, 33, 126, -1000,
	-1000, 298, 125, 7, 298, 205, 214, 130, -1000, -31,
	-33, 34, 61, -1000, 119, 122, -11, -1000, -1000, 229,
	124, 228, -1000, 55, 264, 7, 7, 7, 7, 7,
	7, 160, 174, -1000, 88, 49, 214, 17, 7, -35,
	-1000, 188, -1000, 32, 147, 130, -41, -1000, -1000, -1000,
	123, 162, -56, -48, 122, -23, -1000, -23, -1000, -24,
	49, 49, 165, 165, 88, 79, -1000, 155, 7, -17,
	-42, -1000, 73, -1000, -1000, 195, -1000, 147, 200, -1000,
	-1000, 130, -1000, 248, -1000, 159, 43, -1000, -49, 101,
	-1000, 7, 59, -1000, -1000, 122, -1000, 88, -16, -1000,
	113, 193, -1000, -18, -1000, -24, 157, -1000, 154, -62,
	-1000, -1000, -23, 222, -44, 51, 32, -50, -47, -51,
	-52, 197, 191, 298, -54, -1000, -1000, -1000, -1000, -1000,
	219, -1000, 7, -1000, -1000, -1000, -1000, 186, 7, 121,
	261, -1000, 217, 32, 188, 190, 32, 47, -1000, 7,
	-1000, -1000, 121, 121, 32, 46, 183, -1000, 121, -1000,
	-1000, -1000, 183, -1000,
}

var yyPgo = [...]int{
	0, 332, 273, 331, 330, 9, 329, 328, 14, 17,
	6, 327, 326, 13, 5, 8, 325, 324, 125, 323,
	322, 2, 321, 7, 259, 320, 15, 319, 11, 318,
	317, 0, 12, 316, 315, 314, 313, 3, 312, 10,
	311, 310, 1, 4, 246, 309, 308, 307, 16, 306,
	305, 304,
}

var yyR1 = [...]int{
	0, 1, 2, 2, 51, 51, 3, 3, 3, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 25,
	25, 44, 44, 10, 10, 6, 6, 6, 6, 50,
	50, 49, 49, 48, 11, 11, 13, 13, 14, 9,
	9, 12, 12, 16, 16, 15, 15, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 7, 7, 8, 38,
	38, 45, 45, 46, 46, 46, 5, 22, 22, 19,
	19, 20, 20, 18, 18, 18, 21, 21, 21, 23,
	23, 24, 24, 26, 26, 27, 27, 28, 28, 29,
	30, 30, 32, 32, 36, 36, 33, 33, 37, 37,
	41, 41, 43, 43, 40, 40, 42, 42, 42, 39,
	39, 39, 31, 31, 31, 31, 31, 31, 31, 31,
	34, 34, 34, 47, 47, 35, 35, 35, 35, 35,
	35, 35, 35,
}

var yyR2 = [...]int{
	0, 1, 2, 3, 0, 1, 1, 1, 1, 2,
	1, 1, 3, 3, 4, 11, 8, 9, 6, 0,
	3, 0, 3, 1, 3, 9, 8, 6, 7, 0,
	4, 1, 3, 3, 0, 1, 1, 3, 3, 1,
	3, 1, 3, 0, 1, 1, 3, 1, 1, 1,
	1, 6, 3, 2, 1, 1, 1, 3, 5, 0,
	3, 0, 1, 0, 1, 2, 12, 0, 1, 1,
	1, 2, 4, 1, 4, 4, 1, 3, 5, 3,
	4, 1, 3, 0, 3, 0, 1, 1, 2, 6,
	0, 1, 0, 2, 0, 3, 0, 2, 0, 2,
	0, 3, 0, 4, 2, 4, 0, 1, 1, 0,
	1, 2, 1, 1, 2, 2, 4, 4, 6, 6,
	1, 1, 3, 0, 1, 3, 3, 3, 3, 3,
	3, 3, 4,
}

var yyChk = [...]int{
	-1000, -1, -2, -3, -4, -6, -5, 20, 22, 23,
	4, 5, 15, 24, 25, 28, 29, 34, -51, 77,
	21, 6, 11, 13, 12, 6, 7, 11, 26, 26,
	36, -24, 63, -22, 35, -2, 63, -44, 51, -44,
	13, 63, -25, 8, 63, -24, -24, -24, 30, 76,
	-19, 74, -20, -18, -21, 69, 63, 63, 49, 14,
	-44, -26, 37, 38, 16, 78, 78, -32, 41, -49,
	-48, 63, 63, 36, 71, -39, 63, 48, 78, 76,
	78, 52, 63, 14, 38, 65, 17, -11, -9, 63,
	-9, -43, 5, -31, -34, -35, 49, 73, 52, -18,
	-17, 78, 65, 66, 67, 68, 58, 63, 57, 59,
	56, -32, 71, 62, -23, -24, 78, -18, 63, 74,
	-21, 63, -7, -8, 63, 78, 63, 65, -8, 79,
	71, 79, -37, 44, 13, 72, 73, 75, 74, 61,
	62, 54, -47, 49, -31, -31, 78, -31, 78, 78,
	63, -43, -48, -31, -43, -26, -5, -39, 79, 79,
	76, 71, 64, -9, 78, 27, 63, 27, 65, 14,
	-31, -31, -31, -31, -31, -31, 56, 49, 50, 53,
	-5, 79, -31, 79, -37, -27, -28, -29, -30, 60,
	-39, 79, 63, 18, -8, -38, 80, 79, -9, -13,
	-14, 78, -13, -10, 63, 78, 56, -31, 78, 79,
	48, -32, -28, 39, -39, 19, -46, 56, 49, 65,
	79, -50, 71, 14, -16, -15, -31, -9, -5, -15,
	64, -36, 42, -23, -10, -45, 55, 56, 81, -14,
	31, 79, 71, 79, 79, 79, 79, -33, 40, 43,
	-43, 79, 32, -31, -41, 45, -31, -12, -21, 14,
	33, -37, 43, 71, -31, -40, -21, -21, 71, -42,
	46, 47, -21, -42,
}

var yyDef = [...]int{
	0, -2, 1, 4, 6, 7, 8, 0, 10, 11,
	0, 0, 0, 0, 0, 0, 0, 67, 2, 5,
	9, 0, 21, 21, 0, 0, 19, 0, 0, 0,
	0, 0, 81, 0, 68, 3, 12, 0, 0, 0,
	21, 13, 83, 0, 0, 0, 0, 92, 0, 0,
	0, 69, 70, 109, 73, 0, 76, 0, 0, 0,
	0, 14, 0, 0, 0, 34, 0, 102, 0, 92,
	31, 0, 82, 0, 0, 71, 110, 0, 0, 0,
	0, 22, 0, 0, 0, 20, 0, 0, 35, 39,
	0, 98, 0, 93, -2, 113, 0, 0, 0, 120,
	121, 0, 47, 48, 49, 50, 0, 76, 0, 54,
	55, 102, 0, 0, 102, 83, 0, 109, 111, 0,
	0, 77, 0, 56, 0, 0, 0, 84, 18, 0,
	0, 0, 27, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 124, 114, 115, 0, 0, 0, 0,
	53, 98, 32, 33, -2, 109, 0, 72, 74, 75,
	0, 0, 59, 0, 0, 0, 40, 0, 99, 0,
	125, 126, 127, 128, 129, 130, 131, 0, 0, 0,
	0, 122, 0, 52, 28, 92, 86, -2, 0, 91,
	79, 109, 78, 0, 57, 63, 0, 16, 0, 29,
	36, 43, 26, 103, 23, 0, 132, 116, 0, 117,
	0, 94, 88, 0, 80, 0, 61, 64, 0, 0,
	17, 25, 0, 0, 0, 44, 45, 0, 0, 0,
	0, 96, 0, 102, 0, 58, 62, 65, 60, 37,
	0, 38, 0, 24, 118, 119, 51, 100, 0, 0,
	0, 15, 0, 46, 98, 0, 97, 95, 41, 0,
	30, 66, 0, 0, 89, 101, 106, 42, 0, 104,
	107, 108, 106, 105,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	78, 79, 74, 72, 71, 73, 76, 75, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 80, 3, 81,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 77,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmts = yyDollar[1].stmts
			setResult(yylex, yyDollar[1].stmts)
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []SQLStmt{yyDollar[1].stmt}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmts = append([]SQLStmt{yyDollar[1].stmt}, yyDollar[3].stmts...)
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &BeginTransactionStmt{}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &CommitStmt{}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &RollbackStmt{}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &CreateDatabaseStmt{DB: yyDollar[3].id}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &UseDatabaseStmt{DB: yyDollar[3].id}
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &UseSnapshotStmt{sinceTx: yyDollar[3].number, asBefore: yyDollar[4].number}
		}
	case 15:
		yyDollar = yyS[yypt-11 : yypt+1]
		{
			yyVAL.stmt = &CreateTableStmt{ifNotExists: yyDollar[3].boolean, table: yyDollar[4].id, colsSpec: yyDollar[6].colsSpec, pkColNames: yyDollar[10].ids}
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &CreateIndexStmt{ifNotExists: yyDollar[3].boolean, table: yyDollar[5].id, cols: yyDollar[7].ids}
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &CreateIndexStmt{unique: true, ifNotExists: yyDollar[4].boolean, table: yyDollar[6].id, cols: yyDollar[8].ids}
		}
	case 18:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.stmt = &AddColumnStmt{table: yyDollar[3].id, colSpec: yyDollar[6].colSpec}
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[3].number
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = []string{yyDollar[1].id}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ids = yyDollar[2].ids
		}
	case 25:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &UpsertIntoStmt{isInsert: true, tableRef: yyDollar[3].tableRef, cols: yyDollar[5].ids, rows: yyDollar[8].rows, onConflict: yyDollar[9].onConflict}
		}
	case 26:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &UpsertIntoStmt{tableRef: yyDollar[3].tableRef, cols: yyDollar[5].ids, rows: yyDollar[8].rows}
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.stmt = &DeleteFromStmt{tableRef: yyDollar[3].tableRef, where: yyDollar[4].exp, indexOn: yyDollar[5].ids, limit: int(yyDollar[6].number)}
		}
	case 28:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.stmt = &UpdateStmt{tableRef: yyDollar[2].tableRef, updates: yyDollar[4].updates, where: yyDollar[5].exp, indexOn: yyDollar[6].ids, limit: int(yyDollar[7].number)}
		}
	case 29:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.onConflict = nil
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.onConflict = &OnConflictDo{}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.updates = []*colUpdate{yyDollar[1].update}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.updates = append(yyDollar[1].updates, yyDollar[3].update)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.update = &colUpdate{col: yyDollar[1].id, op: yyDollar[2].cmpOp, val: yyDollar[3].exp}
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ids = nil
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = yyDollar[1].ids
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.rows = []*RowSpec{yyDollar[1].row}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.rows = append(yyDollar[1].rows, yyDollar[3].row)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.row = &RowSpec{Values: yyDollar[2].values}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = []string{yyDollar[1].id}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ids = append(yyDollar[1].ids, yyDollar[3].id)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.cols = []*ColSelector{yyDollar[1].col}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = append(yyDollar[1].cols, yyDollar[3].col)
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.values = nil
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.values = yyDollar[1].values
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.values = []ValueExp{yyDollar[1].exp}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].exp)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Number{val: int64(yyDollar[1].number)}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Varchar{val: yyDollar[1].str}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Bool{val: yyDollar[1].boolean}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Blob{val: yyDollar[1].blob}
		}
	case 51:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.value = &Cast{val: yyDollar[3].exp, t: yyDollar[5].sqlType}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.value = &SysFn{fn: yyDollar[1].id}
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.value = &Param{id: yyDollar[2].id}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Param{id: fmt.Sprintf("param%d", yyDollar[1].pparam), pos: yyDollar[1].pparam}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &NullValue{t: AnyType}
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.colsSpec = []*ColSpec{yyDollar[1].colSpec}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.colsSpec = append(yyDollar[1].colsSpec, yyDollar[3].colSpec)
		}
	case 58:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.colSpec = &ColSpec{colName: yyDollar[1].id, colType: yyDollar[2].sqlType, maxLen: int(yyDollar[3].number), notNull: yyDollar[4].boolean, autoIncrement: yyDollar[5].boolean}
		}
	case 59:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[2].number
		}
	case 61:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 66:
		yyDollar = yyS[yypt-12 : yypt+1]
		{
			yyVAL.stmt = &SelectStmt{
				distinct:  yyDollar[2].distinct,
				selectors: yyDollar[3].sels,
				ds:        yyDollar[5].ds,
				indexOn:   yyDollar[6].ids,
				joins:     yyDollar[7].joins,
				where:     yyDollar[8].exp,
				groupBy:   yyDollar[9].cols,
				having:    yyDollar[10].exp,
				orderBy:   yyDollar[11].ordcols,
				limit:     int(yyDollar[12].number),
			}
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.distinct = false
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.distinct = true
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sels = nil
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sels = yyDollar[1].sels
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyDollar[1].sel.setAlias(yyDollar[2].id)
			yyVAL.sels = []Selector{yyDollar[1].sel}
		}
	case 72:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[3].sel.setAlias(yyDollar[4].id)
			yyVAL.sels = append(yyDollar[1].sels, yyDollar[3].sel)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sel = yyDollar[1].col
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, col: "*"}
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, db: yyDollar[3].col.db, table: yyDollar[3].col.table, col: yyDollar[3].col.col}
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.col = &ColSelector{col: yyDollar[1].id}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.col = &ColSelector{table: yyDollar[1].id, col: yyDollar[3].id}
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.col = &ColSelector{db: yyDollar[1].id, table: yyDollar[3].id, col: yyDollar[5].id}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].tableRef.asBefore = yyDollar[2].number
			yyDollar[1].tableRef.as = yyDollar[3].id
			yyVAL.ds = yyDollar[1].tableRef
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[2].stmt.(*SelectStmt).as = yyDollar[4].id
			yyVAL.ds = yyDollar[2].stmt.(DataSource)
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.tableRef = &tableRef{table: yyDollar[1].id}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.tableRef = &tableRef{db: yyDollar[1].id, table: yyDollar[3].id}
		}
	case 83:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[3].number
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.joins = nil
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = yyDollar[1].joins
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = []*JoinSpec{yyDollar[1].join}
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.joins = append([]*JoinSpec{yyDollar[1].join}, yyDollar[2].joins...)
		}
	case 89:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.join = &JoinSpec{joinType: yyDollar[1].joinType, ds: yyDollar[3].ds, indexOn: yyDollar[4].ids, cond: yyDollar[6].exp}
		}
	case 90:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.joinType = InnerJoin
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joinType = yyDollar[1].joinType
		}
	case 92:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exp = nil
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 94:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.cols = nil
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = yyDollar[3].cols
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exp = nil
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.number = yyDollar[2].number
		}
	case 100:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ordcols = nil
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ordcols = yyDollar[3].ordcols
		}
	case 102:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ids = nil
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.ids = yyDollar[4].ids
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.ordcols = []*OrdCol{{sel: yyDollar[1].col, descOrder: yyDollar[2].opt_ord}}
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.ordcols = append(yyDollar[1].ordcols, &OrdCol{sel: yyDollar[3].col, descOrder: yyDollar[4].opt_ord})
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.opt_ord = false
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = false
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = true
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.id = ""
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.id = yyDollar[1].id
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.id = yyDollar[2].id
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].exp
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].binExp
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = &NotBoolExp{exp: yyDollar[2].exp}
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = &NumExp{left: &Number{val: 0}, op: SUBSOP, right: yyDollar[2].exp}
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exp = &LikeBoolExp{val: yyDollar[1].exp, notLike: yyDollar[2].boolean, pattern: yyDollar[4].exp}
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exp = &ExistsBoolExp{q: (yyDollar[3].stmt).(*SelectStmt)}
		}
	case 118:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.exp = &InSubQueryExp{val: yyDollar[1].exp, notIn: yyDollar[2].boolean, q: yyDollar[5].stmt.(*SelectStmt)}
		}
	case 119:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.exp = &InListExp{val: yyDollar[1].exp, notIn: yyDollar[2].boolean, values: yyDollar[5].values}
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].sel
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].value
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 123:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: ADDOP, right: yyDollar[3].exp}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: SUBSOP, right: yyDollar[3].exp}
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: DIVOP, right: yyDollar[3].exp}
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: MULTOP, right: yyDollar[3].exp}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &BinBoolExp{left: yyDollar[1].exp, op: yyDollar[2].logicOp, right: yyDollar[3].exp}
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &CmpBoolExp{left: yyDollar[1].exp, op: yyDollar[2].cmpOp, right: yyDollar[3].exp}
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &CmpBoolExp{left: yyDollar[1].exp, op: EQ, right: &NullValue{t: AnyType}}
		}
	case 132:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.binExp = &CmpBoolExp{left: yyDollar[1].exp, op: NE, right: &NullValue{t: AnyType}}
		}
	}
	goto yystack /* stack new state and value */
}
