# Pesto compiler
This is a tiny compiler just to learn how compilers work. It takes a `*.pesto` file and runs multiple stages of compiler.

This compiler does Laxer and Parser. It first read the input program, then Laxers breaks down the code and make chunkes called Tokens.
Tokens are objects representing what each peace of code do.
Then it's sent to Parser. Parser takes tokens and verifies them. Parser, parses the token list to Statements. Each statement has a Token, which is it's token with a Type and a Literal (like a `let` which has a LET type and a "let" literal), a Name, which is a identifier, that means it has a Token and a Value.

## How to use
In a machine running golang, run the command bellow to compiler `test.pesto`:

`go run . -f test.pesto`

It runs Laxer & Parser on this code.

For ease of use, I've implemented multiple tests to run on sample code. using the flags bellow each test will run and the results are printed. (keep in mid, this repository is not complete yet. There are more to these tests, does compiler does a bit more; it'll be implemented later on)


`-test_let`: runs _`let x = 10;`_ and the output is _`statement, name: "x", value: 10`_ with _LetStatement_

`-test_prefix`: handles prefixes like !true, it run _`!true`_ and the output is _`statement, operator: "!", right value: true`_ with _ExpressionStatement_ and expression: _PrefixExpression_

`-test_infix`: handles infixes and operations like 5+8, it run_`4 * 8`_ and the output is _`statement, left value: 4, operator: "*", right value: 8`_ with _ExpressionStatement_ and expression: _InfixExpression_, and the Left of operator and Right of operator

`-test_bool`: run _`"true"`_ and the output is _`statement, value: true`_ with _ExpressionStatement_ and expression: _Boolean_

`-test_if`: run _`if (x == y) { x }`_, which first the condition is parsed through Infix expressions with it's condition, and then the output is _`statement, condition (x==y), identifier value: x`_ with _ExpressionStatement_ and expression: _IfExpression_ with Condition
