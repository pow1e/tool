package tests

import (
	"fmt"
	"testing"
	"tools/pkg/utils/users"
)

func TestUserRun(t *testing.T) {
	res, err := users.Run("tasklist /FI \"IMAGENAME eq mysqld.exe\"\n")
	fmt.Println(res)
	fmt.Println(err)
}
