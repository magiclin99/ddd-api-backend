// Package handler provides standard flow of handling http API
package handler

import (
	"context"
	"dddapib/internal/domain/model/aperr"
	"dddapib/internal/infrastructure/transport/http/dto"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handlerFuncWithoutPayload func(ctx context.Context, g *gin.Context) (any, error)
type handlerFuncWithPayload[T any] func(ctx context.Context, g *gin.Context, payload *T) (any, error)

func JsonWithoutPayload(handler handlerFuncWithoutPayload) func(g *gin.Context) {
	return func(g *gin.Context) {

		// http API context start from here
		ctx := newContext()

		rsp, err := handler(ctx, g)
		if err != nil {
			replyError(g, err)
			return
		}
		replyOk(g, rsp)
	}
}

func Json[T any](handler handlerFuncWithPayload[T]) func(g *gin.Context) {
	return func(g *gin.Context) {

		// http API context start from here
		ctx := newContext()

		var payload *T
		if err := g.ShouldBindJSON(&payload); err != nil {
			replyError(g, aperr.InvalidRequest(err.Error()))
			return
		}

		rsp, err := handler(ctx, g, payload)
		if err != nil {
			replyError(g, err)
			return
		}
		replyOk(g, rsp)
	}
}

func replyOk(g *gin.Context, data any) {
	g.JSON(200, dto.NewApiOK(data))
}

func replyError(g *gin.Context, err error) {
	var e *aperr.Error
	switch {
	case errors.As(err, &e):
		g.JSON(400, dto.NewApiError(e.Code, e.Message))
	default:
		e := aperr.UnexpectedServerError()
		g.JSON(500, dto.NewApiError(e.Code, e.Message))
	}

}

func newContext() context.Context {
	// return simple value context with correlation id
	// in real world, should extract correlation id from http header (or generate a new one)
	return context.WithValue(context.Background(), "cid", uuid.New().String())
}
