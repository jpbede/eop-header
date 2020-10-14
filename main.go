package main

import (
	"bytes"
	"flag"
	"github.com/jhillyerd/enmime"
	"github.com/jpbede/eop-header/eop"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var filePath string

func init() {
	flag.StringVar(&filePath, "file", "", "Path to EML file")
}

func main() {
	flag.Parse()
	if filePath == "" {
		log.Fatal("Missing file path")
		os.Exit(1)
	}

	rawMail, _ := ioutil.ReadFile(filePath)

	br := bytes.NewReader(rawMail)
	env, _ := enmime.ReadEnvelope(br)

	parser := eop.NewParserWithEnvelop(env)
	parser.ParseAndRender()
}
