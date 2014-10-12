# Scrapy settings for Sunflower project
#
# For simplicity, this file contains only the most important settings by
# default. All the other settings are documented here:
#
#     http://doc.scrapy.org/en/latest/topics/settings.html
#

BOT_NAME = 'Sunflower'

SPIDER_MODULES = ['Sunflower.spiders']
NEWSPIDER_MODULE = 'Sunflower.spiders'

SCHEDULER_ORDER = 'BFO'
DEFAULT_ITEM_CLASS = 'Sunflower.items.SunflowerItem'

ITEM_PIPELINES = {'Sunflower.pipelines.SunflowerPipeline':1000}

# Download Delay, for not download too quick
DOWNLOAD_DELAY = 2

# random modify the download delay
RANDOMIZE_DOWNLOAD_DELAY = True

# UA
USER_AGENT = ''

# warning deadline of memory usage
MEMUSAGE_WARNING_MB = 2048

# waning notification
MEMUSAGE_NOTIFY_MAIL = ['darren.xyli@gmail.com']

# download timeout
DOWNLOAD_TIMEOUT = 60

# log level
LOG_LEVEL = 'INFO'
