// Code generated by goyacc - DO NOT EDIT.

package parser

import __yyfmt__ "fmt"

import (
	//"fmt"
	"github.com/tesujiro/goa/ast"
)

var defaultExpr = ast.FieldExpr{Expr: &ast.NumExpr{Literal: "0"}}
var defaultExprs = []ast.Expr{&defaultExpr}

//var IN_REGEXP bool

type yySymType struct {
	yys        int
	token      ast.Token
	rule       ast.Rule
	rules      []ast.Rule
	pattern    ast.Pattern
	stmt       ast.Stmt
	stmts      []ast.Stmt
	stmt_if    ast.Stmt
	expr       ast.Expr
	exprs      []ast.Expr
	ident_args []string
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault  = 57379
	yyEofCode  = 57344
	ANDAND     = 57356
	BEGIN      = 57367
	BREAK      = 57374
	CONTINUE   = 57375
	DELETE     = 57365
	DIVEQ      = 57364
	ELSE       = 57372
	END        = 57368
	EQEQ       = 57352
	FALSE      = 57350
	FOR        = 57373
	FUNC       = 57376
	GE         = 57354
	IDENT      = 57346
	IF         = 57371
	IN         = 57366
	LE         = 57355
	LEN        = 57358
	MINUSEQ    = 57362
	MINUSMINUS = 57360
	MULEQ      = 57363
	NEQ        = 57353
	NIL        = 57351
	NUMBER     = 57347
	OROR       = 57357
	PLUSEQ     = 57361
	PLUSPLUS   = 57359
	PRINT      = 57369
	REGEXP     = 57370
	RETURN     = 57377
	STRING     = 57348
	TRUE       = 57349
	UNARY      = 57378
	yyErrCode  = 57345

	yyMaxDepth = 200
	yyTabOfs   = -74
)

