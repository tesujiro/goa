%{
package parser
import (
	//"fmt"
	"github.com/tesujiro/ago/ast"
)

var defaultExpr = ast.FieldExpr{Expr: &ast.NumExpr{Literal: "0"}}
var defaultExprs = []ast.Expr{&defaultExpr}
var inRegExp bool
%}

%union{
	token		ast.Token
	rule		ast.Rule
	rules		[]ast.Rule
	pattern		ast.Pattern
	stmt		ast.Stmt
	stmts		[]ast.Stmt
	expr		ast.Expr
	exprs		[]ast.Expr
	identArgs	[]string
}

%type <rules>		program
%type <rule>		rule
%type <pattern> 	pattern
%type <stmt>		stmt
%type <stmt>		opt_stmt
%type <stmts>		action
%type <stmts>		stmts
%type <stmts>		opt_stmts
%type <stmt>		stmt_if
%type <expr>		expr
%type <expr>		common_expr
%type <expr>		multi_val_expr
%type <expr>		simp_expr
%type <expr>		non_post_simp_expr
%type <expr>		opt_variable
%type <expr>		variable
%type <expr>		simple_variable
%type <expr>		opt_expr
%type <expr>		regexp_literal
%type <expr>		opt_input_redir
%type <exprs>		exprs
%type <exprs>		variables
%type <exprs>		opt_exprs
%type <identArgs>	ident_args

%token <token> IDENT NUMBER STRING TRUE FALSE NIL POW
%token <token> EQEQ NEQ GE LE NOTTILDE ANDAND OROR LEN 
%token <token> PLUSPLUS MINUSMINUS PLUSEQ MINUSEQ MULEQ DIVEQ MODEQ POWEQ
%token <token> DELETE IN
%token <token> BEGIN END PRINT PRINTF REGEXP
%token <token> IF ELSE FOR WHILE DO BREAK CONTINUE
%token <token> FUNC RETURN EXIT NEXT NEXTFILE
%token <token> CONCAT_OP GETLINE

%nonassoc LOWEST
%right PLUSEQ MINUSEQ MULEQ DIVEQ MODEQ POWEQ '='
%right '?' ':'
%left IN
%left OROR
%left ANDAND
%left GETLINE
%nonassoc ','
%left '~' NOTTILDE
%nonassoc EQEQ NEQ
%nonassoc '>' '<' GE LE

%left CONCAT_OP
%left STRING NUMBER
%nonassoc IDENT TRUE FALSE PRINTF FUNC NIL
%left '+' '-'
%left '*' '/' '%'
%right '!' UNARY
%left POW
%left PLUSPLUS MINUSMINUS
%right '$'
%nonassoc '['
%left '(' ')'

%%

program
	: opt_term /* empty */
	{
		$$ = []ast.Rule{}
	}
	| program rule 
	{
		$$ = append($1,$2)
		yylex.(*Lexer).result = $$
	}


rule
	: pattern action opt_term
	{
		$$ = ast.Rule{Pattern: $1, Action: $2}
	}
	| pattern term
	{
		$$ = ast.Rule{Pattern: $1, Action: []ast.Stmt{ &ast.PrintStmt{Exprs: defaultExprs }}}
	}
	| action opt_term
	{
		$$ = ast.Rule{Pattern: &ast.ExprPattern{}, Action: $1}
	}

pattern
	/*
	:
	{
		$$ = &ast.ExprPattern{}
	}
	*/
	: FUNC IDENT '(' ident_args ')'
	{
		//fmt.Println("FUNC RULE")
		$$ = &ast.FuncPattern{Name: $2.Literal, Args: $4}
	}
	| BEGIN
	{
		$$ = &ast.BeginPattern{}
	}
	| END
	{
		$$ = &ast.EndPattern{}
	}
	| expr
	{
		$$ = &ast.ExprPattern{Expr:$1}
	}
	| expr ',' expr
	{
		$$ = &ast.StartStopPattern{
			Start: $1,
			Stop: $3,
		}
	}

