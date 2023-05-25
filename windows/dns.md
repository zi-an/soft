# DNS解析

# DNS-type相关
* https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/nslookup-set-type
* 一般使用A就行了
```
A: Specifies a computer's IP address.
ANY: Specifies a computer's IP address.
CNAME: Specifies a canonical name for an alias.
GID Specifies a group identifier of a group name.
HINFO: Specifies a computer's CPU and type of operating system.
MB: Specifies a mailbox domain name.
MG: Specifies a mail group member.
MINFO: Specifies mailbox or mail list information.
MR: Specifies the mail rename domain name.
MX: Specifies the mail exchanger.
NS: Specifies a DNS name server for the named zone.
PTR: Specifies a computer name if the query is an IP address; otherwise, specifies the pointer to other information.
SOA: Specifies the start-of-authority for a DNS zone.
TXT: Specifies the text information.
UID: Specifies the user identifier.
UINFO: Specifies the user information.
WKS: Describes a well-known service.
```
# nslookup 与 dig
|命令|nslookup|dig||
|-|-|-|-|
|Windows|√|x||
|Linux|√|√|都要安装bind-utils|
|总信息|简单|详细|dig显示type,ttl等<tr><td>资料</td><td colspan="3">https://www.isc.org/bind/<br>https://downloads.isc.org/isc/bind9/9.18.13/doc/arm/Bv9ARM.pdf</td></tr>

# nslookup
* 参考 https://learn.microsoft.com/zh-cn/windows-server/administration/windows-commands/nslookup-server?source=recommendations

|命令|作用||
|-|-|-|
|nslookup qq.com|使用系统默认的服务器解析qq.com||
|nslookup qq.com 223.5.5.5|用阿里dns解析qq.com||
|nslookup - 223.5.5.5|交互模式解析服务器为223.5.5.5|交互阿里的dns解析|
|nslookup - 223.5.5.5<br>set type=mx<br>qq.com|查qq.ccom域名的MX记录||


# dig
* yum install bind-utils -y

|命令|作用||
|-|-|-|
|dig qq.com|使用系统默认的服务器解析qq.com||
|dig @223.5.5.5 qq.com|指定192..200服务器解析qq.com|
dig @223.5.5.5 qq.com mx|查qq.ccom域名的MX记录||

# DNSBench

## 下载
* https://www.grc.com/files/DNSBench.exe

## 配置
* dnsbench.ini
```
echo 192.168.10.1 xiaomi > dnsbench.ini
echo 192.168.10.200 javadns >> dnsbench.ini
echo 223.5.5.5 alidns >> dnsbench.ini
echo 119.29.29.29 qqdns >> dnsbench.ini
```

##
* 谷歌DoH https://developers.google.com/speed/public-dns/docs/doh
* 阿里DoH https://alidns.com/articles/6018321800a44d0e45e90d71

# 摘要

* https://dns.google/dns-query?
* https://dns.google/resolve?

* /dns-query? 可以用get/post需要参数dns={base64url},算法内含特殊字符 不推荐
* /resolve?   返回为json,直接在添加name=xxx.xxx即可，方便

* https://dns.google/resolve?name=qq.com
* https://dns.google/dns-query?dns=uGkBAAABAAAAAAAAB2FsaWJhYmEDY29tAAABAAE


