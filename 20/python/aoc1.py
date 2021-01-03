def two_sum(arr, target):
    hash = {}
    for el in arr:
        diff = target - el
        if diff in hash:
            return (el, diff)
        else:
            hash[el] = True
    return ()


if __name__ == "__main__":
    text_file = open("../input/aoc1.txt", "r")
    string_data = text_file.readlines()
    data = map(lambda x: int(x), string_data)
    ans = two_sum(data, 2020)
    if len(ans) > 0 :
        print("product: {}".format(ans[0] * ans[1]))