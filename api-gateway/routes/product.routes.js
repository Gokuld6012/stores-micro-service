const productController = require('../controllers/product.controller')

function init(router) {
    router.get('/test', productController.testRequest);
}

module.exports.init = init;
