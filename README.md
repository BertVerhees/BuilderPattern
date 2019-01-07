# BuilderPattern

This tool adds a builderpattern to a struct, it saves a lot of boilerplate, I created it for my own convenience, but you can use it or alter it as you like.
But in that case, remove my name from the top of the files.

Use: `go get -u build github.com/bertverhees/builderpattern`

Go to the directory:
`cd $GOPATH/src/github.com/bertverhees/builderpattern`

Run:
`go install`

Use the tool as:
`builderpattern file_with_structs.go`

It will create, per struct a builderpattern file in the same directory, it calls that file structname_builder.go.

If that file already exists then it will refuse to overwrite it and stops execution with an errormessage.
In that case, it will not construct any builder-file.

For example, you have a go-file with this struct in it

    type NameSlices struct{
        S1 []string
        S2 []int
        Sc1 []NameLower
        Scp1 []*NameLower
    }`
    
The tool will create the file NameSlices_builder.go, and that will look like below.
You can then use it in code like:
    
    func aFunc()*NameSlices {
    	nb := NewNameSlicesBuilder()
    	return nb.WithS1(one).WithS2(another).Build()
    }
============================================

    package main
  
    

    type NameSlicesBuilder struct {
	    __nameslices	*NameSlices
    }

    func NewNameSlicesBuilder()*NameSlicesBuilder{
	    b := &NameSlicesBuilder{}
	    b.__nameslices = &NameSlices{}
	    return b
    }

    func newNameSlices(b *NameSlicesBuilder) *NameSlices {
	    s := &NameSlices{}
	    s.S1 = make([]string,0)
	    if b.__nameslices.S1 != nil {
		    for _,v := range b.__nameslices.S1 {
			    s.S1 = append(s.S1,v)
		    }
	    }
	    s.S2 = make([]int,0)
	    if b.__nameslices.S2 != nil {
		    for _,v := range b.__nameslices.S2 {
			    s.S2 = append(s.S2,v)
		    }
	    }
	    s.Sc1 = make([]NameLower,0)
	    if b.__nameslices.Sc1 != nil {
		    for _,v := range b.__nameslices.Sc1 {
			    s.Sc1 = append(s.Sc1,v)
		    }
	    }
	    s.Scp1 = make([]*NameLower,0)
	    if b.__nameslices.Scp1 != nil {
		    for _,v := range b.__nameslices.Scp1 {
			    s.Scp1 = append(s.Scp1,v)
		    }
	    }
	    return s
    }
    
    func (b *NameSlicesBuilder) WithS1(value []string) *NameSlicesBuilder {
	    b.__nameslices.S1 = value
	    return b
    }
    
    func (b *NameSlicesBuilder) WithS2(value []int) *NameSlicesBuilder {
	    b.__nameslices.S2 = value
	    return b
    }
    
    func (b *NameSlicesBuilder) WithSc1(value []NameLower) *NameSlicesBuilder {
	    b.__nameslices.Sc1 = value
	    return b
    }
    
    func (b *NameSlicesBuilder) WithScp1(value []*NameLower) *NameSlicesBuilder {
	    b.__nameslices.Scp1 = value
	    return b
    }
    
    func (b *NameSlicesBuilder) Build() *NameSlices {
	    return newNameSlices(b)
    }
    
        
    