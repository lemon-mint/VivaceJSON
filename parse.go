package main

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
