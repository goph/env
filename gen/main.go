package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var types = []string{
	"bool",
	"float32",
	"float64",
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
}

var formatFuncs = map[string]string{
	"bool":    "strconv.FormatBool(bool(*RECEIVER))",
	"float32": "strconv.FormatFloat(float64(*RECEIVER), 'g', -1, 32)",
	"float64": "strconv.FormatFloat(float64(*RECEIVER), 'g', -1, 64)",
	"int":     "strconv.FormatInt(int64(*RECEIVER), 10)",
	"int8":    "strconv.FormatInt(int64(*RECEIVER), 10)",
	"int16":   "strconv.FormatInt(int64(*RECEIVER), 10)",
	"int32":   "strconv.FormatInt(int64(*RECEIVER), 10)",
	"int64":   "strconv.FormatInt(int64(*RECEIVER), 10)",
	"uint":    "strconv.FormatUint(uint64(*RECEIVER), 10)",
	"uint8":   "strconv.FormatUint(uint64(*RECEIVER), 10)",
	"uint16":  "strconv.FormatUint(uint64(*RECEIVER), 10)",
	"uint32":  "strconv.FormatUint(uint64(*RECEIVER), 10)",
	"uint64":  "strconv.FormatUint(uint64(*RECEIVER), 10)",
}

var parseFuncs = map[string]string{
	"bool":    "strconv.ParseBool(val)",
	"float32": "strconv.ParseFloat(val, 32)",
	"float64": "strconv.ParseFloat(val, 64)",
	"int":     "strconv.ParseInt(val, 0, 64)",
	"int8":    "strconv.ParseInt(val, 0, 8)",
	"int16":   "strconv.ParseInt(val, 0, 16)",
	"int32":   "strconv.ParseInt(val, 0, 32)",
	"int64":   "strconv.ParseInt(val, 0, 64)",
	"uint":    "strconv.ParseInt(val, 0, 64)",
	"uint8":   "strconv.ParseInt(val, 0, 8)",
	"uint16":  "strconv.ParseInt(val, 0, 16)",
	"uint32":  "strconv.ParseInt(val, 0, 32)",
	"uint64":  "strconv.ParseInt(val, 0, 64)",
}

func main() {
	tpl, err := template.ParseFiles("value.tpl")
	if err != nil {
		panic(err)
	}

	for _, t := range types {
		file, err := os.Create(fmt.Sprintf("value_%s.go", t))
		if err != nil {
			panic(err)
		}

		receiver := t[:1]

		if strings.HasPrefix(t, "uint") {
			receiver = "i"
		}

		article := "a"

		if strings.HasPrefix(t, "int") {
			article = "an"
		}

		formatFunc := strings.Replace(formatFuncs[t], "RECEIVER", receiver, 1)

		data := map[string]string{
			"Type":       t,
			"TypeName":   strings.Title(t),
			"Article":    article,
			"Receiver":   receiver,
			"FormatFunc": formatFunc,
			"ParseFunc":  parseFuncs[t],
		}

		err = tpl.Execute(file, data)
		if err != nil {
			file.Close()

			panic(err)
		}

		file.Close()
	}
}
