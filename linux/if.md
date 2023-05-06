https://blog.csdn.net/yippvivian/article/details/119566029



now=`md5sum 123`
last=`md5sum 456`
if [ ${now:0:32} == ${now:0:32} ] ;then
 echo "same"
else
 echo "diff"
fi