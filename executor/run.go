package executor

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"splashcode/lexer"
	"splashcode/parser"
)

// Run will execute the given program (parser.Program), any args
// are passed into the input, and a stack trace maybe eneabled.
func Run(prog *parser.Program, input lexer.Token, stackTrace bool) (parser.Stack, lexer.Token) {
ExecutionLoop:
	// Loop through all tokens in program
	for i := 0; i < len(prog.Tokens); i++ {

		token := prog.Tokens[i]

		if stackTrace {
			fmt.Println("STRACT::STACK", prog.Stack)
			fmt.Println("STRACE::TOKEN", lexer.TokenTypeToString(token.TokenType), token.Value)
		}

		switch token.TokenType {
		case lexer.TypeGOTO:
			// Move execution cursor 'i' to marker
			label := prog.Tokens[i+1].Value.(string)
			i = prog.Markers[label]
			break
		case lexer.TypeMARK:
			// Add/Udate marker to "i + 1"
			label := prog.Tokens[i+1].Value.(string)
			prog.Markers[label] = i + 1
			i++
			break
		case lexer.TypeIF:
			var TokeA, TokeB lexer.Token
			prog.Stack, TokeA = prog.Stack.Pop()
			prog.Stack, TokeB = prog.Stack.Pop()
			if TokeA.TokenType != TokeB.TokenType ||
				TokeA.Value != TokeB.Value {
				i = token.Value.(int)
			}
			break
		case lexer.TypeENDIF:
			//NOTHING
			break
		case lexer.TypeFUNC:
			// Move execution cursor 'i' to marker
			i = token.Value.(int)
			break
		case lexer.TypeENDFUNC:
			// Nothing
			break
		case lexer.TypeDUP:
			val := prog.Stack.Read()
			prog.Stack = prog.Stack.Push(val)
			break
		case lexer.TypeDROP:
			prog.Stack, _ = prog.Stack.Pop()
			break
		case lexer.TypePICK:
			count := prog.Tokens[i+1].Value.(int)
			val := prog.Stack.Pick(count)
			prog.Stack = prog.Stack.Push(val)
			break
		case lexer.TypeROLL:
			count := prog.Tokens[i+1].Value.(int)
			val := prog.Stack.Pick(count)
			prog.Stack = prog.Stack.Push(val)
			prog.Stack = prog.Stack.Delete(count)
			break
		case lexer.TypeADD:
			var TokeA, TokeB lexer.Token
			prog.Stack, TokeA = prog.Stack.Pop()
			prog.Stack, TokeB = prog.Stack.Pop()
			prog.Stack = prog.Stack.Push(tokenAddition(TokeA, TokeB))
			break
		case lexer.TypeSUB:
			var TokeA, TokeB lexer.Token
			prog.Stack, TokeA = prog.Stack.Pop()
			prog.Stack, TokeB = prog.Stack.Pop()
			prog.Stack = prog.Stack.Push(tokenSubtraction(TokeA, TokeB))
			break
		case lexer.TypeMUL:
			var TokeA, TokeB lexer.Token
			prog.Stack, TokeA = prog.Stack.Pop()
			prog.Stack, TokeB = prog.Stack.Pop()
			prog.Stack = prog.Stack.Push(tokenMultiply(TokeA, TokeB))
			break
		case lexer.TypeDIV:
			var TokeA, TokeB lexer.Token
			prog.Stack, TokeA = prog.Stack.Pop()
			prog.Stack, TokeB = prog.Stack.Pop()
			prog.Stack = prog.Stack.Push(tokenDivide(TokeA, TokeB))
			break
		case lexer.TypeHASH:
			var TokeA lexer.Token
			prog.Stack, TokeA = prog.Stack.Pop()
			prog.Stack = prog.Stack.Push(tokenHash(TokeA))

		case lexer.TypeFIN:
			break ExecutionLoop
		case lexer.TypeINPUT:
			prog.Stack = prog.Stack.Push(input)
			break
		case lexer.TypePRINT:
			fmt.Print(prog.Stack.Read().Value)
			break
		case lexer.TypePRINTLN:
			fmt.Println(prog.Stack.Read().Value)
		default:
			prog.Stack = prog.Stack.Push(token)
			break
		}
	}

	// Return last token
	if len(prog.Stack) > 0 {
		return prog.Stack.Pop()
	}
	return prog.Stack, lexer.Token{}
}

