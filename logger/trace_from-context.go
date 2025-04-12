package logger

import "context"

func TraceIdFromContext(ctx context.Context) string {
	traceID, ok := ctx.Value(TraceIDKey).(string)
	if !ok || traceID == "" {
		return GenerateID()
	}
	return traceID
}
