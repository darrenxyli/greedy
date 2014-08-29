/* Universal implementation for generators */

module.exports = function(engine) {

    var range = function range(min, max, step) {
        var arr = [];
        var increaseStep = step ? step : 1;
        for (var i = min; i < max; i += increaseStep) {
            arr.push(i);
        }
        return arr;
    };

    return {
        name: 'Generator',
        handler: {
            range: range
        }
    };
};
