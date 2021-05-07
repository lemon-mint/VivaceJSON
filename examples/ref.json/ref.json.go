// Code generated by "VivaceJSON"; DO NOT EDIT.


package main

import (
    "bytes"
    "io"
    "fmt"
    "strconv"
)
/*
 <Object> <String>abcde</String> :: <String>fg</String>  <String>lemon</String> :: <String>mint</String>  <String>mint</String> :: <Object> <String>choco</String> :: <String>delicious</String>  <String>Drink</String> :: <String>delicious</String>  <String>foo</String> :: <Object> <String>bar</String> :: <String>foo</String>  <String>floatTest</String> :: <Number>3.14</Number>  <String>intTest</String> :: <Number>1024</Number>  <String>bool</String> :: <Bool>true</Bool>  <String>aa</String> :: <Object> <String>abc</String> :: <Bool>false</Bool>  <String>intt2</String> :: <Number>15</Number> </Object> </Object> </Object>  <String>escape_test</String> :: <String>a\"\\\\\"aa</String>  <String>alpha</String> :: <Object> <String>a</String> :: <String>b</String>  <String>c</String> :: <String>d</String> </Object> </Object> 
0 0 469 469 0 0
  abcde : fg
  lemon : mint
  mint : [[109 105 110 116]]
    choco : delicious
    Drink : delicious
    foo : [[109 105 110 116] [102 111 111]]
      bar : foo
      floatTest : 3.14
      intTest : 1024
      bool : true
      aa : [[109 105 110 116] [102 111 111] [97 97]]
        abc : false
        intt2 : 15
  escape_test : a\"\\\\\"aa
  alpha : [[97 108 112 104 97]]
    a : b
    c : d
Field(KeyPath=[], Key=[abcde], Type=[String])
Field(KeyPath=[], Key=[lemon], Type=[String])
Field(KeyPath=[mint], Key=[choco], Type=[String])
Field(KeyPath=[mint], Key=[Drink], Type=[String])
Field(KeyPath=[mint > foo], Key=[bar], Type=[String])
Field(KeyPath=[mint > foo], Key=[floatTest], Type=[Float])
Field(KeyPath=[mint > foo], Key=[intTest], Type=[Int])
Field(KeyPath=[mint > foo], Key=[bool], Type=[Bool])
Field(KeyPath=[mint > foo > aa], Key=[abc], Type=[Bool])
Field(KeyPath=[mint > foo > aa], Key=[intt2], Type=[Int])
Field(KeyPath=[], Key=[escape_test], Type=[String])
Field(KeyPath=[alpha], Key=[a], Type=[String])
Field(KeyPath=[alpha], Key=[c], Type=[String])
*/



func (obj *ExamplesrefjsonrefjsonOOGNTX2B) pack(w io.Writer) {
	w.Write(JSON_TOKEN_OPEN)

	// Field(KeyPath=[], Key=[alpha], Type=[Object])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{97, 108, 112, 104, 97}) //alpha
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	obj.AlphaR3J7NLLI.pack(w)

	w.Write(JSON_TOKEN_SEP)

	// Field(KeyPath=[], Key=[abcde], Type=[String])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{97, 98, 99, 100, 101}) //abcde
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	w.Write(JSON_TOKEN_STRING)
	escapeStr(w, obj.AbcdeG256KDWZ)
	w.Write(JSON_TOKEN_STRING)

	w.Write(JSON_TOKEN_SEP)

	// Field(KeyPath=[], Key=[lemon], Type=[String])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{108, 101, 109, 111, 110}) //lemon
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	w.Write(JSON_TOKEN_STRING)
	escapeStr(w, obj.Lemon6RSNPVY4)
	w.Write(JSON_TOKEN_STRING)

	w.Write(JSON_TOKEN_SEP)

	// Field(KeyPath=[], Key=[escape_test], Type=[String])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{101, 115, 99, 97, 112, 101, 95, 116, 101, 115, 116}) //escape_test
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	w.Write(JSON_TOKEN_STRING)
	escapeStr(w, obj.EscapetestFCG3PUHP)
	w.Write(JSON_TOKEN_STRING)

	w.Write(JSON_TOKEN_SEP)

	// Field(KeyPath=[], Key=[mint], Type=[Object])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{109, 105, 110, 116}) //mint
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	obj.Mint3RXRPO7M.pack(w)

	w.Write(JSON_TOKEN_CLOSE)
}


