package tests

import (
	"testing"
	"tools/model/system/request"
)

func TestName(t *testing.T) {
	scanner, err := request.NewPortScanner(
		request.WithPorts([]int{12, 434, 65535}),
		request.WithIp("127.0.0.1"),
		request.WithTimeOut(123),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(scanner)
}
