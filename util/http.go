package util

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/andrezhz/go-tools/context"
	terrors "github.com/andrezhz/go-tools/errors"
	"github.com/andrezhz/go-tools/log"
	"github.com/parnurzeal/gorequest"
)

func HttpGet(url string, params interface{}, headers map[string]string, timeout time.Duration) (string, error) {
	t := time.Now()
	request := gorequest.New().Get(url).Timeout(timeout)
	for key, value := range headers {
		request.Set(key, value)
	}
	resp, body, errs := request.Query(params).End()
	if errs != nil {
		log.LOGERROR("Get request url(%v),params(%v),errs(%v)", url, Struct2String(params), errs)
		return "", errors.New(Struct2String(errs))
	}
	log.LOGFLOW("latency(%v),Get request url(%v),params(%v),response(%v)", time.Now().Sub(t).Nanoseconds(), resp.Request.URL, Struct2String(params), body)
	return body, nil
}

func HttpPostJson(url string, params interface{}, headers map[string]string, timeout time.Duration) (string, error) {
	t := time.Now()
	request := gorequest.New().Post(url).Timeout(timeout)
	for key, value := range headers {
		request.Set(key, value)
	}
	resp, body, errs := request.Send(params).End()
	if errs != nil {
		log.LOGERROR("PostJson request url(%v),params(%v), headers(%s), errs(%v)", url, Struct2String(params), headers, errs)
		return "", errors.New(Struct2String(errs))
	}
	log.LOGFLOW("latency(%v),PostJson request url(%v),params(%v), headers(%s), reponse(%v)", time.Now().Sub(t).Nanoseconds(), resp.Request.URL, Struct2String(params), headers, body)
	return body, nil
}

func HttpPut(url string, params interface{}, headers map[string]string, timeout time.Duration) (string, error) {
	t := time.Now()
	request := gorequest.New().Put(url).Timeout(timeout)
	for key, value := range headers {
		request.Set(key, value)
	}
	resp, body, errs := request.Send(params).End()
	if errs != nil {
		log.LOGERROR("Put request url(%v),params(%v),errs(%v)", url, Struct2String(params), errs)
		return "", errors.New(Struct2String(errs))
	}
	log.LOGFLOW("latency(%v),Put request url(%v),params(%v),reponse(%v)", time.Now().Sub(t).Nanoseconds(), resp.Request.URL, Struct2String(params), body)
	return body, nil
}

func HttpDelete(url string, params interface{}, headers map[string]string, timeout time.Duration) (string, error) {
	t := time.Now()
	request := gorequest.New().Delete(url).Timeout(timeout)
	for key, value := range headers {
		request.Set(key, value)
	}
	resp, body, errs := request.Send(params).End()
	if errs != nil {
		log.LOGERROR("Delete request url(%v),params(%v),errs(%v)", url, Struct2String(params), errs)
		return "", errors.New(Struct2String(errs))
	}
	log.LOGFLOW("latency(%v),Delete request url(%v),params(%v),response(%v)", time.Now().Sub(t).Nanoseconds(), resp.Request.URL, Struct2String(params), body)
	return body, nil
}

// Parse Result
func ParseResult(body string, out *context.OutPackage, ack interface{}) error {
	if err := json.Unmarshal([]byte(body), out); err != nil {
		terrors.ERR_PARAM_PARSE_ERROR.ToErrMsg(out)
		return err
	}
	if out.GetStatus() != 0 {
		return terrors.ErrCodeToStand(out)
	}
	if ack == nil {
		return nil
	}
	dataBytes, err := json.Marshal(out.GetData())
	if err != nil {
		terrors.ERR_PARAM_PARSE_ERROR.ToErrMsg(out)
		return err
	}
	if err = json.Unmarshal(dataBytes, ack); err != nil {
		terrors.ERR_PARAM_PARSE_ERROR.ToErrMsg(out)
		return err
	}
	// set out.Data
	out.SetData(ack)
	return nil
}
