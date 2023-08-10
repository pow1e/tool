package service

import (
	"fmt"
	"net"
	"sync"
	"time"
	"tools/model/system"
	"tools/model/system/request"
)

var writerLock sync.Mutex

func (s *ScanSrv) ScanPort(target *request.PortScan) ([]*system.PortScanResp, error) {
	var wg sync.WaitGroup
	scanners := make([]*system.PortScanResp, 0)

	// verify the data and parse port slices after verify,
	if err := target.Validate(); err != nil {
		return nil, err
	}

	// get information about this port
	for _, port := range target.Ports {
		wg.Add(1)
		go func(ip string, port int, wg *sync.WaitGroup) {
			defer wg.Done()
			var conn net.Conn
			var err error
			if target.TimeOut != nil {
				conn, err = net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Duration(*target.TimeOut)*time.Second)
			} else {
				conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
			}
			if err == nil {
				defer conn.Close()
				writerLock.Lock()
				target := system.BasePortScanner(true, ip, port).BuildInfoScan()
				scanners = append(scanners, target)
				writerLock.Unlock()
			}
		}(target.Ip, port, &wg)
	}
	wg.Wait()
	return scanners, nil
}
