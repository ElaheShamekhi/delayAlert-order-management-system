package handler

import (
	"delayAlert-order-management-system/internal/locale"
	"delayAlert-order-management-system/internal/serr"
	"delayAlert-order-management-system/internal/validation"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Error struct {
	Message string         `json:"message"`
	Code    serr.ErrorCode `json:"code"`
	TraceID string         `json:"trace_id"`
}

func handleError(ctx *gin.Context, err error) {
	tID := getTraceID(ctx)
	lang := getLanguage(ctx)

	switch err.(type) {
	case *serr.ServiceError:
		var e *serr.ServiceError
		errors.As(err, &e)
		l := log.Error().Str("method", e.Method).Str("code", string(e.ErrorCode)).Str("trace_id", tID)
		if e.Cause != nil {
			l.Err(e.Cause)
		}
		l.Msg(e.Message)
		ctx.AbortWithStatusJSON(
			e.Code,
			Error{Code: e.ErrorCode, Message: locale.Localize(e.Message, lang), TraceID: tID},
		)
		return
	case *validation.Err:
		var e *validation.Err
		errors.As(err, &e)
		l := log.Error().Str("method", e.Method).Str("code", string(e.Code)).Str("trace_id", tID)

		l.Msg(e.Message)
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			Error{Code: serr.ErrorCode(e.Code), Message: locale.Localize(e.Message, lang), TraceID: tID},
		)
		return
	default:
		log.Error().Err(err).Str("trace_id", getTraceID(ctx)).Msg("unknown error")
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			Error{Code: serr.ErrInternal, Message: "internal server error", TraceID: tID},
		)
		return
	}
}

func getTraceID(ctx *gin.Context) string {
	rID, exist := ctx.Get("trace_id")
	if exist {
		switch rID.(type) {
		case string:
			return rID.(string)
		case []byte:
			return string(rID.([]byte))
		}
	}
	return ""
}
