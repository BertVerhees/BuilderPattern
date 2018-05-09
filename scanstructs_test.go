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
