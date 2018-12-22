package lexer

import "fmt"

const (
	TypeINT     = iota // An integer
	TypeFLOAT   = iota // A float
	TypeSTRING  = iota // A string
	TypeBOOLEAN = iota // A Boolean
	TypeGOTO    = iota // Goto a Marker
	TypeMARK    = iota // Marks a position for goto Marker "name"
	TypeIF      = iota // If compares the last 2 elementsin stack, if not equal skip to next end if
	TypeENDIF   = iota // Marks end of IF
	TypeFUNC    = iota // Marks a start of a function
	TypeENDFUNC = iota // Marks the end of a function
	TypeDUP     = iota // Duplicates the last element and adds to stack
	TypeDROP    = iota // Deletes last element from stack
	TypePICK    = iota // Duplicates a previous element from stack e.g PICK 5
	TypeROLL    = iota // moves a previous element from stack and places it at the top e.g ROLL 5
	TypeFIN     = iota // Quits program, often displaying the last value in stack
	TypeADD     = iota // Will add the last two elements in stack and add result to stack
	TypeSUB     = iota // Will subtract the last two elements and add result to stack
	TypeMUL     = iota // Will Multiply the last two elements and add result to stack
	TypeDIV     = iota // Will Divide the last two elements and add result to stack
	TypeHASH    = iota // This will sha256 hash the last element into the stack
	TypeINPUT   = iota // This will read a token from input into stack
	TypePRINT   = iota // This will print the last element in stack
	TypePRINTLN = iota // This will print out a line
)

// TokenTypeToString convert an TokenType int to a string
func TokenTypeToString(tokenType int) string {
	switch tokenType {
	case TypeINT:
		return "INT"
	case TypeFLOAT:
		return "FLOAT"
	case TypeSTRING:
		return "STRING"
	case TypeBOOLEAN:
		return "BOOLEAN"
	case TypeGOTO:
		return "GOTO"
	case TypeMARK:
		return "MARK"
	case TypeIF:
		return "IF"
	case TypeENDIF:
		return "ENDIF"
	case TypeFUNC:
		return "FUNC"
	case TypeENDFUNC:
		return "ENDFUNC"
	case TypeDUP:
		return "DUP"
	case TypeDROP:
		return "DROP"
	case TypePICK:
		return "PICK"
	case TypeROLL:
		return "ROLL"
	case TypeFIN:
		return "FIN"
	case TypeADD:
		return "ADD"
	case TypeSUB:
		return "SUB"
	case TypeMUL:
		return "MUL"
	case TypeDIV:
		return "DIV"
	case TypeHASH:
		return "HASH"
	case TypeINPUT:
		return "INPUT"
	case TypePRINT:
		return "PRINT"
	case TypePRINTLN:
		return "PRINTLN"
	default:
		return "UNKNOWN"
	}

	return "UNKNOWN"
}

func (token Token) String() string {
	s := "{"
	s += TokenTypeToString(token.TokenType) + ": "
	s += fmt.Sprintf("%v", token.Value)
	s += "}"
	return s
}
