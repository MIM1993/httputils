package httputils

import (
	"bytes"
	"errors"
	"github.com/lexkong/log"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func HTTPPost(url string, data []byte, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	//设置请求头
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	log.Debug(string(data))

	resp, err := client.Do(req)
	if err != nil {
		log.Warn(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn(err.Error())
		return nil, err
	}

	log.Debug(string(body))
	return body, nil
}

func HTTPPostWithTimeout(url string, data []byte, timeout int, headers map[string]string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	//设置请求头
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	log.Debug(string(data))

	resp, err := client.Do(req)
	if err != nil {
		log.Warn(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn(err.Error())
		return nil, err
	}

	log.Debug(string(body))
	return body, nil
}

func HttpPostRequestForm(url string, params map[string]string, headers map[string]string) ([]byte, error) {
	var ret []byte
	form, contentType, err := CreateMultipartFormBody(params)
	if form == nil {
		return ret, err
	}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, url, form)
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	req.Close = true
	req.Header.Add("Content-Type", contentType)

	resp, err := client.Do(req)
	if err != nil {
		log.Warn(err.Error())
		return ret, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return body, err
	}

	return body, nil
}

// CreateMultipartFormBody CreateMultipartFormBody
func CreateMultipartFormBody(params map[string]string) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)

	// Add fields
	for key, val := range params {
		if PathExists(val) {
			// Open file
			f, err := os.Open(val)
			if err != nil {
				return nil, "", err
			}
			defer f.Close()

			// Add file fields
			fw, err := w.CreateFormFile(key, val)
			if err != nil {
				return nil, "", err
			}
			if _, err = io.Copy(fw, f); err != nil {
				return nil, "", err
			}
		} else {
			// Add string fields
			fw, err := w.CreateFormField(key)
			if err != nil {
				return nil, "", err
			}
			if _, err = fw.Write([]byte(val)); err != nil {
				return nil, "", err
			}
		}
	}
	w.Close()

	return body, w.FormDataContentType(), nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetParamFormData(r *http.Request, key string) ([]byte, error) {
	if key == "" {
		return nil, errors.New("key is nil")
	}
	file, _, err := r.FormFile(key)
	if err != nil {
		log.Warn(err.Error())
		return nil, err
	}
	var data []byte
	data, err = ioutil.ReadAll(file)
	if err != nil {
		log.Warn(err.Error())
		return nil, err
	}
	return data, nil
}

//捕获异常的函数
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				log.Infof("recover:%v", string(bug.Stack()))
				http.Error(w, e.Error(), http.StatusOK)
			}
		}()
		fn(w, r)
	}
}