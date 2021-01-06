# def toboggan(data):
#     trees = 0
#     curr_pos = 0
#     for line in data:
#         line = line.strip()
#         if line[curr_pos] == '#':
#             trees += 1
#         curr_pos = (curr_pos + 3) % len(line)
#     return trees

# if __name__ == "__main__":
#     text_file = open("../input/aoc3.txt", "r")
#     string_data = text_file.readlines()
#     ans = toboggan(string_data)
#     print("the number of trees encountered is {}".format(ans))

def run_toboggan_sim(data, route):
    trees = 0
    curr_pos_x = 0
    curr_pos_y = 0
    while curr_pos_y < len(data):
        line = data[curr_pos_y]
        line = line.strip()
        if line[curr_pos_x] == '#':
            trees += 1
        curr_pos_x = (curr_pos_x + route[0]) % len(line)
        curr_pos_y += route[1]
    return trees

def toboggan2(data):
    trees = 1
    routes = [(1,1), (3,1), (5,1), (7,1), (1, 2)]
    for route in routes:
        trees *= run_toboggan_sim(data, route)
    return trees

if __name__ == "__main__":
    text_file = open("../input/aoc3.txt", "r")
    string_data = text_file.readlines()
    ans = toboggan2(string_data)
    print("the number of trees encountered is {}".format(ans))

   