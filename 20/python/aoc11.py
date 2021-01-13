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


def create_new_seating_map(data):
    x,y = 0,0
    # for multi-dimensional array use deepcopy
    new_seating_map = copy.deepcopy(data)
    for line in data:
        for seat in line:
            new_seat = check_seat(data, x, y)
            new_seating_map[y][x] = new_seat
            #increment seat counter x
            x += 1
        #increment seat counter y
        y += 1
        x = 0
    return new_seating_map

def seating(data):
    data = matrixize(data)
    prev_seat_count = 0
    current_seat_count = occupied_count(data)
    while True:
        prev_seat_count = current_seat_count
        data = create_new_seating_map(data)
        current_seat_count = occupied_count(data)
        if current_seat_count == prev_seat_count:
            return current_seat_count

if __name__ == "__main__":
    text_file = open("../input/aoc11.txt", "r")
    string_data = text_file.readlines()
    ans = seating(string_data)
    print("the number of occupied is {}".format(ans))
