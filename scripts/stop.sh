pid=$(ps -ef |grep search_engine | grep -v grep |awk '{print $2}')
kill $pid
