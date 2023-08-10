/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-04 16:53:29
 * @LastEditors: William
 * @LastEditTime: 2023-08-08 15:38:21
 */
import request from 'utils/request';

export const getdecrypt = (data: any) => {
  const result = request({
    url: 'api/v1/md5/decrypt',
    method: 'post',
    data: data
  })
  return {
    result
  }
}

export const getencrypt = (data: any) => {
  const result = request({
    url: 'api/v1/encrypt',
    method: 'post',
    data: data
  })
  return {
    result
  }
}
