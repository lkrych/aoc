#include <algorithm>
#include <cstdio>
#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"


// fill matrix
std::vector<std::vector<char>> processSeats(std::vector<std::string> arr) {
    std::vector<std::vector<char>> matrix = {};
    for(int i = 0; i < arr.size(); i++) {
        std::string row = arr[i];
        std::vector<char> matrix_row;
        for (int j = 0; j < row.size(); j++) {
            char c = row[j];
            matrix_row.push_back(c);
        }
        matrix.push_back(matrix_row);
    }
    return matrix;
}

void writeMatrixToFile(std::vector<std::vector<char>> matrix, std::string filename) {
    std::ofstream myfile (filename);
    if (myfile.is_open())
    {
        for(int i = 0; i < matrix.size(); i++ ) {
            std::vector<char> row = matrix[i];
            for(int j = 0; j < row.size(); j++) {
                myfile << row[j];
            }
            myfile << "\n";
        }
        myfile.close();
    }
}

// //part 1
// std::vector<std::vector<char>> boardingRound(std::vector<std::vector<char>> matrix) {
//     std::vector<std::vector<char>> m = copyMatrix(matrix);
//     for (int i = 0; i < matrix.size(); i++) {
//         std::vector<char> row = matrix[i];
//         for (int j = 0; j < row.size(); j++) {
//             int occupied = 0;
//             char c = matrix[i][j];
//             if (c != '#' && c != 'L') {
//                 continue;
//             }
//             // check left
//             // std::cout << "checking left"  << std::endl;
//             if (j - 1 >= 0) {
//                 if (matrix[i][j-1] == '#') {
//                     occupied++;
//                 }
//             }
//             // check right
//             // std::cout << "checking right"  << std::endl;
//             if (j + 1 <= row.size() - 1) {
//                 if (matrix[i][j+1] == '#') {
//                     occupied++;
//                 }
//             }
//             // check up
//             // std::cout << "checking up"  << std::endl;
//             if (i - 1 >= 0) {
//                 if (matrix[i-1][j] == '#') {
//                     occupied++;
//                 }
//                 //diagonal left
//                 //  std::cout << "checking diagonal up and left"  << std::endl;
//                 if (j - 1 >= 0) {
//                     if (matrix[i-1][j-1] == '#') {
//                         occupied++;
//                     }
//                 }
//                 // diagonal right
//                 // std::cout << "checking diagonal up and right"  << std::endl;
//                 if (j + 1 <= row.size() - 1) {
//                     if (matrix[i-1][j+1] == '#') {
//                         occupied++;
//                     }
//                 }
//             }
//             // check down
//             // std::cout << "checking down"  << std::endl;
//             if (i + 1 <= matrix.size() - 1) {
//                 if (matrix[i+1][j] == '#') {
//                     occupied++;
//                 }
//                 // diagonal left
//                 // std::cout << "checking diagonal left and down"  << std::endl;
//                 if (j - 1 >= 0) {
//                     if (matrix[i+1][j-1] == '#') {
//                         occupied++;
//                     }
//                 }
//                 // diagonal right
//                 // std::cout << "checking diagonal right and down"  << std::endl;
//                 if (j + 1 <= row.size() - 1) {
//                     if (matrix[i+1][j+1] == '#') {
//                         occupied++;
//                     }
//                 }
//             }
//             // process rules
//             // manipulate m so that we don't mutate as we are processing
//             if (occupied == 0 && c == 'L') {
//                 m[i][j] = '#';
//             }
//             if (occupied >= 4 && c == '#') {
//                 m[i][j] = 'L';
//             }
//         }
//     }
//     return m;
// }

