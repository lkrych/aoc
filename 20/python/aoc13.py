def shuttle1(data):
    earliest_departure = int(data[0].strip())
    split = data[1].strip().split(",")
    possible_buses = []
    for el in split:
        if el == "x":
            continue
        else:
            possible_buses.append(int(el))
    min_diff = 100000
    min_bus = 0
    for b in possible_buses:
        base = earliest_departure // b
        diff = 0
        if base == 0:
            diff = b - earliest_departure
        else:
            diff = (b * (base + 1)) % earliest_departure
        print (f"bus: {b}, base: {base}, diff: {diff}")
        if diff < min_diff:
            min_diff = diff
            min_bus = b
    
    return min_diff * min_bus
        

if __name__ == "__main__":
    text_file = open("../input/aoc13.txt", "r")
    string_data = text_file.readlines()
    ans = shuttle1(string_data)
    print(f"the earliest shuttle you can take is {ans}")