package service

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
	"tools/global"
	"tools/model/system"
	"tools/model/system/request"
	"tools/pkg/consts"
)

var writerLock sync.Mutex

func (s *ScanSrv) ScanPort(target *request.PortScan) ([]*system.PortScanResp, error) {
	var wg sync.WaitGroup
	scanners := make([]*system.PortScanResp, 0)
	if len(target.Ports) != 0 {
		for _, port := range target.Ports {
			if port > 65535 {
				return nil, errors.New("端口号不能大于65535")
			}
			if port < 1 {
				return nil, errors.New("端口号不能小于1")
			}
		}
	} else {
		target.Ports = global.TargetPorts
	}
	if target.Ip == "" {
		return nil, errors.New("ip不能为空")
	}
	if target.TimeOut != 0 {
		target.TimeOut = consts.TimeOut
	}

	for _, port := range target.Ports {
		wg.Add(1)
		go func(ip string, port int, wg *sync.WaitGroup) {
			defer wg.Done()
			conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Duration(target.TimeOut)*time.Second)
			if err == nil {
				defer conn.Close()
				writerLock.Lock()
				target := system.BasePortScanner(ip, port)
				target.BuildInfoScan()
				scanners = append(scanners, target)
				writerLock.Unlock()
			}
		}(target.Ip, port, &wg)
	}
	wg.Wait()
	return scanners, nil
}
