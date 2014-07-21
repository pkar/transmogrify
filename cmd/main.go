package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkar/transmogrify"
)

func main() {
	transCMDs := flag.String("cmds", "", "transform commands H,V,int... If provided -trans is ignored")
	transPath := flag.String("trans", "", "path to file with transform commands H,V,int... If provided -cmds is ignored")

	textPath := flag.String("text", "STDIN", "path to text to encode, default is stdin")
	flag.Parse()

	// Set the streaming text reader.
	var textStream io.Reader
	if *textPath == "STDIN" {
		textStream = os.Stdin
	} else {
		var err error
		textStream, err = os.Open(*textPath)
		if err != nil {
			log.Fatalf("%s %s", err, *textPath)
		}
	}

	// Set the initial commands.
	cmds := *transCMDs
	if *transPath != "" {
		cmdFile, err := os.Open(*transPath)
		if err != nil {
			log.Fatalf("%s %s", err, *transPath)
		}
		b, err := ioutil.ReadAll(cmdFile)
		if err != nil {
			log.Fatalf("%s %s", err, *transPath)
		}
		cmds = string(b)
	}

	t := transmogrify.New(textStream, cmds)
	err := t.Print()
	if err != nil {
		log.Fatalf("\n\n%s", err)
	}
}
