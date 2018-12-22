# SplashCode W.I.P.
SplashCode is a work in progress

------------

_SplashCode is a very simple forth-like language, inspired by bitcoin Script. Designed to be run in transactions. Eventually this, or something like with will be used with the Splash Blockchain_ 

Features:
- Stack-based language (Last In - First Out).
- Processed from left to right.
- For loops and functions can be disabled. (They should be for transactional use).
- Not Turing Complete.
- Some opperations use implicit variables gathered from the transaction.

</br>

Example:
```
"Starting...", PRINTLN, DROP
0
FUNC, "MyFunction"
    PRINT
    "\,", PRINT, DROP
    DUP, 10, IF
        "Done!"
        FIN
    ENDIF
    1, ADD
    GOTO, "MyFunction"
ENDFUNC
INPUT, TRUE, IF
    GOTO, "MyFunction"
ENDIF
"You chose not to run the function", PRINTLN
```

Outputs:
```
$ go run main.go -file="./lib/test.sc" -input="TRUE"
Starting...
0,1,2,3,4,5,6,7,8,9,10,
[Done 0.09ms]
```

---

## Types
Types are directly added to the stack on execution.
| Type      | Example    |
|:----------|:-----------|
| INT       | `7` `54`   |
| FLOAT     | `10f` `2.5`|
| STRING    | `"hello"` `"one\, two"` |
| BOOLEAN   | `TRUE` `FALSE` |

---

## Arithmetic:
Arithmetic operations pop two previous numbers from stack and pushes a result.
| Word | Opcode | Input | Output | Description |
|:-----|:-------|:-------------|:-------|:------------|
| ADD  | 16     |number, number| number | Pops two numbers from stack, adds them and pushes the result back to the stack.|
| SUB  | 17     |number, number| number | Pops two numbers from stack, subtracts the second from the first and pushes the result back to the stack.|
| MUL  | 18     |number, number| number | Pops two numbers from stack, multiplies them and pushes the result back to the stack.|
| DIV  | 19     |number, number| number | Pops two numbers from stack, divides the second from the first and pushes the result back to the stack. |

---

## Other Key Words:
Key words modify or read or add elements to the stack
| Word | Opcode | Input | Output | Description |
|:-----|:-------|:-------------|:-------|:------------|
| GOTO    | 5      | string    |        | Goto will move the execution to a Supplied Marker or Function. | 
| MARK    | 6      | string    |        | Mark will add a execution Cursor marker to program; Use goto to return the execution to the given marker.|
| IF      | 7      | any, any  |        | Will pop two values from the stack and compare them, if they are equal execution will continue, otherwise program will skip to ENDIF. |
| ENDIF   | 8      |           |        | Marks the end of an if statement. |
| FUNC    | 9      | string    |        | Registers a function; Unless the function is called the execution cursor will skip to ENDFUNC. | 
| ENDFUNC | 10     | string    |        | Marks the end of a function |
| DUP     | 11     | any       | any    | Will Duplicate the last token in the stack. |
| DROP    | 12     | any       |        | Pops a token from the stack and discards it.
| PICK    | 13     | n=integer |   any  | Duplicates an element `n` back in the stack. |
| ROLL    | 14     | n=integer |   any  | Moves an element `n` back in the stack to the top. |
| FIN     | 15     |          |         | Ends the program |
| HASH    | 20     | string   | string  | Pops a string from the stack and applies SHA256 to it and Pushes the result back onto the stack |
