package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/leafney/rose-notify/common/utils"
	"github.com/leafney/rose-notify/common/vars"
	"io/ioutil"
	"net/http"
)

func (r *Robot) send(msg interface{}) error {
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

	var dr = string(data)
	if dr != "ok" {
		return fmt.Errorf("slack notification send failed: [%v]", dr)
	}

	return nil
}
