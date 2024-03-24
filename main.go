package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"

	"Task5/generators/bool"
	"Task5/generators/float"
	"Task5/generators/int"
	"Task5/generators/string"
)

type Generator interface {
	Generate() interface{}
	GenerateSlice() []interface{}
	GenerateWithParam(param int) (interface{}, error)
}

func processGenerator(g Generator) {
	value := g.Generate()
	switch v := value.(type) {
	case int:
		fmt.Println("Generated int:", v)
	case string:
		fmt.Println("Generated string:", v)
	case bool:
		fmt.Println("Generated bool:", v)
	case float64:
		fmt.Println("Generated float:", v)
	default:
		fmt.Println("Generated unknown type:", v)
	}

	slice := g.GenerateSlice()
	fmt.Println("Generated slice:", slice)

	value, err := g.GenerateWithParam(6)
	if err != nil {
		zap.L().Error("Failed to generate value with param", zap.Error(err))
		logrus.WithError(err).Error("Failed to generate value with param")
		return
	}
	fmt.Println("Generated value with param:", value)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	intGen := generatorInt.Generator{}
	stringGen := generatorString.Generator{}
	boolGen := generatorBool.Generator{}
	floatGen := generatorFloat.Generator{}

	processGenerator(intGen)
	processGenerator(stringGen)
	processGenerator(boolGen)
	processGenerator(floatGen)
}
