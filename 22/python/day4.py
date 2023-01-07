# read from file input
file_name = input()
f = open(f"/Users/lkrych/aoc/22/input/{file_name}.txt", "r")

def split_pairs(p):
	return p.split("-")

overlaps = 0
for line in f:
	l = line.strip()
	pairs = l.split(",")
	p1 = pairs[0]
	p2 = pairs[1]
	p1_range = split_pairs(p1)
	p2_range = split_pairs(p2)
	# p1 range encloses p2
	if int(p1_range[0]) <= int(p2_range[0]) and int(p1_range[1]) >= int(p2_range[1]):
		overlaps += 1
	# p2 range encloses p1
	elif int(p2_range[0]) <= int(p1_range[0]) and int(p2_range[1]) >= int(p1_range[1]):
		overlaps += 1

print(overlaps)
	
