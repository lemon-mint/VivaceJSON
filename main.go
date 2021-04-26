package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

const JSON_DoubleQuote byte = '"'
const JSON_Backslash byte = '\\'
const JSON_ObjectBegin byte = '{'
const JSON_ObjectEnd byte = '}'
const JSON_ArrayBegin byte = '['
const JSON_ArrayEnd byte = ']'
const JSON_KV byte = ':'
const JSON_Number_Dot byte = '.'
const JSON_Number_Minus byte = '-'
const JSON_Number_0 byte = '0'
const JSON_Number_1 byte = '1'
const JSON_Number_2 byte = '2'
const JSON_Number_3 byte = '3'
const JSON_Number_4 byte = '4'
const JSON_Number_5 byte = '5'
const JSON_Number_6 byte = '6'
const JSON_Number_7 byte = '7'
const JSON_Number_8 byte = '8'
const JSON_Number_9 byte = '9'
const JSON_Bool_T byte = 't'
const JSON_Bool_R byte = 'r'
const JSON_Bool_U byte = 'u'
const JSON_Bool_E byte = 'e'
const JSON_Bool_F byte = 'f'
const JSON_Bool_A byte = 'a'
const JSON_Bool_L byte = 'l'
const JSON_Bool_S byte = 's'
const JSON_Bool_N byte = 'n'

func main() {
	inputFileName := flag.String("if", "", "input file")
	flag.Parse()
	_ = inputFileName
	RefDoc, err := ioutil.ReadFile("ref.json")
	if err != nil {
		log.Fatalln(err)
	}
	els := Parse(RefDoc)
	for i := range els {
		fmt.Println(els[i].CursorPos1, els[i].CursorPos2)
		fmt.Println(string(RefDoc[els[i].CursorPos1:els[i].CursorPos2]))
	}
}

