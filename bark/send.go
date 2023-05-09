/**
 * @Author:      leafney
 * @Date:        2022-10-09 19:32
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package bark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Response struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func (r *Bark) send(title, body string) error {
	if utils.IsEmpty(r.token) {
		return vars.ErrTokenEmpty
	}

	msg := map[string]string{
		"title":      title,
		"body":       body,
		"device_key": r.token,
	}

	if utils.IsNotEmpty(r.group) {
		msg["group"] = r.group
	}
	if utils.IsNotEmpty(r.copyText) {
		msg["copy"] = r.copyText
	}
	if r.autoCopy {
		msg["autoCopy"] = "1"
	}
	if utils.IsNotEmpty(r.icon) {
		msg["icon"] = r.icon
	}
	if utils.IsNotEmpty(r.url) {
		msg["url"] = r.url
	}
	if utils.IsNotEmpty(r.level) {
		msg["level"] = r.level
	}
	if r.isArchive {
		msg["isArchive"] = "1"
	}
	if utils.IsNotEmpty(r.soundName.String()) {
		msg["sound"] = r.soundName.String()
	}
	if r.badge > 0 {
		msg["badge"] = utils.Int64ToStr(r.badge)
	}

	var method = http.MethodPost
	webURL := r.host
	value := url.Values{}
	var bodyData io.Reader

	if r.isGet {
		webURL, _ = utils.JoinPath(webURL, r.token, title, body)
		method = http.MethodGet
		for k, v := range msg {
			if k == "device_key" || k == "title" || k == "body" {
				continue
			}
			if len(v) == 0 {
				continue
			}
			value.Set(k, v)
		}
	} else {
		webURL, _ = utils.JoinPath(webURL, "push")
		m, err := json.Marshal(msg)
		if err != nil {
			return err
		}
		if r.debug {
			fmt.Printf("[debug] body: [%v] \n", string(m))
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

	if r.debug {
		fmt.Printf("[debug] url: [%v] \n", req.URL.String())
	}

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

	if r.debug {
		fmt.Printf("[debug] response: [%v] \n", string(data))
	}

	var dr Response
	err = json.Unmarshal(data, &dr)
	if err != nil {
		return err
	}
	if dr.Code != 200 {
		return fmt.Errorf("bark notification send failed: [%v]", dr.Message)
	}

	return nil
}
