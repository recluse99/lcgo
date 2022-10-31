/*
 * @lc app=leetcode.cn id=819 lang=golang
 *
 * [819] 最常见的单词
 */
package lcgo

// @lc code=start
func mostCommonWord(paragraph string, banned []string) string {

	//paragraph=strings.Trim(paragraph," ")

	wordMap := map[string]int{}

	concMap := map[byte]int{
		'!': 0,
		'?': 0,
		27:  0,
		'.': 0,
		';': 0,
		' ': 0,
		',': 0,
	}

	ls := len(paragraph)
	var firstIndex int
	for i := 0; i < ls; i++ {

		if _, ok := concMap[paragraph[i]]; !ok {
			firstIndex = i

			// 一个大坑  。 弄混了 for 执行流程
			for ; i < ls; i++ {
				_, o := concMap[paragraph[i]]
				if o {
					break
				}
			}

			var str string = paragraph[firstIndex:i]

			str = convertToLower(str)

			wordMap[str]++

		}

	}

	banMap := map[string]int{}

	for _, value := range banned {
		banMap[value]++
	}

	var count int = 0
	var result string = ""

	for str, strCout := range wordMap {
		//str = convertToLower(str)
		if _, ok := banMap[str]; !ok {
			if strCout > count {
				count = strCout
				result = str
			}
		}

	}

	return result
}

func convertToLower(s string) string {

	var bytes = []byte(s)

	for i := 0; i < len(bytes); i++ {
		if bytes[i] <= 90 {
			bytes[i] += 32
		}
	}

	return string(bytes)

}

// @lc code=end
