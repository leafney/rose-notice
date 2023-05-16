package xizhi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (r *XiZhi) send(title, body string) error {
	if utils.IsEmpty(r.token) {
		return vars.ErrTokenEmpty
	}

	msg := map[string]string{
		"title": title,
	}

	if utils.IsNotEmpty(body) {
		msg["content"] = body
	}

	webURL := fmt.Sprintf("%s/%s.send", r.host, r.token)

	m, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if r.debug {
		fmt.Printf("[debug] body: [%v] \n", string(m))
	}

	req, err := http.NewRequest(http.MethodPost, webURL, bytes.NewReader(m))
	if err != nil {
		return err
	}

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
	if dr.Code != 200 {
		return fmt.Errorf("xizhi notification send failed: [%v]", dr.Msg)
	}

	return nil
}
