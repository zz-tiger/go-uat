package main

import (
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"uat/entity"
)

var maps map[int]int = make(map[int]int)
var resData = 0
var lock sync.Mutex
var group = sync.WaitGroup{}
var index = 10
var channel chan int = make(chan int)

func codeTest() {

	var index int8 = -1

	for i := 1; ; i++ {

		val := index >> i

		fmt.Println(val)

		if i == 5 {
			goto label
		}
	}

label:
	{
		fmt.Println("")
	}

}

type student struct {
	age  int
	name string
}

type stuSlice []student

func (s stuSlice) Len() int {
	return len(s)
}
func (s stuSlice) Less(i, j int) bool {
	return s[i].age > s[j].age
}

func (s stuSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func init() {
	fmt.Println("init......")
	fmt.Println("开启初始化工作......")

	var students = stuSlice{
		{age: 20, name: "tanxiaomi"},
		{age: 22, name: "tandami"},
		{age: 23, name: "tanmi"},
	}
	fmt.Println(students)
	fmt.Println(&students)
	sort.Sort(students)
	fmt.Println(students)
	fmt.Println(&students)
}

func main() {

	codeTest()

	var a int = 645
	var ptr *int = &a
	var scanln string
	fmt.Scanln(&scanln)
	fmt.Println("输入的内容是: ", scanln)

	fmt.Scanf("", &scanln)
	fmt.Println("格式化后的内容: ", scanln)

	fmt.Println(ptr)
	fmt.Println(*ptr)

	fmt.Println("hello world")

	runtime.GOMAXPROCS(8)

	//group.Add(10)

	var num1 string = "65"
	fmt.Println(strconv.ParseInt(num1, 10, 0))
	//strconv.Itoa()

	// base 代表数值进制 10->十进制 2->二进制
	formatUint := strconv.FormatUint(9981, 10)
	fmt.Println(formatUint)
	var strInt int = 64
	sprintln := fmt.Sprintln(strInt)
	fmt.Println(sprintln)
	for i := 1; i <= 10; i++ {
		go goruntineDemo(i)
	}
	//time.Sleep(time.Second * 1)
	//group.Wait()
	for i := range maps {
		fmt.Println(i, "->", maps[i])
	}
	//index <- channel
	//fmt.Println(<-channel)
	go catChannel()
	<-channel

	fmt.Println("block.......")
	go func() {
		channel <- 20
	}()
	for i := range maps {
		fmt.Println(i, "->", maps[i])
	}

}

//func checkone(i int) int {
//	if i > 1 {
//		return checkone(i - 1)
//	}
//	return i
//}

func goruntineDemo(i int) {

	result := 1
	key := i
	for i > 0 {
		result = result * i
		i--
	}
	lock.Lock()
	maps[key] = result
	lock.Unlock()
	resData = result
	//group.Done()

	if key == index {

		channel <- 10
	}
}

func catChannel() {

	var cats chan interface{}
	cats = make(chan interface{}, 2)

	cats <- entity.Cat{"Tom", 5}
	cats <- entity.Cat{"Jerry", 3}

	cats2 := <-cats
	// 类型强转
	cat := cats2.(entity.Cat)

	fmt.Println(cat.Name)

	var f1 = 1
	var f2 = 1

	fmt.Println(sums(float64(f1), float64(f2)))
}

func sums(f1 float64, f2 float64) float64 {
	return f1 + f2
}
