const express = require('express');
const mysql = require('mysql2');  
const bodyParser = require('body-parser');

const app = express();
app.use(bodyParser.json());

const db = mysql.createConnection({
    host: 'mysql',
    user: 'root',
    password: 'root',
    database: 'testdb'
});

db.connect((err) => {
    if (err) throw err;
    console.log('Connected to MySQL Database');
});

app.post('/create', (req, res) => {
    const name = req.body.name;
    const sql = 'INSERT INTO your_table (name) VALUES (?)';
    db.query(sql, [name], (err, result) => {
        if (err) return res.status(500).send(err);
        res.send('Record added!');
    });
});

app.get('/read', (req, res) => {
    const sql = 'SELECT name FROM your_table';
    db.query(sql, (err, results) => {
        if (err) return res.status(500).send(err);
        res.json(results);
    });
});

const PORT = 3000;
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
