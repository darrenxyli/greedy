/* @Todo: Can not save data, need to recheck */
var mongoose = require('mongoose');
var Post = require('./postModel');

/**
* @Class: MonnodeObj
* @Description: connection with mongoDB
* @Todo: auto get conneciton condition by read config
* */
var MonnodeObj = function () {
    this.__db = mongoose.createConnection('mongodb://127.0.0.1:27017/posts', { server: { poolSize: 1 }});
    this.__db.on('error',console.error.bind(console,'连接错误:'));
    this.__db.once('open',function(){
      console.log('open it!');
    });
};

/**
* @Method getPost()
* @Description: query method
* @Param {String} mTitle
* @Param {String} mLink
* @Param {String} mContext
* */
MonnodeObj.prototype.getPost = function getPost(mTitle, mLink, mContext) {
    var mPost = new Post({
        title : mTitle,
        link :  mLink,
        context: mContext
    });
    return mPost;
}

/**
* @Method savePost()
* @Description save method
* @Param {String} mTitle
* @Param {String} mLink
* @Param {String} mContext
* */
MonnodeObj.prototype.savePost = function savePost(mTitle, mLink, mContext) {
    var post = this.getPost(mTitle,mLink,mContext);
    console.log(post);
    post.save(function(err){
        if(err){
            console.error(err.message);
            return;
        }
        console.log(post.title);
        this.__db.close();
    });
}

/**
* @Method addPost()
* @Description save with object
* @Param {Object} post
* */
MonnodeObj.prototype.addPost = function addPost(post) {
    post.save(function(err){
        if(err){
            console.error(err.message);
            return;
        }
        console.log(post.title);
    });
    this.__db.close();
}

/**
* @Method savePosts
* @Description save many posts
* @Param {Array} posts
* */
MonnodeObj.prototype.savePosts = function savePosts(posts) {
    posts.forEach(function(e){
        addPost(e);
    });
    this.__db.close();
}

/**
* @Method Close()
* @Description close connection with mongoDB
* */
MonnodeObj.prototype.Close = function queryPost(key) {
    this.__db.close();
}

/**
* @Description exports part
* */
module.exports.MonnodeObj = MonnodeObj;
