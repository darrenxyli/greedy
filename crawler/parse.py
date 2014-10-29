#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 00:37:42
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 01:36:09

import json


class Setting(object):

    __config = None

    def __init__(self, configpath):
        configFile = open(configpath).read()
        configObj = json.loads(configFile)
        self.config = configObj

    def getConfig(self):
        return self.__config
