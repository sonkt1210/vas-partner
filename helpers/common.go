package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GetCookieValueByName(raw string, name string) (string, error) {
	header := http.Header{}
	header.Add("Cookie", raw)
	request := http.Request{Header: header}
	cookie, err := request.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// ToIntString convert an interface to integer string
func ToIntString(i interface{}) string {
	if i == nil {
		return ""
	}
	switch i.(type) {
	case float32:
		return fmt.Sprintf("%.0f", i)
	case float64:
		return fmt.Sprintf("%.0f", i)
	case int:
		return fmt.Sprintf("%d", i)
	case int8:
		return fmt.Sprintf("%d", i)
	case int16:
		return fmt.Sprintf("%d", i)
	case int32:
		return fmt.Sprintf("%d", i)
	case int64:
		return fmt.Sprintf("%d", i)
	case uint:
		return fmt.Sprintf("%d", i)
	case uint8:
		return fmt.Sprintf("%d", i)
	case uint16:
		return fmt.Sprintf("%d", i)
	case uint32:
		return fmt.Sprintf("%d", i)
	case uint64:
		return fmt.Sprintf("%d", i)
	}
	return fmt.Sprint(i)
}

// ParseInt parse interface to int
func ParseInt(i interface{}) int {
	str := ToIntString(i)
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return int(v)
}

// ParseInt64 parse interface to int64
func ParseInt64(i interface{}) int64 {
	str := ToIntString(i)
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return int64(v)
}

// ParseInt32 parse interface to int32
func ParseInt32(i interface{}) int32 {
	str := ToIntString(i)
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return int32(v)
}

func GetDaysAgo(srcDate string, destDate int32, layoutISO string) string {
	if layoutISO == "" {
		layoutISO = "2006-01-02"
	}
	t, _ := time.Parse(layoutISO, srcDate)

	return t.Add(time.Duration(destDate) * time.Hour * 24).Format(layoutISO)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetTimeNowForDB() string {
	return time.Now().Format(time.RFC3339)
}

func FormatMoney(i interface{}) string {
	n := ParseInt64(i)
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits--
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = '.'
		}
	}
}

func GenerateUserToken(str string) string {
	str = fmt.Sprintf("%s%s", str, time.Now().Format(time.RFC3339))
	h := md5.Sum([]byte(str))
	return hex.EncodeToString(h[:])
}

func GenerateSessionKey(str string) string {
	h := md5.Sum([]byte(str))
	return hex.EncodeToString(h[:])
}

func int64ToIntValue(id int64, tp reflect.Type) reflect.Value {
	var v interface{}
	kind := tp.Kind()

	if kind == reflect.Ptr {
		kind = tp.Elem().Kind()
	}

	switch kind {
	case reflect.Int16:
		temp := int16(id)
		v = &temp
	case reflect.Int32:
		temp := int32(id)
		v = &temp
	case reflect.Int:
		temp := int(id)
		v = &temp
	case reflect.Int64:
		temp := id
		v = &temp
	case reflect.Uint16:
		temp := uint16(id)
		v = &temp
	case reflect.Uint32:
		temp := uint32(id)
		v = &temp
	case reflect.Uint64:
		temp := uint64(id)
		v = &temp
	case reflect.Uint:
		temp := uint(id)
		v = &temp
	}

	if tp.Kind() == reflect.Ptr {
		return reflect.ValueOf(v).Convert(tp)
	}
	return reflect.ValueOf(v).Elem().Convert(tp)
}

func rValue(bean interface{}) reflect.Value {
	return reflect.Indirect(reflect.ValueOf(bean))
}

func rType(bean interface{}) reflect.Type {
	sliceValue := reflect.Indirect(reflect.ValueOf(bean))
	// return reflect.TypeOf(sliceValue.Interface())
	return sliceValue.Type()
}

func structName(v reflect.Type) string {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Name()
}

func sliceEq(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	sort.Sort(sort.StringSlice(left))
	sort.Sort(sort.StringSlice(right))
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func indexName(tableName, idxName string) string {
	return fmt.Sprintf("IDX_%v_%v", tableName, idxName)
}

func eraseAny(value string, strToErase ...string) string {
	if len(strToErase) == 0 {
		return value
	}
	var replaceSeq []string
	for _, s := range strToErase {
		replaceSeq = append(replaceSeq, s, "")
	}

	replacer := strings.NewReplacer(replaceSeq...)

	return replacer.Replace(value)
}

func GetCurrentTime() time.Time {
	return time.Now().Local()
}

func JFilterEntities(str string) string {
	if len(str) == 0 {
		return ""
	}
	rs := strings.Replace(str, "<", "&lt;", -1)
	rs = strings.Replace(rs, ">", "&gt;", -1)
	rs = strings.Replace(rs, "'", "&#39;", -1)
	rs = strings.Replace(rs, "\"", "&quot;", -1)
	rs = strings.Replace(rs, "\\'", "&#39;", -1)
	return rs
}

func JParseFilterEntities(str string) string {
	if len(str) == 0 {
		return ""
	}
	rs := strings.Replace(str, "&lt;", "<", -1)
	rs = strings.Replace(rs, "&gt;", ">", -1)
	rs = strings.Replace(rs, "&#39;", "'", -1)
	rs = strings.Replace(rs, "&quot;", "\"", -1)
	rs = strings.Replace(rs, "&#39;", "'", -1)
	return rs
}

func HReadFile(sfile string) (string, error) {
	content, err := ioutil.ReadFile(sfile)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
