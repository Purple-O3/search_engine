import random
import requests
import json
import datetime


for i in range(0, 1000000000):
    #li = [{"term":"hlrqjaps","union":True,"inter":False},{"term":"oazuvpjq","union":True,"inter":False},{"term":"rxbkjtnl","union":False,"inter":True}] 
    li = [{"term":"hlrqjaps","union":True,"inter":False}] 
    data = {"retreive_terms": li, "title_must":"szv","price_start":0,"price_end":10}
    data = json.dumps(data)
    ret = requests.post("http://127.0.0.1:7788/retrieve", data)
    if i % 100 == 0:
        print(datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S'))
        print(data)
        print(ret.text)
