const express = require('express');
const bodyParser = require('body-parser');
const notifier = require('node-notifier');

const app = express();
const port = process.env.PORT || 9000;

app.use(bodyParser.json())

app.get("/health", (req, res) => {
    res.status(200).send();
});


app.post("/notify", (req, res) => {
    notify(req.body, reply => res.send(reply)) 
});

app.listen(port, () => {
    console.log(`Server is up and running on ${port}`);
});

const notify = ({title, message}, callback) => {
    notifier.notify(
        {
            title: title || "Unknown title",
            message: "Unknown message",
            sound: true,
            wait: true,
            reply: true, 
            closeLable: "Completed?",
            timeout: 15
        },
        (err, response, reply) => {
            callback(reply);
        }
    )
};