type ExamplesrefjsonrefjsonOOGNTX2B struct {
	// Field(KeyPath=[], Key=[alpha], Type=[Object])
	// Field(KeyPath=[], Key=[abcde], Type=[String])
	// Field(KeyPath=[], Key=[lemon], Type=[String])
	// Field(KeyPath=[], Key=[escape_test], Type=[String])
	// Field(KeyPath=[], Key=[mint], Type=[Object])

	AlphaR3J7NLLI      AlphaCMSVNBNU `json:"alpha"`
	AbcdeG256KDWZ      string        `json:"abcde"`
	Lemon6RSNPVY4      string        `json:"lemon"`
	EscapetestFCG3PUHP string        `json:"escape_test"`
	Mint3RXRPO7M       MintDKBZPDCQ  `json:"mint"`
}



func (obj *AlphaCMSVNBNU) pack(w io.Writer) {
	w.Write(JSON_TOKEN_OPEN)

	// Field(KeyPath=[alpha], Key=[a], Type=[String])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{97}) //a
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	w.Write(JSON_TOKEN_STRING)
	escapeStr(w, obj.AFGS2DQZF)
	w.Write(JSON_TOKEN_STRING)

	w.Write(JSON_TOKEN_SEP)

	// Field(KeyPath=[alpha], Key=[c], Type=[String])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{99}) //c
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	w.Write(JSON_TOKEN_STRING)
	escapeStr(w, obj.CE3G3NIX4)
	w.Write(JSON_TOKEN_STRING)

	w.Write(JSON_TOKEN_CLOSE)
}


type AlphaCMSVNBNU struct {
	// Field(KeyPath=[alpha], Key=[a], Type=[String])
	// Field(KeyPath=[alpha], Key=[c], Type=[String])

	AFGS2DQZF string `json:"a"`
	CE3G3NIX4 string `json:"c"`
}



func (obj *MintDKBZPDCQ) pack(w io.Writer) {
	w.Write(JSON_TOKEN_OPEN)

	// Field(KeyPath=[mint], Key=[foo], Type=[Object])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{102, 111, 111}) //foo
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	obj.Foo7WMB5NDO.pack(w)

	w.Write(JSON_TOKEN_SEP)

	// Field(KeyPath=[mint], Key=[choco], Type=[String])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{99, 104, 111, 99, 111}) //choco
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	w.Write(JSON_TOKEN_STRING)
	escapeStr(w, obj.ChocoKGDKVVTW)
	w.Write(JSON_TOKEN_STRING)

	w.Write(JSON_TOKEN_SEP)

	// Field(KeyPath=[mint], Key=[Drink], Type=[String])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{68, 114, 105, 110, 107}) //Drink
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	w.Write(JSON_TOKEN_STRING)
	escapeStr(w, obj.DrinkX4GHDGGP)
	w.Write(JSON_TOKEN_STRING)

	w.Write(JSON_TOKEN_CLOSE)
}


type MintDKBZPDCQ struct {
	// Field(KeyPath=[mint], Key=[foo], Type=[Object])
	// Field(KeyPath=[mint], Key=[choco], Type=[String])
	// Field(KeyPath=[mint], Key=[Drink], Type=[String])

	Foo7WMB5NDO   Foo6V47CVXY `json:"foo"`
	ChocoKGDKVVTW string      `json:"choco"`
	DrinkX4GHDGGP string      `json:"Drink"`
}



func (obj *Foo6V47CVXY) pack(w io.Writer) {
	w.Write(JSON_TOKEN_OPEN)

	// Field(KeyPath=[mint > foo], Key=[bar], Type=[String])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{98, 97, 114}) //bar
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	w.Write(JSON_TOKEN_STRING)
	escapeStr(w, obj.BarAARINFMS)
	w.Write(JSON_TOKEN_STRING)

	w.Write(JSON_TOKEN_SEP)

	// Field(KeyPath=[mint > foo], Key=[floatTest], Type=[Float])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{102, 108, 111, 97, 116, 84, 101, 115, 116}) //floatTest
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	fmt.Fprintf(w, "%f", obj.FloatTest6PH2QOTM)

	w.Write(JSON_TOKEN_SEP)

	// Field(KeyPath=[mint > foo], Key=[intTest], Type=[Int])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{105, 110, 116, 84, 101, 115, 116}) //intTest
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	IntTestUY74FP2E_s := strconv.Itoa(obj.IntTestUY74FP2E)
	writeStr(w, IntTestUY74FP2E_s)

	w.Write(JSON_TOKEN_SEP)

	// Field(KeyPath=[mint > foo], Key=[bool], Type=[Bool])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{98, 111, 111, 108}) //bool
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	if obj.Bool5GMNVIFN {
		w.Write(JSON_BOOL_TRUE)
	} else {
		w.Write(JSON_BOOL_FALSE)
	}

	w.Write(JSON_TOKEN_SEP)

	// Field(KeyPath=[mint > foo], Key=[aa], Type=[Object])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{97, 97}) //aa
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	obj.AaDKIS6D4Y.pack(w)

	w.Write(JSON_TOKEN_CLOSE)
}


