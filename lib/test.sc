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