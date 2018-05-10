package BuilderPattern

type NameLower struct{
	bo bool
	st string
	in int
	in8 int8
	in16 int16
	in32 int32
	in64 int64
	uin uint
	uin8 uint8
	uin16 uint16
	uin32 uint32
	uin64 uint64
	by byte
	ru rune
	fl32 float32
	fl64 float64
	co64 complex64
	co128 complex128
}

type NameUpper struct{
	Bo bool
	St string
	In int
	In8 int8
	In16 int16
	In32 int32
	In64 int64
	Uin uint
	Uin8 uint8
	Uin16 uint16
	Uin32 uint32
	Uin64 uint64
	By byte
	Ru rune
	Fl32 float32
	Fl64 float64
	Co64 complex64
	Co128 complex128
}

type NameComplex struct { // a struct
	Na1 NameLower
	Na2 NameUpper
}

/*
type NameFour struct{

}
*/

type NameSlices struct{
	S1 []string
	S2 []int
	Sc1 []NameLower
	Scp1 []*NameLower
}

/*
type NameSix struct{

}
*/

/* abc */ type NameMaps struct{ /*
abc
*/
	M1 map[string]string
	M2 map[int]string
	M3 map[string]NameLower
	M4 map[uint64]*NameComplex
}
/*
edf
*/