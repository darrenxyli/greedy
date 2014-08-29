/*
 * the redis persistence layer for nodejs
 */

/*
 * Import a third party pakage
 */
var redis = require("redis");

var NodeRedisObj = function NodeRedisObj(Port, Host, Options) {
    this.__port = Port ? Port : 6379;
    this.__host = Host ? Host :'127.0.0.1';
    this.__options = Options ? Options : {};
    this.__client = redis.createClient(this.__port, this.__host, this.__options);
};

/*
 *@name NodeRedisUpdate
 *@param insertSet : Key type
 *@param insertValue : Anytype
 *@description Update a Item into collection
 *@Todo: remove the same data
 */
NodeRedisObj.prototype.NodeRedisUpdate = function NodeRedisUpdate(insertSet, insertValue) {
    client = this.__client;
    client.on("error", function (err) {
        console.log("error event - " + client.host + ":" + client.port + " - " + err);
    });
    client.sadd(insertSet, insertValue);
};

/*
 *@name NodeRedisGet
 *@param insertSet : insertSet type
 *@description Get a Item into collection, but this is for test only
 */
NodeRedisObj.prototype.NodeRedisGet = function NodeRedisGet(insertSet) {
    client = this.__client;
    client.spop(insertSet, function (err, value) {
        if (err) throw err;
        console.log(value);
    });
};

/*
 *@name NodeRedisEnd
 *@description End a connection of Redis
 */
NodeRedisObj.prototype.NodeRedisEnd = function NodeRedisEnd() {
    client = this.__client;
    client.end();
};

/*
 * module export part
 */
module.exports.NodeRedisObj = NodeRedisObj;
