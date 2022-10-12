/**
 * @Author:      leafney
 * @Date:        2022-10-12 10:17
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package pushdeer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/leafney/rose-notify/utils"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Response struct {
	Code    int      `json:"code"`
	Message RespCont `json:"content"`
}
type RespCont struct {
	Result []string `json:"result"`
}

func (r *Robot) send(msgType, title, body string) error {

	msg := map[string]string{
		"type":    msgType,
		"text":    title,
		"desp":    body,
		"pushkey": r.key,
	}

	var method = http.MethodPost
	webURL, _ := utils.JoinPath(r.host, "message/push")
	value := url.Values{}
	var bodyData io.Reader

	if r.isGet {
		method = http.MethodGet
		for k, v := range msg {
			if len(v) == 0 {
				continue
			}
			value.Set(k, v)
		}
	} else {
		m, err := json.Marshal(msg)
		if err != nil {
			return err
		}
		bodyData = bytes.NewReader(m)
	}

	req, err := http.NewRequest(method, webURL, bodyData)
	if err != nil {
		return err
	}

	if r.isGet {
		req.URL.RawQuery = value.Encode()
	} else {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}
	//fmt.Println(req.URL.String()) // TODO

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

	//fmt.Println(string(data)) // TODO

	var dr Response
	err = json.Unmarshal(data, &dr)
	if err != nil {
		return err
	}
	if dr.Code != 0 {
		return fmt.Errorf("pushdeer notification send failed: [%v]", dr.Message)
	}

	return nil
}
