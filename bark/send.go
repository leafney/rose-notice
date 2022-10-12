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
	"github.com/leafney/rose-notify/utils"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Response struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func (r *Robot) send(title, body string) error {

	msg := map[string]string{
		"title":      title,
		"body":       body,
		"device_key": r.key,
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
	if utils.IsNotEmpty(r.soundName) {
		msg["sound"] = r.soundName
	}
	if r.badge > 0 {
		msg["badge"] = strconv.FormatInt(r.badge, 10)
	}

	var method = http.MethodPost
	webURL := r.host
	value := url.Values{}
	var bodyData io.Reader

	if r.isGet {
		webURL, _ = utils.JoinPath(webURL, r.key, title, body)
		method = http.MethodGet
		for k, v := range msg {
			if k == "device_key" || k == "title" || k == "body" {
				continue
			}
			value.Set(k, v)
		}
	} else {
		webURL, _ = utils.JoinPath(webURL, r.key)
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
	if dr.Code != 200 {
		return fmt.Errorf("bark notification send failed: [%v]", dr.Message)
	}

	return nil
}
