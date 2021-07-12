
### 先写个基本框架，能跑起来就行，有空再写了

1. 拷贝整个src目录 到 temp 目录 
2. GetFiles return []filePath, isDir -> return
3. readFile 单个文件 is js/ts ? go on
4. GetHttpModule 读取单个文件 不为空则替换 import 为 /temp/modules/moduleName
5. WriteFileTo 写入处理后的文件内容到 temp 目录 指定位置，找不到文件会创建吗？
6. 新增 http 下载任务到任务池中，如果没有相同下载任务就加入，下载成功写入到 /temp/modules 
7. 在 js.mod 中新增记录

3、4 步似乎可以并发提升效率。

### TODO：

- url —— moduleName —— project 三者需要一一对应
- 全局缓存，一个 url 只被下载一次
- esbuild 接入完善
- cli 实现
- 配置文件实现



