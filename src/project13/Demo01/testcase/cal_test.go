package main

import "testing"

func TestAddUpper (t *testing.T)  {
	res:=addUpper(10)
	if res!=55{
		t.Fatal("%v,%v",55,res)
	}
	t.Logf("nice")
}
