package web

import (
	"context"
	"go-for-devs/context2/db"
)

type CtxKey string

const (
  RequestID CtxKey = "request_id"
)

func WithRequestID(ctx context.Context) context.Context {
  ctx = context.WithValue(ctx, RequestID, "456")

  // cxt = context.WithValue(ctx, db.RequestID, "???")

  return ctx
}
