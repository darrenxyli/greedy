#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 19:47:19
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 19:47:50


def run_in_thread(func, *args, **kwargs):
    from threading import Thread
    thread = Thread(target=func, args=args, kwargs=kwargs)
    thread.daemon = True
    thread.start()
    return thread


def run_in_subprocess(func, *args, **kwargs):
    from multiprocessing import Process
    thread = Process(target=func, args=args, kwargs=kwargs)
    thread.daemon = True
    thread.start()
    return thread