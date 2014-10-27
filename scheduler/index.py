#!/usr/bin/env python
# Website: http://www.darrenxyli.com
# Author: darrenxyli<darren.xyli.com>
# encoding: utf-8

from flask import Flask
from flask.ext import restful
from flask.ext.restful import reqparse
import os
import redis
import json

filePath = os.path.dirname(os.path.abspath(__file__)) + "/../config.json"
with open(filePath) as f:
    config = json.load(f)

pool = redis.ConnectionPool(
    host=config["REDIS_IPADDRESS_TASK"],
    port=config["REDIS_PORT_TASK"],
    db=config["REDIS_DBNAME_BROKER"]
)
r = redis.Redis(connection_pool=pool)

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
