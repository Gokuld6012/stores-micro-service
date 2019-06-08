const authController = require('../controllers/auth.controller')

function init(router) {
    router.get('/test', authController.testRequest);
}

module.exports.init = init;
