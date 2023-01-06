# read from file input
file_name = input()
f = open(f"/Users/lkrych/aoc/22/input/{file_name}.txt", "r")

# helper method for finding shared val in group
def get_shared_val(group1, group2):
	g1_dic = {}
	g2_dic = {}
	
	for char in group1:
		g1_dic[char] = True
	
	for char in group2:
		g2_dic[char] = True

	for k in g1_dic.keys():
		if k in g2_dic:
			return k

sum = 0
for line in f:
	l = line.strip()
	middle = len(l)//2
	first_compartment = l[:middle]
	second_compartment = l[middle:]
	val = get_shared_val(first_compartment, second_compartment)
	if val.isupper():
		# ascii A = 65, 65 - 38 = 27
		sum += ord(val) - 38
	else:
		# ascii a = 97
		sum += ord(val) - 96

print(sum)


