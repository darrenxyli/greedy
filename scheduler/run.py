#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 18:43:37
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 20:29:40

import time

from libs.utils.config import Config
from libs.utils.process import run_in_subprocess
from api import app


class g(object):
    config = Config().getSettings()


def runAPI(g=g):
    host = '0.0.0.0'
    port = g.config['SCHEDULER_PORT']
    app.run(host=host, port=port, debug=True)


def allRun():
    threads = []
    threads.append(run_in_subprocess(runAPI, g=g))

    while True:
        try:
            time.sleep(10)
        except KeyboardInterrupt:
            break

    for thread in threads:
        thread.join()

if __name__ == '__main__':
    allRun()