action
	: '{' opt_stmts '}'
	{
		$$ = $2
	}

opt_stmts
	: /* empty */
	{
		$$ = []ast.Stmt{}
	}
	| stmts opt_term
	{
		$$ = $1
	}
	
stmts
	: opt_term stmt
	{
		$$ = []ast.Stmt{$2}
	}
	| stmts semi opt_term stmt
	{
		$$ = append($1,$4)
	}

stmt
	: expr
	{
		$$ = &ast.ExprStmt{Expr: $1}
	}
	| multi_val_expr
	{
		$$ = &ast.ExprStmt{Expr: $1}
	}
	| DELETE expr
	{
		$$ = &ast.DelStmt{Expr: $2}
	}
	| PRINT 
	{
		$$ = &ast.PrintStmt{Exprs: defaultExprs }
	}
	| PRINT exprs
	{
		$$ = &ast.PrintStmt{Exprs: $2}
	}
	| stmt_if
	{
		$$ = $1
	}
	| FOR '{' opt_stmts '}'
	{
		$$ = &ast.LoopStmt{Stmts: $3}
	}
	| FOR expr '{' opt_stmts '}'
	{
		$$ = &ast.LoopStmt{Stmts: $4, Expr: $2}
	}
	| FOR opt_stmt ';' opt_expr ';' opt_expr '{' opt_stmts '}'
	{
		$$ = &ast.CForLoopStmt{Stmt1: $2, Expr2: $4, Expr3: $6, Stmts: $8}
	}
	| FOR '(' opt_stmt ';' opt_expr ';' opt_expr ')' '{' opt_stmts '}'
	{
		$$ = &ast.CForLoopStmt{Stmt1: $3, Expr2: $5, Expr3: $7, Stmts: $10}
	}
	| FOR '(' opt_stmt ';' opt_expr ';' opt_expr ')' stmt 
	{
		$$ = &ast.CForLoopStmt{Stmt1: $3, Expr2: $5, Expr3: $7, Stmts: []ast.Stmt{$9}}
	}
	| FOR '(' opt_stmt ';' opt_expr ';' opt_expr ')' semi stmt 
	{
		$$ = &ast.CForLoopStmt{Stmt1: $3, Expr2: $5, Expr3: $7, Stmts: []ast.Stmt{$10}}
	}
	| WHILE '{' opt_stmts '}'
	{
		$$ = &ast.LoopStmt{Stmts: $3}
	}
	| WHILE expr '{' opt_stmts '}'
	{
		$$ = &ast.LoopStmt{Stmts: $4, Expr: $2}
	}
/*
	| WHILE '(' expr ')' opt_term stmt opt_term
	{
		$$ = &ast.LoopStmt{Stmts: []ast.Stmt{$6}, Expr: $3}
	}
*/
	| DO '{' opt_stmts '}' WHILE '(' expr ')'
	{
		$$ = &ast.DoLoopStmt{Stmts: $3, Expr: $7}
	}
	| BREAK
	{
		$$ = &ast.BreakStmt{}
	}
	| CONTINUE
	{
		$$ = &ast.ContinueStmt{}
	}
	| NEXT
	{
		$$ = &ast.NextStmt{}
	}
	| NEXTFILE
	{
		$$ = &ast.NextfileStmt{}
	}
	| FOR '(' IDENT IN IDENT ')' '{' opt_stmts '}'
	{
		$$ = &ast.MapLoopStmt{KeyID: $3.Literal, MapID: $5.Literal, Stmts:$8}
	}
	| RETURN opt_exprs
	{
		$$ = &ast.ReturnStmt{Exprs:$2}
	}
	| EXIT opt_expr
	{
		$$ = &ast.ExitStmt{Expr:$2}
	}

