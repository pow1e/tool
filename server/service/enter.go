package service

import "sync"

// 调用方法的单例模式
var (
	UserSrvImpl     *UserSrv
	SysSrvImpl      *SysSrv
	ShortCutSrvImpl *ShortCutSrv
)

var (
	UserSrvOnce     sync.Once
	SysSrvOnce      sync.Once
	ShortCutSrvOnce sync.Once
)

// 单例模式的结构体
type UserSrv struct{}
type SysSrv struct{}
type ShortCutSrv struct{}

// GetUserSrv 获取单例模式的UserSrv对象
func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvImpl = &UserSrv{}
	})
	return UserSrvImpl
}

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
