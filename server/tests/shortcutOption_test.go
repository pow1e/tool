package tests

import (
	"fmt"
	"testing"
	"tools/model/system"
)

func TestOption(t *testing.T) {
	shortcut := &system.Shortcut{
		Id:          "1",
		Name:        "",
		Description: "1",
		Directives:  "",
		Type:        "1",
	}

	err := shortcut.Validate(system.WithShorCutName())
	if err != nil {
		fmt.Println("参数校验失败", err.Error())
	}
	fmt.Println("参数校验成功")
}
