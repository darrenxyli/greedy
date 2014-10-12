var NodeRedisObj = require('../..//Hive/noderedis').NodeRedisObj;

var client = new NodeRedisObj(6379, '127.0.0.1', {});
client.NodeRedisUpdate('name1', 'xinyangli');
client.NodeRedisUpdate('name1', 'd');
client.NodeRedisUpdate('name1', 'f');
client.NodeRedisUpdate('name1', 'x');
client.NodeRedisGet('name1');
console.log(good);