#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 20:42:41
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 23:14:14
import time

from libs.config import Config
from libs.taskDB import taskDB
from libs.process import run_in_subprocess


class g(object):
    config = Config().getSettings()
    taskDBConn = taskDB().connection
    downloadPath = "~/download/"

def runAPI(g=g):
    from scheduler.api import app
    host = '0.0.0.0'
    port = g.config['SCHEDULER_PORT']
    app.run(host=host, port=port, debug=True)

def runScheduler(g=g):
    from scheduler.scheduler import Scheduler
    schedu = Scheduler(g=g)
    schedu.run()

def runAllService():
    threads = []
    #threads.append(run_in_subprocess(runAPI, g=g))
    threads.append(run_in_subprocess(runScheduler, g=g))

    while True:
        try:
            time.sleep(10)
        except KeyboardInterrupt:
            break

    for thread in threads:
        thread.join()

if __name__ == '__main__':
    runAllService()
