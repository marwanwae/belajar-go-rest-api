package test

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(t *testing.T) {

	ctxA := context.Background()
	ctxB := context.WithValue(ctxA, "name", "marwan")
	ctxC := context.WithValue(ctxB, "agama", "islam")
	fmt.Println("ctxA", ctxA)
	fmt.Println("ctxB", ctxB)
	fmt.Println("ctxB", ctxC.Value("name"))
}
