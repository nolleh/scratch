package main

import (
	"context"
	"fmt"
	"go-for-devs/context2/db"
	"go-for-devs/context2/web"
	"log"
)


func main() {
  ctx := context.Background()
  ctx = db.WithRequestID(ctx)
  ctx = web.WithRequestID(ctx)
  id, err := db.RequestIdFrom(ctx)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("db.RequestID:", id)
}
