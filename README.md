# Windows Soft

## 大文件上传europa.zip
* 先上传记录信息.gitattributes
```
git lfs track "europa.zip"
git add .gitattributes
git commit -m "submit gitattributes"
git push -u origin windows 
```
* 再提交大文件
```
git add europa.zip
git commit -m "add europa.zip"
git push -u origin windows
```

* drivergenius.exe 为驱动精灵，在安装好有线/无线网卡驱动后使用