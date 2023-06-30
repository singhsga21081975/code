def double(func):
    def wrapper(x):
        result = func(x)
        return result * 2
    return wrapper

def square(x):
    return x ** 2

import pdb

# Applying the 'double' decorator to the 'square' function
square = double(square)

result = square(5)
print(result)  # Output: 50


def multiply(func):
    pdb.set_trace()
    def wrapper(x1):
        pdb.set_trace()
        val= func(x1)
        print("execute")
        return val
    return wrapper


def multiply1(func):
    pdb.set_trace()
    def wrapper():
        pdb.set_trace()
        val= func()
        print("execute")
        return val
    return wrapper

def m():
    print("mmmmmmmmmmmmmmmmm")
def sq(x):
    print(" execute square")
    return x ** 2

sq = multiply(sq)

m = multiply1(m)

print(" square d",sq(3))
print(" square d",m())
