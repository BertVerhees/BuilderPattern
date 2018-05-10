package main

import (
	"testing"
	"fmt"
	"strings"
)

func TestScanStructs(t *testing.T){
	test := func(str *Struct, name string, fields []Field){
		if str.Name != name {
			t.Errorf("The structname should be %s, but is %s", name, str.Name)
		}
		for i,f := range str.Fields{
			if f.Type.GetTypeName() != fields[i].Type.GetTypeName(){
				t.Errorf("The struct-fields in struct %s are not as expected", name)
			}
		}
	}
	err := readStructs("scanstructnames_test.go")
	if err!= nil{
		fmt.Errorf(err.Error())
	}
	if len(structs)!=5{
		t.Error("There must be five structs discovered")
	}
	fields := make([]Field,0)
	fields = append(fields, Field{"bo", &NormalType{TypeDescription{TypeName: "bool"}}})
	fields = append(fields, Field{"st", &NormalType{TypeDescription{TypeName: "string"}}})
	fields = append(fields, Field{"in", &NormalType{TypeDescription{TypeName: "int"}}})
	fields = append(fields, Field{"in8", &NormalType{TypeDescription{TypeName: "int8"}}})
	fields = append(fields, Field{"in16", &NormalType{TypeDescription{TypeName: "int16"}}})
	fields = append(fields, Field{"in32", &NormalType{TypeDescription{TypeName: "int32"}}})
	fields = append(fields, Field{"in64", &NormalType{TypeDescription{TypeName: "int64"}}})
	fields = append(fields, Field{"uin", &NormalType{TypeDescription{TypeName: "uint"}}})
	fields = append(fields, Field{"uin8", &NormalType{TypeDescription{TypeName: "uint8"}}})
	fields = append(fields, Field{"uin16", &NormalType{TypeDescription{TypeName: "uint16"}}})
	fields = append(fields, Field{"uin32", &NormalType{TypeDescription{TypeName: "uint32"}}})
	fields = append(fields, Field{"uin64", &NormalType{TypeDescription{TypeName: "uint64"}}})
	fields = append(fields, Field{"by", &NormalType{TypeDescription{TypeName: "byte"}}})
	fields = append(fields, Field{"ru", &NormalType{TypeDescription{TypeName: "rune"}}})
	fields = append(fields, Field{"fl32", &NormalType{TypeDescription{TypeName: "float32"}}})
	fields = append(fields, Field{"fl64", &NormalType{TypeDescription{TypeName: "float64"}}})
	fields = append(fields, Field{"co64", &NormalType{TypeDescription{TypeName: "complex64"}}})
	fields = append(fields, Field{"co128", &NormalType{TypeDescription{TypeName: "complex128"}}})
	test(structs[0], "NameLower", fields)
	for i,f := range fields{
		fields[i] = Field{strings.Title(f.Name), &NormalType{TypeDescription{TypeName: f.Type.GetTypeName()}}}
	}
	test(structs[1], "NameUpper", fields)
	fields = make([]Field,0)
	fields = append(fields, Field{"Na1", &NormalType{TypeDescription{"NameLower"}}})
	fields = append(fields, Field{"Na2", &NormalType{TypeDescription{"NameUpper"}}})
	test(structs[2], "NameComplex", fields)
	fields = make([]Field,0)
	fields = append(fields, Field{"S1", &NormalType{TypeDescription{"[]string"}}})
	fields = append(fields, Field{"S2", &NormalType{TypeDescription{"[]int"}}})
	fields = append(fields, Field{"Sc1", &NormalType{TypeDescription{"[]NameLower"}}})
	fields = append(fields, Field{"Scp1", &NormalType{TypeDescription{"[]*NameLower"}}})
	test(structs[3], "NameSlices", fields)
	fields = make([]Field,0)
	fields = append(fields, Field{"M1", &NormalType{TypeDescription{"map[string]string"}}})
	fields = append(fields, Field{"M2", &NormalType{TypeDescription{"map[int]string"}}})
	fields = append(fields, Field{"M3", &NormalType{TypeDescription{"map[string]NameLower"}}})
	fields = append(fields, Field{"M4", &NormalType{TypeDescription{"map[uint64]*NameComplex"}}})
	test(structs[4], "NameMaps", fields)
}



func TestRemoveDoubleSpaces(t *testing.T){
	test := func(s, r string){
		p := removeDoubleSpaces(s)
		if p!=r{
			p = strings.Replace(p," ","_", -1)
			p = strings.Replace(p,"\t","-", -1)
			r = strings.Replace(r," ","_", -1)
			r = strings.Replace(r,"\t","-", -1)
			s = strings.Replace(s," ","_", -1)
			s = strings.Replace(s,"\t","-", -1)
			t.Errorf("*%s* (p)should be equal with *%s* (r), input was *%s* (s)", p, r, s)
		}
	}
	s := "aaaaa  aaaaa" //double space
	r := "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa  aaaaa  aaaaa" //2 double space
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa		aaaaa" //double tab
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa		aaaaa		aaaaa" //2 double tab
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa 	aaaaa" //space tab
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa 	aaaaa 	aaaaa" //2 space tab
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa	 aaaaa" //tab space
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa	 aaaaa	 aaaaa" //2 tab space
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa   aaaaa" //three space
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa   aaaaa   aaaaa" //2 three space
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa			aaaaa" //three tab
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa			aaaaa			aaaaa" //2 three tab
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa 	 aaaaa" //space tab space
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa 	 aaaaa 	 aaaaa" //2 space tab space
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa	 	aaaaa" //tab space tab
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa	 	aaaaa	 	aaaaa" //2 tab space tab
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
}
