package main

// Необходимо написать функцию, которая находит самый длинный общий префикс среди элементов массива строк.
// Если такого префикса нет, программа должна вернуть пустую строку.

func main() {
	strs := []string{"flower", "flow", "float"}
	println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	l := 0

	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}

	for _, last := range strs[0] {
		for _, s := range strs {
			if rune(s[l]) != last {
				return strs[0][:l]
			}
		}

		l++
	}
	return strs[0]
}
