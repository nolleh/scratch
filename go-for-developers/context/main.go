package main

import (
	"context"
	"fmt"
)

type CtxKeyA string
type CtxKeyB string

func A(ctx context.Context) {
	key := CtxKeyA("request_id")
	ctx = context.WithValue(ctx, key, "123")

	B(ctx)
}

func B(ctx context.Context) {
	key := CtxKeyB("request_id")
	ctx = context.WithValue(ctx, key, "456")
	Logger(ctx)
}

func Logger(ctx context.Context) {
	a := ctx.Value(CtxKeyA("request_id"))
	fmt.Println("A\t", "request_id:", a)

	b := ctx.Value(CtxKeyB("request_id"))
	fmt.Println("B\t", "request_id:", b)
}

func main() {
	ctx := context.Background()

	A(ctx)

}
