from flask import Flask
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
    app.run(debug=False)
