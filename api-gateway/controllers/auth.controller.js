const { auth } = require('../services/index');

exports.testRequest = async (req, res) => {
    try {
        const responseMessage = await auth.GetTestResponse('');
        return res.json(responseMessage);
    } catch (err) {
        return res.status(400).json({
            err: err.message || 'Something went wrong'
        })
    }
}
