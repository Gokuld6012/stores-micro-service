const authController = require('../controllers/auth.controller')

function init(router) {
    router.get('/test', authController.testRequest);
    router.post('/register', authController.registerNewUser);
    router.post('/login', authController.login);
}

module.exports.init = init;
