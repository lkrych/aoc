def count_parents(lug_map, target):
    to_search = [target]
    prospective_parents = {}
    while len(to_search) > 0:
        current_target = to_search.pop()
        if current_target in lug_map:
            targets = lug_map[current_target]
        else:
            continue

        for new_target in targets:
            prospective_parents[new_target] = True
            to_search.append(new_target)
    return len(prospective_parents)

def luggage(data):
    luggage_map = {}
    for line in data:
        line = line.strip()
        split_by_contain = line.split("bags contain")
        parent_bag = split_by_contain[0].strip()
        if "no other bags" in split_by_contain[1]:
            luggage_map[parent_bag] = []
            continue
        split_by_comma = split_by_contain[1].split(",")
        for bag in split_by_comma:
            bag = bag.strip()
            bag_split = bag.split(" ")
            child_bag = bag_split[1] + " " + bag_split[2]
            if child_bag in luggage_map:
                luggage_map[child_bag].append(parent_bag)
            else:
                luggage_map[child_bag] = [parent_bag]
        
    return count_parents(luggage_map, "shiny gold")



def count_children(lug_map, target):
    # assume target in map
    if len(lug_map[target]) == 0:
        return 0
    count = 0
    for curr_target in lug_map[target]:
        count += int(curr_target[1])
        count += (int(curr_target[1]) * count_children(lug_map, curr_target[0]))
    return count

def luggage2(data):
    luggage_map = {}
    for line in data:
        line = line.strip()
        split_by_contain = line.split("bags contain")
        parent_bag = split_by_contain[0].strip()
        if "no other bags" in split_by_contain[1]:
            luggage_map[parent_bag] = []
            continue
        split_by_comma = split_by_contain[1].split(",")
        for bag in split_by_comma:
            bag = bag.strip()
            bag_split = bag.split(" ")
            child_bag = bag_split[1] + " " + bag_split[2]
            child_bag_count = bag_split[0]
            if parent_bag in luggage_map:
                luggage_map[parent_bag].append((child_bag, child_bag_count))
            else:
                luggage_map[parent_bag] = [(child_bag, child_bag_count)]
        
    return count_children(luggage_map, "shiny gold")

if __name__ == "__main__":
    text_file = open("../input/aoc7.txt", "r")
    string_data = text_file.readlines()
    ans = luggage2(string_data)
    print("the number of bags a shiny gold bags contains is {}".format(ans))
