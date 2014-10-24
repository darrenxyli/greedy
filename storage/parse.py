# !/usr/bin/python
# -*-coding:utf-8-*-
import json

# @class: Setting
# @Description: parse the configFile
class Setting(object):

    __config = None

    def __init__(self, configpath):
        configFile = open(configpath).read()
        configObj = json.loads(configFile)
        self.config = configObj


    def getConfig(self):
        return self.__config
