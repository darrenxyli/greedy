#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 17:00:26
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 18:20:21

from rq import (Queue, Connection)
from libs.utils.config import Config
from libs.db.redis.taskDB import taskDB
config = Config().getSettings()

taskConn = taskDB().connection

with Connection(taskConn):
    q = Queue('low')
    from count import count_words_at_url
    for x in xrange(3):
        result = q.enqueue(
            count_words_at_url, 'http://nvie.com')
