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

skynet.register_protocol {
    name = "text",
    id = skynet.PTYPE_TEXT,
    unpack = skynet.tostring,
        pack = function (...)
        local n = select ("#" , ...)
        if n == 0 then
            return ""
        elseif n == 1 then
            return tostring(...)
        else
            return table.concat({...}," ")
        end
    end,
}
local addr = skynet.launch("gos")
print("---------------:", skynet.call(addr, "text", "hello, world"))
```



