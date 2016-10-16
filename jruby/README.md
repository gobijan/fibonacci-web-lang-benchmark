# Fibonacci web lang benchmark usinga jRuby + Sinatra + torquebox

## Set up & Run
* Install jruby
* bundle install
* bundle exec torquebox jar --envvar RACK_ENV=production
* java -jar jruby.jar --port 8123
* `wrk -c 64 -d 30s http://localhost:8123/10`


## Results

Let the jvm have 2-3 runs of the bench to warm up....

**7465 Requests/sec** => 224502 requests in 30s
Average Latency 14.36ms

Hardware: MacBook Pro Retina 13" i5@2.6GHz and 8GB RAM
OS X 10.9.5 w/ Security Update 2015-007
