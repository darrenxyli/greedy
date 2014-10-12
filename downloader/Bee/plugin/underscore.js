/* Wrapper for underscore */

module.exports = function(engine) {
    var underscore = require('underscore');

    return {
        name: '_',
        handler: underscore
    };
};
