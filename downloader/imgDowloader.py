#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 16:59:40
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 20:03:59

import requests

def imgDownload(url, filepath):
    r = requests.get(url, stream=True)
    if r.status_code == 200:
        with open(filepath, 'wb') as f:
            for chunk in r.iter_content(1024):
                f.write(chunk)