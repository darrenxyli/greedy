var express = require('express');
var app = express();
var redisClient = require("./noderedis");
var config = require("../config.json");

client = new redisClient(config.REDIS_PORT_TASK, config.REDIS_IPADDRESS_TASK, {}) ;

app.post('/hive', function (req, res) {
    saveType = req.query.urlType;
    saveUrl = req.query.urlContent;
    // client.NodeRedisUpdate(saveType, saveUrl);
    console.log(client.NodeRedisGet(saveType));
    res.status(404);
    res.send("good");
});

app.listen(3000) ;
