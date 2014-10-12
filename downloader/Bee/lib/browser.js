(function() {
    "use strict";

    /**
     * Status: Unstable
     *
     * @Notice: will add more APIs when we meet the special needs and group them in the future
     */

    var fs = require('fs');
    var Constants = require(ROOT + '/lib/constants');
    var BaseEngine = require(ROOT + '/lib/base');

    var createCasper = function(setting) {
        var userAgent = setting.userAgent ? setting.userAgent : 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/537.17 (KHTML, like Gecko) Chrome/24.0.1312.57 Safari/537.17';
        var logLevel = setting.log && setting.log.level ? setting.log.level : 'debug';

        var options = {
            verbose : true,
            logLevel : logLevel,
            pageSettings : {
                userAgent : userAgent
            }
        };

        if (setting.loadImages) {
            options.pageSettings.loadImages = setting.loadImages;
        }
        if (setting.userName) {
            options.pageSettings.userName = setting.userName;
        }
        if (setting.password) {
            options.pageSettings.password = setting.password;
        }
        if (setting.javascriptEnabled) {
            options.pageSettings.javascriptEnabled = setting.javascriptEnabled;
        }
        if (setting.loadPlugins) {
            options.pageSettings.loadPlugins = setting.loadPlugins;
        }
        if (setting.localToRemoteUrlAccessEnabled) {
            options.pageSettings.localToRemoteUrlAccessEnabled = setting.localToRemoteUrlAccessEnabled;
        }
        if (setting.XSSAudit) {
            options.pageSettings.XSSAuditingEnabled = setting.XSSAudit;
        }
        var casper = require('casper').create(options);

        /* @Todo: Bind more events for future */
        return casper;
    };

    var BEngine = function BEngine(setting){
        this.browser = null;
        this.name = 'Browser';
    };

    BEngine.prototype = new BaseEngine();

    BEngine.prototype.create = function(setting) {
        this.browser = createCasper(setting);
    };

    BEngine.prototype.DumpPage = function(path, content, callback) {
        console.log('Dumping Page to: ' + path);
        var e = null;
        try {
            fs.write(this.GetRealPath(path), content, 'w');
        } catch(err) {
            e = err;
        }
        if (callback)
            callback(e);
    };

    BEngine.prototype.Capture = function() {
        arguments[0] = this.GetRealPath(arguments[0]); /* The first parameter is the filename */
        this.browser.capture.apply(this.browser, arguments);
    };

    BEngine.prototype.GetCurrentURL = function() {
        return this.browser.getCurrentUrl();
    };

    BEngine.prototype.GetCurrentPageContent = function(type) {
        if (type == 'html') {
            return this.browser.getHTML();
        } else {
            return this.browser.getPageContent();
        }
    };

    BEngine.prototype.GetCurrentPageHTMLContent = function() {
        return this.GetCurrentPageContent('html');
    };

    BEngine.prototype.Log = function(msg, level) {
        if (level === null) {
            level = 'debug';
        }
        this.browser.log(msg, level, 'browser');
    };

    BEngine.prototype.Download = function(url, filename, method, data) {
        if (method === null) {
            method = 'POST';
        }
        if (data === null) {
            data = {};
        }
        this.browser.download(url, this.GetRealPath(filename), method, data);
    };

    BEngine.prototype.Exists = function(selector) {
        return this.browser.exists(selector);
    };

    BEngine.prototype.SaveFile = function(filename, data) {
        fs.write(this.GetRealPath(filename), data, 'w');
    };

    BEngine.prototype.ReadFile = function(filename) {
        var content = null;
        if (!fs.exists(this.GetRealPath(filename))) {
            return content;
        }
        content = fs.read(this.GetRealPath(filename));
        return content;
    };

    BEngine.prototype.Login = function(url, callback) {
        this.browser.thenOpen(url, callback);
    };

    BEngine.prototype.Reload = function(fn) {
        var callback = (typeof fn === 'undefined') ? function(){} : fn;
        this.browser.reload(callback);
    };

    BEngine.prototype.Fill = function() {
        this.browser.fill.apply(this.browser, arguments);
    };

    BEngine.prototype.Exit = function(code, message) {
        if (code === null) {
            code = 0;
        }
        if (message) {
            console.log(message);
        }
        this.browser.exit(code);
    };

    BEngine.prototype.Wait = function() {
        this.browser.waitFor.apply(this.browser, arguments);
    };

    BEngine.prototype.Sleep = function(timeout) {
        this.Wait(function(){return false;}, function(){}, function(){}, timeout);
    };

    BEngine.prototype.Then = function(callback) {
        this.browser.then(callback);
    };

    BEngine.prototype.Click = function() {
        this.browser.click.apply(this.browser, arguments);
    };

    BEngine.prototype.ClickLabel = function() {
        this.browser.clickLabel.apply(this.browser, arguments);
    };

    BEngine.prototype.ThenOpen = function(url, callback) {
        this.browser.thenOpen(url, (function(engine) {
            return callback;
        })(this));
    };

    /* Template Writers manually invoke the troubleshooting methods */
    BEngine.prototype.Troubleshooting = function(message) {
        if (typeof message !== 'undefined') {
            this.Log(message, 'error');
        }
        this.DumpPage(this.GetTaskDirectory() + '/' + new Date().getTime() + '.html', this.GetCurrentPageHTMLContent());
        this.Capture(this.GetTaskDirectory() + '/' + new Date().getTime() + '.png');
    };

    BEngine.prototype.Open = BEngine.prototype.ThenOpen;

    BEngine.prototype.Start = function() {
        this.browser.start();
    };

    BEngine.prototype.Run = function() {
        this.browser.run();
    };

    BEngine.prototype.EvalInDOM = function() {
        /**
         * Can use JSON parse<->stringify to hide some bad parts
         * Because the variable returned from page.evaluate is "immutable"
         */
        return this.browser.evaluate.apply(this.browser, arguments);
    };

    BEngine.prototype.GetPageContent = function(url, method, params, settings) {
        if (method === null) {
            method = 'GET';
        }
        if (settings === null) {
            settings = {};
        }
        return this.EvalInDOM(function(url, method, params, settings) {
            return __utils__.sendAJAX(url, method, params, false, settings);
        }, {url: url, params: params, method: method, settings: settings});
    };

    BEngine.prototype.on = function() {
        this.browser.on.apply(this.browser, arguments);
    };

    BEngine.prototype.emit = function() {
        this.browser.emit.apply(this.browser, arguments);
    };

    BEngine.prototype.removeAllListeners = function() {
        this.browser.removeAllListeners.apply(this.browser, arguments);
    };

    BEngine.prototype.SaveCookies = function(path) {
        var cookies = phantom.cookies;
        this.SaveFile(this.GetRealPath(path), JSON.stringify(cookies));
    };

    BEngine.prototype.GetCookies = function() {
        return phantom.cookies;
    };

    BEngine.prototype.LoadCookiesFromFile = function(path) {
        var content = fs.read(this.GetRealPath(path));
        if (content) {
            try{
                var obj = JSON.parse(content);
            } catch (err) {
                console.error('Error cookie file to load from disk');
            }
            if ($._.isArray(obj)) {
                $._.each(obj, function(e, idx) {
                    phantom.addCookie(e);
                });
            } else if ($._.isObject(obj)) {
                phantom.addCookie(obj);
            }
        }
    };

    BEngine.prototype.LoadCookies = function(o) {
        var obj = $._.isString(o) ? JSON.parse(o) : o;
        if ($._.isArray(obj)) {
            $._.each(obj, function(e, idx) {
                phantom.addCookie(e);
            });
        } else {
            phantom.addCookie(obj);
        }
    };

    BEngine.prototype.Base64EncodeResource = function() {
        return this.browser.base64encode.apply(this.browser, arguments);
    };

    module.exports = BEngine;
})();