type Foo6V47CVXY struct {
	// Field(KeyPath=[mint > foo], Key=[bar], Type=[String])
	// Field(KeyPath=[mint > foo], Key=[floatTest], Type=[Float])
	// Field(KeyPath=[mint > foo], Key=[intTest], Type=[Int])
	// Field(KeyPath=[mint > foo], Key=[bool], Type=[Bool])
	// Field(KeyPath=[mint > foo], Key=[aa], Type=[Object])

	BarAARINFMS       string     `json:"bar"`
	FloatTest6PH2QOTM float64    `json:"floatTest"`
	IntTestUY74FP2E   int        `json:"intTest"`
	Bool5GMNVIFN      bool       `json:"bool"`
	AaDKIS6D4Y        Aa6UOHQXT6 `json:"aa"`
}



func (obj *Aa6UOHQXT6) pack(w io.Writer) {
	w.Write(JSON_TOKEN_OPEN)

	// Field(KeyPath=[mint > foo > aa], Key=[abc], Type=[Bool])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{97, 98, 99}) //abc
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	if obj.AbcV366M2TD {
		w.Write(JSON_BOOL_TRUE)
	} else {
		w.Write(JSON_BOOL_FALSE)
	}

	w.Write(JSON_TOKEN_SEP)

	// Field(KeyPath=[mint > foo > aa], Key=[intt2], Type=[Int])
	w.Write(JSON_TOKEN_STRING)
	w.Write([]byte{105, 110, 116, 116, 50}) //intt2
	w.Write(JSON_TOKEN_STRING)
	w.Write(JSON_TOKEN_KVSEP)

	InttLNPPK4MG_s := strconv.Itoa(obj.InttLNPPK4MG)
	writeStr(w, InttLNPPK4MG_s)

	w.Write(JSON_TOKEN_CLOSE)
}


type Aa6UOHQXT6 struct {
	// Field(KeyPath=[mint > foo > aa], Key=[abc], Type=[Bool])
	// Field(KeyPath=[mint > foo > aa], Key=[intt2], Type=[Int])

	AbcV366M2TD  bool `json:"abc"`
	InttLNPPK4MG int  `json:"intt2"`
}

var JSON_TOKEN_OPEN = []byte("{")
var JSON_TOKEN_CLOSE = []byte("}")
var JSON_TOKEN_SEP = []byte(", ")
var JSON_TOKEN_KVSEP = []byte(": ")
var JSON_TOKEN_STRING = []byte(`"`)

var JSON_BOOL_TRUE = []byte("true")
var JSON_BOOL_FALSE = []byte("true")

func escapeStr(w io.Writer, a string) {
	io.WriteString(w, a)
}

func writeStr(w io.Writer, a string) {
	io.WriteString(w, a)
}

