const express = require('express'),
    bodyParser = require('body-parser'),
    port = 50050,
    app = express(),
    authRoute = require('./routes/auth.routes'),
    productRoute = require('./routes/product.routes');

app.use(bodyParser.json());

// API Routes
const authApi = express.Router();
app.use('/api', authApi);
authRoute.init(authApi);

const productApi = express.Router();
app.use('/api/product', productApi);
productRoute.init(productApi);

app.listen(port, () => {
    console.log('Server is listening at port', port)
});
