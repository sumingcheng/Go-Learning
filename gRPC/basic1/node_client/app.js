const express = require('express');
const bodyParser = require('body-parser');

// 导入gRPC客户端
const gClient = require('./grpc');
const app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: false}));
app.all('*', (req, res, next) => {
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Header', 'X-Requested-With');
    res.header('Access-Control-Allow-Methods', 'PUT, POST, GET, DELETE, OPTIONS');
    next();
})

list = []

app.post('/add_todo', (req, res) => {
    const {content} = req.body;

    gClient.AddTodo({
        content: content
    }, (err, response) => {
        res.json(response)
    })
})

app.post('/toggle_todo', (req, res) => {
    const {id} = req.body;

    gClient.ToggleTodo({
        id: id
    }, (err, response) => {
        res.json(response)
    })
})

app.post('/remove_todo', (req, res) => {
    const {id} = req.body;

    gClient.RemoveTodo({
        id: id
    }, (err, response) => {
        res.json(response)
    })
})


app.listen(3000, () => {
    console.log('Server is running on port :http://localhost:3000');
})