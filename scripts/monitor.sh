pid=$(ps -ef |grep search_ | grep -v grep |awk '{print $2}')
top -pid $pid
