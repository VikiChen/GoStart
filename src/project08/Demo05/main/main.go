package main

type stu struct {
	Name string
	Age int
	Address string
}
//map 
func main() {
	var testMap map[string]stu
	testMap =make(map[string]stu)
	stu1 :=stu{Name:"ana",Address:"sh",Age:22}
	testMap["stu1"]=stu1
}
