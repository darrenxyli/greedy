#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 17:00:26
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 23:16:35

import time

from rq import (Queue, Connection)


class Scheduler(object):
    UPDATE_PROJECT_INTERVAL = 5 * 60
    LOOP_LIMIT = 1000
    LOOP_INTERVAL = 0.1
    ACTIVE_TASKS = 100
    INQUEUE_LIMIT = 0
    EXCEPTION_LIMIT = 3
    DELETE_TIME = 24 * 60 * 60

    def __init__(self, g):
        self.g = g
        self.config = g.config
        self.taskdb = g.taskDBConn
        # self.resultdb = resultdb
        self.downloadPath = g.downloadPath

        self._quit = False
        self._exceptions = 0
        self.projects = []

    def _load_projects(self):
        self.projects = []
        for project in self.config['PROJECTS']:
            self.projects.append(project)
            print "project: {name} loaded.".format(name=project)

    def _start_control(self):
        with Connection(self.taskdb):
            for project in self.projects:
                q = Queue('{project}_task'.format(project=project))
                from .projects.haixiuzuProject import HaixiuzuPro
                result = q.enqueue(
                    HaixiuzuPro,
                    'http://{ip}:{port}/{version}/{project}'.format(
                        ip=self.config['SCHEDULER_IPADDRESS'],
                        port=self.config['SCHEDULER_PORT'],
                        version='hive',
                        project=project
                    ),
                    self.downloadPath
                )

    def quit(self):
        self._quit = True

    def run(self):
        print "loading projects"
        self._load_projects()

        while not self._quit:
            try:
                time.sleep(self.LOOP_INTERVAL)
                self._start_control()
            except KeyboardInterrupt:
                break
            except Exception as e:
                print e
                self._exceptions += 1
                if self._exceptions > self.EXCEPTION_LIMIT:
                    self.quit()
                    break
                continue

        print "scheduler exiting..."
