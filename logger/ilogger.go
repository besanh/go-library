package logger

import (
	"fmt"
	"reflect"

	"go.uber.org/zap"
)

type ILogger interface {
	Info(msg string, fields ...any)
	Error(msg string, fields ...any)
	Debug(msg string, fields ...any)
	Warn(msg string, fields ...any)
	Fatal(msg string, fields ...any)
}

func (l *impl) Info(msg string, args ...any) {
	l.core.Info(msg, BuildFields(args...)...)
}

func (l *impl) Debug(msg string, args ...any) {
	l.core.Debug(msg, BuildFields(args...)...)
}

func (l *impl) Warn(msg string, args ...any) {
	l.core.Warn(msg, BuildFields(args...)...)
}

func (l *impl) Error(msg string, args ...any) {
	l.core.Error(msg, BuildFields(args...)...)
}

func (l *impl) Fatal(msg string, args ...any) {
	l.core.Fatal(msg, BuildFields(args...)...)
}

func sanitize(v any) any {
	if v == nil {
		return nil
	}

	if errVal, ok := v.(error); ok {
		return errVal.Error()
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface:
		if rv.IsNil() {
			return nil
		}
		return sanitize(rv.Elem().Interface())

	case reflect.Struct:
		out := make(map[string]any)
		rt := rv.Type()
		for i := 0; i < rv.NumField(); i++ {
			field := rt.Field(i)
			if field.PkgPath != "" {
				// Unexported
				continue
			}
			out[field.Name] = sanitize(rv.Field(i).Interface())
		}
		return out

	case reflect.Slice, reflect.Array:
		n := rv.Len()
		out := make([]any, n)
		for i := 0; i < n; i++ {
			out[i] = sanitize(rv.Index(i).Interface())
		}
		return out

	case reflect.Map:
		out := make(map[string]any)
		for _, key := range rv.MapKeys() {
			k := fmt.Sprint(key.Interface())
			out[k] = sanitize(rv.MapIndex(key).Interface())
		}
		return out

	default:
		// everything else (primitives, strings, etc.)
		return v
	}
}

func BuildFields(args ...any) []zap.Field {
	var fields []zap.Field
	argCount := 0
	for i := 0; i < len(args); i++ {
		if key, ok := args[i].(string); ok && i+1 < len(args) {
			fields = append(fields, zap.Any(key, sanitize(args[i+1])))
			i++
		} else {
			v := args[i]
			argCount++
			// Use "arg_N" as fallback key
			key := fmt.Sprintf("arg_%d", argCount)
			fields = append(fields, zap.Any(key, sanitize(v)))
		}
	}
	return fields
}
