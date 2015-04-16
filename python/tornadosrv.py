import sys
from tornado.httpserver import HTTPServer
from tornado.ioloop import IOLoop
import tornado.web

def fib(n):
    if n == 0:
        return 0
    elif n == 1:
        return 1
    else:
        return fib(n - 1) + fib(n - 2)

class FibHandler(tornado.web.RequestHandler):
    @tornado.web.asynchronous
    def get(self, number):
        number = int(number)
        self.write("Python + Flask<hr>fib("+ str(number) + "): " + str(fib(number)))
        self.finish()

application = tornado.web.Application([
    (r"/(.+)", FibHandler),
])

if __name__ == "__main__":
    num_proc = int(sys.argv[1]) if len(sys.argv) > 1 else 1
    print("start tornado with {0} processes".format(num_proc))
    server = HTTPServer(application)
    server.bind(8888)
    server.start(num_proc)
    IOLoop.instance().start()
