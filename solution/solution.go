package main

import (
	"fmt"
)

/*
lc 1005
*/

// func main() {
// 	//[[3,9],[7,12],[3,8],[6,8],[9,10],[2,9],[0,9],[3,9],[0,6],[2,8]]
// 	points := [][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}}

// 	fmt.Println(points)

// 	sort.Slice(points, func(i, j int) bool {
// 		if points[i][0] == points[j][0] {
// 			return points[i][1] < points[j][1]
// 		} else {
// 			return points[i][0] < points[j][0]
// 		}
// 	})

// 	fmt.Println(points)

// 	count := findMinArrowShots(points)

// 	fmt.Println(count)

// }

// func main() {

// 	name := "xuqiang"
// 	number, _ := strconv.ParseInt(name, 10, 64)

// 	fmt.Println(number)

// 	for _, value := range name {

// 		fmt.Printf("%T, %v\n", value, value)

// 	}

// }

func main() {

	bytes := []byte{4, 5, 8, 'a'}

	for _, value := range bytes {

		var a string = string(value)

		fmt.Println(a)

		fmt.Printf("%T, %v,%s,%T\n", value, value, a, a)
	}

}
