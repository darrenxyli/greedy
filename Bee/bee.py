#!/usr/bin/env python
# encoding: utf-8

# import the father path
import sys
sys.path.append('..')

# import packages
from goose import Goose
from goose.text import StopWordsChinese
from mongoengine import *
import click

class PostTable(Document):
    title = StringField(max_length=120, required=True)
    meta_des = StringField()
    meta_tags = StringField()
    content = StringField(required=True)
    url = URLField(verify_exists=True, unique=True)

@click.command()
@click.option('--startUrl', prompt='The url you want crawl', required=True)
def startBee(startUrl):
    url = startUrl
    g = Goose({'stopwords_class': StopWordsChinese})
    article = g.extract(url=url)

    if len(article.title) > 0 and len(article.cleaned_text) > 0:

        print 'TITLE:'
        print article.title
        print 'CONTENT:'
        print article.cleaned_text[:500]
        print 'META DES:'
        print article.meta_description
        print 'META KEYWORDS:'
        print article.meta_keywords
        print 'DOMAIN:'
        print article.domain
        print 'FINAL URL:'
        print article.final_url

        connect('postTable', host='mongodb://localhost/postDB')
        post = PostTable(
            title=article.title,
            meta_des=article.meta_description,
            meta_tags=article.meta_keywords,
            content=article.cleaned_text,
            url=article.final_url
        )
        post.save()

if __name__ == '__main__':
    startBee()
