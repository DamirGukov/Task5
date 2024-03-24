package generatorString

import (
	"fmt"
	"math/rand"
)

type Generator struct{}

func (g Generator) Generate() interface{} {
	return randomString(rand.Intn(10) + 1)
}

func (g Generator) GenerateSlice() []interface{} {
	slice := make([]interface{}, rand.Intn(10)+1)
	for i := 0; i < len(slice); i++ {
		slice[i] = randomString(rand.Intn(10) + 1)
	}
	return slice
}

func (g Generator) GenerateWithParam(param int) (interface{}, error) {
	if param <= 0 {
		return nil, fmt.Errorf("invalid parameter for StringGenerator: %d", param)
	}
	return randomString(param), nil
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
