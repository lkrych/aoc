# read from file input
file_name = input()
f = open(f"/Users/lkrych/aoc/22/input/{file_name}.txt", "r")

def split_pairs(p):
    pairs = p.split("-")
    pairs = [int(el) for el in pairs]
    return pairs

overlaps = 0
for line in f:
    l = line.strip()
    pairs = l.split(",")
    p1 = pairs[0]
    p2 = pairs[1]
    p1_range = split_pairs(p1) #1,4
    p2_range = split_pairs(p2) #6,18
    #if 1 >= 6 and 1 <= 18
    if p1_range[0] >= p2_range[0] and p1_range[0] <= p2_range[1]:
        overlaps += 1
    #  4 >= 6 and 4 <= 18
    elif p1_range[1] >= p2_range[0] and p1_range[1] <= p2_range[1]:
        overlaps += 1
     # 6 >= 1 and 6 <= 4
    elif p2_range[0] >= p1_range[0] and p2_range[0] <= p1_range[1]:
        overlaps += 1
     # 18 >= 1 and 18 <= 4
    elif p2_range[1] >= p1_range[0] and p2_range[1] <= p1_range[1]:
        overlaps += 1
print(overlaps)
	
