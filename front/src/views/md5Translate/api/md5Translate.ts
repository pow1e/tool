import httpInstance from '@/utils/http'

export const getMd5Decrypt = (data: any) => {
  return httpInstance({
    url: '/v1/md5/decrypt',
    method: 'post',
    data
  })
}

export const getMd5Encrypt = (data: any) => {
  return httpInstance({
    url: '/v1/md5/encrypt',
    method: 'post',
    data
  })
}