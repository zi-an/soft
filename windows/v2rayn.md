# v2rayn

## 软件起源
* https://github.com/2dust/v2rayn/releases
* xray的ui界面是v2rayn,v2rayn的core一般使用xray
* 自带路由规则已经很不错了


## 订阅地址
### 网站
* https://v2rayshare.com/

### 订阅地址
```
https://v2rayshare.com/wp-content/uploads/2023/10/20231001.txt
```
* 其中年月日都会变,需要自己调整
* 网络有时候打不开,就用手机流量打开后复制base64的编码再导入到v2rayn

### 快捷键
* ctrl+a 选中全部
* ctrl+t 选中的服务器中测速,一般把测速结果拉到前面
* enter 激活服务器
* del 删除选中的服务器

# 连接
* 配置文件config.json

|端口|协议|客户端|
|-|-|-|
|10808|socks|本地-默认|
|10809|http|本地-默认|
|10808|socks5|局域网-需开启|
|10809|http|局域网-需开启|

### 连接软件
* 浏览器插件 SwitchyOmega
* 手机http代理或者pac自动配置
* 手机软件https://github.hscsec.cn/2dust/v2rayNG/releases



nginx配置
* windows的root配置在location下,这里配置的是server下的
* proxy_ssl_server_name on;proxy_ssl_session_reuse off;防止ssl无法握手问题
* proxy_store永久存储到本地
* location / 匹配路径代理,if当没有文件时代理
```
root html;
proxy_ssl_server_name on;
proxy_ssl_session_reuse off;
proxy_store on;proxy_temp_path html;#proxy_set_header Accept-Encoding '';
location / { if (!-e \$request_filename){proxy_pass https://v2rayshare.com;}}
```

## 批量获取脚本
| https://v2rayshare.com/wp-content/uploads/2023/10/20231015.txt
```
curl http://7.mm/wp-content/uploads/2023/10/202310[01-15].txt
```

## nginx开启关闭脚本
```
taskkill /f /im nginx.exe
CD C:\Users\root\Desktop\n24\;nginx.exe
```

## 历史记录
* 有的已经被删除无法获取,测试结果是只保留最近两个月的记录
```
curl http://7.mm/wp-content/uploads/2023/01/202301[01-31].txt
curl http://7.mm/wp-content/uploads/2023/02/202302[01-30].txt
curl http://7.mm/wp-content/uploads/2023/03/202303[01-31].txt
curl http://7.mm/wp-content/uploads/2023/04/202304[01-30].txt
curl http://7.mm/wp-content/uploads/2023/05/202305[01-31].txt
curl http://7.mm/wp-content/uploads/2023/06/202306[01-30].txt
curl http://7.mm/wp-content/uploads/2023/07/202307[01-31].txt
curl http://7.mm/wp-content/uploads/2023/08/202308[01-31].txt
curl http://7.mm/wp-content/uploads/2023/09/202309[01-30].txt
curl http://7.mm/wp-content/uploads/2023/10/202310[01-15].txt
```