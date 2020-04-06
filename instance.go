package qingyun

import (
	"net/url"
)

// RunInstances
func (cli *Client) RunInstances(data map[string]string) (urlValues url.Values) {
	// urlValues.Set("action", "RunInstances")
	for k, v := range data {
		urlValues.Set(k, v)
	}
	// return urlValues
	// body, err := cli.HttpGET("RunInstances", urlValues, nil)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s", body)

	return
}
