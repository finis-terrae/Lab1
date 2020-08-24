package rot

func customAbs(x int) int {
	if x >= 0 {
		return x % 26
	}
	return 26 + x
}
func rot(r rune, n int) string {
	if r >= 'a' && r <= 'z' {
		return string(customAbs(int(r)+n-97) + 97)
	}
	if r >= 'A' && r <= 'Z' {
		return string(customAbs(int(r)+n-65) + 65)
	}
	return string(r)
}

func Rot(str string, n int) string {
	var strNew string
	for _, r := range str {
		strNew = strNew + rot(r, n)
	}
	return strNew
}
