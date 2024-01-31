# notepad 优化

* https://github.com/notepad-plus-plus/notepad-plus-plus/releases/download/v8.6.2/npp.8.6.2.Installer.x64.exe

## 自动换行+ UTF-8 + Unix
视图->自动换行
设置->首选项->新建->Unix(LF)
设置->首选项->新建->UTF-8(无BOM)

## 暗黑模式
设置->首选项->暗黑模式->开启暗黑模式

## 备份
设置->首选项->备份
```
冗余备份-> 使用自定义备份目录(D:\temp\)
```

## 默认搜索
```
https://cn.bing.com/search?q=$(CURRENT_WORD)
```

## 常用插件

### NppExec
* 在nppexec中不能用交互式命令,所以在tools.bat中补全
* 暗色需要在NppExec-->Advanced Options-->Blackgroundcolor设置为 66 66 66
* 插件下也有个style.css文件,有需要自己改

```shell
npp_save
if $(ext_part) ==.java goto java
if $(ext_part) ==.bat goto bat
exit

:java
env_set java_home=d:\tools\ideaic\jbr
env_set path=$(sys.path);$(sys.java_home)\bin
cd $(current_directory)
javac -encoding UTF-8 -d d:\workspaces\.tmp $(file_name)
cd d:\workspaces\.tmp
//cls
java $(name_part)
exit

:bat
cd $(current_directory)
$(FILE_NAME)
exit

exit
```

### 修改hosts文件
```
npp_open C:\WINDOWS\system32\drivers\etc\hosts
exit
```

### NppMarkdown Panel
* ~~下载NppMarkdownPanel.css,已优化的深色模式css,放在NppMarkdownPanel插件文件夹内~~

 ~~插件->NppMarkdownPanel->Edit Setting->CSS file(选择NppMarkdownPanel.css)(大小看屏幕)~~
* 新版已自带深色模式,调节大小就可以