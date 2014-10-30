#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 19:58:28
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-30 16:37:56

import requests
import json
import os

from libs.process import run_in_thread
from downloader.imgDowloader import imgDownload

def Task(url, path):
    resp = requests.get(url)
    jsonStr = json.loads(resp.text.encode('utf-8'))
    jsonAttrs = json.loads(jsonStr)
    if 'imgs' in jsonAttrs.keys():
        for imgUrl in jsonAttrs['imgs']:
            head, tail = os.path.split(imgUrl)
            filePath = '{path}{filename}'.format(path=path, filename=tail)
            run_in_thread(imgDownload, url=imgUrl, filepath=filePath).join()
    return True