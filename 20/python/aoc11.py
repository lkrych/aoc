import copy

def matrixize(data):
    matrix = []
    for line in data:
        line = line.strip()
        # split string into list
        split = list(line)
        matrix.append(split)
    return matrix

def occupied_count(data):
    seats = 0
    for line in data:
        for seat in line:
            if seat == "#":
                seats += 1
    return seats

def check_seat(data, x, y):
    curr_val = data[y][x]
    occupied = 0
    min_x = 0
    max_x = len(data[y]) - 1
    min_y = 0
    max_y = len(data) - 1
    
    # check upper rows
    if y - 1 >= min_y:
        if x - 1 >= min_x:
            if data[y-1][x-1] == "#":
                occupied += 1
        if data[y-1][x] == "#":
            occupied += 1
        if x +1 <=  max_x:
            if data[y-1][x+1] == "#":
                occupied += 1

    # check left and right

    if x - 1 >= min_x:
        if data[y][x-1] == "#":
                occupied += 1
    if x + 1 <=  max_x:
            if data[y][x+1] == "#":
                occupied += 1

    # check bottom rows

    if y + 1 <= max_y:
        if x - 1 >= min_x:
            if data[y+1][x-1] == "#":
                occupied += 1
        if data[y+1][x] == "#":
            occupied += 1
        if x +1 <=  max_x:
            if data[y+1][x+1] == "#":
                occupied += 1

    # change rules

    if curr_val == "L" and occupied == 0:
        return "#"
    elif curr_val == "#" and occupied >= 4:
        return "L"
    else:
        return curr_val

def check_up(data, x, y):
    #check the value if it's possible
    if y - 1 >= 0:
        val = data[y-1][x]
        if val == "#" or val == "L":
            return val
        else:
            return check_up(data, x, y-1)
    else:
        return ""

def check_down(data, x, y):
    #check the value if it's possible
    if y + 1 < len(data):
        val = data[y+1][x]
        if val == "#" or val == "L":
            return val
        else:
            return check_down(data, x, y+1)
    else:
        return ""

def check_left(data, x, y):
    #check the value if it's possible
    if x - 1 >= 0:
        val = data[y][x - 1]
        if val == "#" or val == "L":
            return val
        else:
            return check_left(data, x - 1, y)
    else:
        return ""

def check_right(data, x, y):
    #check the value if it's possible
    if x + 1 < len(data[y]):
        val = data[y][x+1]
        if val == "#" or val == "L":
            return val
        else:
            return check_right(data, x+1, y)
    else:
        return ""

def check_up_left(data, x, y):
    #check the value if it's possible
    if y - 1 >= 0 and x - 1 >= 0:
        val = data[y-1][x-1]
        if val == "#" or val == "L":
            return val
        else:
            return check_up_left(data, x-1, y-1)
    else:
        return ""

def check_up_right(data, x, y):
    #check the value if it's possible
    if y - 1 >= 0 and x + 1 < len(data[y]):
        val = data[y-1][x+1]
        if val == "#" or val == "L":
            return val
        else:
            return check_up_right(data, x+1, y-1)
    else:
        return ""

def check_down_left(data, x, y):
    if y + 1 < len(data) and x - 1 >= 0:
        val = data[y+1][x-1]
        if val == "#" or val == "L":
            return val
        else:
            return check_down_left(data, x-1, y+1)
    else:
        return ""

def check_down_right(data, x, y):
    if y + 1 < len(data) and x + 1 < len(data[y]):
        val = data[y+1][x+1]
        if val == "#" or val == "L":
            return val
        else:
            return check_down_right(data, x+1, y+1)
    else:
        return ""

def check_seat2(data, x, y):
    curr_val = data[y][x]
    occupied = 0
    
    # check upper rows
    if check_up_left(data, x, y) == "#":
        occupied += 1
    if check_up(data, x, y) == "#":
        occupied += 1
    if check_up_right(data, x, y) == "#":
        occupied += 1

    # check left and right
    if check_left(data, x, y) == "#":
        occupied += 1
    if check_right(data, x, y) == "#":
        occupied += 1

    # check bottom rows

    if check_down_left(data, x, y) == "#":
        occupied += 1
    if check_down(data, x, y) == "#":
        occupied += 1
    if check_down_right(data, x, y) == "#":
        occupied += 1


    # change rules

    if curr_val == "L" and occupied == 0:
        return "#"
    elif curr_val == "#" and occupied >= 5:
        return "L"
    else:
        return curr_val


# to do part 1, just swap line 126 check_seat for check_seat2
def create_new_seating_map(data):
    x,y = 0,0
    # for multi-dimensional array use deepcopy
    new_seating_map = copy.deepcopy(data)
    for line in data:
        for seat in line:
            new_seat = check_seat2(data, x, y)
            new_seating_map[y][x] = new_seat
            #increment seat counter x
            x += 1
        #increment seat counter y
        # print(line)
        y += 1
        x = 0
    # print("")
    return new_seating_map

def seating(data):
    data = matrixize(data)
    prev_seat_count = 0
    current_seat_count = occupied_count(data)
    while True:
        prev_seat_count = current_seat_count
        data = create_new_seating_map(data)
        current_seat_count = occupied_count(data)
        # print("seat count is {}".format(current_seat_count))
        if current_seat_count == prev_seat_count:
            return current_seat_count

if __name__ == "__main__":
    text_file = open("../input/aoc11.txt", "r")
    string_data = text_file.readlines()
    ans = seating(string_data)
    print("the number of occupied is {}".format(ans))
