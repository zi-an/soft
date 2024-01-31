# hugo

| 时间:2023年5月5日

* 建议使用hugo extended (解压hugo_extended_0.111.3_windows-amd64.zip把exe放到C:/Windows下)
* 主题hugo-book下载https://themes.gohugo.io/themes/hugo-book/
* 主题演示地址 https://hugo-book-demo.netlify.app/

# build shell
* 创建文件并把example运行
```
hugo new site mysite
cd mysite
git clone https://github.com/alex-shpak/hugo-book themes/hugo-book
cp ./themes/hugo-book/exampleSite/content.en/* ./content -r
hugo server --theme=hugo-book --buildDrafts
```

# public shell
* 修改 baseURL
* 使用主题生成静态页面文件
* 之后生成./public文件夹,复制到nginx即可以
``` 
hugo --theme=hugo-book 
```


```
git clone --branch=mysite git@github.com:zi-an/zi-an.github.io.git mysite
```


# 80端口测试项目
```
hugo server --contentDir=themes/hugo-book/exampleSite/content.en -p=80
```

content
>linux 1
>windows 4
>java 2
>tools 3