package parser

import (
	"github.com/ShinNakamura/monkey/ast"
	"github.com/ShinNakamura/monkey/lexer"
	"testing"
)

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if l := len(program.Statements); l != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", l)
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if tl := returnStmt.TokenLiteral(); tl != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", tl)
		}
	}
}

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	/*
	input := `
let x 5;
let = 10;
let 838383;
`
	*/
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	checkParserErrors(t, p)
	if l := len(program.Statements); l != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", l)
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if tl := s.TokenLiteral(); tl != "let" {
		t.Errorf("s.TokenLiteral() not 'let'. got=%q", tl)
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if nv := letStmt.Name.Value; nv != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, nv)
		return false
	}

	if nt := letStmt.Name.TokenLiteral(); nt != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", name, nt)
		return false
	}

	return true
}
