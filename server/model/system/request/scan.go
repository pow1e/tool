package request

type PortScan struct {
	TimeOut int    `json:"timeOut"`
	Ports   []int  `json:"ports"`
	Ip      string `json:"ip"`
}
