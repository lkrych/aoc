def two_sum(arr, target):
    hash = {}
    for el in arr:
        diff = target - el
        if diff in hash:
            return (el, diff)
        else:
            hash[el] = True
    return ()

def xmas(data, window):
    for i in range(window, len(data)):
        curr_val = data[i]
        sum = two_sum(data[i - window: i + window - 1], curr_val)
        if len(sum) == 0:
            return curr_val


if __name__ == "__main__":
    text_file = open("../input/aoc9.txt", "r")
    string_data = text_file.readlines()
    data = list(map(lambda x: int(x), string_data))
    ans = xmas(data, 25)
    print("the number that breaks the xmas protocol is {}".format(ans))
