#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 17:53:40
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 18:18:05

import redis

from libs.utils.config import Config


class taskDB(object):

    def __init__(self):
        config = Config().getSettings()

        pool = redis.ConnectionPool(
            host=config["REDIS_IPADDRESS_TASK"],
            port=config["REDIS_PORT_TASK"],
            db=config["REDIS_DBNAME_BROKER"]
        )
        self.connection = redis.Redis(connection_pool=pool)

    def connection(self):
        return self.connection
