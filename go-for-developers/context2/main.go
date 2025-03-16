package main

import (
	"context"
	"fmt"
	"go-for-devs/context2/db"
	"log"
)


func main() {
  ctx := context.Background()
  ctx = db.WithRequestId(ctx)
  ctx = web.WithRequestId(ctx)
  id, err := db.RequestIdFrom(ctx)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("db.RequestID:", id)
}
