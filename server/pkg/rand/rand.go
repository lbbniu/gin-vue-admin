package rand

import "math/rand"

// String 随机字符串
func String(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[Int(0, len(letters))]
	}
	return string(b)
}

func Int(min, max int) int {
	return min + rand.Intn(max-min)
}
