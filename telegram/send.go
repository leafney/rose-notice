package telegram

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
	"strings"
)

type Response struct {
	OK          bool        `json:"ok"`
	ErrorCode   int         `json:"error_code,omitempty"`
	Description string      `json:"description,omitempty"`
	Result      interface{} `json:"result,omitempty"`
}

func (r *Telegram) send(text string) error {
	if utils.IsEmpty(r.token) {
		return vars.ErrTokenEmpty
	}

	if utils.IsEmpty(r.chatId) {
		return vars.ErrTelegramChatIdEmpty
	}

	// Determine chatId format
	if !utils.IsNumber(r.chatId) {
		if !strings.HasPrefix(r.chatId, "@") {
			r.chatId = fmt.Sprintf("@%s", r.chatId)
		}
	}

	msg := map[string]string{
		"text":    text,
		"chat_id": r.chatId,
	}

	if r.isMarkdown {
		msg["parse_mode"] = "MarkdownV2"
	}

	if r.isHtml {
		msg["parse_mode"] = "HTML"
	}

	if r.isSilent {
		msg["disable_notification"] = "true"
	}

	if r.disablePagePreview {
		msg["disable_web_page_preview"] = "true"
	}

	var method = http.MethodPost
	webURL, _ := utils.JoinPath(r.host, fmt.Sprintf("bot%s", r.token), "sendMessage")
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

	// test set proxy
	//proxyUrl, _ := url.Parse("http://127.0.0.1:7890")
	//client := http.Client{
	//	Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
	//}

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
	if !dr.OK {
		return fmt.Errorf("telegram notification send failed: [%v] - [%v]", dr.ErrorCode, dr.Description)
	}

	return nil
}
