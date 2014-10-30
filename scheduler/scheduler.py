#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 17:00:26
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-30 02:28:51

import time
import logging

from rq import (Queue, Connection)
from task_queue import TaskQueue


class Scheduler(object):
    UPDATE_PROJECT_INTERVAL = 10
    LOOP_LIMIT = 1000
    LOOP_INTERVAL = 0.1
    ACTIVE_TASKS = 100
    INQUEUE_LIMIT = 0
    EXCEPTION_LIMIT = 3
    DELETE_TIME = 24 * 60 * 60
    UPDATE_TASKID_INTERVAL = 3

    def __init__(self, g):
        self.g = g
        self.config = g.config
        self.taskdb = g.taskDBConn
        # self.resultdb = resultdb
        self.downloadPath = g.downloadPath

        self._quit = False
        self._exceptions = 0
        self._last_update_project = 0
        self.taskId = 0
        self.projects = []
        self.task_queue = {}

    def _init_task_queue(self, project):
        self.task_queue[project] = TaskQueue(rate=0, burst=0)
        self.task_queue[project].rate = 0.2
        self.task_queue[project].burst = 3

    def _load_projects(self):
        self.projects = []
        for project in self.config['PROJECTS']:
            self.projects.append(project)
            if project not in self.task_queue:
                self._init_task_queue(project)
                print "[project {name}] taks queue initialed".format(name=project)
            # logging.debug("project: {name} loaded.".format(name=project))
            self._update_project(project)
            print "[project {name}] project loaded.".format(name=project)
        self._last_update_project = time.time()
        self.taskId += self.UPDATE_TASKID_INTERVAL

    def _push_rq(self, project, taskid):
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
        # logging.debug('[Project {project}]task {taskid} in work queue'.format(project=project, taskid=taskid))
        print '[Project {project}] task {taskid} in work queue'.format(project=project, taskid=taskid)

    def _update_projects(self):
        now = time.time()
        if self._last_update_project + self.UPDATE_PROJECT_INTERVAL > now:
            return
        for project in self.projects:
            self._update_project(project)
            # logging.debug("[project %s] project updated.", project)
            print "[project {project}] project updated.".format(project=project)
        self._last_update_project = now
        self.taskId += self.UPDATE_TASKID_INTERVAL

    def _update_project(self, project):
        for tId in xrange(self.taskId, self.taskId + self.UPDATE_TASKID_INTERVAL):
            self.task_queue[project].put(tId, 1)
        self.task_queue[project].check_update()

    def _start_control(self):
        with Connection(self.taskdb):
            for project in self.projects:
                taskId = self.task_queue[project].get()
                if taskId is not None:
                    self._push_rq(project, taskId)

    def quit(self):
        self._quit = True

    def run(self):
        # logging.info("loading projects")
        print "loading projects"
        self._load_projects()

        while not self._quit:
            try:
                time.sleep(self.LOOP_INTERVAL)
                self._update_projects()
                self._start_control()
            except KeyboardInterrupt:
                break
            except Exception as e:
                logging.exception(e)
                self._exceptions += 1
                if self._exceptions > self.EXCEPTION_LIMIT:
                    self.quit()
                    break
                continue

        print "scheduler exiting..."
