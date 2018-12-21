package main

import (
	"fmt"
	"sort"
	"math/rand"
)

type Hero struct {
	Name string
	Age int
}

type Interface interface {

}

//heroSice 切片类型
type HeroSlice []Hero

//实现 Interface 接口
func (heroSlice HeroSlice)  Len() int  {
	return len(heroSlice)
}
//less 觉得按照什么标准排序
func (heroSlice HeroSlice)  Less(i,j int) bool  {
	return heroSlice[i].Age>heroSlice[j].Age
}

//swaps 交换元素
func (heroSlice HeroSlice)  Swap(i,j int) {
	temp:=heroSlice[i]
	heroSlice[i]=heroSlice[j]
	heroSlice[j]=temp
}
//接口
func main() {

	//初始化切片
	 intSlice :=[]int{0,-1,10,8,20}
	 //对intSlice排序
	 //方法一
	 fmt.Println(intSlice)
	 sort.Ints(intSlice)
	 fmt.Println(intSlice)
	 //方法二

	var slices HeroSlice
	for i:=0;i<10;i++{
		hero := Hero{}
		hero.Name=fmt.Sprintf("英雄%d",rand.Intn(100))
		hero.Age=rand.Intn(100)
		slices = append(slices, hero)
	}

	for _,v := range slices{
		fmt.Println(v)
	}
	sort.Sort(slices)
	fmt.Println("---------------")
	for _,v := range slices{
		fmt.Println(v)
	}
}
