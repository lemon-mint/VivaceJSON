package main

import (
	"bytes"
	"fmt"
)

type jsonNode struct {
	KeyPath [][]byte
	Key     []byte
	Type    int

	Children []*jsonNode
}

func (x *jsonNode) Add(KeyPath [][]byte, Key []byte) {

}

type field struct {
	Key       []byte
	ValueType int
	KeyPath   [][]byte
}

func (f *field) String() string {
	buf := bytes.NewBuffer(nil)
	for i := range f.KeyPath {
		if buf.Bytes() != nil {
			fmt.Fprint(buf, " > ")
		}
		fmt.Fprint(buf, string(f.KeyPath[i]))
	}
	return fmt.Sprintf("Field(KeyPath=[%v], Key=[%v])", string(buf.Bytes()), string(f.Key))
}
