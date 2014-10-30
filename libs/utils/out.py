#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 19:38:27
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 19:39:18

from colorama import Fore


def red(s):
    return Fore.RED + s + Fore.RESET


def green(s):
    return Fore.GREEN + s + Fore.RESET
