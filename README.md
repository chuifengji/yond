## 写玩具才能用的 nodejs 玩具框架 yond

用来发泄个人对 npm 的“厌恶”~

npm 什么的，最讨厌了！

主要是对 http import 这种语法进行预编译转换， 摆脱 npm install 和 node_modules

### Usage

`install`

go get github.com/chuifengji/yond/v1 (还没发布呢)

`yond init` 

初始化项目目录

`yond dev` 

将 http import 转换到本地的 modules 目录下 

就是 test 文件中 src -> dev

`yond build` 

接入 esbuild 进行进一步的代码转换 


先写个基本框架，有空再写了


### TODO：

- url —— moduleName —— project 三者需要一一对应
- 全局缓存，一个 url 只被下载一次
- esbuild 接入完善
- 配置文件实现

### Problem:
- 需要可靠的 cdn 平台
- 容易产生大量的重复代码

