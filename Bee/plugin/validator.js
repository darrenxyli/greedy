module.exports = function(engine) {

    var Error = {};

    Error.validResponseCode = function validResponseCode(code, callback) {

        if (typeof code !== 'object') {
            code = [code];
        }

        return function(resHeaders, content, resCode) {
            for (var i = 0, l = code.length; i < l; ++i) {
                if (parseInt(resCode) === code[i]) {
                    return false;
                }
            }
            return [true, callback ? callback(resHeaders, content, resCode) : 'response code not matched'];
        };

    };

    Error.validResponseHeader = function validResponseHeader(obj, setting, callback) {
        var strictMode = true;
        if (setting && typeof setting.strict !== 'undefined') {
            strictMode = setting.strict;
        }

        return function(resHeaders, content, resCode) {
            if (!resHeaders) return false;

            var valid = true;

            for (var key in obj) {
                if (obj.hasOwnProperty(key)) {
                    if (!resHeaders[key]) {
                        valid = false;
                        break;
                    }

                    if (strictMode) {
                        if (resHeaders[key] !== obj[key]) {
                            valid = false;
                            break;
                        }
                    } else {
                        if (!resHeaders[key].match(obj[key])) {
                            valid = false;
                            break;
                        }
                    }

                }
            }

            return valid ? false : [true, callback ? callback(resHeaders, content, resCode) : 'Response header not matched'];
        };

    };

    Error.existResponseHeader = function existResponseHeader(header, callback) {

        return function(resHeaders, content, resCode) {
            if (!resHeaders) return [true, 'no headers found'];
            var valid = resHeaders.hasOwnProperty(header) ? false : true;
            if (!valid) return valid;
            return [valid, callback ? callback(resHeaders, content, resCode) : 'header not found'];
        };

    };

    Error.notExistResponseHeader = function notExistResponseHeader(header, callback) {

        return function(resHeaders, content, resCode) {
            if (!resHeaders) return [true, 'no headers found'];
            return !resHeaders.hasOwnProperty(header) ? false : [true, callback ? callback(resHeaders, content, resCode) : 'header exist'];
        };

    };

    Error.validResponseContentType = function validResponseContentType(contentType, setting, callback) {
        var strictMode = true;
        if (setting && typeof setting.strict !== 'undefined') {
            strictMode = setting.strict;
        }

        return function(resHeaders, content, resCode) {
            if (!resHeaders) return [true, 'headers not found'];
            var contentTypeKey;
            var hasContentType = $._.some(resHeaders, function (values, key) {
                if (key.toLowerCase() == 'content-type') {
                    contentTypeKey = key;
                    return true;
                }

                return false;
            });

            if (!hasContentType) {
                return [true, 'Content-Type not found'];
            }

            if (strictMode) {
                return (contentType === resHeaders[contentTypeKey]) ? false : [true, callback ? callback(resHeaders, content, resCode) : 'ContentType not matched'];
            } else {
                return resHeaders[contentTypeKey].match(contentType) ? false : [true, callback ? callback(resHeaders, content, resCode) : 'ContentType not matched'];
            }
        };
    };

    Error.validResponseBodySize = function validResponseBodySize(min, max, callback) {
        if (typeof min === 'undefined' || min < 0) {
            min = 0;
        }

        if (typeof max === 'undefined' || max < 0) {
            max = Infinity;
        }

        return function(resHeaders, content, resCode) {
            return (content.length >= min && content.length <= max) ? false : [true, callback ? callback(resHeaders, content, resCode) : 'Body size not matched'];
        };
    };

    Error.validResponseJSONContent = function(callback) {
        return function validResponseJSONContent(resHeaders, content, resCode) {
            try {
                var obj = JSON.parse(content);
                return false;
            } catch(e) {
                return [true, callback ? callback(resHeaders, content, resCode) : 'malformed JSON body'];
            }
        };
    };

    return {
        name: 'Validator',
        handler: {
            Error: Error
        }
    };
};
