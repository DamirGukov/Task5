package generatorInt

import (
	"fmt"
	"math/rand"
)

type Generator struct{}

func (g Generator) Generate() interface{} {
	return rand.Int()
}

func (g Generator) GenerateSlice() []interface{} {
	slice := make([]interface{}, rand.Intn(10)+1)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Int()
	}
	return slice
}

func (g Generator) GenerateWithParam(param int) (interface{}, error) {
	if param <= 0 {
		return nil, fmt.Errorf("invalid parameter for IntGenerator: %d", param)
	}
	return rand.Intn(param), nil
}
