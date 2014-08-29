var Modes = {
    Testing: 'testing',
    Production: 'prod',
    Troubleshooting: 'troubleshooting'
};

var Progs = {
    CasperJS: 'casperjs',
    Node: 'node'
};

var Engine= {
    API: 'api',
    BROWSER: 'browser'
}

var Files = {
    GlobalSerialized: '.global' /* Symbol G will be serialized to this file and probaraly will be restored in the future */
};

module.exports = {
    Mode: Modes,
    Prog: Progs,
    File: Files,
    Engine: Engine
};
