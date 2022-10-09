/**
 * @Author:      leafney
 * @Date:        2022-10-09 11:19
 * @Project:     rose-notice
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package feishu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type Response struct {
	Code          int    `json:"code"`
	Msg           string `json:"msg"`
	StatusCode    int    `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
}

func (r *Robot) send(msg interface{}) error {
	m, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	webURL := r.baseUrl

	u, err := url.Parse(webURL)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, r.token)
	webURL = u.String()

	req, err := http.NewRequest(http.MethodPost, webURL, bytes.NewReader(m))
	if err != nil {
		return err
	}

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

	if dr.StatusCode != 0 || dr.Code != 0 {
		return fmt.Errorf("feishu notice send failed: [%v]", dr.Msg)
	}

	return nil
}
