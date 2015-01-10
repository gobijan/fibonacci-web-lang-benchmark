import re
from twisted.web.server import Site
from twisted.web.resource import Resource
from twisted.internet import reactor


def fib(n):
 if n == 0:
    return 0
 elif n == 1:
    return 1
 else:
    return fib(n-1) + fib(n-2)


class FibonacciResource(Resource):
    isLeaf = True
    def __init__(self, num):
        self.num = int(num)

    def render_GET(self, request):
         return "Python + Twisted<hr>fib(%s): %s" % (self.num, fib(self.num))


class Router(Resource):
    def getChild(self, path, request):
        if re.match('\d+', path):
            return FibonacciResource(path)
        return self


reactor.listenTCP(5000, Site(Router()))
reactor.run()