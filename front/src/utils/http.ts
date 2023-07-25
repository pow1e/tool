/*
 * @Description: 
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-07-14 09:26:46
 * @LastEditors: William
 * @LastEditTime: 2023-07-24 10:52:30
 */
import axios from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'

let httpInstance = axios.create({
  baseURL: '',
  timeout: 5000
})

if (process.env.NODE_ENV === 'development') {
  httpInstance.defaults.baseURL = '/api'
}
else {
    httpInstance.defaults.baseURL = 'http://localhost:8080/api'
}

// axios请求拦截器
httpInstance.interceptors.request.use(config => {
  return config
}, e => Promise.reject(e))

// axios响应式拦截器
httpInstance.interceptors.response.use(res => res.data, error => {
  // 统一错误提示
  if (!error.response) {
    ElMessageBox.confirm(`
        <p>检测到请求错误</p>
        <p>${error}</p>
        `, '发生错误', {
      dangerouslyUseHTMLString: true,
      distinguishCancelAndClose: true,
      confirmButtonText: '稍后重试',
      cancelButtonText: '取消'
    })
    return
  }

  switch (error.response.status) {
    case 500:
      ElMessageBox.confirm(`
        <p>检测到接口错误${error}</p>
        <p>错误码<span style="color:red"> 500 Interval Server Error</span>：请先查看后台日志，并将错误日志上传至 Issue 处</p>
        `, '发生错误', {
        dangerouslyUseHTMLString: true,
        distinguishCancelAndClose: true,
        confirmButtonText: '我知道了',
        cancelButtonText: '取消'
      })
      break
    case 404:
      ElMessageBox.confirm(`
          <p>检测到接口错误${error}</p>
          <p>错误码<span style="color:red"> 404 </span></p>
          `, '发生错误', {
        dangerouslyUseHTMLString: true,
        distinguishCancelAndClose: true,
        confirmButtonText: '我知道了',
        cancelButtonText: '取消'
      })
      break
  }

  return error
})

export default httpInstance
