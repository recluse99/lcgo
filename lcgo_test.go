package lcgo

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestDemo(t *testing.T) {

	fmt.Println("starting test...")

}

func TestDemo01(t *testing.T) {

	bytes := []byte{91, 66, 77, 'a'}

	var result string

	for _, value := range bytes {

		result += strconv.FormatInt(int64(value), 10)

		fmt.Println(result)

		fmt.Printf("%T, %v, %s, %T\n", value, value, result, result)
	}

	fmt.Println(string(bytes))

}

func TestDemo02(t *testing.T) {

	name := "1xuqinag 徐强"

	for _, value := range name {

		fmt.Println(value)

		fmt.Printf("%T,%v\n", value, value)

	}

}

func TestDemo03(t *testing.T) {

	nums := []int{2, 3, 4, 5, 6}

	for _, value := range nums {

		fmt.Printf("%T,%v\n", value, value)
	}

	fmt.Println()

}

//lc 738

func TestDemo04(t *testing.T) {

	var n int = 10

	result := monotoneIncreasingDigits(n)

	fmt.Println(result)

}

func TestDemo05(t *testing.T) {

	bytes := []byte{'0', '1', '2', '3', '4', '5'}

	fmt.Println(bytes)
	fmt.Println(string(bytes))

	for i := 0; i < len(bytes); i++ {
		bytes[i] += 1
	}

	fmt.Println(bytes)

	fmt.Println(string(bytes))

}

// 63
func TestDemo06(t *testing.T) {

	ob := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	result := uniquePathsWithObstacles(ob)

	fmt.Println(result)

}

func TestDemo07(t *testing.T) {

	//go 中字符的底层数uint8 print会输出ascii
	var result byte = 'a'

	fmt.Println(result, string(result))

	var number = 5
	str := strconv.FormatInt(int64(number), 10)
	fmt.Printf("%T,%v\n", number, number)
	fmt.Printf("%T,%v\n", str, str)

}

func Test08(t *testing.T) {
	s := "dasdadad"

	// string以下标取值是uint（byte）
	fmt.Printf("%T,%v\n", s[1], s[1])

	//bytes := []byte(s)

	for _, value := range s {
		// 以 range 迭代  value为 int32  rune

		// 类型推断为int32 rune
		temp := 'd'
		if value == temp {
			fmt.Println("this is d")
			fmt.Printf("%c ,%T\n", value, value)
			fmt.Printf("%c ,%T\n", temp, temp)
		}
	}

}

// test  for  string 中
func TestString(t *testing.T) {
	s := "dasdadad"

	for i := 0; i < len(s); i++ {
		//直接比的话   uint8和int32 都可以直接和 'd' 进行比较
		if s[i] == 'd' {
			fmt.Println("this is d")
		}
		fmt.Printf("%T,%v\n", s[i], s[i])
	}
}

func Test09(t *testing.T) {

	var a int = 2
	var b float64 = 12.20
	//c defalut is float64
	c := 23.45
	fmt.Println(a / 2.0)
	fmt.Println(b / 2)
	fmt.Printf("%T,%v\n", c, c)

}

// test leetcode 914
func Test10(t *testing.T) {

	deck := []int{1, 1, 1, 2, 2, 2, 3, 3}

	result := hasGroupsSizeX(deck)

	fmt.Println(result)

}

func Test11(t *testing.T) {

	// a:=1
	// b:=a

	// //切片只能和nil比较
	// a:=[]int{1,1,1,}
	// b:=a

	//map只能和nil比较
	// a:=map[int]int{1:1,2:4}
	// b:=a
	// fmt.Println(a==b)
}

// test  指针赋值
func Test12(t *testing.T) {

	var a *int

	fmt.Println(a)
	fmt.Println(&a)

	var number int = 1
	a = &number

	fmt.Println(a)

	b := a

	fmt.Println(b)

	fmt.Println(a, b)

}

// test leetcode 989
func Test13(t *testing.T) {

	num := []int{1, 2, 0, 0}

	k := 34

	result := addToArrayForm(num, k)

	fmt.Println(result)

}

func Test14(t *testing.T) {

	a := '1' - 48

	fmt.Printf("%T,%v\n", a, a)

	fmt.Println(a)

}

// 680  leetcode  test
func Test15(t *testing.T) {

	s := "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"

	resulit := validPalindrome(s)

	fmt.Println(resulit)

}

func Test16(t *testing.T) {
	s := "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"

	fmt.Printf("%c,%c\n", s[19], s[80])

	fmt.Printf("%c,%c\n", s[20], s[79])
	fmt.Printf("%c,%c\n", s[21], s[78])
	fmt.Printf("%c,%c\n", s[22], s[77])

}

func Test17(t *testing.T) {

	s := "dasdasd"
	str := strings.Builder{}

	// 返回 len(s)
	l, _ := str.WriteString(s)

	fmt.Println(l)
}

func Test18(t *testing.T) {

	// "Bob hit a ball, the hit BALL flew far after it was hit."
	//  ["hit"]
	input := " ? Bob hit a ball, the hit BALL flew far after it was hit."
	ban := []string{"hit"}

	result := mostCommonWord(input, ban)

	fmt.Println(result)
}

// leetcode  test 17
func Test19(t *testing.T) {

	digits := "2"

	result := letterCombinations(digits)

	fmt.Println(result)

}

// silce  append test
func Test20(t *testing.T) {

	a := []int{3, 4, 5}

	b := a
	fmt.Printf("%T,%v,%p\n", b, b, b)
	fmt.Printf("%T,%v,%p\n", a, a, a)

	b = b[0:1]

	fmt.Printf("%T,%#v,%p\n", a, a, a)
	fmt.Printf("%T,%+v,%p\n", b, b, b)

}

// struct   ==  test   //直接覆盖
func Test21(t *testing.T) {

	a := struct {
		age  int
		name string
	}{
		age:  42,
		name: "dsadads",
	}

	b := struct {
		age  int
		name string
	}{
		age:  33,
		name: "xxxx",
	}

	fmt.Printf("%v,%p,\n", a, &a)

	fmt.Printf("%v,%p,\n", b, &b)

	b = a

	fmt.Printf("%v,%p,\n", b, &b)

}

// leetcode 204  test
func Test22(t *testing.T) {

	n := 444

	result := countPrimes(n)

	fmt.Println(result)

}
