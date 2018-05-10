package BuilderPattern

import (
	"testing"
	"fmt"
	"strings"
)

func TestScanStructs(t *testing.T){
	result, err := readStructLines("scanstructnames.go")
	if err!= nil{
		fmt.Errorf(err.Error())
	}
	r := strings.Join(result," ")
	if "package BuilderPattern type NameOne struct{ } type NameTwo struct{ } type NameThree struct{  } type NameFive struct{ }  type NameSeven struct{  }" != r {
		t.Error("strings should be equal")
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
	s := "aaaaa  aaaaa" //double aaaaa
	r := "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa  aaaaa  aaaaa" //2 double aaaaa
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa		aaaaa" //double tab
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa		aaaaa		aaaaa" //2 double tab
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa 	aaaaa" //aaaaa tab
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa 	aaaaa 	aaaaa" //2 aaaaa tab
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa	 aaaaa" //tab aaaaa
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa	 aaaaa	 aaaaa" //2 tab aaaaa
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa   aaaaa" //three aaaaa
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa   aaaaa   aaaaa" //2 three aaaaa
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa			aaaaa" //three tab
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa			aaaaa			aaaaa" //2 three tab
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa 	 aaaaa" //aaaaa tab aaaaa
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa 	 aaaaa 	 aaaaa" //2 aaaaa tab aaaaa
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
	s = "aaaaa	 	aaaaa" //tab aaaaa tab
	r = "aaaaa aaaaa"
	test(s,r)
	s = "aaaaa	 	aaaaa	 	aaaaa" //2 tab aaaaa tab
	r = "aaaaa aaaaa aaaaa"
	test(s,r)
}
