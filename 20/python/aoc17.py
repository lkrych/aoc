
def active_cubes(data):
    # needs to represent three dimensions
    
    current_state = []
    for row in data:



if __name__ == "__main__":
    text_file = open("../input/aoc17test.txt", "r")
    string_data = text_file.readlines()
    ans = active_cubes(string_data)
    print(f"the number of active cubes is {ans}")