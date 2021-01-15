def parse_dir(direction):
    direction = direction.strip()
    dir = direction[0]
    mag = int(direction[1:])
    return (dir, mag)

def calculate_dir(ship, dir, mag):
    dir_translate = {"N": 0, "E": 1, "S": 2, "W": 3}
    mag_translate = {90: 1, 180: 2, 270: 3, 360: 4}
    translate_dir = {0: "N", 1: "E", 2: "S", 3: "W"}
    dir_int = dir_translate[ship.curr_dir]
    if dir == "R":
        dir_int += mag_translate[mag]
    elif dir == "L":
        dir_int -= mag_translate[mag]
    return translate_dir[dir_int % 4]
    
    

def move_dir(ship, dir, mag):

    # the easy ones
    if dir == "N":
        ship.vertical_pos += mag
    elif dir == "S":
         ship.vertical_pos -= mag
    elif dir == "E":
         ship.horizontal_pos += mag
    elif dir == "W":
         ship.horizontal_pos -= mag
    elif dir == "F":
        ship = move_dir(ship, ship.curr_dir, mag)

    # modify the direction of the ship
    if dir == "R" or dir == "L":
        ship.curr_dir = calculate_dir(ship, dir, mag)

    return ship

class Ship:
    def __init__(self, curr_dir):
        self.curr_dir = curr_dir
        self.vertical_pos = 0
        self.horizontal_pos = 0

def navigation(data):
    s = Ship("E")
    for line in data:
        (new_dir, magnitude) = parse_dir(line)
        # print("{}{}".format(new_dir, magnitude))
        s = move_dir(s, new_dir, magnitude)
        # print('dir: {}, vert: {}, hor: {}'.format(s.curr_dir, s.vertical_pos, s.horizontal_pos))
    
    return abs(s.horizontal_pos) + abs(s.vertical_pos)
        

if __name__ == "__main__":
    text_file = open("../input/aoc12.txt", "r")
    string_data = text_file.readlines()
    ans = navigation(string_data)
    print("the manhattan distance is {}".format(ans))
