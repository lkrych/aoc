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

def sum_xmas(data, low, high):
    sum = 0
    # actually include the high index because range is exclusive
    i = low
    while i <= high:
        sum += data[i]
        i+=1
    return sum

def xmas2(data, target):
    running_sum = 0
    low_idx = 0
    high_idx = 0
    while True:
        running_sum = sum_xmas(data, low_idx, high_idx)
        if running_sum == target:
            return min(data[low_idx:high_idx]) + max(data[low_idx:high_idx])
        elif running_sum < target:
            high_idx += 1
        elif running_sum > target:
            low_idx += 1
    
    

if __name__ == "__main__":
    text_file = open("../input/aoc9.txt", "r")
    string_data = text_file.readlines()
    data = list(map(lambda x: int(x), string_data))
    ans = xmas2(data, 22477624)
    print("the number that breaks the xmas protocol is {}".format(ans))
