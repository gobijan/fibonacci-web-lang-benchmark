var cluster = require('cluster');

if (cluster.isMaster) {
	// Code to run if we're in the master process

	var numCPUs = require('os').cpus().length - 1; // leave 1 core for wrk
	// Fork workers.
	for (var i = 0; i < numCPUs; i++) {
		cluster.fork();
	}
} else {
	// Code to run if we're in a worker process
	var express = require("express");
	var app = express();

	app.get("/:number", function(req, res) {
		var number = req.param('number');
		var result = fib(number);
		res.send("Node + Express<hr> fib(" + number + "): " + result);
	});
	app.listen(3000);
	console.log("Express Worker running.")
}


var fib = function(n) {
	if (n === 0) {
		return 0;
	} else if (n == 1) {
		return 1;
	} else {
		return fib(n - 1) + fib(n - 2)
	}
};