#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 19:58:28
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 20:11:37

import requests
import json
import os

from libs.utils.process import run_in_subprocess
from downloader.imgDowloader import imgDownload

def HaixiuzuTask(url):
    resp = requests.get(url)
    jsonStr = json.loads(resp.text.encode('utf-8'))
    jsonAttrs = json.loads(jsonStr)
    if 'imgs' in jsonAttrs.keys():
        for imgUrl in jsonAttrs['imgs']:
            head, tail = os.path.split(imgUrl)
            filePath = '/Users/darrenxyli/playground/github/greedy/{filename}'.format(filename=tail)
            run_in_subprocess(imgDownload, url=imgUrl, filepath=filePath).join()
    return True