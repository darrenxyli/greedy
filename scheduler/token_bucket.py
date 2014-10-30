#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-30 00:36:51
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-30 00:37:00

import time
try:
    import threading as _threading
except ImportError:
    import dummy_threading as _threading

class Bucket(object):
    '''
    traffic flow control with token bucket
    '''

    update_interval = 30
    def __init__(self, rate=1, burst=None):
        self.rate = float(rate)
        if burst is None:
            self.burst = float(rate)*10
        else:
            self.burst = float(burst)
        self.mutex = _threading.Lock()
        self.bucket = self.burst
        self.last_update = time.time()

    def get(self):
        now = time.time()
        if self.bucket >= self.burst:
            self.last_update = now
            return self.bucket
        bucket = self.rate * (now - self.last_update)
        self.mutex.acquire()
        if bucket > 1:
            self.bucket += bucket
            if self.bucket > self.burst:
                self.bucket = self.burst
            self.last_update = now
        self.mutex.release()
        return self.bucket

    def set(self, value):
        self.bucket = value

    def desc(self, value=1):
        self.bucket -= value