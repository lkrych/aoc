# https://stackoverflow.com/a/33639875
x = 0
y = 0
d = 1
m = 1

while x < 3:
    while 2 * x * d < m:
        print(x, y)
        x = x + d
    while 2 * y * d < m :
        print(x, y)
        y = y + d
    d = -1 * d
    m = m + 1
