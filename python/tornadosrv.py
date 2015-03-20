import tornado.ioloop
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
    application.listen(8888)
    tornado.ioloop.IOLoop.instance().start()
