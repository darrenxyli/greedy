/**
 * API Based Scraping Engine
 *
 * @Todo: Write short comments to demenstrate how this engine architectured
 */
(function() {
    "use strict";

    var BaseEngine = require(ROOT + '/lib/base');
    var EventEmitter = require('events').EventEmitter;
    var Constants = require(ROOT + '/lib/constants');
    var request = require('request');
    var fs = require('fs');
    var zlib = require('zlib');
    var Fiber = require('fibers');

    var APIEngine = function APIEngine() {
        this.name = Constants.Engine.API;
        this.__steps = [];
        this.__ee = new EventEmitter;
        this.__started = false;
        this.__running = false;
        this.__extraSteps = []; /* Special for install sequential steps in one single step!!! */
    };

    /* init the api engine with BaseEngine */
    APIEngine.prototype = new BaseEngine();

    APIEngine.prototype.create = function (setting, mode) {};

    /*
    * @Method: Exit()
    * @Description: exit the process of scraper, and throw the error out
    * @Output: error code, error message
    * */
    APIEngine.prototype.Exit = function(code, message) {
        code = code === null ? 0 : code;
        if (message) {
            console.log(message);
        }
        process.exit(code);
    };

    /*
    * @Method:Then()
    * @Description: achieve the sequential excute function, to solve the character of NodeJS
    * */
    APIEngine.prototype.Then = function(callback) {
        if (this.__running) {
            this.__extraSteps.push(callback); /* @Todo: Add comments to explain why implemented like this */
        } else {
            this.__steps.push(callback);
        }
    };

    APIEngine.prototype.on = function() {
        this.__ee.on.apply(this.__ee, arguments);
    };

    APIEngine.prototype.emit = function() {
        this.__ee.emit.apply(this.__ee, arguments);
    };

    APIEngine.prototype.removeAllListeners = function() {
        this.__ee.removeAllListeners.apply(this.__ee, arguments);
    };

    /* @Warning: Differences with Browser Based Engine */
    APIEngine.prototype.Sleep = function(timeout) {
        var fiber = Fiber.current;
        setTimeout(function() {
            fiber.run();
        }, timeout);
        Fiber.yield();
    };

    /*
    * @Method: Log()
    * @Description: log out info
    * */
    APIEngine.prototype.Log = function(msg, level) {
        level = level ? level.toLowerCase() : 'log';
        console[level](msg);
    };

    /*
    * @Method: Troubleshooting()
    * @Description: throw the error messages
    * */
    APIEngine.prototype.Troubleshooting = function(message) {
        /* Troubleshooting in API based engine will only log critical logs */
        this.Log(message, 'error');
    };

    /**
     * Request Remote Resource
     *
     * @Warning: Current implementation relies Fiber to send the request one by one
     * This will have conflicts with our concurrent downloading mechanism
     *
     * @param {String} uri Resource to request
     * @param {String} verb HTTP method
     * @param {Object} data Data to send to remote server
     * @param {Object} setting Extra settings like headers, enable global proxy, custom success conditions
     * @param {Object} onSuccess Function to call when this action performed successfully
     * @param {Object} onError Function to call when this action performed failed
     *
     */
    APIEngine.prototype.Request = function(uri, verb, data, setting, onSuccess, onError) {
        var fiber = Fiber.current;

        var headers = setting.headers ? setting.headers : {};
        var extraOptions = setting.options ? setting.options : {};
        var options = {
            uri: uri,
            headers: headers
        };

        /* handle form like submit */
        if (setting.form) {
            options.form = data;
        } else {
            options.qs = data;
        }

        /* Merge the extra options with builtin options */
        for (var key in extraOptions) {
            options[key] = extraOptions[key];
        }

        // if (this.IsEnabledProxy()) {
        //     var oneProxy = this.GetDefaultProxy();
        //     if (oneProxy.valid()) {
        //         options.proxy = oneProxy.getAuthWithServerInfo(true);
        //     } else {
        //         this.emit(Constants.Event.NO_MORE_PROXY);
        //     }
        // }

        var req = request[verb.toLowerCase()](options);

        req.on('error', function(err) {
            if (onError) {
                onError(err);
            }
            fiber.run();
        });

        var checkErrorConditions = function(headers, data, code) {
            if (!setting.errorConditions) return null;
            var error, errormsg;
            var conditions = $._.isArray(setting.errorConditions) ? setting.errorConditions : [setting.errorConditions];
            for (var i = 0, l = setting.errorConditions.length; i < l; ++i) {
                var condition = setting.errorConditions[i];
                var a = condition(headers, data, code);
                var r = (typeof a === 'boolean') ? a : a[0];
                if (r) {
                    error = r;
                    errormsg = (typeof a !== 'boolean') ? a[1] : 'unknown error message';
                    break;
                }
            }
            return error ? {code: '', errno: errormsg} : null;
        };

        req.on('response', function(res) {
            var chunks = [];
            req.on('data', function(chunk) {
                chunks.push(chunk);
            });
            req.on('end', function() {
                var buffer = Buffer.concat(chunks);
                var data;
                var error;
                if (setting.gzipped) {
                    zlib.gunzip(buffer, function(err, decoded) {
                        if (err) {
                            onError(err);
                            fiber.run();
                        } else {
                            data = decoded.toString();
                            error = checkErrorConditions(res.headers, data, res.statusCode);
                            if (error) {
                                onError(error);
                            } else {
                                onSuccess(data);
                            }
                            fiber.run();
                        }
                    });
                } else {
                    error = checkErrorConditions(res.headers, buffer.toString(), res.statusCode);
                    if (error) {
                        onError(error);
                    } else {
                        onSuccess(setting.binary ? buffer : buffer.toString(), res.statusCode);
                    }
                    //fiber.run();
                }
            });
        });
        //Fiber.yield();
    };

    /*
     * @Method: Start()
    * @Description: start the api engine
    * */
    APIEngine.prototype.Start = function() {
        this.__started = true;
        this.Log('I\'m begining to work...');
    };

    /**
     * @Method: SaveFile()
     * @Description: save the data into file
     * @param {String} filename
     * @param {String} data
     * */
    APIEngine.prototype.SaveFile = function(filename, data) {
        /* Please notice the difference between NodeJS and PhantomJS */
        this.Log('Saving file to ' + this.GetRealPath(filename));
        fs.writeFileSync(this.GetRealPath(filename), data);
    };

    /**
     * @Method: ReadFile()
     * @Description: load file
     * @param {String} filename
     * */
    APIEngine.prototype.ReadFile = function(filename) {
        var content = null;
        if (!fs.existsSync(this.GetRealPath(filename))) {
            return content;
        }
        content = fs.readFileSync(this.GetRealPath(filename));
        return content.toString();
    };

    /**
    * @Method: Run()
    * @Description: run the scraper
    * */
    APIEngine.prototype.Run = function() {
        if (!this.__started) {
            throw new EngineException('Give me the battery!');
            return;
        }
        if (!this.__steps.length) {
            throw new EngineException('No steps defined, are you kidding me? Exiting...');
            return;
        }
        this.__running = true;

        Fiber((function(context) {
            return function() {
                var step;
                while (true) {
                    if (!context.__steps.length) {
                        break;
                    }
                    step = context.__steps.shift();
                    if (step) {
                        step();
                        /* @Warning: The fix progress is pretty bullshit, but it works! Refactor later! */
                        context.__steps = context.__extraSteps.concat(context.__steps);
                        context.__extraSteps = [];
                    } else {
                        context.Log('What happened? Illegal step found!');
                    }
                }
            };
        })(this)).run();
    };
    module.exports = APIEngine;
})();
