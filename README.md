# qingyun-sdk-go
qingyun sdk go 简单鉴权版本


## Usage

```go
package qingyun

import (
	"fmt"
	"testing"
)

const (
	authFile = `/path/2/.qingcloud/config.yaml`
)

// Test_DescribeReservedResources 根据预留合约ID 查询绑定信息
func Test_DescribeReservedResources(t *testing.T) {
	cli := NewWithFile(authFile)

	data := map[string]string{
		"contracts.1": "rc-kkkkkkkk",
		"zone":        "pek3c",
	}
	bodyByte, _ := cli.Get("DescribeReservedResources", data)
	fmt.Printf("%s\n", bodyByte)
}

```

```json
{"action":"DescribeReservedResourcesResponse","total_count":1,"reserved_resource_set":[{"user_id":"usr-PPPPPPP","zone_id":"pek3","resource_id":"vol-abcdefg","root_user_id":"usr-PPPPPPP","create_time":"2020-04-01T14:33:40Z","entry_id":"rce-234adf23f","contract_id":"rc-kkkkkkkk"}],"zone_id":"pek3","ret_code":0}
```

## 天坑

青云验签文档有个坑， 参数和结果对不上。请求demo `zone=pek3b` ， 实际结果用的是 `zone=pek1`。