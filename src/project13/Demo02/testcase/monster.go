package testcase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Skill string
}

func (this *Monster) Store() bool {
	filePath := "d:/1.txt"
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		return false
	}
return true
}


func (this *Monster) reStore() bool {
	filePath := "d:/1.txt"
	bytes, e := ioutil.ReadFile(filePath)
	if e != nil {
		fmt.Println(e)
		return false
	}
	err:=json.Unmarshal(bytes,this)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}