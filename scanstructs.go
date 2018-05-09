package BuilderPattern

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


func FindStructNamesInFile(fileName string)([]string, error){
	//readLine("/home/verhees/OpenEhr/development/GO/src/BuilderPattern/scanstructnames.go")
	//data, err := ioutil.ReadFile(fileName)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error while reading file: %s\n", err)
	//	return nil, err
	//}
	return nil, nil
}

func readStructLines(path string)([]string, error) {
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while reading file: %s\n", err)
		return nil, err
	} else {
		defer inFile.Close()
	}
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	result = make([]string,0)
	for scanner.Scan() {
		value := scanner.Text()
		addString(value)
	}
	fmt.Println(names)
	return result,nil
}

func addString(s string){
	if strings.Index(s, "//")>-1 {
		s = s[:strings.Index(s, "//")]
	}
	s = readLine(s)
	if s != "" {
		result = append(result, s)
		if names==nil{
			names = make([]string,0)
		}
		if ((strings.HasPrefix(s, "type ")) && ((strings.Index(s," struct " )>-1)||(strings.Index(s," struct{" )>-1))){
			names = append(names, strings.TrimSpace(s[len("type "):strings.Index(s," struct")]))
		}
	}
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

var inComment bool
var result []string
var names []string

