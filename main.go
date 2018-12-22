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
