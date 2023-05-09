package wochat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (r *Robot) send(msg interface{}) error {
	if utils.IsEmpty(r.key) {
		return vars.ErrTokenEmpty
	}

	m, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if r.debug {
		fmt.Printf("[debug] body: [%v] \n", string(m))
	}

	webURL := r.host
	value := url.Values{}
	value.Set("key", r.key)

	req, err := http.NewRequest(http.MethodPost, webURL, bytes.NewReader(m))
	if err != nil {
		return err
	}

	req.URL.RawQuery = value.Encode()
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

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

	if dr.ErrCode != 0 || dr.ErrMsg != "ok" {
		return fmt.Errorf("wowechat notification send failed: [%v]", dr.ErrMsg)
	}

	return nil
}
