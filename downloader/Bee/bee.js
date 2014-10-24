/*init the engine*/
var engine = require('./lib/engine');
var Getopt = require('node-getopt');
//var mondb = require('../Hive/postService').MonnodeObj;
var utils = require('./lib/spiderUtil');
var cheerio = require('cheerio');

/* get starturl from cmd */
var getopt = new Getopt([
    ['u', 'url=ARG', ''],
    ['', 'overwrite', 'overwrite mode']
]);

getopt.setHelp(
	"Sample Option: node bee.js -u http://starturl"
);

var opt = getopt.parseSystem();

if (!Object.keys(opt.options).length) {
    getopt.showHelp();
    console.log('\nYou need give me options!');
    process.exit();
}

/* init the obj */
var E = engine({'loadImages':false, 'localToRemoteUrlAccessEnabled':false}, '/Users/xinyang/Documents', '/Users/xinyang/Documents');
//var DB = new mondb();

/*crawl the data*/
if (opt.options.url) {
	var URL = opt.options.url;
	E.Start();
	E.Request(
		URL,
		'GET',
		{},
		{},
		function onSuccess(res) {
			var $ = cheerio.load(res);
			var title = $('title').text();
			console.log(title);
			var content = $('body').text();
			console.log(content);
            //DB.savePost(title, URL, content);
		}
	);
    //DB.Close();
	E.Then(function() {
		E.Exit(1, 'bye');
	})
};
