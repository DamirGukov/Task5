package generatorFloat

import (
	"fmt"
	"math/rand"
)

type Generator struct{}

func (g Generator) Generate() interface{} {
	return rand.Float64()
}

func (g Generator) GenerateSlice() []interface{} {
	slice := make([]interface{}, rand.Intn(10)+1)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Float64()
	}
	return slice
}

func (g Generator) GenerateWithParam(param int) (interface{}, error) {
	if param <= 0 {
		return nil, fmt.Errorf("invalid parameter for FloatGenerator: %d", param)
	}
	return rand.Float64() * float64(param), nil
}