// {ExamplesrefjsonrefjsonOOGNTX2B ExamplesrefjsonrefjsonOOGNTX2B ExamplesrefjsonrefjsonOOGNTX2B ExamplesrefjsonrefjsonOOGNTX2B ExamplesrefjsonrefjsonOOGNTX2B [{2 [{[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO IntTestUY74FP2E rawKeyIntTest6EL7Z26T}] [{[105 110 116 84 101 115 116] rawKeyIntTest6EL7Z26T}]} {3 [{[{rawKeyPathMintRB2TD3EA 0} {rawKeyPathFooNMWRXXHG 1} {rawKeyPathAaC4C5NY6G 2}] .Mint3RXRPO7M.Foo7WMB5NDO.AaDKIS6D4Y InttLNPPK4MG rawKeyIntt6RBFR6FA}] [{[109 105 110 116] rawKeyPathMintRB2TD3EA} {[102 111 111] rawKeyPathFooNMWRXXHG} {[97 97] rawKeyPathAaC4C5NY6G} {[105 110 116 116 50] rawKeyIntt6RBFR6FA}]}] [{2 [{[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO FloatTest6PH2QOTM rawKeyFloatTest6HBUR6Y4}] [{[102 108 111 97 116 84 101 115 116] rawKeyFloatTest6HBUR6Y4}]}] [{2 [{[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO Bool5GMNVIFN rawKeyBoolWFAMSTIB}] [{[98 111 111 108] rawKeyBoolWFAMSTIB}]} {3 [{[{rawKeyPathMintRB2TD3EA 0} {rawKeyPathFooNMWRXXHG 1} {rawKeyPathAaC4C5NY6G 2}] .Mint3RXRPO7M.Foo7WMB5NDO.AaDKIS6D4Y AbcV366M2TD rawKeyAbcKIHD2Z2I}] [{[97 98 99] rawKeyAbcKIHD2Z2I}]}] [{0 [{[]  AbcdeG256KDWZ rawKeyAbcdeBKDYOUYQ} {[]  Lemon6RSNPVY4 rawKeyLemonOOQMMOLL} {[]  EscapetestFCG3PUHP rawKeyEscapetestJJG6R4JP}] [{[97 98 99 100 101] rawKeyAbcdeBKDYOUYQ} {[108 101 109 111 110] rawKeyLemonOOQMMOLL} {[101 115 99 97 112 101 95 116 101 115 116] rawKeyEscapetestJJG6R4JP}]} {1 [{[{rawKeyPathMintVCB2BUTA 0}] .Mint3RXRPO7M ChocoKGDKVVTW rawKeyChocoWRJ3KBCN} {[{rawKeyPathMintVCB2BUTA 0}] .Mint3RXRPO7M DrinkX4GHDGGP rawKeyDrinkN2MSLU22} {[{rawKeyPathAlphaPRCRSEA2 0}] .AlphaR3J7NLLI AFGS2DQZF rawKeyALG3VE3OW} {[{rawKeyPathAlphaPRCRSEA2 0}] .AlphaR3J7NLLI CE3G3NIX4 rawKeyCL6SSH2Y4}] [{[109 105 110 116] rawKeyPathMintVCB2BUTA} {[99 104 111 99 111] rawKeyChocoWRJ3KBCN} {[68 114 105 110 107] rawKeyDrinkN2MSLU22} {[97 108 112 104 97] rawKeyPathAlphaPRCRSEA2} {[97] rawKeyALG3VE3OW} {[99] rawKeyCL6SSH2Y4}]} {2 [{[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO BarAARINFMS rawKeyBarJ5QX3LJC}] [{[109 105 110 116] rawKeyPathMint5D4JJ2ZC} {[102 111 111] rawKeyPathFooGHVBMOPI} {[98 97 114] rawKeyBarJ5QX3LJC}]}]}

// KeyPaths

// IntKeys

// [{[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO IntTestUY74FP2E rawKeyIntTest6EL7Z26T}]

var rawKeyIntTest6EL7Z26T []byte = []byte{105, 110, 116, 84, 101, 115, 116}

// [{[{rawKeyPathMintRB2TD3EA 0} {rawKeyPathFooNMWRXXHG 1} {rawKeyPathAaC4C5NY6G 2}] .Mint3RXRPO7M.Foo7WMB5NDO.AaDKIS6D4Y InttLNPPK4MG rawKeyIntt6RBFR6FA}]

var rawKeyPathMintRB2TD3EA []byte = []byte{109, 105, 110, 116}
var rawKeyPathFooNMWRXXHG []byte = []byte{102, 111, 111}
var rawKeyPathAaC4C5NY6G []byte = []byte{97, 97}
var rawKeyIntt6RBFR6FA []byte = []byte{105, 110, 116, 116, 50}

// FloatKeys

// [{[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO FloatTest6PH2QOTM rawKeyFloatTest6HBUR6Y4}]

var rawKeyFloatTest6HBUR6Y4 []byte = []byte{102, 108, 111, 97, 116, 84, 101, 115, 116}

// StringKeys

// [{[]  AbcdeG256KDWZ rawKeyAbcdeBKDYOUYQ} {[]  Lemon6RSNPVY4 rawKeyLemonOOQMMOLL} {[]  EscapetestFCG3PUHP rawKeyEscapetestJJG6R4JP}]

