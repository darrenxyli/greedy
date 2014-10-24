var NodeMongoObj = require('../../Hive/postService').MonnodeObj;

var q = new NodeMongoObj();
q.savePost('title', 'www.baidu.com', 'helloworld');
