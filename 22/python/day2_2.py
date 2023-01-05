# read from file input
file_name = input()
f = open(f"/Users/lkrych/aoc/22/input/{file_name}.txt", "r")

#set up variables
score = 0
ROCK = "rock"
PAPER = "paper"
SCISSORS = "scissors"
LOSE = "lose"
DRAW = "draw"
WIN = "win"

rps_dict = {
    "A": ROCK,
    "B": PAPER,
    "C": SCISSORS,
    "X": LOSE,
    "Y": DRAW,
    "Z": WIN
}

for line in f:
    l = line.strip()
    hands = l.split(" ")
    opponent_hand = rps_dict.get(hands[0])
    result = rps_dict.get(hands[1])

    if opponent_hand == ROCK and result == LOSE:
        score += 0
        score += 3
    elif opponent_hand == ROCK and result == DRAW:
        score += 3
        score += 1
    elif opponent_hand == ROCK and result == WIN:
        score += 6
        score += 2
    elif opponent_hand == PAPER and result == DRAW:
        score += 3
        score += 2
    elif opponent_hand == PAPER and result == LOSE:
        score += 0
        score += 1
    elif opponent_hand == PAPER and result == WIN:
        score += 6
        score += 3
    elif opponent_hand == SCISSORS and result == LOSE:
        score += 0
        score += 2
    elif opponent_hand == SCISSORS and result == DRAW:
        score += 3
        score += 3
    elif opponent_hand == SCISSORS and result == WIN:
        score += 6
        score += 1

print(score)
