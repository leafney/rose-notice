/**
 * @Author:      leafney
 * @Date:        2022-10-10 11:18
 * @Project:     rose-notice
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package utils

import (
	"net/url"
	"path"
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

func IsNotEmpty(s string) bool {
	return len(strings.TrimSpace(s)) > 0
}
