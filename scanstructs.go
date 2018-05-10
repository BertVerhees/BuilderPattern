package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

const(
	OPEN_COMMENT = "/*"
	CLOSE_COMMENT = "*/"
	)

func main(){

}


func FindStructNamesInFile(fileName string)([]string, error){
	//readLine("/home/verhees/OpenEhr/development/GO/src/BuilderPattern/scanstructnames_test.go")
	//data, err := ioutil.ReadFile(fileName)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error while reading file: %s\n", err)
	//	return nil, err
	//}
	return nil, nil
}

func readStructs(path string)(error) {
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while reading file: %s\n", err)
		return err
	} else {
		defer inFile.Close()
	}
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	lines = make([]string,0)
	for scanner.Scan() {
		value := scanner.Text()
		addString(value)
	}
	return nil
}

func addString(s string){
	if strings.Index(s, "//")>-1 {
		s = s[:strings.Index(s, "//")]
	}
	s = readLine(s)
	if s != "" {
		lines = append(lines, s)
		if structs==nil{
			structs = make([]*Struct,0)
		}
		if ((strings.HasPrefix(s, "type ")) && ((strings.Index(s," struct " )>-1)||(strings.Index(s," struct{" )>-1))){
			currentStruct = &Struct{
				Name:strings.TrimSpace(s[len("type "):strings.Index(s," struct")]),
			}
			structs = append(structs, currentStruct)
			inStruct = true
		}else if inStruct {
			if s == "}" {
				inStruct = false
			}else{
				s = removeDoubleSpaces(s)
				sl := strings.Split(s," ")
				if len(sl)>1{
					f := &Field{
						Name:sl[0],
					}
					if strings.HasPrefix(sl[1],"map"){
						f.Type = &Map{}
					}else if strings.HasPrefix(sl[1],"[]"){
						f.Type = &Slice{}
					}else {
						f.Type = &NormalType{}
					}
					f.Type.SetTypeName(sl[1])
					if currentStruct.Fields == nil {
						currentStruct.Fields = make([]*Field,0)
					}
					currentStruct.Fields = append(currentStruct.Fields, f)
				}
			}
		}
	}
}

var inComment bool
var inStruct bool
var structs []*Struct
var currentStruct *Struct

func removeDoubleSpaces(line string)string{
	r := ""
	for i,s := range line{
		if s == '\t' {
			s = ' '
		}
		if i < len(line)-1 {
			if ((s == ' ') && (line[i+1]!=' ' && line[i+1]!='\t')) || (s!=' '){
				r = r + string(s)
			}
		}else{
			if s!=' ' && s!='\t' {
				r = r + string(s)
			}
		}
	}
	return r
}

func readLine(line string)string{
	r := ""
	for i,s := range line{
		if s == '/'{
			if i < len(line)-1 {
				if line[i+1] == '*' {
					inComment = true
				}
			}
		}else {
			if inComment && s == '*' {
				if i < len(line)-1 {
					if line[i+1] == '/' {
						inComment = false
					}
					if len(line) > i+2 {
						return readLine(line[i+2:])
					}
				}
			} else {
				if !inComment {
					r = r + string(s)
				}
			}
		}
	}
	return strings.TrimSpace(r)
}

var lines []string

type Struct struct {
	Name string
	Fields []*Field
}

type Field struct{
	Name string
	Type IType
}

type Map struct{
	NormalType
	KeyType IType
	ValueType IType
}

type Slice struct {
	NormalType
	ValueType IType
}

type NormalType struct{
	TypeDescription
}

type TypeDescription struct{
	TypeName string
}

func (t *TypeDescription)SetTypeName(s string) {
	t.TypeName = s
}

func (t TypeDescription)GetTypeName()string{
	return t.TypeName
}

type IType interface{
	SetTypeName(s string)
	GetTypeName()string
}

