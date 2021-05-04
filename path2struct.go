package main

import (
	"bytes"
)

type DepthKeyPath struct {
	Paths     []PathName
	GoKeyPath string
	GoKeyName string
	Key       string
}

type PathName struct {
	Path string
	Idx  int
}

type Path2StructMap struct {
	ObjectName  string
	IntName     string
	FloatName   string
	BoolName    string
	StringName  string
	IntDepth    []keydepthPath
	FloatDepth  []keydepthPath
	BoolDepth   []keydepthPath
	StringDepth []keydepthPath
}

type keydepthPath struct {
	Depth    int
	KeyPaths []DepthKeyPath
	Varinfo  []struct {
		Bytes  []byte
		GoName string
	}
}

func generatePath2StructMap(RootElement *TempStruct, fields []*field) []byte {
	keyvarOnce := new(onceCounterString)
	fieldTypes := make(map[string][]*field)
	fieldTypes["string"] = nil
	fieldTypes["int"] = nil
	fieldTypes["float64"] = nil
	fieldTypes["bool"] = nil
	for _, v := range fields {
		switch v.GoType {
		case "string":
			fieldTypes[v.GoType] = append(fieldTypes[v.GoType], v)
		case "int":
			fieldTypes[v.GoType] = append(fieldTypes[v.GoType], v)
		case "bool":
			fieldTypes[v.GoType] = append(fieldTypes[v.GoType], v)
		case "float64":
			fieldTypes[v.GoType] = append(fieldTypes[v.GoType], v)
		}
	}
	keypatharr := make(map[string][]keydepthPath)
	keypatharr["string"] = nil
	keypatharr["int"] = nil
	keypatharr["float64"] = nil
	keypatharr["bool"] = nil
	for key := range fieldTypes {
		depthclass := make(map[int][]*field)
		for i := range fieldTypes[key] {
			pathlen := len(fieldTypes[key][i].KeyPath)
			_, ok := depthclass[pathlen]
			if !ok {
				depthclass[pathlen] = nil
			}
			depthclass[pathlen] = append(depthclass[pathlen], fieldTypes[key][i])
		}
		for i := range depthclass {
			kp, varinfo := cvtDepthKeyPath(depthclass[i], keyvarOnce)
			keypatharr[key] = append(keypatharr[key],
				keydepthPath{
					Depth:    i,
					KeyPaths: kp,
					Varinfo:  varinfo,
				})
		}
	}
	geninfo := Path2StructMap{
		ObjectName:  RootElement.StructName,
		IntName:     RootElement.StructName,
		IntDepth:    keypatharr["int"],
		FloatName:   RootElement.StructName,
		FloatDepth:  keypatharr["float64"],
		StringName:  RootElement.StructName,
		StringDepth: keypatharr["string"],
		BoolName:    RootElement.StructName,
		BoolDepth:   keypatharr["bool"],
	}
	//fmt.Println(geninfo)
	var buf *bytes.Buffer = bytes.NewBuffer(nil)
	tpl.ExecuteTemplate(buf, "path2structmap.gotemplate", geninfo)
	return CleanCode(buf.Bytes())
}

func cvtDepthKeyPath(f []*field, c *onceCounterString) ([]DepthKeyPath, []struct {
	Bytes  []byte
	GoName string
}) {
	retstruct := []struct {
		Bytes  []byte
		GoName string
	}{}
	arr := make([]DepthKeyPath, len(f))
	for i := range arr {
		NewKeyPath := []PathName{}
		for j := range f[i].KeyPath {
			VarGoName := "rawKeyPath" + GetName(f[i].KeyPath[j], f[i].KeyPath, "keypathvar", "rawKeyPath")
			NewKeyPath = append(NewKeyPath, PathName{
				Idx:  j,
				Path: VarGoName,
			})
			if c.AssertOnce(VarGoName) {
				retstruct = append(retstruct,
					struct {
						Bytes  []byte
						GoName string
					}{
						Bytes:  f[i].KeyPath[j],
						GoName: VarGoName,
					})
			}
		}
		var gokeypath string
		if len(f[i].KeyPath) > 0 {
			for j := range f[i].KeyPath {
				gokeypath += "." + GetName(f[i].KeyPath[j], f[i].KeyPath[:j])
			}
		}
		VarGoName := "rawKey" + GetName(f[i].Key, f[i].KeyPath, "keyvar", "rawKey")
		if c.AssertOnce(VarGoName) {
			retstruct = append(retstruct,
				struct {
					Bytes  []byte
					GoName string
				}{
					Bytes:  f[i].Key,
					GoName: VarGoName,
				})
		}
		arr[i] = DepthKeyPath{
			Paths:     NewKeyPath,
			GoKeyPath: gokeypath,
			GoKeyName: GetName(f[i].Key, f[i].KeyPath),
			Key:       VarGoName,
		}
	}
	return arr, retstruct
}
