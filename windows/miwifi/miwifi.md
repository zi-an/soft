# CR8808

## 设备json
[{"mac":"2A:83:D0:58:88:7D","tag":2,"name":"苹果12mini","ip":"192.168.10.19"},{"mac":"00:16:11:00:8D:B4","tag":2,"name":"wifi网卡","ip":"192.168.10.8"},{"mac":"D4:35:38:6C:CB:59","tag":2,"name":"小米音箱","ip":"192.168.10.10"},{"mac":"64:2C:0F:CA:58:63","tag":2,"name":"Y5s","ip":"192.168.10.15"},{"mac":"58:B6:23:78:79:37","tag":2,"name":"智能插座","ip":"192.168.10.4"},{"mac":"00:24:54:96:A5:78","tag":2,"name":"三星主板网卡","ip":"192.168.10.9"},{"mac":"28:16:7F:A6:B9:09","tag":2,"name":"红米note8","ip":"192.168.10.18"},{"mac":"9C:2E:A1:28:4A:77","tag":2,"name":"红米5","ip":"192.168.10.14"},{"mac":"4C:63:71:0B:44:EB","tag":2,"name":"K30","ip":"192.168.10.17"},{"mac":"78:8B:2A:C7:A0:8B","tag":2,"name":"chuangmi_camera_ipc020","ip":"192.168.10.3"},{"mac":"70:BB:E9:BE:09:75","tag":2,"name":"小米8青春版","ip":"192.168.10.13"},{"mac":"2C:60:0C:77:D9:63","tag":2,"name":"K610D","ip":"192.168.10.5"},{"mac":"02:09:00:00:00:00","tag":2,"name":"209","ip":"192.168.10.209"},{"mac":"AE:E1:95:C4:89:31","tag":2,"name":"苹果12mini","ip":"192.168.10.45"},{"mac":"34:1C:F0:C9:C2:5E","tag":2,"name":"K30S","ip":"192.168.10.11"},{"mac":"74:4C:A1:6A:23:07","tag":2,"name":"redmibook","ip":"192.168.10.7"},{"mac":"4C:63:71:09:A5:BA","tag":2,"name":"K30","ip":"192.168.10.12"},{"mac":"D4:35:38:45:32:B3","tag":2,"name":"CR6608","ip":"192.168.10.2"}]

## 拨号
* 18719189027/189027

## wifi
* 2.4G:x/xxxxxxxx
* 5G:∞/xxxxxxxx

## 静态DHCP
* 浏览器调试中复制powershell命令,修改最后的data,可以urlencode,也可以直接json
* 修改data时注意大小引号,json中的tag没用也没影响
* 也可以用来修改已固定的记录,甚至直接把上述设备json接入
* 部分功能用浏览器的调试模式强制改参数
```powershell
$session = New-Object Microsoft.PowerShell.Commands.WebRequestSession
$session.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
$session.Cookies.Add((New-Object System.Net.Cookie("__guid", "23904762.1721885128700282000.1679089760279.2747", "/", "192.168.10.1")))
$session.Cookies.Add((New-Object System.Net.Cookie("monitor_count", "1", "/", "192.168.10.1")))
Invoke-WebRequest -UseBasicParsing -Uri "http://192.168.10.1/cgi-bin/luci/;stok=d3dbf991fd1043faaae8b291cdbd1787/api/xqnetwork/mac_bind" `
-Method "POST" `
-WebSession $session `
-Headers @{
"Accept"="application/json, text/javascript, */*; q=0.01"
  "Accept-Encoding"="gzip, deflate"
  "Accept-Language"="zh-CN,zh;q=0.9"
  "Origin"="http://192.168.10.1"
  "Referer"="http://192.168.10.1/cgi-bin/luci/;stok=d3dbf991fd1043faaae8b291cdbd1787/web/prosetting/dhcpipmacband"
  "X-Requested-With"="XMLHttpRequest"
} `
-ContentType "application/x-www-form-urlencoded; charset=UTF-8" `
-Body 'data=[{"ip":"192.168.10.209","mac":"02:09:00:00:00:00","name":"777"}]'
      -----------------------------------------------------------------------
```

## 旧版固件
* 旧版(6.2.11)有自定义Hosts,设置完后在更新依旧生效(APP上看不到此功能)
* 从新版到旧版会出现连上WiFi上不了网的问题,hosts设置后记得升级
* 不建议使用(6.2.220),会有插件问题
```
https://cdn.cnbj1.fds.api.mi-img.com/xiaoqiang/rom/cr8808/miwifi_cr8808_firmware_9d216_6.2.11.bin
https://cdn.cnbj1.fds.api.mi-img.com/xiaoqiang/rom/cr8808/miwifi_cr8808_firmware_0fbd7_6.2.147.bin
https://cdn.cnbj1.fds.api.mi-img.com/xiaoqiang/rom/cr8808/miwifi_cr8808_firmware_a3144_6.2.220.bin
``` 

# AX1800

## 部分功能可以复制CR8808的地址,比如设置固定IP
## Mesh未解决
## 无自定义Hosts

```
https://cdn.cnbj1.fds.api.mi-img.com/xiaoqiang/rom/ra71/miwifi_ra71_firmware_e8042_1.0.88.bin
https://cdn.cnbj1.fds.api.mi-img.com/xiaoqiang/rom/ra71/miwifi_ra71_firmware_f91f7_1.0.93.bin
```

* cr6608固件
* 但是AX1800不能刷回CR6608
```
https://cdn.cnbj1.fds.api.mi-img.com/xiaoqiang/rom/cr6608/miwifi_cr6608_all_c11bb_1.0.100.bin
```

192.168.1.1/cgi-bin/telnetenable.cgi?telnetenable=1&key=74E29A3A3443


D4:35:38:BF:98:19