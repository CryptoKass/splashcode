package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"splashcode/executor"
	"splashcode/lexer"
	"splashcode/parser"
	"time"
)

func main() {

	/* Flags */
	filename := flag.String("file", "", "path to a *.sb or .sbc file")
	stackTrace := flag.Bool("trace", false, "display stack trace on run")
	debug := flag.Bool("debug", false, "print additional debuging messages")
	input := flag.String("input", "-1", "set input for program,  the input will be parsed into a Token")
	compile := flag.String("compile", "", "compile to a given a filepath")

	// Parse Flags
	flag.Parse()

	//Read file and Tokenize
	var prog parser.Program
	buf, err := ioutil.ReadFile(*filename)

	// If file is of type '.scb' splash code bytes then unmarshal
	// Otherwise assume file is of type '.sc' then tokenize and parse
	if filepath.Ext(*filename) == ".scb" {
		prog = loadProgFrom(buf)
	} else {
		if err != nil {
			panic(err)
		}
		data := string(buf)

		//Tokenize the data
		tokens := lexer.Tokenize(data, *debug)

		//Parse the tokens into a program
		prog = parser.Parse(tokens)
	}

	//Run or Compile
	started := time.Now()
	if *compile != "" {
		if *debug {
			fmt.Println("\nCompiling to", *compile, "...\n ")
		}
		saveProgTo(prog, *compile)
	} else {
		if *debug {
			fmt.Println("\nRunning file", *filename, "...\n ")
		}
		executor.Run(&prog, lexer.StringToToken(*input), *stackTrace)
	}

	duration := time.Since(started)
	fmt.Println("\n[Done", fmt.Sprintf("%.2fms]", duration.Seconds()*1000))

}

func saveProgTo(prog parser.Program, filePath string) {

	//Encode prog
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	err := enc.Encode(prog)
	if err != nil {
		panic(err)
	}

	//Hex encode
	buf := make([]byte, hex.EncodedLen(len(data.Bytes())))
	hex.Encode(buf, data.Bytes())

	//Write to file
	err = ioutil.WriteFile(filePath, buf, 0644)
	if err != nil {
		panic(err)
	}
}

func loadProgFrom(data []byte) (prog parser.Program) {

	//Hex decode
	buf := make([]byte, hex.DecodedLen(len(data)))
	hex.Decode(buf, data)

	var dataBuf bytes.Buffer
	dataBuf.Write(buf)
	dec := gob.NewDecoder(&dataBuf)
	err := dec.Decode(&prog)
	if err != nil {
		panic(err)
	}
	return
}
