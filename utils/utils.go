package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

var (
	// ShortenToMax is a setting for shortening strings.
	ShortenToMax = 400
)

// NewUUID generates a random UUID according to RFC 4122
func NewUUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return ""
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
func NewShotUUID() string {
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s

}

//ConvertToISO ... Convert Time to ISO
func ConvertToISO(normalTime string) string {
	//2019-06-28T095706Z - GOT
	//2018-08-13T07:02:44.039Z - Expected

	layout := "2006-01-02T15:04:05.000Z"
	time, err := time.Parse(layout, normalTime)

	if err != nil {
		return ""
	}
	return time.String()
}

//ConvertDateToISODate ... Parse and ISO formatted date
func ConvertDateToISODate(isoFormat string) (string, error) {
	val, err := time.Parse("2006-01-02T150405Z", isoFormat)
	if err == nil {
		return val.Format(time.RFC3339), nil
	}
	return "", errors.New("date_parse_error")
}

//ConvertTimeFormat ... Converting time format
func ConvertTimeFormat(normalTime string, fromLayout string, toLayout string) string {
	t, err := time.Parse(fromLayout, normalTime)
	if err != nil {
		return ""
	}
	return t.Format(toLayout)
}

//ConvertDateFormat ... Converting time to ISO
func ConvertDateFormat(normalDate string, fromLayout string, toLayout string) string {
	t, err := time.Parse(fromLayout, normalDate)
	if err != nil {
		return ""
	}
	return t.Format(toLayout)
}

//ConvertTimeToUTC ... Converting Date time to UTC
func ConvertTimeToUTC(timezone string, dateTime string) string {
	t, err := time.Parse(time.RFC3339, dateTime)
	if err == nil {
		// "Asia/Dubai"
		location, err1 := time.LoadLocation(timezone)
		if err1 == nil {
			gst := t.In(location)
			return gst.Format(time.RFC3339)
		}
		fmt.Print("ConvertTimeToUTC error2: ", err1)
	}
	fmt.Print("ConvertTimeToUTC error1: ", err)
	return ""
}

//GetDateTimeFromUTCString ... Get date time from string
func GetDateTimeFromUTCString(timeString string) (string, string) {
	timeStamp, err := time.Parse(time.RFC3339, timeString)
	fmt.Print("ERROR", err)
	timevalue := fmt.Sprintf("%02d:%02d:%02d", timeStamp.Hour(), timeStamp.Minute(), timeStamp.Second())

	dateValue := strings.Split(timeString, "T")[0]
	return dateValue, timevalue
}

// Short shortens reader content safely and returns tring.
func Short(reader io.Reader) string {
	b, _ := ioutil.ReadAll(reader)
	return Shortb(b)
}

// Shortb shortens reader content safely and returns tring.
func Shortb(b []byte) string {
	str := string(b)

	if len(str) <= ShortenToMax {
		return str
	}
	return str[:ShortenToMax]
}

// Errorfy creates a new errorf function to use in a certain package/service.
func Errorfy(pkgName string) func(msg string, args ...interface{}) error {
	return func(msg string, args ...interface{}) error {
		return errors.New(pkgName + ": " + fmt.Sprintf(msg, args...))
	}
}

// RoundFloat64 returns a float64 of x with prec digits of precision after the decimal point.
func RoundFloat64(x float64, prec int) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}

	sign := 1.0
	if x < 0 {
		sign = -1
		x *= -1
	}

	var rounder float64
	pow := math.Pow(10, float64(prec))
	xpow := x * pow
	_, f := math.Modf(xpow)

	if f >= 0.5 {
		rounder = math.Ceil(xpow)
	} else {
		rounder = math.Floor(xpow)
	}

	return rounder / pow * sign
}

// ToCents convers float amount to cents.
func ToCents(amount float64) int64 {
	return int64(amount * 100)
}

// ToCents32 convers float amount to cents.
func ToCents32(amount float32) int64 {
	return int64(amount * 100)
}

// ToCentsFloat64 convers float amount to cents.
func ToCentsFloat64(amount float64) int64 {
	return int64(amount * 100)
}
