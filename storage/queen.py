# !/usr/bin/python
# -*-coding:utf-8-*-
'''
the controler of whole scraper engine like Queen Bee.
'''
import parse
import os
import utils
from whip import Whip

# CONSTANTS
CONFIG_PATH = 'config.json'
SETTING = parse.Setting(CONFIG_PATH).config
ROOT = os.path.dirname(os.path.abspath(__file__))


# Alarms
class Alarm(Exception):
    pass


def alarm_handler(signo, frame):
    raise Alarm()


# init the environment
def initEnvironment():
    print utils.green("\nStart init environment...")
    os.system('bash initEnvironment.sh')

# bud receive starturl and make many flowers
def initScrapyTask():
    # put project into scrapyd
    print utils.green('\nStart URL list status:\n')
    for starturl in SETTING['SUNFLOWER_START_URLS']:
        CMD = 'curl http://' + SETTING['SCRAPYD_ADDRESS'] + '/schedule.json -d project=Sunflower -d spider=flower -d starturl=' + starturl
        os.system(CMD)
    print utils.green('Spiders are running now...')

# start whip to work
def startWhip():
    worker = Whip()
    print utils.red('\nCaptain is whipping ...')
    print utils.red('Press CTRL+C to stop whipping')
    print utils.red('Or run "bash stopEnvironment.sh" to kill all process')
    while(1):
        worker.fight()

# main funciton
def main():
    initEnvironment()
    initScrapyTask()
    startWhip()


if __name__ == '__main__':
    main()