var rawKeyAbcdeBKDYOUYQ []byte = []byte{97, 98, 99, 100, 101}
var rawKeyLemonOOQMMOLL []byte = []byte{108, 101, 109, 111, 110}
var rawKeyEscapetestJJG6R4JP []byte = []byte{101, 115, 99, 97, 112, 101, 95, 116, 101, 115, 116}

// [{[{rawKeyPathMintVCB2BUTA 0}] .Mint3RXRPO7M ChocoKGDKVVTW rawKeyChocoWRJ3KBCN} {[{rawKeyPathMintVCB2BUTA 0}] .Mint3RXRPO7M DrinkX4GHDGGP rawKeyDrinkN2MSLU22} {[{rawKeyPathAlphaPRCRSEA2 0}] .AlphaR3J7NLLI AFGS2DQZF rawKeyALG3VE3OW} {[{rawKeyPathAlphaPRCRSEA2 0}] .AlphaR3J7NLLI CE3G3NIX4 rawKeyCL6SSH2Y4}]

var rawKeyPathMintVCB2BUTA []byte = []byte{109, 105, 110, 116}
var rawKeyChocoWRJ3KBCN []byte = []byte{99, 104, 111, 99, 111}
var rawKeyDrinkN2MSLU22 []byte = []byte{68, 114, 105, 110, 107}
var rawKeyPathAlphaPRCRSEA2 []byte = []byte{97, 108, 112, 104, 97}
var rawKeyALG3VE3OW []byte = []byte{97}
var rawKeyCL6SSH2Y4 []byte = []byte{99}

// [{[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO BarAARINFMS rawKeyBarJ5QX3LJC}]

var rawKeyPathMint5D4JJ2ZC []byte = []byte{109, 105, 110, 116}
var rawKeyPathFooGHVBMOPI []byte = []byte{102, 111, 111}
var rawKeyBarJ5QX3LJC []byte = []byte{98, 97, 114}

// BoolKeys

// [{[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO Bool5GMNVIFN rawKeyBoolWFAMSTIB}]

var rawKeyBoolWFAMSTIB []byte = []byte{98, 111, 111, 108}

// [{[{rawKeyPathMintRB2TD3EA 0} {rawKeyPathFooNMWRXXHG 1} {rawKeyPathAaC4C5NY6G 2}] .Mint3RXRPO7M.Foo7WMB5NDO.AaDKIS6D4Y AbcV366M2TD rawKeyAbcKIHD2Z2I}]

var rawKeyAbcKIHD2Z2I []byte = []byte{97, 98, 99}

// Set Functions

func (obj *ExamplesrefjsonrefjsonOOGNTX2B) ExamplesrefjsonrefjsonOOGNTX2B_SetInt(KeyPath [][]byte, RawKey []byte, value int) {
	KeyPathlen := len(KeyPath)

	if KeyPathlen == 2 {

		// {[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO IntTestUY74FP2E rawKeyIntTest6EL7Z26T}
		if bytes.Equal(rawKeyPathMint5D4JJ2ZC, KeyPath[0]) && bytes.Equal(rawKeyPathFooGHVBMOPI, KeyPath[1]) && bytes.Equal(rawKeyIntTest6EL7Z26T, RawKey) {
			obj.Mint3RXRPO7M.Foo7WMB5NDO.IntTestUY74FP2E = value
			return
		}

		return
	}

	if KeyPathlen == 3 {

		// {[{rawKeyPathMintRB2TD3EA 0} {rawKeyPathFooNMWRXXHG 1} {rawKeyPathAaC4C5NY6G 2}] .Mint3RXRPO7M.Foo7WMB5NDO.AaDKIS6D4Y InttLNPPK4MG rawKeyIntt6RBFR6FA}
		if bytes.Equal(rawKeyPathMintRB2TD3EA, KeyPath[0]) && bytes.Equal(rawKeyPathFooNMWRXXHG, KeyPath[1]) && bytes.Equal(rawKeyPathAaC4C5NY6G, KeyPath[2]) && bytes.Equal(rawKeyIntt6RBFR6FA, RawKey) {
			obj.Mint3RXRPO7M.Foo7WMB5NDO.AaDKIS6D4Y.InttLNPPK4MG = value
			return
		}

		return
	}

}

