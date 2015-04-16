from flask import Flask
from gevent.wsgi import WSGIServer
app = Flask(__name__)

@app.route('/<int:number>')
def index(number=1):
    return "Python + Flask<hr>fib("+ str(number) + "): " + str(fib(number))

def fib(n):
    if n == 0:
        return 0
    elif n == 1:
        return 1
    else:
        return fib(n - 1) + fib(n - 2)

if __name__ == '__main__':
	WSGIServer(('127.0.0.1', 5000), app, None, 'default', None).serve_forever()
