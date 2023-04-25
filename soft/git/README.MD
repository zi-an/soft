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
```
git clone git@github.com:zi-an/soft
```

## git config
```
git config --global user.email zian@zi.an
git config --global user.name zian
git config --global gui.encoding utf-8
git config --global -l
```