# scrcpy
https://github.com/Genymobile/scrcpy/tree/master/doc


# scrcpy-noconsole.vbs
* 无控制台界面
* 使用scrcpy-noconsole.vbs创建桌面快捷方式,换图标
* 快捷方式添加参数 --turn-screen-off --stay-awake
D:\tools\scrcpy\scrcpy-noconsole.vbs --turn-screen-off --stay-awake
--turn-screen-off 关闭屏幕
--stay-awake 保持唤醒



# adb connect 
* 无线匹配
```
adb pair 14.mm:port 
code
```

* 之后的脚本
```
@echo off
 
set /p port="set 14.mm port:"

D:\tools\sdk\platform-tools\adb connect 14.mm:%port%
D:\tools\sdk\platform-tools\adb tcpip 5555
D:\tools\sdk\platform-tools\adb kill-server
D:\tools\sdk\platform-tools\adb connect 14.mm
D:\tools\sdk\platform-tools\adb shell wm size 720x1280
```