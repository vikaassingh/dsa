package dsa

func solve(A string) string {
	str := []rune(A)
	e := len(str) - 1
	for s := 0; s < e; s++ {
		temp := str[s]
		str[s] = str[e]
		str[e] = temp
		e--
	}
	return string(str)
}
