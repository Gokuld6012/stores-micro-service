const { auth } = require('../services/index');

exports.testRequest = async (req, res) => {
    try {
        const responseMessage = await auth.GetTestResponse('');
        return res.json(responseMessage);
    } catch (err) {
        return res.status(400).json({
            message: err.details
        })
    }
}

exports.registerNewUser = async (req, res) => {
    try {
        const userData = {
            firstname: req.body.firstname,
            lastname: req.body.lastname,
            username: req.body.username,
            password: req.body.password,
        }
        const userResponseData = await auth.RegisterNewUser(userData, {})
        return res.json(userResponseData)
    } catch (err) {
        return res.status(400).json({
            message: err.details
        })
    }
}

exports.login = async (req, res) => {
    try {
        const loginData = {
            username: req.body.username,
            password: req.body.password
        }
        const loginResponseData = await auth.Login(loginData, {})
        return res.json(loginResponseData)
    } catch (err) {
        return res.status(400).json({
            message: err.details
        })
    }
}
