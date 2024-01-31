# 原生ASOP
## ROM
* https://crdroid.net/downloads
* https://sourceforge.net/projects/crdroid/files/
* https://havoc-os.com/download
* https://sourceforge.net/projects/havoc-os/files/
* sourceforge.net上部分文件会下载不了,用其它文件拼接后缀

https://sourceforge.net/projects/crdroid/files/

## TWRP & ROM
* https://dl.twrp.me/rolex/twrp-3.3.1-0-rolex.img.html(crdroid7.6/android11)

* https://dl.twrp.me/rosy/twrp-3.7.0_9-0-rosy.img.html
* https://sourceforge.net/projects/crdroid/files/rosy/
* https://gigenet.dl.sourceforge.net/project/crdroid/rosy/7.x/crDroidAndroid-11.0-20210721-rosy-v7.8.zip
* https://gigenet.dl.sourceforge.net/project/crdroid/rosy/9.x/crDroidAndroid-13.0-20230216-rosy-v9.2.zip

 
## ADB
* MiFlash里都有
```
E:\mi\MiFlash2020-3-14-0\Source\ThirdParty\Google\Android
```

## 刷机
https://miuiver.com/how-to-flash-twrp/

1.解锁,关机,开机+音量下
2.刷入TWRP,如果没刷入rom会被重置
fastboot flash recovery twrp-3.3.1-0-rolex.img
fastboot boot twrp-3.3.1-0-rolex.img 
fastboot reboot
3.第二部后手机进入recovery,清空除SD卡外的操作,并刷入rom
4.重设联网验证与NTP服务器
``shell
adb shell settings put global captive_portal_http_url http://connect.rom.miui.com/generate_204

adb shell settings put global captive_portal_https_url https://connect.rom.miui.com/generate_204

adb shell settings put global ntp_server ntp.aliyun.com
``


## wifi debug
> 用miflash的adb版本只有39,不能直连开发者模式中打开的wifi调试,要下载新的
* https://dl.google.com/android/repository/platform-tools_r34.0.1-windows.zip?hl=zh-cn

> 用usb打开后的wifi调试,可以用旧版的adb连(Centos7的adb是更旧的31版,无法升级,只能用此方法连)
* https://mirrors.tuna.tsinghua.edu.cn/epel/7/x86_64/Packages/a/android-tools-20130123git98d0789-5.el7.x86_64.rpm
* rpm -i android-tools-20130123git98d0789-5.el7.x86_64.rpm

## 安卓13 
* adb pair 192.168.10.14:41857 匹配码
```
#查看应用
adb shell pm list package

adb shell dumpsys package com.android.launcher3

adb -d shell pm grant package android.permission.CAMERA
```
 
# 小米运动健康 
* com.mi.health
* (通知权限)android.service.notification.NotificationListenerService
* 以下均不可用,别试了
```
adb -d shell pm grant com.mi.health android.service.notification.NotificationListenerService
adb shell appops set --uid com.mi.health android.service.notification.NotificationListenerService allow
```

