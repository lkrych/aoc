# read from file input
file_name = input()
f = open(f"/Users/lkrych/aoc/22/input/{file_name}.txt", "r")

# helper method for finding shared val in group
def get_shared_val(group1, group2, group3):
	g1_dic = {}
	g2_dic = {}
	g3_dic = {}

	for char in group1:
		g1_dic[char] = True
	
	for char in group2:
		g2_dic[char] = True
	
	for char in group3:
		g3_dic[char] = True

	for k in g1_dic.keys():
		if k in g2_dic:
			if k in g3_dic:
				return k

sum = 0
lines = []
for line in f:
    l = line.strip()
    lines.append(l)
    if len(lines) == 3:
	    val = get_shared_val(lines[0], lines[1], lines[2])
		# reset lines
	    lines = []
	    if val.isupper():
		    # ascii A = 65, 65 - 38 = 27
		    sum += ord(val) - 38
	    else:
		    # ascii a = 97
		    sum += ord(val) - 96

print(sum)


