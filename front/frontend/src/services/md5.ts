/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-04 16:53:29
 * @LastEditors: William
 * @LastEditTime: 2023-08-04 17:24:26
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