// tokenAddition will add two tokens together and will return the
// result.
func tokenAddition(tokenA lexer.Token, tokenB lexer.Token) (result lexer.Token) {

	aIsInt := tokenA.TokenType == lexer.TypeINT
	bIsInt := tokenB.TokenType == lexer.TypeINT

	if aIsInt && bIsInt {
		result.TokenType = lexer.TypeINT
		result.Value = tokenA.Value.(int64) + tokenB.Value.(int64)
		return
	}
	if !aIsInt && !bIsInt {
		result.TokenType = lexer.TypeFLOAT
		result.Value = tokenA.Value.(float64) + tokenB.Value.(float64)
		return
	}
	if aIsInt && !bIsInt {
		result.TokenType = lexer.TypeFLOAT
		result.Value = float64(tokenA.Value.(int64)) + tokenB.Value.(float64)
		return
	}

	result.TokenType = lexer.TypeFLOAT
	result.Value = tokenA.Value.(float64) + float64(tokenB.Value.(int64))
	return
}

// tokenSubtraction will subtract the second token by the first and
// return the result.
func tokenSubtraction(tokenB lexer.Token, tokenA lexer.Token) (result lexer.Token) {

	aIsInt := tokenA.TokenType == lexer.TypeINT
	bIsInt := tokenB.TokenType == lexer.TypeINT

	if aIsInt && bIsInt {
		result.TokenType = lexer.TypeINT
		result.Value = tokenA.Value.(int64) - tokenB.Value.(int64)
		return
	}
	if !aIsInt && !bIsInt {
		result.TokenType = lexer.TypeFLOAT
		result.Value = tokenA.Value.(float64) - tokenB.Value.(float64)
		return
	}
	if aIsInt && !bIsInt {
		result.TokenType = lexer.TypeFLOAT
		result.Value = float64(tokenA.Value.(int64)) - tokenB.Value.(float64)
		return
	}

	result.TokenType = lexer.TypeFLOAT
	result.Value = tokenA.Value.(float64) - float64(tokenB.Value.(int64))
	return
}

// tokenMultiply will subtract the second token by the first and
// return the result.
func tokenMultiply(tokenA lexer.Token, tokenB lexer.Token) (result lexer.Token) {

	aIsInt := tokenA.TokenType == lexer.TypeINT
	bIsInt := tokenB.TokenType == lexer.TypeINT

	if aIsInt && bIsInt {
		result.TokenType = lexer.TypeINT
		result.Value = tokenA.Value.(int64) * tokenB.Value.(int64)
		return
	}
	if !aIsInt && !bIsInt {
		result.TokenType = lexer.TypeFLOAT
		result.Value = tokenA.Value.(float64) * tokenB.Value.(float64)
		return
	}
	if aIsInt && !bIsInt {
		result.TokenType = lexer.TypeFLOAT
		result.Value = float64(tokenA.Value.(int64)) * tokenB.Value.(float64)
		return
	}

	result.TokenType = lexer.TypeFLOAT
	result.Value = tokenA.Value.(float64) * float64(tokenB.Value.(int64))
	return
}

// tokenDivide will subtract the second token by the first and
// return the result.
func tokenDivide(tokenB lexer.Token, tokenA lexer.Token) (result lexer.Token) {

	aIsInt := tokenA.TokenType == lexer.TypeINT
	bIsInt := tokenB.TokenType == lexer.TypeINT

	if aIsInt && bIsInt {
		result.TokenType = lexer.TypeINT
		result.Value = tokenA.Value.(int64) / tokenB.Value.(int64)
		return
	}
	if !aIsInt && !bIsInt {
		result.TokenType = lexer.TypeFLOAT
		result.Value = tokenA.Value.(float64) / tokenB.Value.(float64)
		return
	}
	if aIsInt && !bIsInt {
		result.TokenType = lexer.TypeFLOAT
		result.Value = float64(tokenA.Value.(int64)) / tokenB.Value.(float64)
		return
	}

	result.TokenType = lexer.TypeFLOAT
	result.Value = tokenA.Value.(float64) / float64(tokenB.Value.(int64))
	return
}

// tokenHash will apply sha256 to input string token and return
// the result as a string token.
func tokenHash(tokenA lexer.Token) (result lexer.Token) {
	result.TokenType = lexer.TypeSTRING
	buf := []byte(tokenA.Value.(string))
	hash := sha256.Sum256(buf)
	result.Value = hex.EncodeToString(hash[:])
	return
}
