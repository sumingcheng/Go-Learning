const express = require('express');
const bodyParser = require('body-parser');
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

    const todo = {
        id: new Date().getTime(),
        content,
        completed: false,
    }

    list.push(todo);
    console.log(list)
    res.json(todo)
})

app.post('/toggle_todo', (req, res) => {
    const {id} = req.body;

    let target = null;
    list = list.map(todo => {
        if (todo.id == id) {
            todo.completed = !todo.completed;
        }
        target = todo;
        return todo;
    })
    console.log(list)
    res.json(target)
})

app.post('/remove_todo', (req, res) => {
    const {id} = req.body;

    const target = list.find(todo => todo.id == id);
    list = list.filter(todo => {
        return todo.id != id
    })

    console.log(list)
    res.json(target)
})


app.listen(3000, () => {
    console.log('Server is running on port :http://localhost:3000');
})