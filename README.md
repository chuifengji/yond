## 写玩具才能用的 nodejs 玩具框架 yond

主要是用来发泄个人对 npm 的“厌恶”~

npm 什么的，最讨厌了！

主要是对 http import 这种语法进行预编译转换， 摆脱 npm install 和 node_modules

### Usage

`install`

go get github.com/chuifengji/yond/v1 (还没发布呢)

`yond init` 

初始化项目目录

`yond dev` 

将 http import 转换到本地的 modules 目录下 

`yond build` 

接入 esbuild 进行进一步的代码转换 


先写个基本框架，有空再写了

### Steps：
1. 拷贝整个src目录 到 dev 目录 
2. GetFiles return []filePath, isDir -> return
3. readFile 单个文件 is js/ts ? go on
4. GetHttpModule 读取单个文件 不为空则替换 import 为 /dev/modules/moduleName
5. WriteFileTo 写入处理后的文件内容到 dev 目录 指定位置，找不到文件会创建吗？
6. 新增 http 下载任务到任务池中，如果没有相同下载任务就加入，下载成功写入到 /dev/modules 
7. 在 js.mod 中新增记录

3、4 步似乎可以并发提升效率。

### TODO：

- url —— moduleName —— project 三者需要一一对应
- 全局缓存，一个 url 只被下载一次
- esbuild 接入完善
- cli 实现
- 配置文件实现

### Problem:
- 需要可靠的 cdn 平台
- 容易产生大量的重复代码

