
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

if __name__ == "__main__":
    text_file = open("../input/aoc10.txt", "r")
    string_data = text_file.readlines()
    data = list(map(lambda x: int(x), string_data))
    ans = joltage(data)
    print("the product of three_j and one_j differences is {}".format(ans))
