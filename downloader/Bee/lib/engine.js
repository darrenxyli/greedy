(function() {
    "use strict";
    var fs = require('fs');
    require('./bootstrap');

    function Engine(setting, outputDirectory, taskDirectory) {
        var settings = $._.isObject(setting) ? (setting.length > 0 ? setting[0] : {}) : {};
        var Engine;
        var engineFile, engine, path, plugin, exporter;

        engineFile = ROOT + '/lib/api';

        Engine = require(engineFile);

        engine = new Engine();
        engine.create(settings);

        engine.SetOutputDirectory(outputDirectory);
        engine.SetTaskDirectory(taskDirectory);

        /* Load all plugins and register to the engine */
        engine.Plugin = {};

        $.readdirSync(ROOT + '/plugin/').forEach(function(file, idx) {
            if (file[0] === '.') {
                return;
            }
            path = ROOT + '/plugin/' + file;
            plugin = require(path)(engine);
            if (!plugin || !plugin.name || !plugin.handler) {
                console.error('Failed to load plugin: ' + path);
            } else {
                console.log('[Plugin] ' + plugin.name + ' loaded!');
                engine.Plugin[plugin.name] = plugin.handler;
                if (plugin.exceptions) {
                    $._.forEach(plugin.exception, function(element, idx) {
                        $.registerGlobal(element.name, element.exception);
                    });
                }
            }
        });
        return engine;
    }
    module.exports = Engine;
})();
