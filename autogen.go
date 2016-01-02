package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"text/template"
	"strings"
	"unicode"
)

func main() {
	flag.Parse()
	templFile := flag.Arg(0)
	t := template.New(templFile)
	t.Funcs(map[string]interface{}{
		"camel":  camelCase,
		"recv":   firstLetter,
		"type":   toType,
		"suffix": toTypeSuffix,
		"doc":    toGoDoc,
	})
	t, err := t.ParseFiles(templFile)
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

	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}

func camelCase(x string) string {
	return camelCaseImpl(x, '_')
}

func toTypeSuffix(x string) string {
	return camelCaseImpl(x, ' ')
}
func camelCaseImpl(x string, sep rune) string {
	cc := ""
	needUp := true
	for _, c := range x {
		if c == '/' {
			continue
		}
		if c != sep {
			if needUp {
				cc += string(unicode.ToUpper(c))
				needUp = false
			} else {
				cc += string(c)
			}
		} else {
			needUp = true
		}
	}
	return cc
}

func firstLetter(x string) string {
	return string(x[:1])
}

var typeMap = map[string]string{
	"string array":    "[]string",
	"string selector": "string",
}

func toType(x string) string {
	if t, ok := typeMap[x]; ok {
		return t
	}
	return x
}

func toGoDoc(x string) string {
	if strings.HasPrefix(x, "-") {
		return " " + x
	}
	return x
}
