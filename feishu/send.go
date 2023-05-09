/**
 * @Author:      leafney
 * @Date:        2022-10-09 11:19
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package feishu

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
	Code          int    `json:"code"`
	Msg           string `json:"msg"`
	StatusCode    int    `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
}

func (r *FeiShu) send(msg interface{}) error {
	if utils.IsEmpty(r.token) {
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

	webURL, err = utils.JoinPath(webURL, r.token)
	if err != nil {
		return err
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

	if dr.StatusCode != 0 || dr.Code != 0 {
		return fmt.Errorf("feishu notification send failed: [%v]", dr.Msg)
	}

	return nil
}
