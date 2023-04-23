//create a basoc express server

const express = require('express');
const app = express();
//barse body of the request
const bodyParser = require('body-parser');
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
const port = 3000;

app.get('/', (req, res) => res.send('Hello World!'));

app.get('/get', (req, res) => res.status(200).send({ message: 'get request' }));

app.post('/post', (req, res) => {
    res.status(200).send({ message: req.body });
});
//send a stringified json
app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify({ message: req.body }));
});


app.listen(port, () => console.log(`Example app listening on port ${port}!`));
