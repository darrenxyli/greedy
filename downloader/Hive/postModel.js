var mongoose = require('mongoose');

/**
* @Name PostSchema
* @Description: the model save in mongoDB
* @Type {Schema}
*/
var PostSchema = new mongoose.Schema({
    title: {type: String, unique: true},
    link : String,
    description: String,
    context: String,
    pubDate: {
        type: Date,
        'default': Date.now
    },
    source: String,
    typeId: Number
});

/**
* @Description: init the model in mongoose
*/
var Post = mongoose.model('Post', PostSchema);

/* exports */
module.exports = Post;
