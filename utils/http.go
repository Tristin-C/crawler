package utils

import (
	"bytes"
	"context"
	"net"
	"net/http"
	"time"
)

func HttpGet(ctx context.Context, path string, body []byte, header map[string]string) (resp *http.Response, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*3) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(5 * time.Second)) //设置发送接收数据超时
				return c, nil
			},
		},
	}
	req, err := http.NewRequest("GET", path, bytes.NewReader(body))
	if err != nil {
		return
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	return
}
