## cservice

用`golang`为[skynet](https://github.com/cloudwu/skynet)写服务。
## use

* 编译
```
make
```

* 拷贝gos.so到cservice目录
* 启动服务
```lua
local skynet = require "skynet"
skynet.launch("gos")
```



