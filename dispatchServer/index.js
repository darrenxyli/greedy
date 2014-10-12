var express = require('express');
var app = express();
var redisClient = require("./noderedis");

client = new redisClient();

app.post('/hive', function (req, res) {
	saveType = req.query.urlType;
	saveUrl = req.query.urlContent;
    client.NodeRedisUpdate(saveType, saveUrl);
    res.send("");
});

app.listen(3000);