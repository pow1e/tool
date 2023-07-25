/*
 * @Description:
 * @version:
 * @Author: William
 * @Date: 2023-07-12 15:07:24
 * @LastEditors: William
 * @LastEditTime: 2023-07-12 15:47:28
 */
import { createApp } from 'vue'
import router from './router/index'
import App from './App.vue'
import ElementPlus from 'element-plus'
import "element-plus/dist/index.css";
// 增加暗黑模式样式文件
import 'element-plus/theme-chalk/dark/css-vars.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
// import Message from './utils/message'
const app = createApp(App)
app.use(ElementPlus)
app.use(router)
// app.use(Message)
// app.provide('$showMessage', Message.install)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
app.mount('#app').$nextTick(() => {
  postMessage({ payload: 'removeLoading' }, '*')
})
