# tools

## 技术栈

- 后端：golang,gin
- 前端：Electron-Egg + React + TypeScript + TDesign + Vite 

## 功能模块

+ [x] 快捷指令
    + [x] 新建快捷指令
    + [x] 修改快捷指令
    + [x] 删除快捷指令
    + [x] 运行快捷指令
    + [x] 获取快捷指令
+ [x] md5功能
    +[x] md5加密
    +[x] md5解密

## 后端项目结构

```
├─api
│  └─v1              # handler入口
├─cmd                # 程序入口   
├─config             # 配置入口
├─core               
│  └─internal        # 核心组件的初始化
├─docs               # swagger文档目录
├─global             # 全局对象
├─log                # 日志文件
├─middleware         # 中间件
├─model
│  ├─common          # 公共结构体
│  │  ├─request     
│  │  └─response    
│  └─system          # 请求响应结构体
│      ├─request    
│      └─response   
├─pkg               
│  ├─consts          # 全局常量
│  ├─utils           # 全局工具包
│  │  ├─admin       
│  │  └─users       
│  └─window          # window原生函数
├─router             # 路由
├─service            # 服务层
├─static             # 存放静态文件
│  ├─dict           
│  │  ├─ciphertext
│  │  └─plaintext
│  └─imgs
└─tests              # 单元测试
```

## 前端项目结构

```
├── build                   # Electron-Egg 构建文件
├── data                    # Electron-Egg 配置文件
├── electron                # Electron-Egg 主程序及插件配置文件
|  ├── addon                #
|  |  ├── autoUpdater       #
|  |  ├── awaken            #
|  |  ├── chromeExtension   #
|  |  ├── javaServer        #
|  |  ├── security          #
|  |  └── tray              #
|  ├── config               #
|  ├── controller           #
|  ├── index.js             #
|  ├── jobs                 #
|  ├── preload              #
|  └── service              #
├── frontend                # 前端源文件
|  ├── index.html           # 首页文件
|  ├── LICENSE              # 开源协议
|  ├── package.json         # 前端项目启动脚本及所含依赖
|  ├── public               # 前端项目公共文件
|  ├── README-zh_CN.md      # 前端所用的模版中文说明
|  ├── README.md            # 前端所用的模版英文说明
|  ├── src                  # 核心源文件
|  |  ├── assets            # 项目所用到的媒体文件
|  |  ├── components        # 前端组件
|  |  |  ├── Board
|  |  |  ├── DatePicker
|  |  |  └── ErrorPage
|  |  ├── configs           # 配置文件
|  |  |  └── color.ts
|  |  ├── global.d.ts       # 全局命名文件
|  |  ├── hooks             # 钩子函数
|  |  ├── layouts           # 前端布局组件
|  |  ├── main.tsx          # React 框架主入口
|  |  ├── modules           # 组件所使用的模组
|  |  ├── pages             # 页面文件
|  |  ├── router            # 页面路由
|  |  ├── services          # api调用
|  |  ├── styles            # 全局样式
|  |  ├── types             # 暗黑主题模式切换
|  |  └── utils             # 工具文件
|  ├── tsconfig.json        # TypeScript 设置
|  ├── vite.config.js       # Vite 设置
├── LICENSE
├── main.js                 # Electron-Egg 主入口
├── package.json            # Electron-Egg 启动脚本及所含依赖
├── public                  # 公共文件
├── README.md               # Electron-Egg 英文说明文档
├── README.zh-CN.md         # Electron-Egg 中文说明文档
```