
def joltage(data):
    data.sort()
    one_j, two_j, three_j = 0, 0, 0
    i = 0
    while i < len(data):
        current_j = data[i]
        if i == 0:
            prev_j = 0
        else:
            prev_j = data[i-1]

        diff = current_j - prev_j
        if diff == 3:
            three_j += 1
        elif diff == 2:
            two_j += 1
        elif diff == 1:
            one_j += 1
        i += 1
    # device's built in adapter
    three_j += 1
    return three_j * one_j

def calculate_combos(el, map):
    combos = 0
    diff_one, diff_two, diff_three = el - 1, el - 2, el - 3
    if diff_one in map:
        combos += map[diff_one]
    if diff_two in map:
        combos += map[diff_two]
    if diff_three in map:
        combos += map[diff_three]
    map[el] = combos
    return map

def joltage2(data):
    data.sort()
    last_el = max(data) + 3
    data.append(last_el)
    combos = {0: 1}
    for el in data:
        combos = calculate_combos(el, combos)
    
    return combos[last_el]

if __name__ == "__main__":
    text_file = open("../input/aoc10.txt", "r")
    string_data = text_file.readlines()
    data = list(map(lambda x: int(x), string_data))
    ans = joltage2(data)
    print("the total number of distinct ways to arrange the adapters is {}".format(ans))
