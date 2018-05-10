package BuilderPattern

import (
	"testing"
	"fmt"
	"strings"
)

func TestScanStructs(t *testing.T){
	err := readStructLines("scanstructnames.go")
	if err!= nil{
		fmt.Errorf(err.Error())
	}
	for _,s := range structs{
		fmt.Println(s.Name)
		for _,f := range s.Fields{
			fmt.Println("\t"+f.Name+" "+f.Type.GetTypeName())
		}
	}
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
