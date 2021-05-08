package main

import (
	"fmt"
	"io"
	"log"
)

type Arr2DGenInfo struct {
	VarName string
	Data    [][]byte
}

func genUnicodeEscape(w io.Writer) {
	r := make([][]byte, 32)
	for i := 0; i < 32; i++ {
		r[i] = []byte(fmt.Sprintf(`\u%04x`, i))
	}
	err := tpl.ExecuteTemplate(w, "2darr.gotemplate", Arr2DGenInfo{
		VarName: "jsonUnicodeEscapeChar",
		Data:    r,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
