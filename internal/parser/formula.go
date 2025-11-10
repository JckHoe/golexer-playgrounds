package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Formula struct {
	Expression *Expression `@@`
}

type Expression struct {
	Or []*OrTerm `@@ ( "OR" @@ )*`
}

type OrTerm struct {
	And []*AndTerm `@@ ( "AND" @@ )*`
}

type AndTerm struct {
	Not        bool        `@"NOT"?`
	Comparison *Comparison `@@`
}

type Comparison struct {
	Left     *Addition `@@`
	Operator string    `( @( ">=" | "<=" | "==" | "!=" | ">" | "<" )`
	Right    *Addition `  @@ )?`
}

type Addition struct {
	Left     *Multiplication `@@`
	Operator string          `( @("+" | "-")`
	Right    *Addition       `  @@ )?`
}

type Multiplication struct {
	Left     *Unary          `@@`
	Operator string          `( @("*" | "/" | "%")`
	Right    *Multiplication `  @@ )?`
}

type Unary struct {
	FunctionCall *FunctionCall `  @@`
	Variable     *Variable     `| @@`
	Number       *float64      `| @Float | @Int`
	String       *string       `| @String`
	Boolean      *bool         `| (@"true" | "false")`
	SubExpr      *Expression   `| "(" @@ ")"`
}

type FunctionCall struct {
	Name      string        `@Ident`
	Arguments []*Expression `"(" ( @@ ( "," @@ )* )? ")"`
}

type Variable struct {
	Parts []string `@Ident ( "." @Ident )*`
}

var (
	formulaLexer = lexer.MustSimple([]lexer.SimpleRule{
		{Name: "Float", Pattern: `\d+\.\d+`},
		{Name: "Int", Pattern: `\d+`},
		{Name: "String", Pattern: `"[^"]*"|'[^']*'`},
		{Name: "Ident", Pattern: `[a-zA-Z_][a-zA-Z0-9_]*`},
		{Name: "Operators", Pattern: `>=|<=|==|!=|[><=+\-*/%(),.]`},
		{Name: "whitespace", Pattern: `\s+`},
	})

	Parser = participle.MustBuild[Formula](
		participle.Lexer(formulaLexer),
		participle.Unquote("String"),
		participle.CaseInsensitive("NOT", "AND", "OR", "true", "false"),
	)
)
