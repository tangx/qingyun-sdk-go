package client

import (
	"fmt"
	"testing"
)

const (
	authFile = `/Users/tangxin/.qingcloud/config.yaml`
)

func Test_Get(t *testing.T) {
	cli := NewWithFile(authFile)

	action := `DescribeZones`
	bodyByte, err := cli.requestGET(action, nil, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", bodyByte)
}
