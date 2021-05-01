package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base32"
	"io"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
)

var GoNameRule = regexp.MustCompile(`[a-zA-Z]+`)
var CodeHeader = []byte("// Code generated by \"VivaceJSON\"; DO NOT EDIT.\n\n")

func GetName(Key []byte, KeyPath [][]byte, AdditionalInfo ...string) string {
	hash := sha256.New()
	hash.Write(Key)
	for i := range KeyPath {
		hash.Write(KeyPath[i])
	}
	for i := range AdditionalInfo {
		hash.Write([]byte(AdditionalInfo[i]))
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

type TempStruct struct {
	StructName string
	RawKeyName []byte
	Locals     []*field
	Children   []*TempStruct
}

func LastDepthEq(a []byte, b [][]byte, depth int) bool {
	if len(b) == 0 {
		return bytes.Equal(a, nil)
	}
	return bytes.Equal(a, b[depth-1])
}

func (ts *TempStruct) Add(keyPaths map[string][]*field, depth int) {
	var found string
	ctr := new(onceCounter)
	for key := range keyPaths {
		keydepth := len(keyPaths[key][0].KeyPath)
		if keydepth == depth {
			if LastDepthEq(ts.RawKeyName, keyPaths[key][0].KeyPath, depth) {
				ts.Locals = keyPaths[key]
				found = key
			}
		} else if keydepth > depth && LastDepthEq(ts.RawKeyName, keyPaths[key][0].KeyPath, depth) {
			if ctr.AssertOnce(keyPaths[key][0].KeyPath[depth]) {
				ts.Children = append(ts.Children, &TempStruct{
					StructName: GetName(keyPaths[key][0].KeyPath[depth], keyPaths[key][0].KeyPath[:depth], "structName"),
					RawKeyName: keyPaths[key][0].KeyPath[depth],
				})
				ts.Locals = append(ts.Locals, &field{
					Key:       keyPaths[key][0].KeyPath[depth],
					ValueType: 1000,
					KeyPath:   keyPaths[key][0].KeyPath[:depth],
					GoName:    GetName(keyPaths[key][0].KeyPath[depth], keyPaths[key][0].KeyPath[:depth]),
					GoType:    GetName(keyPaths[key][0].KeyPath[depth], keyPaths[key][0].KeyPath[:depth], "structName"),
					RawName:   string(keyPaths[key][0].KeyPath[depth]),
				})
			}
		}
	}
	delete(keyPaths, found)
	for i := range ts.Children {
		ts.Children[i].Add(keyPaths, depth+1)
	}
}

func genMixedStruct(StructName string, fields []*field) []byte {
	keyPaths := make(map[string][]*field)
	for _, f := range fields {
		keyPathHash := HashKeyPath(f.KeyPath)
		_, ok := keyPaths[keyPathHash]
		if !ok {
			keyPaths[keyPathHash] = nil
		}
		keyPaths[keyPathHash] = append(keyPaths[keyPathHash], f)
	}
	var RootElement *TempStruct = new(TempStruct)
	var todel string
	ctr := new(onceCounter)
	for key := range keyPaths {
		if len(keyPaths[key][0].KeyPath) == 0 {
			todel = key
			RootElement.Locals = keyPaths[key]
		} else if ctr.AssertOnce(keyPaths[key][0].KeyPath[0]) {
			RootElement.Children = append(RootElement.Children, &TempStruct{
				StructName: GetName(keyPaths[key][0].KeyPath[0], nil, "structName"),
				RawKeyName: keyPaths[key][0].KeyPath[0],
			})
			RootElement.Locals = append(RootElement.Locals, &field{
				Key:       keyPaths[key][0].KeyPath[0],
				ValueType: 1000,
				KeyPath:   nil,
				GoName:    GetName(keyPaths[key][0].KeyPath[0], nil),
				GoType:    GetName(keyPaths[key][0].KeyPath[0], nil, "structName"),
				RawName:   string(keyPaths[key][0].KeyPath[0]),
			})
		}
	}
	delete(keyPaths, todel)
	for i := range RootElement.Children {
		RootElement.Children[i].Add(keyPaths, 1)
	}
	buf := bytes.NewBuffer(nil)
	RootElement.StructName = StructName
	RootElement.Print(buf)
	buf.WriteRune('\n')
	buf.Write(generatePath2StructMap(RootElement, fields))
	return buf.Bytes()
}

func (ts *TempStruct) Print(w io.Writer) {
	w.Write(genStruct(ts.StructName, ts.Locals))
	for i := range ts.Children {
		ts.Children[i].Print(w)
	}
}

func KeyPathInOps(a []byte, b [][]byte) (found bool, idx int) {
	found = false
	for i := range b {
		if bytes.Equal(a, b[i]) {
			found = true
			idx = i
		}
	}
	return
}

func HashKeyPath(keyPath [][]byte) string {
	h := sha256.New()
	for i := range keyPath {
		h.Write(keyPath[i])
	}
	return base32.StdEncoding.EncodeToString(h.Sum(nil))
}

func genStruct(StructName string, fields []*field) []byte {
	buf := bytes.NewBuffer(nil)
	StructTexts := make([]string, len(fields))
	for i := range fields {
		StructTexts[i] = fields[i].String()
	}
	tpl.ExecuteTemplate(buf, "genstruct.gotemplate", map[string]interface{}{
		"StructName":   StructName,
		"StructFields": fields,
		"StructText":   StructTexts,
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
