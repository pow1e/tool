/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-07-14 09:56:36
 * @LastEditors: William
 * @LastEditTime: 2023-07-24 11:00:18
 */
import httpInstance from '@/utils/http'
/**
 * @name: 获取快捷命令列表
 * @test: test font
 * @msg:
 * @param {any} data
 * @return {*}
 */
export const getQuickCommandList = () => {
  return httpInstance({
    url: '/v1/shortcuts',
    method: 'get'
  })
}
/**
 * @name: 新增快捷指令
 * @test: test font
 * @msg:
 * @param {any} data
 * @return {*}
 */
export const addQuickCommand = (data: any) => {
  return httpInstance({
    url: '/v1/shortcuts',
    method: 'post',
    data
  })
}
/**
 * @name: 更新快捷指令
 * @test: test font
 * @msg:
 * @param {any} data
 * @return {*}
 */
export const updateQuickCommand = (data: any) => {
  return httpInstance({
    url: '/v1/shortcuts',
    method: 'put',
    data
  })
}
/**
 * @name: 运行快捷指令
 * @test: test font
 * @msg:
 * @param {string} id
 * @return {*}
 */
export const runQuickCommand = (id: string) => {
  return httpInstance({
    url: '/v1/shortcuts/run',
    method: 'get',
    params: {
      id
    }
  })
}

/**
 * @name: 删除快捷指令
 * @test: test font
 * @msg: 
 * @param {string} id
 * @return {*}
 */
export const deleteQuickCommand = (id: string) => {
  return httpInstance({
    url: `/v1/shortcuts/${id}`,
    method: 'delete'
  })
}