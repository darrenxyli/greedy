var NodeMongoObj = require('../../Hive/nodemongo').NodeMongoObj;

var mongodb = new NodeMongoObj('127.0.0.1', 27017, 'test', "fuck");
var db = mongodb.nodeMongoInsert({name:'xinyangli', value:'xin'});
var a = 0;
