/**
 * Bootstrap for the whole scraping engine
 */
(function() {
    "use strict";

    var fs = require('fs');

    var isBrowser = function() {
        return (typeof GLOBAL !== 'undefined') ? false : true;
    };

    /* Currently We only support 2 Host Environments: Node.JS & Browser */
    var G = isBrowser() ? window : GLOBAL;

    /* Hack for our internal usage and bind to $ */
    G.$ = {
        _: require('underscore'),
        isBrowser: isBrowser,
        slice: [].slice
    };

    G.$.createFunctionWithCtx = function() {
        var fn = arguments[0];
        var arg = (arguments.length > 1) ? [].slice.call(arguments, 1) : [];
        return function() {
            return fn.apply(null, arg);
        };
    };

    G.$.registerGlobal = function(key, value) {
        G[key] = value;
    };

    /* Eliminate differences among different host environments */
    G.$.readdirSync = isBrowser() ? fs.list : fs.readdirSync;
    G.$.utils = isBrowser() ? require('utils') : require('util');
    G.$.isDirectory = isBrowser() ? fs.isDirectory : function(path) {
        return fs.statSync(path).isDirectory();
    };

    G.$.mergeObject = function(obj1, obj2) {
        var x = {};
        var key;
        if (obj1) {
            for (key in obj1) {
                x[key] = obj1[key];
            }
        }
        if (obj2) {
            for (key in obj2) {
                x[key] = obj2[key];
            }
        }
        return x;
    };

    /* Register the root directory */
    G.$.registerGlobal('ROOT', isBrowser() ? fs.workingDirectory : process.cwd());

})();
