
def seat_id(input):
    upper_bound = pow(2, len(input))
    lower_bound = 0
    for c in input:
        diff = (upper_bound - lower_bound) // 2
        if c == "F" or c == "L":
            upper_bound = upper_bound - diff
        elif c == "B" or c == "R":
            lower_bound = lower_bound + diff
    return upper_bound - 1

def get_seat_id(seat):
    row_id = seat_id(seat[:7])
    col_id = seat_id(seat[7:])
    seat_calc = (row_id * 8) + col_id
    return seat_calc
    

def boarding_pass(data):
    highest_seat_id = 0
    for line in data:
        line = line.strip()
        seat_id = get_seat_id(line)
        if seat_id > highest_seat_id:
            highest_seat_id = seat_id
    return highest_seat_id


if __name__ == "__main__":
    text_file = open("../input/aoc5.txt", "r")
    string_data = text_file.readlines()
    ans = boarding_pass(string_data)
    print("the boarding_pass ID is {}".format(ans))
