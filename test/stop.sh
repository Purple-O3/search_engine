pid=$(ps -ef |grep test_ | grep -v grep |awk '{print $2}')
kill $pid
