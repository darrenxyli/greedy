# !/usr/bin/python
# -*-coding:utf-8-*-
import os
import signal
import subprocess
from celery.task import task

# CONSTANT DEFIEND
NODE_PATH = 'node'
CASPER_PATH = 'casperjs'
PYTHON_PATH = 'python'

# Alarms


class Alarm(Exception):
    pass


def alarm_handler(signo, frame):
    raise Alarm()


# @Method: start_scraping_engine()
# @param startUrl: sent by whip, and send to bee to scrape.
# @return returncode
@task
def start_scraping_engine(startUrl, task_id=None, timeout=0, proxies=[], **kwargs):
    ROOT = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'Bee')
    BEE_PATH = ROOT + '/bee.js'
    if timeout:
        signal.signal(signal.SIGALRM, alarm_handler)
        signal.alarm(timeout)

    args = [NODE_PATH, BEE_PATH]

    try:
        stdout = kwargs.get('stdout', None)
        proc = subprocess.Popen(args,
                                stdin=subprocess.PIPE,
                                stdout=stdout,
                                preexec_fn=os.setsid,
                                cwd=ROOT)

        proc.communicate(input=startUrl + '\n')

        if timeout:
            signal.alarm(0)

    except Alarm:
        print 'Try to kill the scrape process, pid:%s' % (proc.pid)
        os.killpg(proc.pid, signal.SIGKILL)

    proc.wait()
    return proc.returncode


@task
def start_bee_engine(startUrl, task_id=None, timeout=0, proxies=[], **kwargs):
    ROOT = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'Bee')
    BEE_PATH = ROOT + '/bee.py'
    if timeout:
        signal.signal(signal.SIGALRM, alarm_handler)
        signal.alarm(timeout)

    args = [PYTHON_PATH, BEE_PATH, '--startUrl', startUrl]

    try:
        stdout = kwargs.get('stdout', None)
        proc = subprocess.Popen(args,
                                stdin=subprocess.PIPE,
                                stdout=stdout,
                                preexec_fn=os.setsid,
                                cwd=ROOT)

        if timeout:
            signal.alarm(0)

    except Alarm:
        print 'Try to kill the scrape process, pid:%s' % (proc.pid)
        os.killpg(proc.pid, signal.SIGKILL)

    proc.wait()
    return proc.returncode
