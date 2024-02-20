# Git

## installed
* https://git-scm.com/download/win
* https://github.com/git-for-windows/git/releases/download/v2.40.0.windows.1/Git-2.40.0-64-bit.exe
* Installed by default

## sinicization
* cmd Open with administrator privileges
```
mkdir "C:\Program Files\Git\mingw64\share\git-gui\lib\msgs"
cd "C:\Program Files\Git\mingw64\share\git-gui\lib\msgs"
curl https://raw.githubusercontent.com/stayor/git-gui-zh/master/zh_cn.msg > zh_cn.msg
```

## git clone git
* 只要最新版本 *--depth 1*
* 只要某个分支 *--single-branch master*
```
git clone --depth 1 --single-branch master git@github.com:zi-an/soft
```
 

## git config
```
git config --global core.autocrlf false
git config --global user.email zian@zi.an
git config --global user.name zian
git config --global gui.encoding utf-8
git config --global -l
```
## github
* 如果git版本太旧会无法更新,需要更新软件
* 一个密钥只能一个用户用,一个密钥给多个用户用会导致之前的用户密钥全部失效,密钥只留给最后一个用户

 
