#include <algorithm>
#include <cmath>
#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"

// compile instructions into map
std::map<std::string, int> runDirections(std::vector<std::pair<std::string, int>> parsedData) {
    std::map<std::string, int> directionsMap;
    for (int i = 0; i < parsedData.size(); i++) {
        std::pair<std::string, int> el = parsedData[i];
        if (directionsMap.find(el.first) != directionsMap.end()) {
            directionsMap[el.first] += el.second;
        } else {
            directionsMap[el.first] = el.second;
        }
    }
    return directionsMap
}

// organize raw instructions into pairs
std::vector<std::pair<std::string, int>> parseDirections(std::vector<std::string> rawData) {
    std::vector<std::pair<std::string, int>> parsed = {};
    for(int i = 0; i < rawData.size(); i++) {
        std::string datum = rawData[i];
        std::string instruction = rawData[0];
        int amount = std::stoi(datum.substr(1, datum.size() - 1));
        parsed.push_back(std::pair<std::string, int>(instruction, amount));
    }
    return parsed;
}

int manhatDist(std::vector<std::string> data) {
    //decode input
    std::vector<std::pair<std::string, int>> parsed = parseDirections(data);
    //run directions
    std::map<std::string, int> directionMap = runDirections(parsed);
    std::map<std::string, int>::iterator it;
    int vertical = 0;
    int horizontal = 0;
    for (it = directionMap.begin(); it != directionMap.end(); it++)
    {
        if (it->first == "N") {
            vertical += it->second;
        } else if (it->first == "S") {
            vertical -= it->second;
        } else if (it->first == "E") {
            horizontal += it->second;
        } else if (it->first == "W") {
            horizontal -= it->second;
        }
    }
    return abs(vertical) + abs(horizontal);

    //return
}

int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc12test.txt");

    int manhattan_distance = manhatDist(data);
    std::cout << "the manhattan_distance is " << manhattan_distance << std::endl;
}