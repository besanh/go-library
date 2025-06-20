package util

import (
	"bytes"
	"time"
)

type (
	IUtil interface {
		// Excel
		HandleExcelStreamWriter(headers, rows [][]string, mergedColumns ...string) (buffer *bytes.Buffer, err error)
		HandleCSVStreamWriter(fileName string, headers [][]string, rows [][]string, mergedColumns ...string) (buffer *bytes.Buffer, err error)

		ParseAnyToString(value any) (string, error)
		ConvertMillisToTimeString(millis int) string
		MustParseAnyToString(value any) string
		ParseStringToAny(value string, dest any) error
		ParseAnyToAny(value any, dest any) (err error)
		ParseString(value any) string
		GenerateRandomString(n int, letterRunes []rune) string
		ParseStringToTime(t string, timezone ...string) *time.Time
		CheckFromAndToDateValid(from, to time.Time, isAllowZero bool) (isOk bool, err error)
		ParseFloat64(value any) float64
		ParseInt64(value any) int64
		GetEndOfDay(t time.Time) time.Time
		ParseFloat64With2Decimal(value float64) float64
	}
	Util struct{}
)

func NewUtil() IUtil {
	return &Util{}
}
