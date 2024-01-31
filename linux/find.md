find / ! -path "/proc/*" ! -path "/sys/*" -exec touch -m -d "2010-10-10 10:10:10" {} \;


find / ! -path "/proc/*" ! -path "/sys/*" -type f -exec touch -m -d "2010-10-10 10:10:10" {} \;

-maxdepth 1 目录最大深度1
-mtime 内容修改时间 
-ctime 状态修改时间
-atime 最近访问时间 ls --time=atime
! -path 规则要以输出内容为准,前缀不匹配不行

http://linux.51yip.com/search/find

wget