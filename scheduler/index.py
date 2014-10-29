#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: darrenxyli <www.darrenxyli.com>
# @Date:   2014-10-29 01:32:29
# @Last Modified by:   darrenxyli
# @Last Modified time: 2014-10-29 18:21:32

from flask import Flask
from flask.ext import restful
from flask.ext.restful import reqparse
from libs.db.redis.taskDB import taskDB

r = taskDB().connection

app = Flask(__name__)
api = restful.Api(app)

parser = reqparse.RequestParser()
parser.add_argument('data', type=str)
parser.add_argument('project', type=str)


class taskPost(restful.Resource):

    def get(self):
        return "POST ONLY", 400

    def post(self):
        args = parser.parse_args()
        data = args['data']
        project = args['project']
        r.sadd(project, data)
        return data, 200


class taskGet(restful.Resource):

    def get(self, project):
        return r.spop(project), 200

api.add_resource(taskPost, '/hive')
api.add_resource(taskGet, '/hive/<string:project>')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3344, debug=True)
