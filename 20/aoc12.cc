#include <algorithm>
#include <cmath>
#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"

// compile instructions into map
std::map<char, int> runDirections(std::vector<std::pair<char, int>> parsedData) {
    std::map<char, int> directionsMap;
    std::map<char, int> dirToNum = {
        {'N', 0},
        {'E', 1},
        {'S', 2},
        {'W', 3},
    };
    std::map<int, char> numToDir = {
        {0, 'N'},
        {1, 'E'},
        {2, 'S'},
        {3, 'W'},
    };
    //initial starting state from prompt
    char currentDirection = 'E';
    for (int i = 0; i < parsedData.size(); i++) {
        std::pair<char, int> el = parsedData[i];
        
        // change directions
        if (el.first == 'L') {
            //counterclockwise
            int currentIntDirection = dirToNum[currentDirection];
            if (el.second == 90) {
                currentIntDirection -= 1;
            } else if (el.second == 180) {
                currentIntDirection -= 2;
            } else if (el.second == 270) {
                currentIntDirection -= 3;
            }
            currentIntDirection = mod(currentIntDirection, 4);
            currentDirection = numToDir[currentIntDirection];
            continue;
        } else if (el.first == 'R') {
            //clockwise
            int currentIntDirection = dirToNum[currentDirection];
            if (el.second == 90) {
                currentIntDirection += 1;
            } else if (el.second == 180) {
                currentIntDirection += 2;
            } else if (el.second == 270) {
                currentIntDirection += 3;
            }
            currentIntDirection = mod(currentIntDirection, 4);
            currentDirection = numToDir[currentIntDirection];
            continue;
        } else if (el.first == 'F') {
            // std::cout << currentDirection << std::endl;
            directionsMap[currentDirection] += el.second;
            continue;
        }

        // if directions equal N,S,E, or W
        if (directionsMap.find(el.first) != directionsMap.end()) {
            directionsMap[el.first] += el.second;
        } else {
            directionsMap[el.first] = el.second;
        }
    }
    return directionsMap;
}

// organize raw instructions into pairs
std::vector<std::pair<char, int>> parseDirections(std::vector<std::string> rawData) {
    std::vector<std::pair<char, int>> parsed = {};
    for(int i = 0; i < rawData.size(); i++) {
        std::string datum = rawData[i];
        char instruction = datum[0];
        int amount = std::stoi(datum.substr(1, datum.size() - 1));
        // std::cout << instruction << ", " << amount << std::endl;
        parsed.push_back(std::pair<char, int>(instruction, amount));
    }
    return parsed;
}

int manhatDist(std::vector<std::string> data) {
    //decode input
    std::vector<std::pair<char, int>> parsed = parseDirections(data);
    //run directions
    std::map<char, int> directionMap = runDirections(parsed);
    std::map<char, int>::iterator it;
    int vertical = 0;
    int horizontal = 0;
    for (it = directionMap.begin(); it != directionMap.end(); it++)
    {
        std::cout << it->first << ", " << it->second << std::endl;
        if (it->first == 'N') {
            vertical += it->second;
        } else if (it->first == 'S') {
            vertical -= it->second;
        } else if (it->first == 'E') {
            horizontal += it->second;
        } else if (it->first == 'W') {
            horizontal -= it->second;
        }
    }
    return abs(vertical) + abs(horizontal);

    //return
}

int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc12.txt");

    int manhattan_distance = manhatDist(data);
    std::cout << "the manhattan_distance is " << manhattan_distance << std::endl;
}