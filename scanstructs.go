/*
 * component:   "BuilderPattern"
 * description: "file to generate builderpattern boilerplate"
 * keywords:    "builderpattern"
 *
 * author:      "Bert Verhees"
 * copyright:   "Copyright (c) 2017 ROSA Software Netherlands"
 * license:     "See notice at bottom of class"
 *
 */

package main

import (
	"os"
	"bufio"
	"strings"
	"log"
	"reflect"
	"fmt"
)

func main(){
	if len(os.Args)>1{
		fileName := os.Args[1]
		readStructs(fileName)
		createBuilders()
	}else{
		fmt.Println("Use like this: BuilderPattern fileWithStrucs.go")
	}
}

func createBuilders() {
	for _, s := range structs {
		filename := s.Name + "_builder.go"
		if file_is_exists(filename) {
			log.Fatalf("File %s already exist, please move out of the way.", filename)
		}
	}
	for _, str := range structs {
		filename := str.Name + "_builder.go"
		writeBuilderFile(filename, str)
	}
}

func writeBuilderFile(s string, str *Struct){
	f, err := os.Create(s)
	if err!=nil{
		log.Fatal("Error while creating file %s", s)
	}
	defer f.Close()
	f.WriteString(packageHeader+"\n")
	if err!=nil{
		log.Fatal("Error while writing in file %s", s)
	}
	f.WriteString("\n")
	writePublicConstructor(f, str)
	writePrivateConstructor(f, str)
	writeWithFunctions(f, str)
	writeBuildFunction(f, str)
}

func isInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}

func writeWithFunctions(f *os.File, str *Struct){
	for _,p := range str.Fields{
		f.WriteString("func (b *" + str.Name + "Builder) With" + p.Name + "(value "+ p.Type.GetTypeName() + ") *" + str.Name + "Builder {" + "\n")
		f.WriteString("	b." + str.VarName + "." + p.Name + " = value" + "\n")
		f.WriteString("	return b" + "\n")
		f.WriteString("}" + "\n")
		f.WriteString("\n")
	}
}

func writeBuildFunction(f *os.File, str *Struct){
	f.WriteString("func (b *" + str.Name + "Builder) Build() *" + str.Name + " {" + "\n")
	f.WriteString("	return new" + str.Name + "(b)" + "\n")
	f.WriteString("}" + "\n")
}

func writePublicConstructor(f *os.File, str *Struct){
	f.WriteString("type " + str.Name + "Builder struct {" + "\n")
	f.WriteString("	" + str.VarName + "	*" + str.Name + "\n")
	f.WriteString("}" + "\n")
	f.WriteString("\n")
	f.WriteString("func New" + str.Name + "Builder()*" + str.Name + "Builder{" + "\n")
	f.WriteString("	b := &" + str.Name + "Builder{}" + "\n")
	f.WriteString("	b." + str.VarName + " = &" + str.Name + "{}\n")
	f.WriteString("	return b" + "\n")
	f.WriteString("}" + "\n")
	f.WriteString("\n")
}

func writePrivateConstructor(f *os.File, str *Struct){
	f.WriteString("func new"+str.Name+"(b *"+str.Name+"Builder) *"+str.Name+" {"+"\n")
	f.WriteString("	s := &"+str.Name+"{}"+"\n")
	for _,p := range str.Fields{
		if isInstanceOf(p.Type, (*NormalType)(nil)) {
			f.WriteString("	s." + p.Name + " = b." + str.VarName + "." + p.Name + "\n")
		}else if  isInstanceOf(p.Type, (*Map)(nil)){
			f.WriteString("	s." + p.Name + " = make("+p.Type.GetTypeName()+")"+ "\n")
			f.WriteString("	if b." + str.VarName + "." + p.Name +" != nil {" + "\n")
			f.WriteString("		for k,v := range b." + str.VarName + "." + p.Name +  " {"+ "\n")
			f.WriteString("			s." + p.Name + "[k] = v"+ "\n")
			f.WriteString("		}"+ "\n")
			f.WriteString("	}"+ "\n")
		}else if  isInstanceOf(p.Type, (*Slice)(nil)){
			f.WriteString("	s." + p.Name + " = make("+p.Type.GetTypeName()+",0)"+ "\n")
			f.WriteString("	if b." + str.VarName + "." + p.Name +" != nil {" + "\n")
			f.WriteString("		for _,v := range b." + str.VarName + "." + p.Name +  " {"+ "\n")
			f.WriteString("			s." + p.Name + " = append(s." + p.Name + ",v)"+ "\n")
			f.WriteString("		}"+ "\n")
			f.WriteString("	}"+ "\n")
		}
	}
	f.WriteString("	return s"+"\n")
	f.WriteString("}" + "\n")
	f.WriteString("\n")
}

func file_is_exists(f string) bool {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func readStructs(path string) {
	inFile, err := os.Open(path)
	if err != nil {
		log.Fatal(os.Stderr, "Error while reading file: %s\n", err)
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
			name := strings.TrimSpace(s[len("type "):strings.Index(s," struct")])
			currentStruct = &Struct{
				Name: name,
				VarName: "__" + strings.ToLower(name),
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
		}else if strings.HasPrefix(s, "package ") {
			packageHeader = s
		}
	}
}

var inComment bool
var inStruct bool
var structs []*Struct
var currentStruct *Struct
var packageHeader string

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
	VarName string
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

/*
 * ***** BEGIN LICENSE BLOCK *****

BSD 2-Clause License

Copyright (c) 2017, Bert Verhees, Rosa Software
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 *
 * ***** END LICENSE BLOCK *****
 */