func (obj *ExamplesrefjsonrefjsonOOGNTX2B) ExamplesrefjsonrefjsonOOGNTX2B_SetFloat(KeyPath [][]byte, RawKey []byte, value float64) {
	KeyPathlen := len(KeyPath)

	if KeyPathlen == 2 {

		// {[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO FloatTest6PH2QOTM rawKeyFloatTest6HBUR6Y4}
		if bytes.Equal(rawKeyPathMint5D4JJ2ZC, KeyPath[0]) && bytes.Equal(rawKeyPathFooGHVBMOPI, KeyPath[1]) && bytes.Equal(rawKeyFloatTest6HBUR6Y4, RawKey) {
			obj.Mint3RXRPO7M.Foo7WMB5NDO.FloatTest6PH2QOTM = value
			return
		}

		return
	}

}

func (obj *ExamplesrefjsonrefjsonOOGNTX2B) ExamplesrefjsonrefjsonOOGNTX2B_SetString(KeyPath [][]byte, RawKey []byte, value string) {
	KeyPathlen := len(KeyPath)

	if KeyPathlen == 0 {

		// {[]  AbcdeG256KDWZ rawKeyAbcdeBKDYOUYQ}
		if bytes.Equal(rawKeyAbcdeBKDYOUYQ, RawKey) {
			obj.AbcdeG256KDWZ = value
			return
		}

		// {[]  Lemon6RSNPVY4 rawKeyLemonOOQMMOLL}
		if bytes.Equal(rawKeyLemonOOQMMOLL, RawKey) {
			obj.Lemon6RSNPVY4 = value
			return
		}

		// {[]  EscapetestFCG3PUHP rawKeyEscapetestJJG6R4JP}
		if bytes.Equal(rawKeyEscapetestJJG6R4JP, RawKey) {
			obj.EscapetestFCG3PUHP = value
			return
		}

		return
	}

	if KeyPathlen == 1 {

		// {[{rawKeyPathMintVCB2BUTA 0}] .Mint3RXRPO7M ChocoKGDKVVTW rawKeyChocoWRJ3KBCN}
		if bytes.Equal(rawKeyPathMintVCB2BUTA, KeyPath[0]) && bytes.Equal(rawKeyChocoWRJ3KBCN, RawKey) {
			obj.Mint3RXRPO7M.ChocoKGDKVVTW = value
			return
		}

		// {[{rawKeyPathMintVCB2BUTA 0}] .Mint3RXRPO7M DrinkX4GHDGGP rawKeyDrinkN2MSLU22}
		if bytes.Equal(rawKeyPathMintVCB2BUTA, KeyPath[0]) && bytes.Equal(rawKeyDrinkN2MSLU22, RawKey) {
			obj.Mint3RXRPO7M.DrinkX4GHDGGP = value
			return
		}

		// {[{rawKeyPathAlphaPRCRSEA2 0}] .AlphaR3J7NLLI AFGS2DQZF rawKeyALG3VE3OW}
		if bytes.Equal(rawKeyPathAlphaPRCRSEA2, KeyPath[0]) && bytes.Equal(rawKeyALG3VE3OW, RawKey) {
			obj.AlphaR3J7NLLI.AFGS2DQZF = value
			return
		}

		// {[{rawKeyPathAlphaPRCRSEA2 0}] .AlphaR3J7NLLI CE3G3NIX4 rawKeyCL6SSH2Y4}
		if bytes.Equal(rawKeyPathAlphaPRCRSEA2, KeyPath[0]) && bytes.Equal(rawKeyCL6SSH2Y4, RawKey) {
			obj.AlphaR3J7NLLI.CE3G3NIX4 = value
			return
		}

		return
	}

	if KeyPathlen == 2 {

		// {[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO BarAARINFMS rawKeyBarJ5QX3LJC}
		if bytes.Equal(rawKeyPathMint5D4JJ2ZC, KeyPath[0]) && bytes.Equal(rawKeyPathFooGHVBMOPI, KeyPath[1]) && bytes.Equal(rawKeyBarJ5QX3LJC, RawKey) {
			obj.Mint3RXRPO7M.Foo7WMB5NDO.BarAARINFMS = value
			return
		}

		return
	}

}

