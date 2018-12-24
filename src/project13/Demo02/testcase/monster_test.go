package testcase

import "testing"

func TestStore(t *testing.T)  {
	monster :=Monster{
		Name:"hon",
		Age:2,
		Skill:"huo",
	}
	res:=monster.Store()
	if !res{
		t.Fatal(true,res)
	}
	t.Logf("success")

}


func TestReStore(t *testing.T)  {
	var monster Monster
	res:=monster.reStore()
	if !res{
		t.Fatal(true,res)
	}
	t.Logf("success")

}