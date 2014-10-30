#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 17:55:50
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 20:46:33

import os
import json


class Config(object):
    settings = {}

    def __init__(self):
        filePath = os.path.dirname(
            os.path.abspath(__file__)) + "/../config.json"
        with open(filePath) as f:
            self.settings = json.load(f)

    def get(self, key):
        return self.settings[key]

    def getSettings(self):
        return self.settings
