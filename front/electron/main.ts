/*
 * @Description: 
 * @version: 
 * @Author: William
 * @Date: 2023-07-12 15:07:24
 * @LastEditors: William
 * @LastEditTime: 2023-07-14 10:33:57
 */
import { app, BrowserWindow, dialog } from 'electron'
import { ElMessage } from 'element-plus/es/components/index.js'
import path from 'node:path'
import { windowsStore } from 'node:process'

// The built directory structure
//
// ├─┬─┬ dist
// │ │ └── index.html
// │ │
// │ ├─┬ dist-electron
// │ │ ├── main.js
// │ │ └── preload.js
// │
process.env.DIST = path.join(__dirname, '../dist')
process.env.PUBLIC = app.isPackaged ? process.env.DIST : path.join(process.env.DIST, '../public')
process.env.ELECTRON_DISABLE_SECURITY_WARNINGS = 'true'

let win: BrowserWindow | null
// 🚧 Use ['ENV_NAME'] avoid vite:define plugin - Vite@2.x
const VITE_DEV_SERVER_URL = process.env['VITE_DEV_SERVER_URL']


function createWindow() {
  win = new BrowserWindow({
    icon: path.join(process.env.PUBLIC, 'electron-vite.ico'),
    webPreferences: {
      preload: path.join(__dirname, 'preload.js'),
    },
    width: 1080,
    height: 625,
    resizable: false,
  })

  // Test active push message to Renderer-process.
  win.webContents.on('did-finish-load', () => {
    win?.webContents.send('main-process-message', (new Date).toLocaleString())
  })
  // win.removeMenu()
  if (VITE_DEV_SERVER_URL) {
    win.loadURL(VITE_DEV_SERVER_URL)
  } else {
    // win.loadFile('dist/index.html')
    win.loadFile(path.join(process.env.DIST, 'index.html'))
  }
  win.on('close', (event) => {
    event.preventDefault()
    dialog.showMessageBox({
      type: 'info',
      title: '提示',
      message: '您确定要退出吗',
      buttons: ['确定', '取消']
    }).then(res => {
      if (res.response === 0) {
        app.exit()
      }
    })
  })
    
}




app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit()
  }
})

app.whenReady().then(createWindow)