func Parse(input []byte) []*Element {
	var stateObject, stateString int

	var Cursor int
	var lenIN int = len(input)

	var objCount int
	var stringCount int

	var isNumber bool

	var BoolPos1 int
	var BoolLast byte
	var WhatBool int
	var isBool bool

	var CursorStack []int
	var output []*Element
	push := func(pos1 int) {
		CursorStack = append(CursorStack, pos1)
	}
	pop := func() int {
		lenCursorStack := len(CursorStack)
		if lenCursorStack == 0 {
			log.Fatalln("Parsing Error (Stack crash)")
		}
		val := CursorStack[lenCursorStack-1]
		CursorStack = CursorStack[:lenCursorStack-1]
		return val
	}

	for ; Cursor < lenIN; Cursor++ {
		if stateString > 0 {
			if input[Cursor] == JSON_DoubleQuote && !isEscape(input, Cursor) {
				fmt.Print("</String> ")
				stateString--
				output = append(output, NewElement(vJSON_String, pop()+1, Cursor))
				continue
			} else {
				fmt.Print(string(rune(input[Cursor])))
			}
		} else {
			if stateObject > 0 {
				if input[Cursor] == JSON_KV {
					fmt.Print("::")
					output = append(output, NewElement(vJSON_KeyValue_Separator, Cursor, Cursor))
				}
				if input[Cursor] == JSON_ObjectEnd && !isEscape(input, Cursor) {
					stateObject--
					fmt.Print("</Object> ")
					output = append(output, NewElement(vJSON_Object_End, pop()+1, Cursor))
				}
			}
			if isNumberChar(input[Cursor]) {
				if !isNumber {
					fmt.Print(" <Number>")
					isNumber = true
					push(Cursor)
				}
				fmt.Print(string(rune(input[Cursor])))
			} else {
				if isNumber {
					fmt.Print("</Number> ")
					isNumber = false
					output = append(output, NewElement(vJSON_Number, pop(), Cursor))
				}
			}
			if !isBool {
				//Process true type=1
				if input[Cursor] == JSON_Bool_T {
					isBool = true
					BoolLast = JSON_Bool_T
					WhatBool = 1
					BoolPos1 = Cursor
				}
				//Process false type=2
				if input[Cursor] == JSON_Bool_F {
					isBool = true
					BoolLast = JSON_Bool_F
					WhatBool = 2
					BoolPos1 = Cursor
				}
				//Process null type=3
				if input[Cursor] == JSON_Bool_N {
					isBool = true
					BoolLast = JSON_Bool_N
					WhatBool = 3
					BoolPos1 = Cursor
				}
			} else {
				if WhatBool == 1 {
					if BoolLast == JSON_Bool_T {
						if input[Cursor] == JSON_Bool_R {
							BoolLast = JSON_Bool_R
						} else {
							isBool = false
						}
					} else if BoolLast == JSON_Bool_R {
						if input[Cursor] == JSON_Bool_U {
							BoolLast = JSON_Bool_U
						} else {
							isBool = false
						}
					} else if BoolLast == JSON_Bool_U {
						if input[Cursor] == JSON_Bool_E {
							fmt.Print(" <Bool>true</Bool> ")
							output = append(output, NewElement(vJSON_Bool_True, BoolPos1, Cursor))
						} else {
							isBool = false
						}
					}
				} else if WhatBool == 2 {
					if BoolLast == JSON_Bool_F {
						if input[Cursor] == JSON_Bool_A {
							BoolLast = JSON_Bool_A
						} else {
							isBool = false
						}
					} else if BoolLast == JSON_Bool_A {
						if input[Cursor] == JSON_Bool_L {
							BoolLast = JSON_Bool_L
						} else {
							isBool = false
						}
					} else if BoolLast == JSON_Bool_L {
						if input[Cursor] == JSON_Bool_S {
							BoolLast = JSON_Bool_S
						} else {
							isBool = false
						}
					} else if BoolLast == JSON_Bool_S {
						if input[Cursor] == JSON_Bool_E {
							fmt.Print(" <Bool>false</Bool> ")
							output = append(output, NewElement(vJSON_Bool_False, BoolPos1, Cursor))
						} else {
							isBool = false
						}
					}
				} else if WhatBool == 3 {
					if BoolLast == JSON_Bool_N {
						if input[Cursor] == JSON_Bool_U {
							BoolLast = JSON_Bool_U
						} else {
							isBool = false
						}
					} else if BoolLast == JSON_Bool_U {
						if input[Cursor] == JSON_Bool_L {
							BoolLast = JSON_Bool_L
						} else {
							isBool = false
						}
					} else if BoolLast == JSON_Bool_L {
						if input[Cursor] == JSON_Bool_L {
							fmt.Print(" <Null>null</Null> ")
							output = append(output, NewElement(vJSON_Null, BoolPos1, Cursor))
						} else {
							isBool = false
						}
					}
				}
			}
		}
		if input[Cursor] == JSON_ObjectBegin && !isEscape(input, Cursor) {
			stateObject++
			fmt.Print(" <Object>")
			push(Cursor)
			output = append(output, NewElement(vJSON_Object_Begin, Cursor, Cursor))
		}
		if input[Cursor] == JSON_DoubleQuote && !isEscape(input, Cursor) {
			stateString++
			fmt.Print(" <String>")
			push(Cursor)
		}
	}
	fmt.Println()
	fmt.Println(stateObject, stateString, Cursor, lenIN, objCount, stringCount)
	if stateObject|stateString != 0x00 {
		log.Fatalln("JSON: Invaild Syntex")
	}
	return output
}

func isEscape(input []byte, Cursor int) bool {
	if Cursor > 0 && input[Cursor-1] == JSON_Backslash {
		return !isEscape(input, Cursor-1)
	}
	return false
}

func isNumberChar(input byte) bool {
	if input == JSON_Number_0 ||
		input == JSON_Number_1 ||
		input == JSON_Number_2 ||
		input == JSON_Number_3 ||
		input == JSON_Number_4 ||
		input == JSON_Number_5 ||
		input == JSON_Number_6 ||
		input == JSON_Number_7 ||
		input == JSON_Number_8 ||
		input == JSON_Number_9 ||
		input == JSON_Number_Minus ||
		input == JSON_Number_Dot {
		return true
	}
	return false
}
