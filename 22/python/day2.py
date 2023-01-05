# read from file input
file_name = input()
f = open(f"/Users/lkrych/aoc/22/input/{file_name}.txt", "r")

#set up variables
score = 0
ROCK = "rock"
PAPER = "paper"
SCISSORS = "scissors"

rps_dict = {
    "A": ROCK,
    "B": PAPER,
    "C": SCISSORS,
    "X": ROCK,
    "Y": PAPER,
    "Z": SCISSORS
}

for line in f:
    l = line.strip()
    hands = l.split(" ")
    opponent_hand = rps_dict.get(hands[0])
    your_hand = rps_dict.get(hands[1])

    if opponent_hand == ROCK and your_hand == ROCK:
        score += 3
        score += 1
    elif opponent_hand == ROCK and your_hand == PAPER:
        score += 6
        score += 2
    elif opponent_hand == ROCK and your_hand == SCISSORS:
        score += 0
        score += 3
    elif opponent_hand == PAPER and your_hand == PAPER:
        score += 3
        score += 2
    elif opponent_hand == PAPER and your_hand == ROCK:
        score += 0
        score += 1
    elif opponent_hand == PAPER and your_hand == SCISSORS:
        score += 6
        score += 3
    elif opponent_hand == SCISSORS and your_hand == ROCK:
        score += 6
        score += 1
    elif opponent_hand == SCISSORS and your_hand == PAPER:
        score += 0
        score += 2
    elif opponent_hand == SCISSORS and your_hand == SCISSORS:
        score += 3
        score += 3

print(score)
