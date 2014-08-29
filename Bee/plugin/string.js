/* Wrapper for string */

module.exports = function(engine) {
    var S = require('string');

    return {
        name: 'S',
        handler: S
    };
};
