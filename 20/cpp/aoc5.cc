#include <algorithm>
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <math.h>
#include "aoc_helper.h"

int calculateByHalves(std::string seatString) {
    int upperRange = (int)pow(2, seatString.size()) - 1;
    int lowerRange = 0;

    for(int i = 0; i < seatString.size(); i++) {
        // std::cout << "Upper: " << upperRange << " Lower: " << lowerRange << std::endl;
        int diff = (upperRange - lowerRange) / 2 + 1;

        // std::cout << "Diff: " << diff << std::endl;

        char el = seatString[i];
        if (el == 'F' || el == 'L') {
            //take lower half
            upperRange = upperRange - diff;
        } else if (el == 'B' || el == 'R') {
            //take upper half
            lowerRange = lowerRange + diff;
        } else {
            std::cout << "Invalid input: " << el << std::endl;
            exit(1);
        }
    }
    // std::cout << "Upper: " << upperRange << " Lower: " << lowerRange << std::endl;
    return (upperRange + lowerRange) / 2;
}

int calculateSeatId(std::string seat) {
    int row = calculateByHalves(seat.substr(0, 7));
    int col = calculateByHalves(seat.substr(7, 3));

    return (row * 8) + col;
}

// part 1
// int highestSeat(std::vector<std::string> arr) {
//     int max = 0;
//     for (int i = 0; i < arr.size(); i++) {
//         std::string seat = arr[i];
//         int seatId = calculateSeatId(seat);
//         if (seatId > max) {
//             max = seatId;
//         }
//     }
//     return max;
// }

// part 2
void mySeat(std::vector<std::string> arr) {
    std::vector<int> seats;
    for (int i = 0; i < arr.size(); i++) {
        std::string seat = arr[i];
        int seatId = calculateSeatId(seat);
        seats.push_back(seatId);
    }
    std::sort(seats.begin(), seats.end());

    int lastSeat;
    int currentSeat;
    for(int i = 1; i < seats.size(); i++) {
        lastSeat = seats[i-1];
        currentSeat = seats[i];
        if (currentSeat - lastSeat != 1) {
            std::cout << "currentSeat " << currentSeat << " lastSeat " << lastSeat << std::endl;
        }
    }
}

int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc5.txt");

    mySeat(data);
}