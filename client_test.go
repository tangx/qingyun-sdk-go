package qingyun

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

func Test_RunInstance(t *testing.T) {
	cli := NewWithFile(authFile)

	runInstancesData := map[string]string{
		// 操作系统镜像ID
		"image_id": "xenial6x64",
		// cpu
		"cpu": "2",
		// 内存 (GB)
		"memory": "4096",
		// 实例类型 202: 企业e2
		"instance_class": "202",
		// 系统盘大小(GB)
		"os_disk_size": "40",
		// 登录方式
		"login_mode":    "keypair",
		"login_keypair": "kp-2kodyll8",
		// 实例名称
		"instance_name": "qingyun-sdk-go-running-d",
		// 所属网络
		"vxnets.1": "vxnet-kj0thr5",
		// 所在可用区
		"zone": "pek3d",
	}
	// urlValue := cli.RunInstances(runInstancesData)

	bodyByte, _ := cli.Do("RunInstances", runInstancesData, nil, nil)
	fmt.Printf("%s", bodyByte)
}
