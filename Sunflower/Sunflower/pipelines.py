# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: http://doc.scrapy.org/en/latest/topics/item-pipeline.html

import redis
pool = redis.ConnectionPool(host='127.0.0.1', port=6379, db=1)
r = redis.Redis(connection_pool=pool)


# @class: SunflowerPipeline
# @Description: save the url from spiders
class SunflowerPipeline(object):
    def process_item(self, item, spider):
        r.sadd('free_table', item['URL'])
