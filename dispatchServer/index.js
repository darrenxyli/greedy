var express = require('express');
var app = express();
var redisClient = require("./noderedis");

client = new redisClient(6379, '199.231.211.82', {});

app.post('/hive', function (req, res) {
	saveType = req.query.urlType;
	saveUrl = req.query.urlContent;
    client.NodeRedisUpdate(saveType, saveUrl);
    console.log(client.NodeRedisGet(saveType));
    res.send("");
});

app.listen(3000);