
* https://www.grc.com/files/DNSBench.exe

* 谷歌DoH https://developers.google.com/speed/public-dns/docs/doh
* 阿里DoH https://alidns.com/articles/6018321800a44d0e45e90d71

# 摘要

* https://dns.google/dns-query?
* https://dns.google/resolve?

* /dns-query? 可以用get/post需要参数dns={base64url},算法内含特殊字符 不推荐
* /resolve?   返回为json,直接在添加name=xxx.xxx即可，方便

* https://dns.google/resolve?name=qq.com
* https://dns.google/dns-query?dns=uGkBAAABAAAAAAAAB2FsaWJhYmEDY29tAAABAAE

