package service

import "sync"

// 调用方法的单例模式
var (
	SysSrvImpl      *SysSrv
	ShortCutSrvImpl *ShortCutSrv
	ScanImpl        *ScanSrv
	AlgorithmImpl   *AlgorithmSrv
)

var (
	SysSrvOnce      sync.Once
	ShortCutSrvOnce sync.Once
	ScanSrvOnce     sync.Once
	AlgorithmOnce   sync.Once
)

type ScanSrv struct{}
type SysSrv struct{}
type ShortCutSrv struct{}
type AlgorithmSrv struct{}

// GetSysSrv 获取单例模式的GetSysSrv对象
func GetSysSrv() *SysSrv {
	SysSrvOnce.Do(func() {
		SysSrvImpl = &SysSrv{}
	})
	return SysSrvImpl
}

// GetShortCutSrv 获取单例模式的GetSysSrv对象
func GetShortCutSrv() *ShortCutSrv {
	ShortCutSrvOnce.Do(func() {
		ShortCutSrvImpl = &ShortCutSrv{}
	})
	return ShortCutSrvImpl
}

func GetScanSrv() *ScanSrv {
	ScanSrvOnce.Do(func() {
		ScanImpl = &ScanSrv{}
	})
	return ScanImpl
}

func GetAlgorithmSrv() *AlgorithmSrv {
	AlgorithmOnce.Do(func() {
		AlgorithmImpl = &AlgorithmSrv{}
	})
	return AlgorithmImpl
}