stmt_if
	: IF expr '{' opt_stmts '}'
	{
	    //fmt.Println("stmt_if:1")
	    $$ = &ast.IfStmt{If: $2, Then: $4, Else: nil}
	}
	| IF expr opt_semi stmt opt_semi opt_semi
	{
	    $$ = &ast.IfStmt{If: $2, Then: []ast.Stmt{$4}, Else: nil}
	}
/*
	| IF '(' expr ')' stmt opt_semi
	{
	    $$ = &ast.IfStmt{If: $3, Then: []ast.Stmt{$5}, Else: nil}
	}
	| IF '(' expr ')' semi stmt
	{
	    $$ = &ast.IfStmt{If: $3, Then: []ast.Stmt{$6}, Else: nil}
	}
	| IF '(' expr ')' semi stmt semi opt_semi
	//| IF '(' expr ')' semi stmt opt_semi
	{
	    $$ = &ast.IfStmt{If: $3, Then: []ast.Stmt{$6}, Else: nil}
	}
*/
	| stmt_if ELSE IF expr '{' opt_stmts '}'
	{
	        $$.(*ast.IfStmt).ElseIf = append($$.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: $4, Then: $6} )
	}
/*
	//| stmt_if ELSE IF '(' expr ')' stmt 
	| stmt_if ELSE IF '(' expr ')' stmt opt_term
	//| stmt_if ELSE IF '(' expr ')' stmt opt_semi
	{
	        $$.(*ast.IfStmt).ElseIf = append($$.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: $5, Then: []ast.Stmt{$7}} )
	}
	| stmt_if ELSE IF '(' expr ')' opt_semi stmt opt_semi opt_semi
	//| stmt_if ELSE IF '(' expr ')' semi stmt opt_semi opt_semi
	//| stmt_if ELSE IF '(' expr ')' opt_semi stmt opt_term
	//| stmt_if ELSE IF '(' expr ')' opt_semi stmt opt_semi opt_semi
	{
	        $$.(*ast.IfStmt).ElseIf = append($$.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: $5, Then: []ast.Stmt{$8}} )
	}
*/
	| stmt_if ELSE IF expr opt_semi stmt opt_semi opt_semi
	{
	        $$.(*ast.IfStmt).ElseIf = append($$.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: $4, Then: []ast.Stmt{$6}} )
	}
	| stmt_if ELSE '{' opt_stmts '}'
	{
		if $$.(*ast.IfStmt).Else != nil {
			yylex.Error("multiple else statement")
		} else {
			//$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
			$$.(*ast.IfStmt).Else = $4
		}
	}
/*
	| stmt_if ELSE stmt opt_term
	{
		if $$.(*ast.IfStmt).Else != nil {
			yylex.Error("multiple else statement")
		} else {
			$$.(*ast.IfStmt).Else = []ast.Stmt{$3}
		}
	}
*/
	| stmt_if ELSE opt_semi stmt opt_term
	{
		if $$.(*ast.IfStmt).Else != nil {
			yylex.Error("multiple else statement")
		} else {
			$$.(*ast.IfStmt).Else = []ast.Stmt{$4}
		}
	}

multi_val_expr
	: variables '=' exprs
	{
		$$ = &ast.AssExpr{Left: $1, Right: $3}
	}

opt_exprs
	: %prec LOWEST/* empty */
	{
		$$ = []ast.Expr{}
	}
	| exprs %prec LOWEST
	{
		$$ = $1
	}

exprs
	: expr %prec LOWEST
	{
		$$ = []ast.Expr{$1}
	}
	| exprs ',' opt_term expr
	{
		$$ = append($1,$4)
	}

