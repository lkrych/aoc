# day1 of advent of code
f = open("/Users/lkrych/aoc/22/input/day1.txt", "r")
calories = 0
most_calories = 0
second_most_calories = 0
third_most_calories = 0
elf_num = 1
elf_most = 1

# iterate over every line
for line in f:
    line = line.strip()
    # if line is empty we need to check to see if this elf is carrying the most calories
    if not line:
        print(f"elf {elf_num}: {calories}")
        if calories > most_calories:
            third_most_calories = second_most_calories
            second_most_calories = most_calories
            most_calories = calories
            elf_most = elf_num
        elif calories > second_most_calories:
            third_most_calories = second_most_calories
            second_most_calories = calories
        elif calories > third_most_calories:
            third_most_calories = calories
        # set calories back to zero and start counting for new elf
        calories = 0
        elf_num += 1
        print()
    else:
        # increment calories
        calories += int(line)

print(f"elf {elf_most} is carrying the most calories: {most_calories}")
print(f"sum of top three calories is {most_calories + second_most_calories + third_most_calories}")

