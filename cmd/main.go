package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

var fileInput string
var pathOutput string

func init() {
	flag.StringVar(&fileInput, "file-input", "radios.csv", "path to input file")
	flag.StringVar(&pathOutput, "path-output", "./", "path to output file")
}

func handlerString(instr string) string {
	outstr := ""
	for _, v := range instr {
		if v == '"' {
			continue
		} else {
			outstr += string(v)
		}

	}
	return outstr
}

func main() {
	flag.Parse()

	fmt.Println("### Read as reader ###")
	fmt.Printf("filePath %s", fileInput)

	fi, err := os.Open(fileInput)
	if err != nil {
		log.Panicln("Unable to open file:", err)
		os.Exit(1)
	}
	defer fi.Close()

	fo, err := os.Create(pathOutput + "rad.csv")
	if err != nil {
		log.Panicln("Unable to create file:", err)
		os.Exit(1)
	}
	defer fo.Close()

	// Чтение файла с ридером
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(fi)
	for sc.Scan() {
		wr.WriteString(sc.Text())
		//fmt.Println(wr.String())
		//fmt.Println(handlerString(wr.String()))
		fo.WriteString(handlerString(wr.String()) + "\n")
		wr.Reset()
	}

	fmt.Println("Done.")
}
