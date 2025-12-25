package ctxmeta

import "context"

type key string

const (
	JobIDKey         key = "job_id"
	UserIDKey        key = "user_id"
	CorrelationIDKey key = "correlation_id"
)

func WithJobMetadata(
	ctx context.Context,
	jobID, userID, correlationID string,
) context.Context {
	ctx = context.WithValue(ctx, JobIDKey, jobID)
	ctx = context.WithValue(ctx, UserIDKey, userID)
	ctx = context.WithValue(ctx, CorrelationIDKey, correlationID)
	return ctx
}

func FromContext(ctx context.Context) map[string]interface{} {
	meta := map[string]interface{}{}

	if v := ctx.Value(JobIDKey); v != nil {
		meta["job_id"] = v
	}
	if v := ctx.Value(UserIDKey); v != nil {
		meta["user_id"] = v
	}
	if v := ctx.Value(CorrelationIDKey); v != nil {
		meta["correlation_id"] = v
	}

	return meta
}
