package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Element struct {
	Type       int
	CursorPos1 int
	CursorPos2 int
}

const vJSON_Object_Begin int = 10
const vJSON_Object_End int = 11
const vJSON_KeyValue_Separator int = 20
const vJSON_String int = 30
const vJSON_Number int = 31
const vJSON_Null int = 32
const vJSON_Bool_True int = 40
const vJSON_Bool_False int = 41
const vJSON_Number_Int int = 50
const vJSON_Number_Float int = 51

/*
func (e *Element) GoString() string {
	switch e.Type {

	}
	return fmt.Sprintf("")
}
*/

func NewElement(Type, pos1, pos2 int) *Element {
	el := new(Element)
	el.Type = Type
	el.CursorPos1 = pos1
	el.CursorPos2 = pos2
	return el
}

var Regexp_Float = regexp.MustCompile(`[-+]?[0-9]+\.[0-9]+`)
var Regexp_Int = regexp.MustCompile(`[-+]?[0-9]+`)

func ProcessNumber(input []*Element, rawData []byte) {
	for i := range input {
		if input[i].Type == vJSON_Number {
			value := rawData[input[i].CursorPos1:input[i].CursorPos2]
			idx := Regexp_Float.FindIndex(value)
			if len(idx) == 2 && idx[0] == 0 && idx[1] == len(value) {
				input[i].Type = vJSON_Number_Float
			} else {
				idx = Regexp_Int.FindIndex(value)
				if len(idx) == 2 && idx[0] == 0 && idx[1] == len(value) {
					input[i].Type = vJSON_Number_Int
				} else {
					log.Fatalln("Invaild Number : ", string(value))
				}
			}
		}
	}
}

func GetKeyType(input []*Element, rawData []byte) []*field {
	var ObjectCount int
	var lenInput int = len(input)

	var KeyPath [][]byte
	push := func(key []byte) {
		KeyPath = append(KeyPath, key)
	}
	pop := func() []byte {
		lenCursorStack := len(KeyPath)
		if lenCursorStack == 0 {
			return nil
		}
		val := KeyPath[lenCursorStack-1]
		KeyPath = KeyPath[:lenCursorStack-1]
		return val
	}
	var output []*field
	for i := 0; i < lenInput; i++ {
		switch input[i].Type {
		case vJSON_Object_Begin:
			ObjectCount++
		case vJSON_Object_End:
			ObjectCount--
			pop()
		}
		if ObjectCount > 0 {
			if input[i].Type == vJSON_KeyValue_Separator {
				key := input[i-1]
				value := input[i+1]
				fmt.Print(strings.Repeat("  ", ObjectCount))
				if value.Type == vJSON_Object_Begin {
					push(rawData[key.CursorPos1:key.CursorPos2])
					fmt.Println(string(rawData[key.CursorPos1:key.CursorPos2]), ":", KeyPath)
					continue
				}

				fmt.Println(string(rawData[key.CursorPos1:key.CursorPos2]), ":", string(rawData[value.CursorPos1:value.CursorPos2]))
				output = append(output, &field{
					Key:       rawData[key.CursorPos1:key.CursorPos2],
					ValueType: value.Type,
					KeyPath:   KeyPath,
				})
				i++
			}
		}
	}
	if ObjectCount != 0 {
		log.Fatalln("GetKeyType: invalid syntax", ObjectCount)
	}
	return output
}
