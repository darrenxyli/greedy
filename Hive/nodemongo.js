/*
 * the mongodb persistence layer for nodejs
 * @Description: pure mongodb method
 */

/*
 *import mongo package
 */
var MongoClient = require('mongodb').MongoClient;

var NodeMongoObj = function NodeMongoObj(Host, Port, DbName, Collection){
    this.__host = Host ? Host : '127.0.0.1';
    this.__port = Port ? Port : 27017;
    this.__dbname = DbName ? DbName : 'test';
    this.__collectionname = Collection;
};
/*
 *@name nodeMongoInitDB
 *@param dbName
 *@description Init a DB object and return it.
 */
NodeMongoObj.prototype.nodeMongoInitDB = function nodeMongoInitDB() {
    var dbConnection = MongoClient.connect('mongodb://'+this.__host+':'+this.__port+'/' + this.__dbname, function(err, db) {
        if (err) throw err;
        this.__db = db;
        return db;
    });
    console.log(dbConnection);
};

/*
 *@name nodeMongoSetDB
 *@param db : DB object
 *@description Set DB value.
 */
NodeMongoObj.prototype.nodeMongoSetDB = function nodeMongoSetDB(db) {
    this.__db = db;
};

/*
 *@name nodeMongoCloseDB
 *@description Close DB .
 */
NodeMongoObj.prototype.nodeMongoCloseDB = function nodeMongoCloseDB() {
    this.__db.close();
};

 /*
 *@name nodeMongoInitCollection
 *@param db : MongoClient Object
 *@param collectionName : String
 *@description Init a collection from a DB object and return it.
 */
NodeMongoObj.prototype.nodeMongoInitCollection = function nodeMongoInitCollection(collectionName) {
    var db = this.__db;
    var collection = db.collection(collectionName);
    return collection;
};

/*
 *@name nodeMongoInsert
 *@param collection : MongoClient Collection Object
 *@param insertData : Dictionary
 *@description Insert a Item into collection
 *@Todo: remove the same data
 */
NodeMongoObj.prototype.nodeMongoInsert = function nodeMongoInsert(obj, insertData) {
    var host = this.__host;
    var port = this.__port;
    var dbname = this.__dbname;
    var collectionname = this.__collectionname;

    var dbConnection = MongoClient.connect('mongodb://' + host + ':' + port + '/' + dbname, function(err, db) {
        if (err) throw err;
        var collection = db.collection(collectionname);
        collection.save(insertData, function(err, docs){
            if (err) throw err;
            console.log("Insert!");
        });
        db.close();
    });
};

/*
 * module export part
 */
module.exports.NodeMongoObj = NodeMongoObj;
