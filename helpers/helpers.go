package helpers

import (
	"fmt"

	"github.com/ncrypt96/gena/constants"
	"github.com/ncrypt96/gena/utils"
)

func ChangeDefaultVarMap(v map[string]string) {
	c := 0
	for k, val := range v {
		c++
		var ip string
		fmt.Printf("\n%d/%d\n", c, len(v))
		fmt.Printf(constants.EnterValForVarMsg, k)
		fmt.Printf(constants.DefaultValMsg, val)
		fmt.Scanf("%s\n", &ip)
		if utils.IsNotEmptyStr(ip) {
			// NOTE: the line below modifies the original map
			v[k] = ip
		}
	}
}
