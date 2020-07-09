import sys
def get_caller_module_dict(levels):
    f = sys._getframe(levels)
    print(sys._getframe(0).f_code.co_name)  #当前函数名
    print(sys._getframe(1).f_code.co_name) #调用该函数的函数的名字，如果没有被调用，则返回<module>，貌似call stack的栈低
    return { **f.f_globals, **f.f_locals }

def a():
    return get_caller_module_dict(1)


def b():
    return get_caller_module_dict(2)

def c():
    print(a())
    print(b())

c()