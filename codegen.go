package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base32"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
)

var GoNameRule = regexp.MustCompile(`[a-zA-Z]+`)

func GetName(Key []byte, KeyPath [][]byte) string {
	hash := sha256.New()
	hash.Write(Key)
	for i := range KeyPath {
		hash.Write(KeyPath[i])
	}
	UseableName := []byte{}
	parts := GoNameRule.FindAll(Key, -1)
	for i := range parts {
		UseableName = append(UseableName, parts[i]...)
	}
	var stringKey string
	if len(UseableName) > 0 {
		stringKey = strings.ToUpper(string(UseableName[:1])) + string(UseableName[1:]) + strings.ToUpper(base32.StdEncoding.EncodeToString(hash.Sum(nil)[:5]))
	} else {
		stringKey = "Key" + strings.ToUpper(base32.StdEncoding.EncodeToString(hash.Sum(nil)[:5]))
	}
	return stringKey
}

func GenName(fields []*field) {
	for _, f := range fields {
		f.RawName = string(f.Key)
		f.GoName = GetName(f.Key, f.KeyPath)
		switch f.ValueType {
		case vJSON_Bool_True:
			f.GoType = "bool"
		case vJSON_Bool_False:
			f.GoType = "bool"
		case vJSON_Number_Float:
			f.GoType = "float64"
		case vJSON_Number_Int:
			f.GoType = "int"
		case vJSON_String:
			f.GoType = "string"
		case vJSON_Null:
			f.GoType = "*bool"
		}
	}
}

var tpl = template.Must(template.ParseGlob("templates/*"))

func genStruct(StructName string, fields []*field) []byte {
	buf := bytes.NewBuffer(nil)

	tpl.ExecuteTemplate(buf, "genstruct.gotemplate", map[string]interface{}{
		"StructName":   StructName,
		"StructFields": fields,
	})
	return CleanCode(buf.Bytes())
}

func CleanCode(in []byte) []byte {
	buf := bytes.NewBuffer(in)
	outbuf := bytes.NewBuffer(nil)
	gfmt := exec.Command("gofmt")
	gfmt.Stdin = buf
	gfmt.Stdout = outbuf
	gfmt.Start()
	gfmt.Wait()
	return outbuf.Bytes()
}
