package main

import (
	"bytes"
	"fmt"
)

type field struct {
	Key       []byte
	ValueType int
	KeyPath   [][]byte
	GoName    string
	GoType    string
	RawName   string
	IsLast    bool
}

func (f *field) String() string {
	buf := bytes.NewBuffer(nil)
	for i := range f.KeyPath {
		if buf.Bytes() != nil {
			fmt.Fprint(buf, " > ")
		}
		fmt.Fprint(buf, string(f.KeyPath[i]))
	}
	var Type string
	switch f.ValueType {
	case vJSON_Bool_True:
		Type = "Bool"
	case vJSON_Bool_False:
		Type = "Bool"
	case vJSON_String:
		Type = "String"
	case vJSON_Number_Float:
		Type = "Float"
	case vJSON_Number_Int:
		Type = "Int"
	case 1000:
		Type = "Object"
	}
	return fmt.Sprintf("Field(KeyPath=[%v], Key=[%v], Type=[%v])", buf.String(), string(f.Key), Type)
}
