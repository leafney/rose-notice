/**
 * @Author:      leafney
 * @Date:        2022-10-09 10:03
 * @Project:     rose-notice
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (r *Robot) send(msg interface{}) error {
	m, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	webURL := r.baseUrl
	value := url.Values{}
	value.Set("access_token", r.token)

	if len(r.secret) != 0 {
		t := time.Now().UnixMilli()
		signMsg := fmt.Sprintf("%d\n%s", t, r.secret)
		value.Set("timestamp", fmt.Sprintf("%d", t))
		value.Set("sign", sign(signMsg, r.secret))
	}

	req, err := http.NewRequest(http.MethodPost, webURL, bytes.NewReader(m))
	if err != nil {
		return err
	}

	req.URL.RawQuery = value.Encode()
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var dr Response
	err = json.Unmarshal(data, &dr)
	if err != nil {
		return err
	}
	if dr.ErrCode != 0 {
		return fmt.Errorf("dingtalk notice send failed: [%v]", dr.ErrMsg)
	}

	return nil
}

func sign(message, secret string) string {
	key := []byte(secret)
	hmac256 := hmac.New(sha256.New, key)
	hmac256.Write([]byte(message))
	data := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(data)
}
