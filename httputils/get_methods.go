package httputils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func HttpGet(url string, params map[string]string) ([]byte, error) {
	var pstr = "?"
	for k, v := range params {
		pstr = pstr + fmt.Sprintf("%v=%v&", k, v)
	}
	url = url + pstr
	res, err := http.Get(url[:len(url)-1])
	var ret []byte
	if err != nil {
		log.Fatalf("http get %s failed:%s", url, err.Error())
		return nil, err
	}
	res.Body.Close()

	ret, _ = ioutil.ReadAll(res.Body)

	return ret, nil
}

func HttpGetWithTimeout(url string, params map[string]string, timeout time.Duration) ([]byte, error) {
	var pstr = "?"
	for k, v := range params {
		pstr = pstr + fmt.Sprintf("%v=%v&", k, v)
	}
	url = url + pstr

	req, err := http.NewRequest(http.MethodGet, url[:len(url)-1], nil)
	client := http.Client{
		Timeout: timeout * time.Second,
	}
	res, err := client.Do(req)

	var ret []byte
	if err != nil {
		log.Fatalf("http get %s failed:%s", url, err.Error())
		return nil, err
	}
	res.Body.Close()

	ret, _ = ioutil.ReadAll(res.Body)

	return ret, nil
}

// GetParam 获取querystring 中的value
func GetParam(m url.Values, key string, defVal string) string {
	v, ok := m[key]
	if ok && len(v[0]) > 0 {
		return v[0]
	}
	return defVal
}

// GetParamInt 获取querystring 中的int type value
func GetParamInt(m url.Values, key string, defVal int) int {
	str := GetParam(m, key, "")
	if str == "" {
		return defVal
	}

	val, err := strconv.Atoi(str)
	if err != nil {
		return defVal
	}

	return val
}

// GetParamInt 获取querystring 中的int type value
func GetParamInt64(m url.Values, key string, defVal int64) int64 {
	str := GetParam(m, key, "")
	if str == "" {
		return defVal
	}

	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defVal
	}

	return val
}