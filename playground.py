#!/usr/bin/env python
# encoding: utf-8

from captain import start_bee_engine

for x in range(1, 3):
    start_bee_engine.delay('http://www.chinanews.com/ty/2014/04-30/6122731.shtml')
