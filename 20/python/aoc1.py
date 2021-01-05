def two_sum(arr, target):
    hash = {}
    for el in arr:
        diff = target - el
        if diff in hash:
            return (el, diff)
        else:
            hash[el] = True
    return ()

def three_sum(arr, target):
    arr = sorted(arr)
    bottom_idx = 0
    top_idx = len(arr) - 1
    while bottom_idx < top_idx :
        for i in range(bottom_idx + 1, top_idx):
            sum = arr[bottom_idx] + arr[i] + arr[top_idx]
            if sum == target:
                return (arr[bottom_idx], arr[i], arr[top_idx])
            elif sum > target:
                if (i - bottom_idx > 1):
                    bottom_idx += 1
                else:
                    top_idx -= 1
                break
            elif top_idx - i <= 1:
                bottom_idx += 1
                break
        
    print("No three sum solution was found :(")
    return ()

# part 1 - two sum
# if __name__ == "__main__":
#     text_file = open("../input/aoc1.txt", "r")
#     string_data = text_file.readlines()
#     data = map(lambda x: int(x), string_data)
#     ans = two_sum(data, 2020)
#     if len(ans) > 0 :
#         print("product: {}".format(ans[0] * ans[1]))

if __name__ == "__main__":
    text_file = open("../input/aoc1.txt", "r")
    string_data = text_file.readlines()
    data = map(lambda x: int(x), string_data)
    ans = three_sum(data, 2020)
    if len(ans) > 0:
        print("product: {}".format(ans[0] * ans[1] * ans[2]))