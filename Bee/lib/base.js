(function() {
    "use strict";

    var Constants = require(ROOT + '/lib/constants');

    var BaseEngine = function BaseEngine() {
        this.__loopIndex = 0;
        this.__outputDirectory = null;
        this.__taskDirectory = null; /* For storing internal debugging support files */
        this.__proxyPool = null;
        this.name = 'base';
        this.__values = {}; /* tempornary variable for internal purpose */
    };

    BaseEngine.prototype.GetLoopIndex = function() {return this.__loopIndex;};
    BaseEngine.prototype.IncLoopIndex = function() {this.__loopIndex += 1;};

    /**
    * @Method: Touch()
    * @Description: create an empty file without data
    * @param: {String} filename
    * */
    BaseEngine.prototype.Touch = function(filename) {
        this.SaveFile(filename, '');
    };

    /**
    * @Method: SetOutputDirectory()
    * @Description: set outputDirectory
    * @param: {String} outputDirectory
    * */
    BaseEngine.prototype.SetOutputDirectory = function(outputDirectory) {
        this.__outputDirectory = outputDirectory;
    };

    /**
    * @Method: GetOutputDirectory()
    * @Description: return the outputDirectory
    * */
    BaseEngine.prototype.GetOutputDirectory = function() {
        return this.__outputDirectory;
    };

    BaseEngine.prototype.SetProxyPool = function(pool) {
        this.__proxyPool = pool;
    };

    BaseEngine.prototype.GetProxyPool = function() {
        return this.__proxyPool;
    };

    /* @Todo: Try to support relative path and absolute path in the easiest way */
    BaseEngine.prototype.GetRealPath = function(path) {
        return (path[0] === '/') ? path : this.GetOutputDirectory() + path;
    };

    /* For Internal Debugging Support */
    BaseEngine.prototype.SetTaskDirectory = function(directory) {
        this.__taskDirectory = directory;
    };
    BaseEngine.prototype.GetTaskDirectory = function() {
        return this.__taskDirectory;
    };
    /* End */

    module.exports = BaseEngine;
})();
