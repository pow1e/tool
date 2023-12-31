# tools

## 技术栈

- 后端：golang,gin
- 前端：Vue 3 + TypeScript + Vite

## 功能模块

+ [x] 快捷指令
    + [x] 新建快捷指令
    + [x] 修改快捷指令
    + [x] 删除快捷指令
    + [x] 运行快捷指令
    + [x] 获取快捷指令
+ [x] 加密算法功能
    + [x] 算法加密
    + [ ] 算法解密
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
