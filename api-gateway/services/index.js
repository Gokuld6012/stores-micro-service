const caller = require('grpc-caller')
const { startCase } = require('lodash');
const serviceConfig = require('./service-list');
let serviceList = {};

serviceConfig.forEach(proto => {
    const PROTO_PATH = `${__dirname}/../proto/${proto.name}.proto`
    const client = caller(`localhost:${proto.port}`, PROTO_PATH, startCase(proto.name))
    serviceList[proto.name] = client;
})

module.exports = serviceList;
