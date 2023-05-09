/**
 * @Author:      leafney
 * @Date:        2022-10-12 12:29
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package chanify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Response struct {
	Res        int    `json:"res"`
	Message    string `json:"msg"`
	RequestUid string `json:"request-uid"`
}

func (r *Chanify) send(title, body, urlPath string) error {
	if utils.IsEmpty(r.token) {
		return vars.ErrTokenEmpty
	}

	if r.isText && !utils.IsNotEmpty(body) {
		return errors.New("body can not empty")
	}
	// when link empty the body can not empty
	if r.isLink && !utils.IsNotEmpty(urlPath) {
		return errors.New("link can not empty")
	}

	msg := map[string]string{}

	if r.isText {
		msg["text"] = body
	}

	if r.isLink {
		msg["link"] = urlPath
	}

	if utils.IsNotEmpty(title) {
		msg["title"] = title
	}

	if utils.IsNotEmpty(r.copyText) {
		msg["copy"] = r.copyText
	}
	if r.autoCopy {
		msg["autocopy"] = "1"
	}

	if utils.IsNotEmpty(r.level) {
		msg["interruption-level"] = r.level
	}

	if r.priority > 0 && r.priority <= 10 {
		msg["priority"] = utils.IntToStr(r.priority)
	}

	if utils.IsNotEmpty(r.soundName) {
		msg["sound"] = r.soundName
	}

	var method = http.MethodPost
	webURL, _ := utils.JoinPath(r.host, "v1/sender", r.token)
	value := url.Values{}
	var bodyData io.Reader

	if r.isGet {
		// method get ,text must in url path
		webURL, _ = utils.JoinPath(webURL, body)

		method = http.MethodGet
		for k, v := range msg {
			if k == "text" {
				continue
			}
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
	if dr.Res != 0 {
		return fmt.Errorf("chanify notification send failed: [%v]", dr.Message)
	}

	return nil
}
