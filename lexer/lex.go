package lexer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const numbers = "-0123456789"

// Token - tokens contain type and value data
type Token struct {
	TokenType int
	Value     interface{}
}

// Tokenize - Converts some utf-8 *.sc string to splashcode tokens
func Tokenize(data string, debug bool) []Token {

	if debug {
		fmt.Println("DEBUG:: Tokenizing...")
	}

	//Replace new lines
	data = strings.Replace(data, "\n", ",", -1)

	//Replace escaped Commands
	data = strings.Replace(data, "\\,", "{COMMA}", -1)

	//Seperate targets
	buf := strings.Split(data, ",")

	//Make empty token array
	tokens := make([]Token, len(buf))

	// Loop through all the []string and convert each
	// token to string
	for i := 0; i < len(buf); i++ {

		//Convert string to token
		target := strings.TrimSpace(buf[i])
		token := StringToToken(target)

		if token.Value == nil && token.TokenType == 0 {
			continue
		}

		//Add token to array
		tokens[i] = token

		//print token if debug was enabled
		if debug {
			fmt.Println("   Registering Token::", token)
		}

	}

	return tokens
}

// StringToToken convert a string to a token
// This will panic in the event of an unknown
// token.
func StringToToken(target string) (token Token) {
	if target == "" {
		return
	} else if token.tokenizeString(target) {
	} else if token.tokenizeNumber(target) {
	} else if token.tokenizeBoolean(target) {
	} else if token.tokenizeKeywords(target) {
	} else {
		panic(errors.New("Unknown syntax: `" + target + "`"))
	}
	return
}

func (token *Token) tokenizeString(target string) bool {
	if string(target[0]) == "\"" {
		target = strings.Replace(target, "\"", "", 2)
		target = strings.Replace(target, "{COMMA}", ",", -1)
		token.TokenType = TypeSTRING
		token.Value = target
		return true
	}
	return false
}
func (token *Token) tokenizeNumber(target string) bool {
	if strings.Contains(numbers, string(target[0])) {

		//check if number is float:
		if strings.ContainsAny(target, "f") || strings.ContainsAny(target, ".") {
			token.TokenType = TypeFLOAT
			// convert float to bytes:
			f, err := strconv.ParseFloat(target, 64)
			if err != nil {
				panic(err)
			}
			token.Value = f
		} else {
			// Otherwise number is int
			token.TokenType = TypeINT
			// convert int to bytes:
			i, err := strconv.ParseInt(target, 10, 64)
			if err != nil {
				panic(err)
			}
			token.Value = i
		}

		return true
	}
	return false
}
func (token *Token) tokenizeBoolean(target string) bool {
	if target == "TRUE" || target == "FALSE" {
		token.TokenType = TypeBOOLEAN
		token.Value, _ = strconv.ParseBool(target)
		return true
	}
	return false
}
func (token *Token) tokenizeKeywords(target string) bool {
	switch target {
	case "GOTO":
		token.TokenType = TypeGOTO
		break
	case "MARK":
		token.TokenType = TypeMARK
		break
	case "IF":
		token.TokenType = TypeIF
		break
	case "ENDIF":
		token.TokenType = TypeENDIF
		break
	case "FUNC":
		token.TokenType = TypeFUNC
		break
	case "ENDFUNC":
		token.TokenType = TypeENDFUNC
		break
	case "DUP":
		token.TokenType = TypeDUP
		break
	case "DROP":
		token.TokenType = TypeDROP
		break
	case "PICK":
		token.TokenType = TypePICK
		break
	case "ROLL":
		token.TokenType = TypeROLL
		break
	case "FIN":
		token.TokenType = TypeFIN
		break
	case "ADD":
		token.TokenType = TypeADD
		break
	case "SUB":
		token.TokenType = TypeSUB
		break
	case "MUL":
		token.TokenType = TypeMUL
		break
	case "DIV":
		token.TokenType = TypeDIV
		break
	case "HASH":
		token.TokenType = TypeHASH
		break
	case "INPUT":
		token.TokenType = TypeINPUT
		break
	case "PRINT":
		token.TokenType = TypePRINT
		break
	case "PRINTLN":
		token.TokenType = TypePRINTLN
		break
	case "":
		return false
	default:
		panic(errors.New("Unknown keywork:" + target))
	}
	return true
}
