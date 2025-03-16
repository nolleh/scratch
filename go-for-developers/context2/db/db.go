package db

import (
	"context"
	"fmt"
)

type CtxKey string

const (
  requestID CtxKey = "request_id"
)

func RequestIdFrom(ctx context.Context) (string, error) {
  // get the request_id from the context
  s, ok := ctx.Value(requestID).(string)
  if !ok {
    return "", fmt.Errorf("request_id not found in context")
  }

  return s, nil
}

func WithRequestID(ctx context.Context) context.Context {
  ctx = context.WithValue(ctx, requestID, "123")
  return ctx
}
