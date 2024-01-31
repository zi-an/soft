// ==UserScript==
// @name         我的
// @namespace    http://tampermonkey.net/
// @version      0.1
// @author       You
// @description  try to take over the world!
// @match        */enter/index.html
// @match        */shipin/*
// @match        */thread.php*
// @match        */read.php*
// @match        *5.mm:9090
// @match        */cgi-bin/luci/web
// @icon         https://cn.bing.com/favicon.ico
// @grant        none
// ==/UserScript==

(function() {
    'use strict';
    var url = window.location.href;
    if(url.indexOf('enter\/index.html') >=0){
        window.location.href=url.replace('enter\/index.html','index\/home.html');
    }
    if(url.indexOf('shipin\/detail') >=0){
        setTimeout( function(){close_discor();},1000 );
    }
    if(url.indexOf('orderway=postdate') <0 && url.indexOf('thread.php') >0){
        window.location.href=url+'&orderway=postdate';
    }
    if(url.indexOf('read.php') >0){
        document.getElementsByClassName('my_wrap')[0].style.display='none';
    }
    if(url.indexOf('9090') >=0){
        setTimeout( function(){document.getElementById("login-user-input").value="root";
                               document.getElementById("login-password-input").value="1";
                               document.getElementById("login-button").click();},500 ) ;
    }
    if(url.indexOf('/cgi-bin/luci/web') >=0){
        document.getElementById("password").value="******";
        document.getElementById("btnRtSubmit").click();
    }
})();
