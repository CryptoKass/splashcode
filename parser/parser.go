// Copyright 2018 <kassCrypto@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this
// software and associated documentation files (the "Software"), to deal in the Software
// without restriction, including without limitation the rights to use, copy, modify,
// merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be included in all copies
// or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
// PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF
// CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package parser

import (
	"fmt"
	"splashcode/lexer"
)

type Program struct {
	Markers map[string]int
	Stack   Stack
	Tokens  []lexer.Token
}

type Stack []lexer.Token

func (stack Stack) Push(token lexer.Token) Stack {
	return append(stack, token)
}

func (stack Stack) Pop() (Stack, lexer.Token) {
	l := len(stack)
	v := stack[l-1]
	stack = stack[:l-1]
	return stack, v
}

func (stack Stack) Read() lexer.Token {
	l := len(stack)
	v := stack[l-1]
	return v
}

func (stack Stack) Pick(i int) lexer.Token {
	l := len(stack)
	v := stack[l-i]
	return v
}

func (stack Stack) Delete(i int) Stack {
	s := stack
	l := len(s)
	s = append(s[:l-i], s[l-i+1:]...)
	return s
}

func (stack Stack) String() string {
	s := "["
	for i := 0; i < len(stack); i++ {
		s += "{"
		s += lexer.TokenTypeToString(stack[i].TokenType)
		s += ":"
		s += fmt.Sprintf("%v", stack[i].Value)
		s += "}, "
	}
	s += "]"
	return s
}

// Parse will take some splashgo tokens convert them to
// a *Program object which can be executed.
func Parse(tokens []lexer.Token) (prog Program) {
	prog.Markers = make(map[string]int)
	prog.Tokens = tokens
	prog.Stack = make(Stack, 0)
	prog.Stack.Push(lexer.Token{0, 0})

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token.TokenType {
		case lexer.TypeMARK:
			labelToken := tokens[i+1]
			prog.Markers[labelToken.Value.(string)] = i + 1
			break
		case lexer.TypeFUNC:
			labelToken := tokens[i+1]
			prog.Markers[labelToken.Value.(string)] = i + 1
			endPoint := prog.findNext(i, lexer.TypeENDFUNC)
			prog.Tokens[i].Value = endPoint
			break
		case lexer.TypeIF:
			endPoint := prog.findNext(i, lexer.TypeENDIF)
			prog.Tokens[i].Value = endPoint
		default:
			// nothing
			break
		}
	}

	return
}

func (prog *Program) findNext(index int, tokenType int) int {
	for cursor := index; cursor < len(prog.Tokens); cursor++ {
		if prog.Tokens[cursor].TokenType == tokenType {
			return cursor
		}
	}
	return -1
}
