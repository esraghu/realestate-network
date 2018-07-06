const express = require("express");
const app = express();
const { exec } = require("child_process");

app.get("/", (req, res) => {
    console.log(`Request received from ${req.connection.remoteAddress} and ${req.headers['x-forwarded-for']}`);
    res.send("Welcome to the RealEstate Blockchain Network");
})

app.get("/query", (req, res) => {
    //const queryCmd = "ls -l ../byfn.sh"
    const queryCmd = "../byfn.sh query";
    exec( queryCmd, (err, stdout, stderr) => {
        if (err) {
            //throw new Error(`${queryCmd} couldn't be executed!`);
            res.send(`${queryCmd} couldn't be executed due to ${err}`);
            
        }
        console.log(`stdout: ${stdout}`);
        console.log(`stderr: ${stderr}`);
        
        if (stdout != null) {
            res.send(stdout);
        } else {
            res.send(stderr);
        }
    })
})

app.listen(3000, () => console.log('Listening on port 3000'));
