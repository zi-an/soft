@echo off
TITLE Tools
:: color black
:: C:\Users\user\tools.bat
ECHO. ===================Tools=======================
echo   as     
echo   hosts  
echo   java   
echo   mysql
echo   redis
ECHO. ===============================================

set /p key="select:"
cls

if "%key%"=="as" goto as
if "%key%"=="java" goto java
if "%key%"=="hosts" goto hosts
if "%key%"=="mysql" goto mysql
if "%key%"=="redis" goto redis
goto end

:java
set JAVA_HOME=D:\tools\idea\jbr
set path=%JAVA_HOME%\bin\
echo "java_home is ready"
goto end 

:as
echo "arthas is start"
D:\tools\idea\jbr\bin\java -Dfile.encoding=utf-8 -jar D:\tools\arthas\arthas-boot.jar
goto end

:hosts
echo "please use administrator privileges "
echo "github https://github.com/jianboy/github-host/blob/master/hosts"
"C:\Program Files\Notepad\notepad++.exe" "C:\WINDOWS\system32\drivers\etc\hosts"
goto end

:mysql 
mysql -h 201.mm -u zian -pzian
goto end

:redis
redis -h 202.mm -a zian
goto end

:end 
