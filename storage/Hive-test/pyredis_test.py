# !/usr/bin/python
# -*-coding:utf-8-*-
import sys
sys.path.append("../../Hive/")
import pyredis

client = pyredis.pyRedisDbObject('127.0.0.1', 6379, 0)
client.save('standby', 'xinyangli')
client.save('standby', 'x')
client.save('standby', 'a')
client.save('standby', 'e')
client.save('standby', 'g')
print client.get('standby')
client.ismember('standby', 'xinyangli')