char searchLeft(std::vector<std::vector<char>> *matrix, int y, int x, int limitX) {
    // std::cout << "searchLeft" << std::endl;
    if (x < limitX) {
        // dummy character
        return 'X';
    }
    if ((*matrix)[y][x] == 'L' || (*matrix)[y][x] == '#') {
        return (*matrix)[y][x];
    }
    return searchLeft(matrix, y, x-1, limitX);
}
char searchRight(std::vector<std::vector<char>> *matrix, int y, int x, int limitX) {
    // std::cout << "searchRight" << std::endl;
    if (x > limitX) {
        // dummy character
        return 'X';
    }
    if ((*matrix)[y][x] == 'L' || (*matrix)[y][x] == '#') {
        return (*matrix)[y][x];
    }
    return searchRight(matrix, y, x+1, limitX);
}
char searchUp(std::vector<std::vector<char>> *matrix, int y, int x, int limitY) {
    // std::cout << "searchUp" << std::endl;
    if (y < limitY) {
        // dummy character
        return 'Y';
    }
    if ((*matrix)[y][x] == 'L' || (*matrix)[y][x] == '#') {
        return (*matrix)[y][x];
    }
    return searchUp(matrix, y-1, x, limitY);
}
char searchDown(std::vector<std::vector<char>> *matrix, int y, int x, int limitY) {
    // std::cout << "searchDown" << std::endl;
    if (y > limitY) {
        // dummy character
        return 'Y';
    }
    
    if ((*matrix)[y][x] == 'L' || (*matrix)[y][x] == '#') {
        return (*matrix)[y][x];
    }
    return searchDown(matrix, y+1, x, limitY);
}
char searchUpLeft(std::vector<std::vector<char>> *matrix, int y, int x, int limitX, int limitY) {
    // std::cout << "searchUpLeft" << std::endl;
    if (x < limitX) {
        // dummy character
        return 'X';
    }
    if (y < limitY) {
        // dummy character
        return 'Y';
    }
    if ((*matrix)[y][x] == 'L' || (*matrix)[y][x] == '#') {
        return (*matrix)[y][x];
    }
    return searchUpLeft(matrix, y-1, x-1, limitX, limitY);
}
char searchUpRight(std::vector<std::vector<char>> *matrix, int y, int x, int limitX, int limitY) {
    // std::cout << "searchRight" << std::endl;
    if (x > limitX) {
        // dummy character
        return 'X';
    }
    if (y < limitY) {
        // dummy character
        return 'Y';
    }
    if ((*matrix)[y][x] == 'L' || (*matrix)[y][x] == '#') {
        return (*matrix)[y][x];
    }
    return searchUpRight(matrix, y-1, x+1, limitX, limitY);
}
char searchDownLeft(std::vector<std::vector<char>> *matrix, int y, int x, int limitX, int limitY) {
    // std::cout << "searchDownLeft" << std::endl;
    if (x < limitX) {
        // dummy character
        return 'X';
    }
    if (y > limitY) {
        // dummy character
        return 'Y';
    }
    if ((*matrix)[y][x] == 'L' || (*matrix)[y][x] == '#') {
        return (*matrix)[y][x];
    }
    return searchDownLeft(matrix, y+1, x-1, limitX, limitY);
}
char searchDownRight(std::vector<std::vector<char>> *matrix, int y, int x, int limitX, int limitY) {
    // std::cout << "searchDownRight" << std::endl;
    if (x > limitX) {
        // dummy character
        return 'X';
    }
    if (y > limitY) {
        // dummy character
        return 'Y';
    }
    if ((*matrix)[y][x] == 'L' || (*matrix)[y][x] == '#') {
        return (*matrix)[y][x];
    }
    return searchDownRight(matrix, y+1, x+1, limitX, limitY);
}

