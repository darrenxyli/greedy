/* Plugin for solving date related problems */

module.exports = function(engine) {
    "use strict";

    var moment = require('moment');

    var ADate = function ADate() {};

    /**
     * Check whether the string passed in is a valid date format
     * @param {String} s date string to verify
     * @return {Boolean} true | false
     */
    ADate.valid = function(s) {
        return Date.parse(s) ? true : false;
    };

    ADate.parse = function(s) {
        /**
         * Fix the timezone issue
         * This function will always return back the local time extracted from s
         */
        var d, n;
        d = new Date();
        if (engine.name.toUpperCase() === 'BROWSER') {
            /**
             * Currently, PhantomJS doesn't support international date api.
             * Fortunately ```Date.parse``` in PhantomJS return the Date we expected
             */
            d.setTime(Date.parse(s));
        } else {
            d.setTime(Date.parse(s + ' GMT'));
        }
        n = new Date(d.getUTCFullYear(), d.getUTCMonth(), d.getUTCDate(),
                     d.getUTCHours(), d.getUTCMinutes(), d.getUTCSeconds(), d.getUTCMilliseconds());
        return n;
    };

    ADate.getUTCTs = function(dateObj) {
        return Math.floor(moment(dateObj).utc().valueOf() / 1000);
    };

    /**
     * Adjust a new Date object by add/sub the days
     */
    ADate.addDays = function(date, days) {
        var d = new Date();
        d.setTime(date.getTime() + days * 86400000);
        return d;
    };

    ADate.minusDays = function(date, days) {
        return this.addDays(date, -days);
    };

    ADate.format = function(date, format) {
        return moment(date).format(format);
    };

    ADate.getCurrentTs = function() {
        return Math.floor(new Date().getTime() / 1000);
    };

    ADate.getDayDiffFromNow = function(date) {
        var today = moment();
        var startDate = moment(date);
        return today.diff(startDate, 'days');
    };

    return {
        name: 'Date',
        handler: ADate
    };
};