expr
	: variable '=' expr
	{
		$$ = &ast.AssExpr{Left: []ast.Expr{$1}, Right: []ast.Expr{$3}}
	}
	/* RELATION EXPRESSION */
	| expr EQEQ expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "==", Right: $3}
	}
	| expr NEQ expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "!=", Right: $3}
	}
	| expr '>' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: ">", Right: $3}
	}
	| expr GE expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: ">=", Right: $3}
	}
	| expr '<' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "<", Right: $3}
	}
	| expr LE expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "<=", Right: $3}
	}
	| expr IN IDENT
	{
		$$ = &ast.ContainKeyExpr{KeyExpr: $1, MapID: $3.Literal}
	}
	/* REGEXP */
	//TODO:
	//| expr '~' expr
	| expr '~' regexp_literal
	{
		$$ = &ast.MatchExpr{Expr: $1, RegExpr: $3}
	}
        | expr '~' common_expr
	{
		$$ = &ast.MatchExpr{Expr: $1, RegExpr: $3}
	}
	//TODO:
	//| expr NOTTILDE expr
	| expr NOTTILDE regexp_literal
	{
		$$ = &ast.UnaryExpr{Operator: "!", Expr: &ast.MatchExpr{Expr: $1, RegExpr: $3}}
	}
	| expr NOTTILDE common_expr
	{
		$$ = &ast.UnaryExpr{Operator: "!", Expr: &ast.MatchExpr{Expr: $1, RegExpr: $3}}
	}
/*	//TODO: remove regexp_literal
*/
	| regexp_literal
	{
		$$ = &ast.MatchExpr{Expr: &defaultExpr, RegExpr: $1}
	}
	/* COMPOSITE ASSIGN OPERATION */
	| variable PLUSEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "+=", Right: $3}
	}
	| variable MINUSEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "-=", Right: $3}
	}
	| variable MULEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "*=", Right: $3}
	}
	| variable DIVEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "/=", Right: $3}
	}
	| variable MODEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "%=", Right: $3}
	}
	| variable POWEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "^=", Right: $3}
	}
	/* TERNARY OPERATOR */
	//| expr '?' expr ':' expr
	| expr '?' expr ':' expr
	{
		$$ = &ast.TriOpExpr{Cond: $1, Then: $3, Else: $5}
	}
	/* BOOL EXPRESSION */
	//| expr OROR expr
	| expr OROR expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "||", Right: $3}
	}
	//| expr ANDAND expr
	| expr ANDAND expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "&&", Right: $3}
	}
	| PRINTF opt_exprs
	{
		$$ = &ast.CallExpr{Name: "printf", SubExprs:$2}
	}
	| common_expr %prec LOWEST
	{
		$$ = $1
	}
	/* GETLINE */
	| common_expr '|' GETLINE opt_variable
	{
		$$ = &ast.GetlineExpr{Command: $1, Var: $4}
	}
	| GETLINE opt_variable opt_input_redir
	{
		$$ = &ast.GetlineExpr{Var: $2, Redir: $3}
	}

regexp_literal
	: '/' /* REGEXP contiunes next */
	{
		//fmt.Println("YACC: want regexp!!")
		inRegExp=true
	}
	REGEXP
	{
		//fmt.Println("FINISH")
		$$ = &ast.StringExpr{Literal: $3.Literal}
	}

common_expr
	: simp_expr %prec LOWEST
	{
		$$ = $1
	}
	/* CONCATENATE */
	| common_expr simp_expr %prec CONCAT_OP
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "CAT", Right: $2}
	}

simp_expr
	: non_post_simp_expr %prec LOWEST
	{
		$$ = $1
	}
	/* ARITHMETIC EXPRESSION */
	| simp_expr '+' simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "+", Right: $3}
	}
	| simp_expr '-' simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "-", Right: $3}
	}
	| simp_expr '*' simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "*", Right: $3}
	}
	| simp_expr '/' simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "/", Right: $3}
	}
	| simp_expr '%' simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "%", Right: $3}
	}
	| simp_expr POW simp_expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "^", Right: $3}
	}
	/* COMPOSITE EXPRESSION */
	| simp_expr PLUSPLUS
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "++", After:true}
	}
	| simp_expr MINUSMINUS
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "--", After:true}
	}

