package lcgo

import (
	"fmt"
	"strconv"
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
	str := strconv.FormatInt(int64(number),10)
	fmt.Printf("%T,%v\n", number, number)
	fmt.Printf("%T,%v\n", str, str)

}
