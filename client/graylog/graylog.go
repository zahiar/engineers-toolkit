package graylog

import (
	"github.com/suzuki-shunsuke/go-graylog/graylog/client"
)

func GetClient() (*client.Client, error) {
	return client.NewClientV3(
		"http://localhost:9000/api",
		"admin",
		"admin",
	)
}
