#include <algorithm>
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

std::vector<std::vector<char>> boardingRound(std::vector<std::vector<char>> matrix) {
    std::vector<std::vector<char>> m = copyMatrix(matrix);
    for (int i = 0; i < matrix.size(); i++) {
        std::vector<char> row = matrix[i];
        for (int j = 0; j < row.size(); j++) {
            int occupied = 0;
            char c = matrix[i][j];
            if (c != '#' && c != 'L') {
                continue;
            }
            // check left
            // std::cout << "checking left"  << std::endl;
            if (j - 1 >= 0) {
                if (matrix[i][j-1] == '#') {
                    occupied++;
                }
            }
            // check right
            // std::cout << "checking right"  << std::endl;
            if (j + 1 <= row.size() - 1) {
                if (matrix[i][j+1] == '#') {
                    occupied++;
                }
            }
            // check up
            // std::cout << "checking up"  << std::endl;
            if (i - 1 >= 0) {
                if (matrix[i-1][j] == '#') {
                    occupied++;
                }
                //diagonal left
                //  std::cout << "checking diagonal up and left"  << std::endl;
                if (j - 1 >= 0) {
                    if (matrix[i-1][j-1] == '#') {
                        occupied++;
                    }
                }
                // diagonal right
                // std::cout << "checking diagonal up and right"  << std::endl;
                if (j + 1 <= row.size() - 1) {
                    if (matrix[i-1][j+1] == '#') {
                        occupied++;
                    }
                }
            }
            // check down
            // std::cout << "checking down"  << std::endl;
            if (i + 1 <= matrix.size() - 1) {
                if (matrix[i+1][j] == '#') {
                    occupied++;
                }
                // diagonal left
                // std::cout << "checking diagonal left and down"  << std::endl;
                if (j - 1 >= 0) {
                    if (matrix[i+1][j-1] == '#') {
                        occupied++;
                    }
                }
                // diagonal right
                // std::cout << "checking diagonal right and down"  << std::endl;
                if (j + 1 <= row.size() - 1) {
                    if (matrix[i+1][j+1] == '#') {
                        occupied++;
                    }
                }
            }
            // process rules
            // manipulate m so that we don't mutate as we are processing
            if (occupied == 0 && c == 'L') {
                m[i][j] = '#';
            }
            if (occupied >= 4 && c == '#') {
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
    std::map<int, int> stability_map;
    int prev_seats = 0;
    while (true) {
        matrix = boardingRound(matrix);
        int occupied_seats = seatCount(matrix);
        if (stability_map.find(occupied_seats) != stability_map.end()) {
            stability_map[occupied_seats] += 1;
        } else {
            stability_map[occupied_seats] = 1;
            prev_seats = occupied_seats;
            continue;
        }
        if (occupied_seats == prev_seats && stability_map[occupied_seats] > 10) {
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