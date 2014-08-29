'''
The redis persistence layer for python
'''
import redis

# @Description: create a ConnectionPool
pool = redis.ConnectionPool(host='127.0.0.1', port=6379, db=1)
r = redis.Redis(connection_pool=pool)


'''
@class: pyRedisDbObject
@Method: __init__(), save(), get()
@Description: consitentce layer of redis in python.
'''
class pyRedisDbObject(object):
    def __init__(self):
       self.client = r

    # @Description: I use set to save url to make every url be unique.
    def save(self, Set, Value):
        self.client.sadd(Set, Value)

    # @Description: use pop method to get url, remove item once find a free url
    def get(self, Set):
        return self.client.spop(Set)
