
def process(line, acc, idx):
    split = line.split(" ")
    action = split[0]
    sign = split[1][0]
    amount = int(split[1][1:])
    print("action: {}, sign: {}, amt: {}".format(action, sign, amount))
    if action == "acc":
        if sign == "+":
            acc += amount
        elif sign == "-":
            acc -= amount
        idx += 1
    elif action == "jmp":
        if sign == "+":
            idx += amount
        elif sign == "-":
            idx -= amount
    elif action == "nop":
        idx += 1
    return (acc, idx)

def halting(data):
    acc = 0
    idx_map = {}
    idx = 0
    while True:
        line = data[idx]
        line = line.strip()
        (new_acc, new_idx) = process(line, acc, idx)
        if new_idx in idx_map:
            return acc
        else:
            idx_map[new_idx] = True
            acc = new_acc
            idx = new_idx

if __name__ == "__main__":
    text_file = open("../input/aoc8.txt", "r")
    string_data = text_file.readlines()
    ans = halting(string_data)
    print("the number before the infinite loop is {}".format(ans))
