package request

import (
	"errors"
	"sort"
	"tools/global"
)

type PortScan struct {
	TimeOut *int   `json:"timeOut"`
	Ports   []int  `json:"ports"`
	Ip      string `json:"ip"`
}

type PortScanOption func(*PortScan) error

func WithIp(ip string) PortScanOption {
	return func(scanner *PortScan) error {
		scanner.Ip = ip
		return nil
	}
}

func WithPorts(ports []int) PortScanOption {
	return func(scanner *PortScan) error {
		// verify ports is legal
		sort.Ints(ports)
		if ports[len(ports)-1] > 65535 {
			return errors.New("端口号不能大于65535")
		}
		if ports[0] < 1 {
			return errors.New("端口号不能小于1")
		}
		scanner.Ports = ports
		return nil
	}
}

func WithTimeOut(timeOut int) PortScanOption {
	return func(scanner *PortScan) error {
		if timeOut == 0 {
			return errors.New("")
		}
		*scanner.TimeOut = timeOut
		return nil
	}
}

func (p *PortScan) SetPortScanner(opts ...PortScanOption) error {
	for _, opt := range opts {
		err := opt(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PortScan) Validate() error {
	if len(p.Ports) == 0 {
		p.Ports = global.TargetPorts
		return nil
	}
	if len(p.Ports) == 1 {
		return nil
	}
	sort.Ints(p.Ports)
	if p.Ports[len(p.Ports)-1] > 65535 {
		return errors.New("端口号不能大于65535")
	}
	if p.Ports[0] < 1 {
		return errors.New("端口号不能小于1")
	}
	// calculate the length of the ports
	end := p.Ports[len(p.Ports)-1]
	start := p.Ports[0]
	count := end - start + 1
	res := make([]int, count)
	for i := range res {
		res[i] = start + i
	}
	// cover the ports
	p.Ports = res
	return nil
}

func (p *PortScan) ParserPorts(ports []int) {

}

func NewPortScanner(opts ...PortScanOption) (*PortScan, error) {
	portScanner := &PortScan{}
	for _, opt := range opts {
		err := opt(portScanner)
		if err != nil {
			return nil, err
		}
	}
	return portScanner, nil
}
