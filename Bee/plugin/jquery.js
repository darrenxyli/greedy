/* Wrapper for jQuery */

module.exports = function(engine) {
    var jQuery;

    if ($.isBrowser()) {
        jQuery = {};
    } else {
        var jsdom = require('jsdom');
        window = jsdom.jsdom().createWindow();
        jQuery = require('jquery');
    }

    return {
        name: '$',
        handler: jQuery
    };
};
