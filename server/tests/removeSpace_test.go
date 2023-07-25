package tests

import (
	"fmt"
	"testing"
	"tools/pkg/utils"
)

func TestRemoveSpace(t *testing.T) {
	space := utils.RemoveSpace("redis-server.exe             20720 Services                   0     51,416 K")
	for _, v := range space {
		fmt.Println(v)
	}
}
