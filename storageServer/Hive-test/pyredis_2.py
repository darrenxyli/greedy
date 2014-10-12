import redis
import json
import re

pool = redis.ConnectionPool(host='localhost', port=6379, db=1)
r = redis.Redis(connection_pool=pool)


def parse(respCmtJson):
    respCmtJson = re.sub(r"(,?)(\w+?)\s+?:", r"\1'\2' :", respCmtJson)
    respCmtJson = respCmtJson.replace("'", "\"")
    print respCmtJson
    # print type(respCmtJson)
    # temp = json.dumps(respCmtJson)
    # cmtDict = json.loads(temp);
    # print cmtDict
    # print type(cmtDict)

for x in range(1, 1000000):
    temp = r.spop('free_table')
    if temp is not None:
        parse(temp)
    tmp = r.spop('runnig_table')
    if tmp is not None:
        parse(tmp)
