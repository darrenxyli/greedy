#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 17:00:26
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 20:11:47

from rq import (Queue, Connection)
from libs.utils.config import Config
from libs.db.redis.taskDB import taskDB

config = Config().getSettings()

taskConn = taskDB().connection

with Connection(taskConn):
    q = Queue('low')
    from scheduler.tasks.haixiuzuTask import HaixiuzuTask
    result = q.enqueue(
        HaixiuzuTask,
        'http://{ip}:{port}/{version}/{project}'.format(
            ip=config['SCHEDULER_IPADDRESS'],
            port=config['SCHEDULER_PORT'],
            version='hive',
            project='haixiuzu'
        )
    )