func (obj *ExamplesrefjsonrefjsonOOGNTX2B) ExamplesrefjsonrefjsonOOGNTX2B_SetBool(KeyPath [][]byte, RawKey []byte, value bool) {
	KeyPathlen := len(KeyPath)

	if KeyPathlen == 2 {

		// {[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO Bool5GMNVIFN rawKeyBoolWFAMSTIB}
		if bytes.Equal(rawKeyPathMint5D4JJ2ZC, KeyPath[0]) && bytes.Equal(rawKeyPathFooGHVBMOPI, KeyPath[1]) && bytes.Equal(rawKeyBoolWFAMSTIB, RawKey) {
			obj.Mint3RXRPO7M.Foo7WMB5NDO.Bool5GMNVIFN = value
			return
		}

		return
	}

	if KeyPathlen == 3 {

		// {[{rawKeyPathMintRB2TD3EA 0} {rawKeyPathFooNMWRXXHG 1} {rawKeyPathAaC4C5NY6G 2}] .Mint3RXRPO7M.Foo7WMB5NDO.AaDKIS6D4Y AbcV366M2TD rawKeyAbcKIHD2Z2I}
		if bytes.Equal(rawKeyPathMintRB2TD3EA, KeyPath[0]) && bytes.Equal(rawKeyPathFooNMWRXXHG, KeyPath[1]) && bytes.Equal(rawKeyPathAaC4C5NY6G, KeyPath[2]) && bytes.Equal(rawKeyAbcKIHD2Z2I, RawKey) {
			obj.Mint3RXRPO7M.Foo7WMB5NDO.AaDKIS6D4Y.AbcV366M2TD = value
			return
		}

		return
	}

}

// Get Functions

func (obj *ExamplesrefjsonrefjsonOOGNTX2B) ExamplesrefjsonrefjsonOOGNTX2B_GetInt(KeyPath [][]byte, RawKey []byte) int {
	KeyPathlen := len(KeyPath)

	if KeyPathlen == 2 {

		// {[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO IntTestUY74FP2E rawKeyIntTest6EL7Z26T}
		if bytes.Equal(rawKeyPathMint5D4JJ2ZC, KeyPath[0]) && bytes.Equal(rawKeyPathFooGHVBMOPI, KeyPath[1]) && bytes.Equal(rawKeyIntTest6EL7Z26T, RawKey) {
			return obj.Mint3RXRPO7M.Foo7WMB5NDO.IntTestUY74FP2E
		}

		return 0
	}

	if KeyPathlen == 3 {

		// {[{rawKeyPathMintRB2TD3EA 0} {rawKeyPathFooNMWRXXHG 1} {rawKeyPathAaC4C5NY6G 2}] .Mint3RXRPO7M.Foo7WMB5NDO.AaDKIS6D4Y InttLNPPK4MG rawKeyIntt6RBFR6FA}
		if bytes.Equal(rawKeyPathMintRB2TD3EA, KeyPath[0]) && bytes.Equal(rawKeyPathFooNMWRXXHG, KeyPath[1]) && bytes.Equal(rawKeyPathAaC4C5NY6G, KeyPath[2]) && bytes.Equal(rawKeyIntt6RBFR6FA, RawKey) {
			return obj.Mint3RXRPO7M.Foo7WMB5NDO.AaDKIS6D4Y.InttLNPPK4MG
		}

		return 0
	}

	return 0
}

func (obj *ExamplesrefjsonrefjsonOOGNTX2B) ExamplesrefjsonrefjsonOOGNTX2B_GetFloat(KeyPath [][]byte, RawKey []byte) float64 {
	KeyPathlen := len(KeyPath)

	if KeyPathlen == 2 {

		// {[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO FloatTest6PH2QOTM rawKeyFloatTest6HBUR6Y4}
		if bytes.Equal(rawKeyPathMint5D4JJ2ZC, KeyPath[0]) && bytes.Equal(rawKeyPathFooGHVBMOPI, KeyPath[1]) && bytes.Equal(rawKeyFloatTest6HBUR6Y4, RawKey) {
			return obj.Mint3RXRPO7M.Foo7WMB5NDO.FloatTest6PH2QOTM
		}

		return 0
	}

	return 0
}

