# sqlite3

## 与CentIS7版本同步3.7.17
* https://www.sqlite.org/2013/sqlite-shell-win32-x86-3071700.zip 

## 网站历史记录查询
* 旧版软件下载 https://www.sqlite.org/chronology.html

## 基本语法
* https://www.runoob.com/sqlite/sqlite-tutorial.html

|语法|作用|备注|
|-|-|-|
|.help|查看帮助|基本语法都有|
|.database|查看数据库||
|.table|查看表|.table %h%|
||||

# java使用内存模式
* cmd命令行直接sqlite3不带参数进入的就是内存模式
* 仅在内存中使用,非文件,多连接池可能会连不上,单线程正常