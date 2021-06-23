package pkg

import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const cliChars = "abcdefghijklmnopqrstuvwxyz"

func RandNumber() int {
	return 20 + rand.Intn(20)
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func randCliChar() string {
	return string(cliChars[rand.Intn(len(cliChars))])
}
