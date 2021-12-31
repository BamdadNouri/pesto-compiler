package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"sandbox/pesto/token"
	"strings"
)

var (
	LaxerOnly  bool = false
	FilePath   string
	TestLet    bool = false
	TestPrefix bool = false
	TestInfix  bool = false
	TestBool   bool = false
	TestIf     bool = false
)

func init() {
	flag.StringVar(&FilePath, "f", "", "path to a *.pesto file")

	flag.BoolVar(&TestLet, "test_let", false, "test Let")
	flag.BoolVar(&TestPrefix, "test_prefix", false, "test Prefix")
	flag.BoolVar(&TestInfix, "test_infix", false, "test Infix")
	flag.BoolVar(&TestBool, "test_bool", false, "test Bool")
	flag.BoolVar(&TestIf, "test_if", false, "test If")
	flag.Parse()
}

func main() {
	if TestLet {
		testLet()
	}
	if TestPrefix {
		testPrefix()
	}
	if TestInfix {
		testInfix()
	}
	if TestBool {
		testBool()
	}
	if TestIf {
		testIf()
	}
	if TestLet || TestPrefix || TestInfix || TestBool || TestIf {
		return
	}

	file, err := readPestoFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	input := string(file)
	laxer := NewLaxer(input)
	if LaxerOnly {
		for {
			l := laxer.NextToken()

			fmt.Println(l)
			if l.Type == token.EOF {
				break
			}
		}
		return
	}

	program := NewParser(laxer).ParseProgram()
	for _, s := range program.Statements {
		fmt.Println(s.String())
		fmt.Println(s.TokenLiteral())
		fmt.Println(s.(*LetStatement).Name)
		fmt.Println(s.(*LetStatement).Value)

	}
}

func readPestoFile() (string, error) {
	if FilePath == "" {
		return "", errors.New("there is no file provided. use -f flag to compile your *.pesto file")
	}
	a := strings.Split(FilePath, ".")
	if a[len(a)-1] != "pesto" {
		return "", errors.New("wrong file time. only *.pesto files are allowd")
	}
	bs, err := ioutil.ReadFile(FilePath)
	if err != nil {
		return "", fmt.Errorf("error in reading file: %s", err)

	}
	return string(bs), nil
}

func testLet() {
	program := NewParser(
		NewLaxer(`
	let x = 10;
	`),
	).ParseProgram()

	if len(program.Statements) != 1 {
		fmt.Println("test fail in testLet")
	}
	fmt.Println("statement:", program.Statements[0].String())
	fmt.Println("name:", program.Statements[0].(*LetStatement).Name)
	fmt.Println("value:", program.Statements[0].(*LetStatement).Value)
}

func testPrefix() {
	program := NewParser(
		NewLaxer(`
	!true;
	`),
	).ParseProgram()

	if len(program.Statements) != 1 {
		fmt.Println("test fail in testLet")
	}
	fmt.Println("statement:", program.Statements[0].String())
	fmt.Println("operator:", program.Statements[0].(*ExpressionStatement).Expression.(*PrefixExpression).Operator)
	fmt.Println("right value:", program.Statements[0].(*ExpressionStatement).Expression.(*PrefixExpression).Right)
}

func testInfix() {
	program := NewParser(
		NewLaxer(`
	4 * 8;;
	`),
	).ParseProgram()

	if len(program.Statements) == 0 {
		fmt.Println("test fail in testLet")
	}
	fmt.Println("statement:", program.Statements[0].String())
	fmt.Println("left value:", program.Statements[0].(*ExpressionStatement).Expression.(*InfixExpression).Left)
	fmt.Println("operator:", program.Statements[0].(*ExpressionStatement).Expression.(*InfixExpression).Operator)
	fmt.Println("right value:", program.Statements[0].(*ExpressionStatement).Expression.(*InfixExpression).Right)
}

func testBool() {
	program := NewParser(
		NewLaxer(`
	true;
	`),
	).ParseProgram()

	if len(program.Statements) != 1 {
		fmt.Println("test fail in testLet")
	}
	fmt.Println("statement:", program.Statements[0].String())
	fmt.Println("value:", program.Statements[0].(*ExpressionStatement).Expression.(*Boolean).Value)
}

func testIf() {
	program := NewParser(
		NewLaxer(`
	if (x == y) { x };
	`),
	).ParseProgram()

	if len(program.Statements) != 1 {
		fmt.Println("test fail in testLet")
	}
	fmt.Println("statement:", program.Statements[0].String())
	fmt.Println("condition:", program.Statements[0].(*ExpressionStatement).Expression.(*IfExpression).Condition)
	fmt.Println("identifier value:", program.Statements[0].(*ExpressionStatement).Expression.(*IfExpression).Consequence.Statements[0].(*ExpressionStatement).Expression.(*Identifier).Value)
}