std::vector<std::vector<char>> boardingRound2(std::vector<std::vector<char>> matrix) {
    std::vector<std::vector<char>> m = copyMatrix(matrix);
    for (int i = 0; i < matrix.size(); i++) {
        std::vector<char> row = matrix[i];
        for (int j = 0; j < row.size(); j++) {
            int occupied = 0;
            char c = matrix[i][j];
            if (c != '#' && c != 'L') {
                continue;
            }
            // std::cout << "for i " << i << ", j " << j << std::endl;
            char left = searchLeft(&matrix, i, j-1, 0);
            // std::cout << "searchLeft "  << left << std::endl;
            if (left == '#') {
                occupied++;
            }
            char right = searchRight(&matrix, i, j+1, row.size() - 1);
            // std::cout << "searchRight "  << right << std::endl;
            if (right == '#') {
                occupied++;
            }
            char up = searchUp(&matrix, i-1, j, 0);
            // std::cout << "searchUp "  << up << std::endl;
            if (up == '#') {
                occupied++;
            }
            char upleft = searchUpLeft(&matrix, i-1,j-1, 0, 0);
            // std::cout << "upLeft "  << upleft << std::endl;
            if (upleft == '#') {
                occupied++;
            }
            char upright = searchUpRight(&matrix, i-1, j+1, row.size() - 1, 0);
            // std::cout << "upright "  << upright << std::endl;
            if (upright == '#') {
                occupied++;
            }
            char down = searchDown(&matrix, i+1,j, matrix.size() - 1);
            // std::cout << "down "  << down << std::endl;
            if (down == '#') {
                occupied++;
            }
            char downleft = searchDownLeft(&matrix, i+1,j-1, 0, matrix.size() - 1);
            // std::cout << "downleft "  << downleft << std::endl;
            if (downleft == '#') {
                occupied++;
            }
            char downright = searchDownRight(&matrix, i+1,j+1, row.size() - 1, matrix.size() - 1);
            // std::cout << "downright "  << downright << std::endl;
            if (downright == '#') {
                occupied++;
            }
            
            
            // process rules
            // manipulate m so that we don't mutate as we are processing
            if (occupied == 0 && c == 'L') {
                m[i][j] = '#';
            }
            if (occupied >= 5 && c == '#') {
                m[i][j] = 'L';
            }
        }
    }
    return m;
}

int seatCount (std::vector<std::vector<char>> matrix) {
    int occupied = 0;
    int unoccupied = 0;
    int floor = 0;
    for (int i = 0; i < matrix.size(); i++) {
        std::vector<char> row = matrix[i];
        for (int j = 0; j < row.size(); j++) {
            char c = matrix[i][j];
            if (c == '#') {
                occupied++;
            } else if (c == 'L') {
                unoccupied++;
            } else if (c == '.') {
                floor++;
            }
        }
    }
    // std::cout << "occ: " << occupied << std::endl;
    // std::cout << "unocc: " << unoccupied << std::endl;
    // std::cout << "floor: " << floor << std::endl;
    // std::cout << "----------" << std::endl;
    return occupied;
}

int stableSeats(std::vector<std::string> arr) {
    std::vector<std::vector<char>> matrix = processSeats(arr);
    seatCount(matrix);
    writeMatrixToFile(matrix, "matrix0.txt");
    std::map<int, int> stability_map;
    int prev_seats = 0;
    int round = 0;
    while (true) {
        round++;
        // std::cout << "-------- round " << round << "------------";
        matrix = boardingRound2(matrix);
        int occupied_seats = seatCount(matrix);
        // char filename[80];
        // sprintf(filename, "matrix%d.txt", round);
        // writeMatrixToFile(matrix, filename);

        if (stability_map.find(occupied_seats) != stability_map.end()) {
            stability_map[occupied_seats] += 1;
        } else {
            stability_map[occupied_seats] = 1;
            prev_seats = occupied_seats;
            continue;
        }
        if (occupied_seats == prev_seats && stability_map[occupied_seats] > 5) {
            return occupied_seats;
        }
        prev_seats = occupied_seats;

    }
}

int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc11.txt");

    int stability = stableSeats(data);
    std::cout << "the number of stable seats is " << stability << std::endl;
}