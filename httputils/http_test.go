package httputils

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttpGet(t *testing.T) {
	var data = map[string]string{
		"1": "2",
		"3": "4",
		"5": "6",
	}
	ret, err := HttpGet("123456", data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ret)
}

func TestHttpGetWithTimeout(t *testing.T) {
	var data = map[string]string{
		"1": "2",
		"3": "4",
		"5": "6",
	}
	ret, err := HttpGetWithTimeout("123456", data, 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ret)
}

func TestAddAndSetHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://10.4.176.79:8910/hello/world", nil)

	//req.Header.Set("hello", "world")
	req.Header.Add("hello", "hello")
	req.Header.Set("hello", "world")

	fmt.Println(req.Header["Hello"][0])
}
