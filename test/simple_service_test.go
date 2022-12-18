package test

import (
	"fmt"
	"go-restful-api/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService(true)

	fmt.Println(err)
	fmt.Println(simpleService)
}
