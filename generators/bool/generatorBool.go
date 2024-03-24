package generatorBool

import (
	"fmt"
	"math/rand"
)

type Generator struct{}

func (g Generator) Generate() interface{} {
	return rand.Intn(2) == 1
}

func (g Generator) GenerateSlice() []interface{} {
	slice := make([]interface{}, rand.Intn(10)+1)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(2) == 1
	}
	return slice
}

func (g Generator) GenerateWithParam(param int) (interface{}, error) {
	if param < 0 || param > 1 {
		return nil, fmt.Errorf("invalid parameter for BoolGenerator: %d", param)
	}
	return param == 1, nil
}
