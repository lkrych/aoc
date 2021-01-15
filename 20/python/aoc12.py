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
        ship = calculate_dir(ship, dir, mag)

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


def calculate_dir2(ship, dir, mag):
    dir_translate = {"N": 0, "E": 1, "S": 2, "W": 3}
    mag_translate = {90: 1, 180: 2, 270: 3, 360: 4}
    translate_dir = {0: "N", 1: "E", 2: "S", 3: "W"}
    dir_int_1 = dir_translate[ship.waypoint_curr_dir["h"]]
    dir_int_2 = dir_translate[ship.waypoint_curr_dir["v"]]
    if dir == "R":
        dir_int_1 += mag_translate[mag]
        dir_int_2 += mag_translate[mag]
    elif dir == "L":
        dir_int_1 -= mag_translate[mag]
        dir_int_2 -= mag_translate[mag]
        
    new_dir1 = translate_dir[dir_int_1 % 4]
    new_dir2 = translate_dir[dir_int_2 % 4]
    # print("new_dir_1: {}, new_dir_2: {}".format(new_dir1, new_dir2))

    if new_dir1 == "E" or new_dir1 == "W":
        # if the dir1 variable didn't change orientation, 
        # we are rotating by 180 or 360 degrees

        # handle changing the waypoint dir
        ship.waypoint_curr_dir["h"] = new_dir1
        ship.waypoint_curr_dir["v"] = new_dir2

    else:
       
        temporary_horizontal = ship.waypoint_h_pos
        ship.waypoint_h_pos = ship.waypoint_v_pos
        # if the horizontal didn't change, the vertical had to 
        ship.waypoint_v_pos = temporary_horizontal

        ship.waypoint_curr_dir["h"] = new_dir2
        ship.waypoint_curr_dir["v"] = new_dir1
    
    ship = align_waypoint_dir(ship)
    
    return ship

def align_waypoint_dir_by_pos(ship):
    if ship.waypoint_v_pos > 0:
        ship.waypoint_curr_dir["v"] = "N"
    elif ship.waypoint_v_pos < 0:
        ship.waypoint_curr_dir["v"] = "S"

    if ship.waypoint_h_pos > 0:
        ship.waypoint_curr_dir["h"] = "E"
    elif ship.waypoint_h_pos < 0:
        ship.waypoint_curr_dir["h"] = "W"
    
    return ship

def align_waypoint_dir(ship):
    # print("align_waypoint: {}".format(ship.waypoint_curr_dir))
    # print("align_waypoint: v_pos: {}, h_pos: {}".format(ship.waypoint_v_pos, ship.waypoint_h_pos))
    if ship.waypoint_curr_dir["v"] == "N":
        ship.waypoint_v_pos = abs(ship.waypoint_v_pos)
    else:
        if ship.waypoint_v_pos > 0:
            ship.waypoint_v_pos *= -1
    
    if ship.waypoint_curr_dir["h"] == "E":
        ship.waypoint_h_pos = abs(ship.waypoint_h_pos)
    else:
        if ship.waypoint_h_pos > 0:
            ship.waypoint_h_pos *= -1
    
    return ship


def move_dir2(ship, dir, mag):

    # the easy ones
    if dir == "N":
        ship.waypoint_v_pos += mag
        ship = align_waypoint_dir_by_pos(ship)
    elif dir == "S":
        ship.waypoint_v_pos -= mag
        ship = align_waypoint_dir_by_pos(ship)
    elif dir == "E":
        ship.waypoint_h_pos += mag
        ship = align_waypoint_dir_by_pos(ship)
    elif dir == "W":
        ship.waypoint_h_pos -= mag
        ship = align_waypoint_dir_by_pos(ship)

    # handle forward
    if dir == "F":
        ship.horizontal_pos += (ship.waypoint_h_pos * mag) 
        ship.vertical_pos += (ship.waypoint_v_pos * mag)

    # modify the direction of the ship
    if dir == "R" or dir == "L":
        ship.curr_dir = calculate_dir2(ship, dir, mag)

    return ship

class Ship2:
    def __init__(self, waypoint_h_pos, waypoint_v_pos):
        self.vertical_pos = 0
        self.horizontal_pos = 0
        self.waypoint_curr_dir = {"h": "E", "v": "N"}
        self.waypoint_v_pos = waypoint_v_pos
        self.waypoint_h_pos = waypoint_h_pos

def navigation2(data):
    # waypoint starts 10 units east and 1 unit north
    s = Ship2(10, 1)
    for line in data:
        (new_dir, magnitude) = parse_dir(line)
        # print("{}{}".format(new_dir, magnitude))
        s = move_dir2(s, new_dir, magnitude)
        # print('dir: {}, vert: {}, hor: {}'.format(s.waypoint_curr_dir, s.waypoint_v_pos, s.waypoint_h_pos))
        # print("curr_pos: vert: {}, hor:{}".format(s.vertical_pos, s.horizontal_pos))
    return s.horizontal_pos + abs(s.vertical_pos)
        

if __name__ == "__main__":
    text_file = open("../input/aoc12.txt", "r")
    string_data = text_file.readlines()
    ans = navigation2(string_data)
    print("the manhattan distance is {}".format(ans))
