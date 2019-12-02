
def problem_1(total, module):
    total += (module // 3) - 2
    return total

def problem_2(total, module):
    while module > 0:
        module = (module // 3) - 2
        if module > 0:
            total += module
    return total


def day_1():
    total = 0
    f = open("aoc_day1.txt", "r")
    for module in f:
        total = problem_2(total, int(module))
    print("The total fuel is ", total)


if __name__ == "__main__":
    day_1()