var (
	yyPrec = map[int]int{
		'=':        0,
		OROR:       1,
		ANDAND:     2,
		IDENT:      3,
		EQEQ:       4,
		NEQ:        4,
		'>':        5,
		'<':        5,
		GE:         5,
		LE:         5,
		'+':        6,
		'-':        6,
		PLUSPLUS:   6,
		MINUSMINUS: 6,
		'*':        7,
		'/':        7,
		'%':        7,
		UNARY:      8,
		'$':        9,
		'(':        10,
		')':        10,
	}

	yyXLAT = map[int]int{
		40:    0,  // '(' (129x)
		43:    1,  // '+' (127x)
		45:    2,  // '-' (127x)
		57346: 3,  // IDENT (127x)
		33:    4,  // '!' (121x)
		36:    5,  // '$' (121x)
		57350: 6,  // FALSE (121x)
		57376: 7,  // FUNC (121x)
		57351: 8,  // NIL (121x)
		57347: 9,  // NUMBER (121x)
		57348: 10, // STRING (121x)
		57349: 11, // TRUE (121x)
		57371: 12, // IF (80x)
		125:   13, // '}' (79x)
		57374: 14, // BREAK (79x)
		57375: 15, // CONTINUE (79x)
		57365: 16, // DELETE (79x)
		57373: 17, // FOR (79x)
		57369: 18, // PRINT (79x)
		57377: 19, // RETURN (79x)
		10:    20, // '\n' (64x)
		123:   21, // '{' (59x)
		59:    22, // ';' (55x)
		44:    23, // ',' (51x)
		41:    24, // ')' (48x)
		37:    25, // '%' (46x)
		42:    26, // '*' (46x)
		47:    27, // '/' (46x)
		60:    28, // '<' (46x)
		62:    29, // '>' (46x)
		57356: 30, // ANDAND (46x)
		57364: 31, // DIVEQ (46x)
		57352: 32, // EQEQ (46x)
		57354: 33, // GE (46x)
		57355: 34, // LE (46x)
		57362: 35, // MINUSEQ (46x)
		57360: 36, // MINUSMINUS (46x)
		57363: 37, // MULEQ (46x)
		57353: 38, // NEQ (46x)
		57357: 39, // OROR (46x)
		57361: 40, // PLUSEQ (46x)
		57359: 41, // PLUSPLUS (46x)
		57381: 42, // expr (45x)
		61:    43, // '=' (38x)
		93:    44, // ']' (38x)
		57382: 45, // exprs (15x)
		57344: 46, // $end (12x)
		57367: 47, // BEGIN (12x)
		57368: 48, // END (12x)
		57391: 49, // stmt (9x)
		57392: 50, // stmt_if (9x)
		57393: 51, // stmts (9x)
		57372: 52, // ELSE (4x)
		57384: 53, // nls (4x)
		57385: 54, // opt_nls (4x)
		91:    55, // '[' (2x)
		57383: 56, // ident_args (2x)
		57386: 57, // opt_semi (2x)
		57390: 58, // semi (2x)
		57380: 59, // action (1x)
		57366: 60, // IN (1x)
		57387: 61, // pattern (1x)
		57388: 62, // program (1x)
		57389: 63, // rule (1x)
		57379: 64, // $default (0x)
		57345: 65, // error (0x)
		57358: 66, // LEN (0x)
		57370: 67, // REGEXP (0x)
		57378: 68, // UNARY (0x)
	}

	yySymNames = []string{
		"'('",
		"'+'",
		"'-'",
		"IDENT",
		"'!'",
		"'$'",
		"FALSE",
		"FUNC",
		"NIL",
		"NUMBER",
		"STRING",
		"TRUE",
		"IF",
		"'}'",
		"BREAK",
		"CONTINUE",
		"DELETE",
		"FOR",
		"PRINT",
		"RETURN",
		"'\\n'",
		"'{'",
		"';'",
		"','",
		"')'",
		"'%'",
		"'*'",
		"'/'",
		"'<'",
		"'>'",
		"ANDAND",
		"DIVEQ",
		"EQEQ",
		"GE",
		"LE",
		"MINUSEQ",
		"MINUSMINUS",
		"MULEQ",
		"NEQ",
		"OROR",
		"PLUSEQ",
		"PLUSPLUS",
		"expr",
		"'='",
		"']'",
		"exprs",
		"$end",
		"BEGIN",
		"END",
		"stmt",
		"stmt_if",
		"stmts",
		"ELSE",
		"nls",
		"opt_nls",
		"'['",
		"ident_args",
		"opt_semi",
		"semi",
		"action",
		"IN",
		"pattern",
		"program",
		"rule",
		"$default",
		"error",
		"LEN",
		"REGEXP",
		"UNARY",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {62, 0},
		2:  {62, 2},
		3:  {63, 2},
		4:  {61, 0},
		5:  {61, 1},
		6:  {61, 1},
		7:  {61, 1},
		8:  {59, 5},
		9:  {51, 0},
		10: {51, 4},
		11: {49, 3},
		12: {49, 3},
		13: {49, 1},
		14: {49, 2},
		15: {49, 1},
		16: {49, 2},
		17: {49, 1},
		18: {49, 4},
		19: {49, 5},
		20: {49, 1},
		21: {49, 1},
		22: {49, 9},
		23: {49, 2},
		24: {50, 5},
		25: {50, 7},
		26: {50, 5},
		27: {45, 1},
		28: {45, 4},
		29: {42, 1},
		30: {42, 1},
		31: {42, 1},
		32: {42, 1},
		33: {42, 1},
		34: {42, 2},
		35: {42, 1},
		36: {42, 4},
		37: {42, 8},
		38: {42, 7},
		39: {42, 4},
		40: {42, 4},
		41: {42, 2},
		42: {42, 2},
		43: {42, 3},
		44: {42, 3},
		45: {42, 3},
		46: {42, 3},
		47: {42, 3},
		48: {42, 3},
		49: {42, 3},
		50: {42, 3},
		51: {42, 3},
		52: {42, 3},
		53: {42, 3},
		54: {42, 3},
		55: {42, 2},
		56: {42, 2},
		57: {42, 2},
		58: {42, 3},
		59: {42, 3},
		60: {42, 3},
		61: {42, 3},
		62: {42, 3},
		63: {42, 3},
		64: {56, 0},
		65: {56, 1},
		66: {56, 4},
		67: {53, 1},
		68: {53, 2},
		69: {54, 0},
		70: {54, 1},
		71: {58, 1},
		72: {57, 0},
		73: {57, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [150][]uint16{
		// 0
		{73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 21: 73, 46: 73, 73, 73, 62: 75},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 21: 70, 42: 80, 46: 74, 78, 79, 61: 77, 63: 76},
		{72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 21: 72, 46: 72, 72, 72},
		{21: 219, 59: 218},
		{21: 69},
		// 5
		{21: 68},
		{94, 110, 111, 21: 67, 25: 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95},
		{177, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 43: 45, 45, 55: 176},
		{44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 43: 44, 44},
		{43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43: 43, 43},
		// 10
		{42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 43: 42, 42},
		{41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 43: 41, 41},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 217},
		{39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 43: 39, 39},
		{145, 3: 144},
		// 15
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 143},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 142},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 141},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 93},
		{94, 110, 111, 24: 109, 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95},
		// 20
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 132, 45: 133},
		{33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 43: 33, 33},
		{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 43: 32, 32},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 131},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 130},
		// 25
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 129},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 128},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 127},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 126},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 125},
		// 30
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 124},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 123},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 122},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 121},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 120},
		// 35
		{16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 43: 16, 16},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 119},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 118},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 117},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 116},
		// 40
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 115},
		{94, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 100, 11, 11, 11, 98, 11, 99, 11, 11, 97, 11, 43: 11, 11},
		{94, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 100, 12, 12, 12, 98, 12, 99, 12, 12, 97, 12, 43: 12, 12},
		{94, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 100, 13, 13, 13, 98, 13, 99, 13, 13, 97, 13, 43: 13, 13},
		{94, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 114, 112, 113, 14, 14, 14, 100, 14, 14, 14, 98, 14, 99, 14, 14, 97, 14, 43: 14, 14},
		// 45
		{94, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 114, 112, 113, 15, 15, 15, 100, 15, 15, 15, 98, 15, 99, 15, 15, 97, 15, 43: 15, 15},
		{94, 110, 111, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 114, 112, 113, 105, 103, 20, 100, 101, 104, 106, 98, 96, 99, 102, 20, 97, 95, 43: 20, 20},
		{94, 110, 111, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 21, 97, 95, 43: 21, 21},
		{94, 110, 111, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 114, 112, 113, 22, 22, 22, 100, 22, 22, 22, 98, 96, 99, 22, 22, 97, 95, 43: 22, 22},
		{94, 110, 111, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 114, 112, 113, 23, 23, 23, 100, 23, 23, 23, 98, 96, 99, 23, 23, 97, 95, 43: 23, 23},
		// 50
		{94, 110, 111, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 114, 112, 113, 24, 24, 24, 100, 24, 24, 24, 98, 96, 99, 24, 24, 97, 95, 43: 24, 24},
		{94, 110, 111, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 114, 112, 113, 25, 25, 25, 100, 25, 25, 25, 98, 96, 99, 25, 25, 97, 95, 43: 25, 25},
		{94, 110, 111, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 114, 112, 113, 105, 103, 26, 100, 26, 104, 106, 98, 96, 99, 26, 26, 97, 95, 43: 26, 26},
		{94, 110, 111, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 114, 112, 113, 105, 103, 27, 100, 27, 104, 106, 98, 96, 99, 27, 27, 97, 95, 43: 27, 27},
		{94, 110, 111, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95, 43: 28, 28},
		// 55
		{94, 110, 111, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95, 43: 29, 29},
		{94, 110, 111, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95, 43: 30, 30},
		{94, 110, 111, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95, 43: 31, 31},
		{94, 110, 111, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 22: 47, 47, 47, 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95, 44: 47},
		{23: 134, 135},
		// 60
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20: 137, 53: 138, 136},
		{34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 43: 34, 34},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 140},
		{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 46: 7, 7, 7},
		{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 139, 4, 46: 4, 4, 4},
		// 65
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 46: 6, 6, 6},
		{94, 110, 111, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 22: 46, 46, 46, 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95, 43: 46, 46},
		{94, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 100, 17, 17, 17, 98, 17, 99, 17, 17, 97, 17, 43: 17, 17},
		{94, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 100, 18, 18, 18, 98, 18, 99, 18, 18, 97, 18, 43: 18, 18},
		{94, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 100, 19, 19, 19, 98, 19, 99, 19, 19, 97, 19, 43: 19, 19},
		// 70
		{211},
		{3: 147, 23: 10, 10, 56: 146},
		{23: 149, 148},
		{23: 9, 9},
		{21: 152},
		// 75
		{3: 5, 20: 137, 53: 138, 150},
		{3: 151},
		{23: 8, 8},
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 51: 153},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 164, 165, 161, 162, 157, 160, 158, 163, 42: 155, 45: 156, 49: 154, 159},
		// 80
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 22: 208, 57: 207, 209},
		{94, 110, 111, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 22: 61, 47, 25: 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95, 43: 205},
		{23: 134, 43: 203},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 202},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 59, 59, 59, 59, 59, 59, 59, 59, 59, 22: 59, 42: 132, 45: 201},
		// 85
		{57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 22: 57, 52: 192},
		{173, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 21: 171, 42: 172},
		{54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 22: 54},
		{53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 22: 53},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 132, 45: 170},
		// 90
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 166},
		{36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 43: 36, 36},
		{94, 110, 111, 21: 167, 25: 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95},
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 51: 168},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 164, 169, 161, 162, 157, 160, 158, 163, 42: 155, 45: 156, 49: 154, 159},
		// 95
		{50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 22: 50, 52: 50},
		{51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 22: 51, 134},
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 51: 190},
		{94, 110, 111, 21: 187, 25: 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95},
		{92, 89, 90, 174, 91, 86, 84, 88, 85, 82, 87, 83, 42: 93},
		// 100
		{177, 45, 45, 24: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 55: 176, 60: 175},
		{3: 182},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 132, 45: 180},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 132, 45: 178},
		{23: 134, 179},
		// 105
		{35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 43: 35, 35},
		{23: 134, 44: 181},
		{38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 43: 38, 38},
		{24: 183},
		{21: 184},
		// 110
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 51: 185},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 164, 186, 161, 162, 157, 160, 158, 163, 42: 155, 45: 156, 49: 154, 159},
		{52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 22: 52},
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 51: 188},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 164, 189, 161, 162, 157, 160, 158, 163, 42: 155, 45: 156, 49: 154, 159},
		// 115
		{55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 22: 55},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 164, 191, 161, 162, 157, 160, 158, 163, 42: 155, 45: 156, 49: 154, 159},
		{56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 22: 56},
		{12: 193, 21: 194},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 197},
		// 120
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 51: 195},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 164, 196, 161, 162, 157, 160, 158, 163, 42: 155, 45: 156, 49: 154, 159},
		{48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 22: 48, 52: 48},
		{94, 110, 111, 21: 198, 25: 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95},
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 51: 199},
		// 125
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 164, 200, 161, 162, 157, 160, 158, 163, 42: 155, 45: 156, 49: 154, 159},
		{49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 22: 49, 52: 49},
		{58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 22: 58, 134},
		{94, 110, 111, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 22: 60, 25: 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 132, 45: 204},
		// 130
		{62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 22: 62, 134},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 42: 206},
		{94, 110, 111, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 22: 63, 25: 114, 112, 113, 105, 103, 108, 100, 101, 104, 106, 98, 96, 99, 102, 107, 97, 95},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 137, 53: 138, 210},
		{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 46: 3, 3, 3},
		// 135
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 46: 1, 1, 1},
		{64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64},
		{3: 147, 23: 10, 10, 56: 212},
		{23: 149, 213},
		{21: 214},
		// 140
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 51: 215},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 164, 216, 161, 162, 157, 160, 158, 163, 42: 155, 45: 156, 49: 154, 159},
		{37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 43: 37, 37},
		{94, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 100, 40, 40, 40, 98, 40, 99, 40, 40, 97, 40, 43: 40, 40},
		{71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 21: 71, 46: 71, 71, 71},
		// 145
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 51: 220},
		{92, 89, 90, 81, 91, 86, 84, 88, 85, 82, 87, 83, 164, 221, 161, 162, 157, 160, 158, 163, 42: 155, 45: 156, 49: 154, 159},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 20: 2, 2, 208, 46: 2, 2, 2, 57: 222, 209},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20: 137, 5, 46: 5, 5, 5, 53: 138, 223},
		{66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 21: 66, 46: 66, 66, 66},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 65

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
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
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			yyVAL.rules = []ast.Rule{}
		}
	case 2:
		{
			yyVAL.rules = append(yyS[yypt-1].rules, yyS[yypt-0].rule)
			yylex.(*Lexer).result = yyVAL.rules
		}
	case 3:
		{
			yyVAL.rule = ast.Rule{Pattern: yyS[yypt-1].pattern, Action: yyS[yypt-0].stmts}
		}
	case 4:
		{
			yyVAL.pattern = &ast.ExprPattern{}
		}
	case 5:
		{
			yyVAL.pattern = &ast.BeginPattern{}
		}
	case 6:
		{
			yyVAL.pattern = &ast.EndPattern{}
		}
	case 7:
		{
			yyVAL.pattern = &ast.ExprPattern{Expr: yyS[yypt-0].expr}
		}
	case 8:
		{
			yyVAL.stmts = yyS[yypt-3].stmts
		}
	case 9:
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 10:
		{
			yyVAL.stmts = append(yyS[yypt-3].stmts, yyS[yypt-2].stmt)
		}
	case 11:
		{
			yyVAL.stmt = &ast.AssStmt{Left: []ast.Expr{yyS[yypt-2].expr}, Right: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 12:
		{
			yyVAL.stmt = &ast.AssStmt{Left: yyS[yypt-2].exprs, Right: yyS[yypt-0].exprs}
		}
	case 13:
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
		}
	case 14:
		{
			yyVAL.stmt = &ast.DelStmt{Expr: yyS[yypt-0].expr}
		}
	case 15:
		{
			yyVAL.stmt = &ast.PrintStmt{Exprs: defaultExprs}
		}
	case 16:
		{
			yyVAL.stmt = &ast.PrintStmt{Exprs: yyS[yypt-0].exprs}
		}
	case 17:
		{
			yyVAL.stmt = yyS[yypt-0].stmt_if
		}
	case 18:
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyS[yypt-1].stmts}
		}
	case 19:
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyS[yypt-1].stmts, Expr: yyS[yypt-3].expr}
		}
	case 20:
		{
			yyVAL.stmt = &ast.BreakStmt{}
		}
	case 21:
		{
			yyVAL.stmt = &ast.ContinueStmt{}
		}
	case 22:
		{
			yyVAL.stmt = &ast.HashLoopStmt{Key: yyS[yypt-6].token.Literal, Hash: yyS[yypt-4].token.Literal, Stmts: yyS[yypt-1].stmts}
		}
	case 23:
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
		}
	case 24:
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts, Else: nil}
		}
	case 25:
		{
			yyVAL.stmt_if.(*ast.IfStmt).ElseIf = append(yyVAL.stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts})
		}
	case 26:
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				//$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
				yyVAL.stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 27:
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 28:
		{
			yyVAL.exprs = append(yyS[yypt-3].exprs, yyS[yypt-0].expr)
		}
	case 29:
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 30:
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 31:
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 32:
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 33:
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 34:
		{
			yyVAL.expr = &ast.FieldExpr{Expr: yyS[yypt-0].expr}
		}
	case 35:
		{
			yyVAL.expr = &ast.StringExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 36:
		{
			yyVAL.expr = &ast.ItemExpr{Literal: yyS[yypt-3].token.Literal, Index: yyS[yypt-1].exprs}
		}
	case 37:
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].token.Literal, Args: yyS[yypt-4].ident_args, Stmts: yyS[yypt-1].stmts}
		}
	case 38:
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].ident_args, Stmts: yyS[yypt-1].stmts}
		}
	case 39:
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].token.Literal, SubExprs: yyS[yypt-1].exprs}
		}
	case 40:
		{
			yyVAL.expr = &ast.AnonymousCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
		}
	case 41:
		{
			yyVAL.expr = &ast.CompExpr{Left: yyS[yypt-1].expr, Operator: "++"}
		}
	case 42:
		{
			yyVAL.expr = &ast.CompExpr{Left: yyS[yypt-1].expr, Operator: "--"}
		}
	case 43:
		{
			yyVAL.expr = &ast.CompExpr{Left: yyS[yypt-2].expr, Operator: "+=", Right: yyS[yypt-0].expr}
		}
	case 44:
		{
			yyVAL.expr = &ast.CompExpr{Left: yyS[yypt-2].expr, Operator: "-=", Right: yyS[yypt-0].expr}
		}
	case 45:
		{
			yyVAL.expr = &ast.CompExpr{Left: yyS[yypt-2].expr, Operator: "*=", Right: yyS[yypt-0].expr}
		}
	case 46:
		{
			yyVAL.expr = &ast.CompExpr{Left: yyS[yypt-2].expr, Operator: "/=", Right: yyS[yypt-0].expr}
		}
	case 47:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "==", Right: yyS[yypt-0].expr}
		}
	case 48:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "!=", Right: yyS[yypt-0].expr}
		}
	case 49:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: ">", Right: yyS[yypt-0].expr}
		}
	case 50:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: ">=", Right: yyS[yypt-0].expr}
		}
	case 51:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "<", Right: yyS[yypt-0].expr}
		}
	case 52:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "<=", Right: yyS[yypt-0].expr}
		}
	case 53:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "||", Right: yyS[yypt-0].expr}
		}
	case 54:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "&&", Right: yyS[yypt-0].expr}
		}
	case 55:
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyS[yypt-0].expr}
		}
	case 56:
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
		}
	case 57:
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyS[yypt-0].expr}
		}
	case 58:
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 59:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "+", Right: yyS[yypt-0].expr}
		}
	case 60:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "-", Right: yyS[yypt-0].expr}
		}
	case 61:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "*", Right: yyS[yypt-0].expr}
		}
	case 62:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "/", Right: yyS[yypt-0].expr}
		}
	case 63:
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyS[yypt-2].expr, Operator: "%", Right: yyS[yypt-0].expr}
		}
	case 64:
		{
			yyVAL.ident_args = []string{}
		}
	case 65:
		{
			yyVAL.ident_args = []string{yyS[yypt-0].token.Literal}
		}
	case 66:
		{
			yyVAL.ident_args = append(yyS[yypt-3].ident_args, yyS[yypt-0].token.Literal)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
