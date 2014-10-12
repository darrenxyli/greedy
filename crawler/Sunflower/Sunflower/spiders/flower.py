#!/usr/bin/python
#-*-coding:utf-8-*-

from scrapy.spider import BaseSpider
from scrapy.selector import HtmlXPathSelector
from scrapy.http import Request
from Sunflower.items import SunflowerItem


class DroneSpider(BaseSpider):
    name = "flower"
    allowed_domains = []
    start_urls = []

    def __init__(self, starturl, *args, **kwargs):
        super(DroneSpider, self).__init__(*args, **kwargs)
        self.start_urls.append(starturl)

    def parse(self, response):
        hxs = HtmlXPathSelector(response)
        links = hxs.select(u'//a/img/@src').extract()
        for link in links:
            item = SunflowerItem()
            item['URL'] = link
            yield item
            yield Request(url=link, callback=self.parse)
