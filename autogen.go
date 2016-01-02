package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"text/template"
)

func main() {
	flag.Parse()
	templFile := flag.Arg(0)
	t, err := template.ParseFiles(templFile)
	if err != nil {
		log.Fatal(err)
	}

	jsonFile := flag.Arg(1)
	f, err := os.Open(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var data interface{}
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(os.Stdout, data)
}
