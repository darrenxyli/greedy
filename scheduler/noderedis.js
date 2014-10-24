/*
 * the redis persistence layer for nodejs
 */
(function() {
    /*
     * Import a third party pakage
     */
    "use strict";

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
        this.__client.on("error", function (err) {
            console.log("error event - " + this.__host + ":" + this.__port + " - " + err);
        });
        this.__client.sadd(insertSet, insertValue);
    };

    /*
     *@name NodeRedisGet
     *@param insertSet : insertSet type
     *@description Get a Item into collection, but this is for test only
     */
    NodeRedisObj.prototype.NodeRedisGet = function NodeRedisGet(insertSet) {
        this.__client.spop(insertSet, function (err, value) {
            if (err) throw err;
            console.log(value);
            return value;
        });
    };

    /*
     *@name NodeRedisEnd
     *@description End a connection of Redis
     */
    NodeRedisObj.prototype.NodeRedisEnd = function NodeRedisEnd() {
        this.__client.end();
    };

    /*
     * module export part
     */
    module.exports = NodeRedisObj;
})();