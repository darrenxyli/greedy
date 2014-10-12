# !/usr/bin/python
# -*-coding:utf-8-*-
'''
This whip means find any free url in redis and
whip them and find captain to work
'''

from captain import start_bee_engine
import Hive.pyredis as pyredis


class Whip(object):

    __client = None

    def __init__(self):
        self.__client = pyredis.pyRedisDbObject()


    # @Method: search()
    # @Description: find any free url from redis
    def search(self, Set):
        return self.__client.get(Set)


    # @Method: fight()
    # @Description: send free url to bee to crawl.
    def fight(self):

        # Step 1: search
        starturl = self.search('free_table')

        if starturl is not None:

            # Step 2: exchange item to running table
            # print starturl
            # self.__client.save('running_table', starturl)

            # Step 3: add celery, whip bee to run
            start_bee_engine.delay(starturl)
