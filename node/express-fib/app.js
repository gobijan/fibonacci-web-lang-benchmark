var express = require("express");
var app = express();

var fib = function (n) {
    if (n === 0) {
        return 0;
    } else if (n == 1) {
        return 1;
    } else {
        return fib(n - 1) + fib(n - 2)
    }
};

app.get("/:number", function (req, res) {
        var number = req.param('number');            
        var result = fib(number);
        res.send("Node + Express<hr> fib("+number+"): "+result);        
});

app.listen(3000);