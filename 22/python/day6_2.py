from collections import deque

# read from file input
file_name = input()
f = open(f"/Users/lkrych/aoc/22/input/{file_name}.txt", "r")

seen_chars = 0
window = deque()
for line in f:
	l = line.strip()
	for c in l:
		seen_chars += 1
		# if the window size is too large remove an item
		if len(window) > 13:
			window.pop()
		window.appendleft(c)
		
		# check if there are any repeats
		duplicate_dict = {}
		duplicates = False
		if len(window) == 14:
			for el in window:
				if el in duplicate_dict:
					duplicates = True
					break
				else:
					duplicate_dict[el] = True
			
			if not duplicates:
				print(f"seen_chars = {seen_chars} for input: {l[:25]}")
				# reset variables
				seen_chars = 0
				window = deque()
				break

