from itertools import filterfalse
from pathlib import Path

file_dir = Path(__file__).resolve().parent

def get_lines(filename):
    with open(f"{file_dir}/../input/{filename}") as f:
        return list(f)

# The score of the winning board can now be calculated. 
# Start by finding the sum of all unmarked numbers on that board; 
# Then, multiply that sum by the number that was just called when the board won, to get the final score, 

def bingo(lines):
    bingo_input = lines[0].strip()
    boards = []
    new_board = {}
    row = 0
    for line in lines[1:]:
        # iterate until break
        if not line:
            boards.append(new_board)
            new_board = {}
            row = 0
        else:
            for idx, val in enumerate(line.strip().split()):
                if not val:
                    continue
                # row, col, marked
                new_board[val] = [row, idx, False]
            row += 1
    # catch the board that's being filled up at the end
    boards.append(new_board)

    #filter boards to make sure we aren't including empty boards
    boards = [board for board in boards if len(board) > 0]
    print(f"boards: {boards}")

    # play out bingo
    for bingo_draw in bingo_input.split(","):
        # print(f"bingo_draw: {bingo_draw}")
        for board in boards:
            # print(f"checking board: {board}")
            if bingo_draw in board:
                board_num = board.get(bingo_draw)
                board_num[2] = True
                board[bingo_draw] = board_num

        # part 2 filter until we have the last board
        if len(boards) == 1:
            board = boards[0]
            sum = 0
            for k,v in board.items():
                if not v[2]:
                    sum += int(k)
            return sum * int(bingo_draw)
        boards = list(filterfalse(check_if_bingo, boards))

# board is a hash_map of bingo_numbers mapped to a list of (row, col, idx)
def check_if_bingo(board):
    # populate board
    pos_board = [[False for _ in range(5)] for _ in range(5)]
    for _, v in board.items():
        pos_board[v[0]][v[1]] = v[2]
    
    # check all rows
    for row_val in range(5):
        if all(pos_board[row_val]):
            # print(f"bingo! :{pos_board[row_val]}")
            return True
            
    
    # check all columns
    col = []
    for col_val in range(5):
        for row_val in range(5):
            col.append(pos_board[row_val][col_val])

        if all(col):
            # print(f"bingo! : {col}")
            return True

    return False


if __name__ == "__main__":
    lines = get_lines("day4.txt")
    lines = [line.strip() for line in lines]
    print(bingo(lines))