opt_input_redir
	: /* empty */
	{
		$$ = nil
	}
	| '<' simp_expr
	{
		$$ = $2
	}

non_post_simp_expr
	: '!' simp_expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "!", Expr:$2}
	}
/*
	| regexp_literal
	{
		//fmt.Println("regexp_literal in non_post_simp_expr")
		$$ = $1
	}
*/
	/* FUNCTION CALL */
	| IDENT '(' opt_exprs ')'
	{
		$$ = &ast.CallExpr{Name: $1.Literal, SubExprs:$3}
	}
	| PRINTF '(' opt_exprs ')'
	{
		$$ = &ast.CallExpr{Name: $1.Literal, SubExprs:$3}
	}
	| non_post_simp_expr '(' opt_exprs ')'
	{
		$$ = &ast.AnonymousCallExpr{Expr: $1, SubExprs:$3}
	}
	/* FUNCTION DEFINITION */
	| FUNC '(' ident_args ')' '{' opt_stmts '}'
	{
		$$ = &ast.FuncExpr{Args: $3, Stmts: $6}
	}
	/* ARITHMETIC EXPRESSION */
	| '(' expr ')'
	{
		$$ = &ast.ParentExpr{SubExpr: $2}
	}
	/* COMPOSITE EXPRESSION */
	| PLUSPLUS simp_expr
	{
		$$ = &ast.CompExpr{Left: $2, Operator: "++"}
	}
	| MINUSMINUS simp_expr
	{
		$$ = &ast.CompExpr{Left: $2, Operator: "--"}
	}
	/* LITERAL */
	| NUMBER
	{
		$$ = &ast.NumExpr{Literal: $1.Literal}
	}
	| TRUE
	{
		$$ = &ast.ConstExpr{Literal: $1.Literal}
	}
	| FALSE
	{
		$$ = &ast.ConstExpr{Literal: $1.Literal}
	}
	| NIL
	{
		$$ = &ast.ConstExpr{Literal: $1.Literal}
	}
	| STRING
	{
		$$ = &ast.StringExpr{Literal: $1.Literal}
	}
	/* var */
	| variable
	{
		$$ = $1
	}
	/* UNARY EXPRESSION */
	| '+' simp_expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "+", Expr:$2}
	}
	| '-' simp_expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "-", Expr:$2}
	}

variables
	: variable
	{
		$$ = []ast.Expr{$1}
	}
	| variables ',' opt_term variable
	{
		$$ = append($1,$4)
	}

opt_variable
	: /* empty */
	{
		$$ = nil
	}
	| variable
	{
		$$ = $1
	}
	;

variable
	: simple_variable
	{
		$$ = $1
	}
	| '$' non_post_simp_expr
	{
		$$ = &ast.FieldExpr{Expr: $2}
	}

simple_variable
	: non_post_simp_expr '[' exprs ']'
	{
		$$ = &ast.ItemExpr{Expr: $1, Index:$3}
	}
	| IDENT %prec LOWEST
	{
		$$ = &ast.IdentExpr{Literal: $1.Literal}
	}

ident_args
	: /* empty */
	{
		$$ = []string{}
	}
	| IDENT
	{
		$$ = []string{$1.Literal}
	}
	| ident_args ',' opt_term IDENT
	{
		$$ = append($1,$4.Literal)
	}

opt_stmt
	: /* empty */
	{
		$$ = nil
	}
	| stmt
	{
		$$ = $1
	}

opt_expr
	: /* empty */
	{
		$$ = nil
	}
	| expr
	{
		$$ = $1
	}

opt_term
	: /* empty */
	| term

term
	: semi
	| term semi

opt_semi
	: /* empty */
	| semi

semi
	: ';'  /* go/scanner return semi when EOL */

/*
a_slash
	: '/'
opt_nls
	: 
	| nls

nls
	: '\n'
	| nls '\n'
*/

%%
