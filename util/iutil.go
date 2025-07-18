package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/dromara/carbon/v2"
)

type IUtil interface {
	// Excel
	HandleExcelStreamWriter(headers, rows [][]string, mergedColumns ...string) (buffer *bytes.Buffer, err error)
	HandleCSVStreamWriter(fileName string, headers [][]string, rows [][]string, mergedColumns ...string) (buffer *bytes.Buffer, err error)

	// Util
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

	// Oauth2
	GenerateCodeVerifier() (string, error)
	GenerateCodeChallenge(verifier string) string
	Encrypt(plaintext string) (string, error)
	Decrypt(encryptedText string) (string, error)
}

var LETTER_RUNES = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

var NUMBER_RUNES = []rune("1234567890")

func (i *Util) ConvertToBytes(message any) ([]byte, error) {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}
	return messageBytes, err
}

func (i *Util) ConvertMillisToTimeString(millis int) string {
	duration := time.Duration(millis) * time.Millisecond
	t := time.Time{}.Add(duration)

	// Format the string as hh:mm:ss.milliseconds
	// Use .Format to format up to seconds, and manually append milliseconds
	return fmt.Sprintf("%s.%03d", t.Format("15:04:05"), millis%1000)
}

func (i *Util) ParseAnyToString(value any) (string, error) {
	ref := reflect.ValueOf(value)
	if ref.Kind() == reflect.String {
		return value.(string), nil
	} else if InArray(ref.Kind(), []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64}) {
		return fmt.Sprintf("%d", value), nil
	} else if InArray(ref.Kind(), []reflect.Kind{reflect.Float32, reflect.Float64}) {
		return fmt.Sprintf("%.3f", value), nil
	} else if ref.Kind() == reflect.Bool {
		return fmt.Sprintf("%t", value), nil
	} else if ref.Kind() == reflect.Slice {
		return fmt.Sprintf("%v", value), nil
	}
	bytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (i *Util) MustParseAnyToString(value any) string {
	str, err := i.ParseAnyToString(value)
	if err != nil {
		return ""
	}
	return str
}

func (i *Util) ParseStringToAny(value string, dest any) error {
	if err := json.Unmarshal([]byte(value), dest); err != nil {
		return err
	}
	return nil
}

func (i *Util) ParseAnyToAny(value any, dest any) (err error) {
	ref := reflect.ValueOf(value)
	var bytes []byte
	if ref.Kind() == reflect.String {
		bytes = []byte(value.(string))
	} else {
		bytes, err = json.Marshal(value)
		if err != nil {
			return err
		}
	}
	if err := json.Unmarshal(bytes, dest); err != nil {
		return err
	}
	return nil
}

func (i *Util) ParseString(value any) string {
	str, ok := value.(string)
	if !ok {
		return str
	}
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Trim(str, "\r\n")
	str = strings.TrimSpace(str)
	return str
}

func (i *Util) GenerateRandomString(n int, letterRunes []rune) string {
	if len(letterRunes) < 1 {
		letterRunes = LETTER_RUNES
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (i *Util) ParseStringToTime(t string, timezone ...string) *time.Time {
	if len(t) == 0 {
		return nil
	}
	c := carbon.Parse(t, timezone...)
	if c.Error != nil {
		return nil
	}
	tPtr := c.StdTime()
	return &tPtr
}

func (i *Util) CheckFromAndToDateValid(from, to time.Time, isAllowZero bool) (isOk bool, err error) {
	if !isAllowZero {
		if from.IsZero() {
			return false, errors.New("is zero")
		}
		if to.IsZero() {
			return false, errors.New("is zero")
		}
	} else if from.After(to) {
		return false, errors.New("from date must be before to date")
	}
	isOk = true
	return
}

func (i *Util) ParseFloat64(value any) float64 {
	if value == nil {
		return 0
	}
	// convert to string
	str := i.MustParseAnyToString(value)
	// convert to float
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return result
}

func (i *Util) ParseInt64(value any) int64 {
	if value == nil {
		return 0
	}
	// convert to string
	str := i.MustParseAnyToString(value)
	// convert to int
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return result
}

// func (i *Util)  to get end of day
func (i *Util) GetEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

// func (i *Util)  to parse float64 only with 2 decimal
func (i *Util) ParseFloat64With2Decimal(value float64) float64 {
	return math.Round(value*100) / 100
}