func (obj *ExamplesrefjsonrefjsonOOGNTX2B) ExamplesrefjsonrefjsonOOGNTX2B_GetString(KeyPath [][]byte, RawKey []byte) string {
	KeyPathlen := len(KeyPath)

	if KeyPathlen == 0 {

		// {[]  AbcdeG256KDWZ rawKeyAbcdeBKDYOUYQ}
		if bytes.Equal(rawKeyAbcdeBKDYOUYQ, RawKey) {
			return obj.AbcdeG256KDWZ
		}

		// {[]  Lemon6RSNPVY4 rawKeyLemonOOQMMOLL}
		if bytes.Equal(rawKeyLemonOOQMMOLL, RawKey) {
			return obj.Lemon6RSNPVY4
		}

		// {[]  EscapetestFCG3PUHP rawKeyEscapetestJJG6R4JP}
		if bytes.Equal(rawKeyEscapetestJJG6R4JP, RawKey) {
			return obj.EscapetestFCG3PUHP
		}

		return ""
	}

	if KeyPathlen == 1 {

		// {[{rawKeyPathMintVCB2BUTA 0}] .Mint3RXRPO7M ChocoKGDKVVTW rawKeyChocoWRJ3KBCN}
		if bytes.Equal(rawKeyPathMintVCB2BUTA, KeyPath[0]) && bytes.Equal(rawKeyChocoWRJ3KBCN, RawKey) {
			return obj.Mint3RXRPO7M.ChocoKGDKVVTW
		}

		// {[{rawKeyPathMintVCB2BUTA 0}] .Mint3RXRPO7M DrinkX4GHDGGP rawKeyDrinkN2MSLU22}
		if bytes.Equal(rawKeyPathMintVCB2BUTA, KeyPath[0]) && bytes.Equal(rawKeyDrinkN2MSLU22, RawKey) {
			return obj.Mint3RXRPO7M.DrinkX4GHDGGP
		}

		// {[{rawKeyPathAlphaPRCRSEA2 0}] .AlphaR3J7NLLI AFGS2DQZF rawKeyALG3VE3OW}
		if bytes.Equal(rawKeyPathAlphaPRCRSEA2, KeyPath[0]) && bytes.Equal(rawKeyALG3VE3OW, RawKey) {
			return obj.AlphaR3J7NLLI.AFGS2DQZF
		}

		// {[{rawKeyPathAlphaPRCRSEA2 0}] .AlphaR3J7NLLI CE3G3NIX4 rawKeyCL6SSH2Y4}
		if bytes.Equal(rawKeyPathAlphaPRCRSEA2, KeyPath[0]) && bytes.Equal(rawKeyCL6SSH2Y4, RawKey) {
			return obj.AlphaR3J7NLLI.CE3G3NIX4
		}

		return ""
	}

	if KeyPathlen == 2 {

		// {[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO BarAARINFMS rawKeyBarJ5QX3LJC}
		if bytes.Equal(rawKeyPathMint5D4JJ2ZC, KeyPath[0]) && bytes.Equal(rawKeyPathFooGHVBMOPI, KeyPath[1]) && bytes.Equal(rawKeyBarJ5QX3LJC, RawKey) {
			return obj.Mint3RXRPO7M.Foo7WMB5NDO.BarAARINFMS
		}

		return ""
	}

	return ""
}

func (obj *ExamplesrefjsonrefjsonOOGNTX2B) ExamplesrefjsonrefjsonOOGNTX2B_GetBool(KeyPath [][]byte, RawKey []byte) bool {
	KeyPathlen := len(KeyPath)

	if KeyPathlen == 2 {

		// {[{rawKeyPathMint5D4JJ2ZC 0} {rawKeyPathFooGHVBMOPI 1}] .Mint3RXRPO7M.Foo7WMB5NDO Bool5GMNVIFN rawKeyBoolWFAMSTIB}
		if bytes.Equal(rawKeyPathMint5D4JJ2ZC, KeyPath[0]) && bytes.Equal(rawKeyPathFooGHVBMOPI, KeyPath[1]) && bytes.Equal(rawKeyBoolWFAMSTIB, RawKey) {
			return obj.Mint3RXRPO7M.Foo7WMB5NDO.Bool5GMNVIFN
		}

		return false
	}

	if KeyPathlen == 3 {

		// {[{rawKeyPathMintRB2TD3EA 0} {rawKeyPathFooNMWRXXHG 1} {rawKeyPathAaC4C5NY6G 2}] .Mint3RXRPO7M.Foo7WMB5NDO.AaDKIS6D4Y AbcV366M2TD rawKeyAbcKIHD2Z2I}
		if bytes.Equal(rawKeyPathMintRB2TD3EA, KeyPath[0]) && bytes.Equal(rawKeyPathFooNMWRXXHG, KeyPath[1]) && bytes.Equal(rawKeyPathAaC4C5NY6G, KeyPath[2]) && bytes.Equal(rawKeyAbcKIHD2Z2I, RawKey) {
			return obj.Mint3RXRPO7M.Foo7WMB5NDO.AaDKIS6D4Y.AbcV366M2TD
		}

		return false
	}

	return false
}

