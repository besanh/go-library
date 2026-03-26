package context

import (
	"context"
	"fmt"
)

// contextKey is unexported to prevent collisions with other packages using context
type contextKey string

const userIDKey contextKey = "user_id"

// Put the user ID into the context
func InjectUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// Retrieve the user ID from the context
func ExtractUserID(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(userIDKey).(string)
	if !ok || userID == "" {
		return "", fmt.Errorf("user ID not found in context")
	}
	return userID, nil
}
