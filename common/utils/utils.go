/**
 * @Author:      leafney
 * @Date:        2022-10-10 11:18
 * @Project:     rose-notify
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package utils

import (
	"net/url"
	"path"
	"strconv"
	"strings"
)

func JoinPath(basePath string, elem ...string) (string, error) {
	u, err := url.Parse(basePath)
	if err != nil {
		return "", err
	}

	if len(elem) > 0 {
		elem = append([]string{u.Path}, elem...)
		u.Path = path.Join(elem...)
	}

	return u.String(), nil
}

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsNotEmpty(s string) bool {
	return len(strings.TrimSpace(s)) > 0
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func IsNumber(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
