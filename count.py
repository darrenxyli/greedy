#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 16:59:40
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 17:00:17

import requests

def count_words_at_url(url):
    resp = requests.get(url)
    return len(resp.text.split())