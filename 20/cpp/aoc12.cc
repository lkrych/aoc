#include <algorithm>
#include <cmath>
#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"

// compile instructions into map
int runDirections(std::vector<std::pair<char, int>> parsedData) {
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
    //keeps track of direction and magnitude
    std::pair<char, int> currentWaypointHorizontal = std::pair<char, int>('E', 10);
    std::pair<char, int> currentWaypointVertical = std::pair<char, int>('N', 1);
    int horizontal = 0;
    int vertical = 0;

    for (int i = 0; i < parsedData.size(); i++) {
        std::pair<char, int> el = parsedData[i];
        // make sure the directions align with the magnitude
        if (currentWaypointVertical.first == 'N' && currentWaypointVertical.second < 0) {
            currentWaypointVertical.second *= -1;
        }
        if (currentWaypointHorizontal.first == 'E' && currentWaypointHorizontal.second < 0) {
            currentWaypointHorizontal.second *= -1;
        }
         if (currentWaypointVertical.first == 'S' && currentWaypointVertical.second > 0) {
            currentWaypointVertical.second *= -1;
        }
        if (currentWaypointHorizontal.first == 'W' && currentWaypointHorizontal.second > 0) {
            currentWaypointHorizontal.second *= -1;
        }

        std::cout << "-----------------" << std::endl;
        std::cout << "Directions: " << el.first << " , " << el.second << std::endl;
        std::cout << "WPHorizontal: " << currentWaypointHorizontal.first << " , " << currentWaypointHorizontal.second << std::endl;
        std::cout << "WPVertical: " << currentWaypointVertical.first << " , " << currentWaypointVertical.second << std::endl;
        std::cout << "h: " << horizontal << std::endl;
        std::cout << "v: " << vertical << std::endl;
        std::cout << "-----------------" << std::endl;
       
        char hdir = currentWaypointHorizontal.first;
        char vdir = currentWaypointVertical.first;
        int hval = currentWaypointHorizontal.second;
        int vval = currentWaypointVertical.second;
        // change directions
        if (el.first == 'L') {
            //counterclockwise
            if (el.second == 90) {
                //flip magnitude
                int temp = currentWaypointVertical.second;
                currentWaypointVertical.second = currentWaypointHorizontal.second;
                currentWaypointHorizontal.second = temp;
                //flip dir
                if (hdir == 'E') {
                    if (vdir == 'N') {
                        currentWaypointHorizontal.first = 'W';
                    } else {
                        currentWaypointVertical.first = 'N';
                    }  
                } else if (hdir == 'W') {
                    if (vdir == 'N') {
                        currentWaypointVertical.first = 'S';
                    } else {
                        currentWaypointHorizontal.first = 'E';
                    }  
                }

            } else if (el.second == 180) {
                if (hdir == 'E') {
                    currentWaypointHorizontal.first = 'W';
                } else if (hdir == 'W') {
                    currentWaypointHorizontal.first = 'E';
                }
                if (vdir == 'N') {
                    currentWaypointVertical.first = 'S';
                } else if (vdir == 'S') {
                    currentWaypointVertical.first = 'N';
                }

            } else if (el.second == 270) {
                //flip magnitude
                int temp = currentWaypointVertical.second;
                currentWaypointVertical.second = currentWaypointHorizontal.second;
                currentWaypointHorizontal.second = temp;
                //flip dir
                if (hdir == 'E') {
                    if (vdir == 'N') {
                        currentWaypointVertical.first = 'S';
                    } else {
                        currentWaypointHorizontal.first = 'W';
                    }  
                } else if (hdir == 'W') {
                    if (vdir == 'N') {
                        currentWaypointHorizontal.first = 'E';
                    } else {
                        currentWaypointVertical.first = 'N';
                    }  
                }
            }
            
            continue;
        } else if (el.first == 'R') {
            //clockwise
            if (el.second == 90) {
                //flip magnitude
                int temp = currentWaypointVertical.second;
                currentWaypointVertical.second = currentWaypointHorizontal.second;
                currentWaypointHorizontal.second = temp;
                //flip dir
                if (hdir == 'E') {
                    if (vdir == 'N') {
                        currentWaypointVertical.first = 'S';
                    } else {
                        currentWaypointHorizontal.first = 'W';
                    }  
                } else if (hdir == 'W') {
                    if (vdir == 'N') {
                        currentWaypointHorizontal.first = 'E';
                    } else {
                        currentWaypointVertical.first = 'N';
                    }  
                }
                
            } else if (el.second == 180) {
                std::cout << "R 180 with " << hdir << " " << vdir << std::endl;
                if (hdir == 'E') {
                    currentWaypointHorizontal.first = 'W';
                } else if (hdir == 'W') {
                    currentWaypointHorizontal.first = 'E';
                }
                if (vdir == 'N') {
                    currentWaypointVertical.first = 'S';
                } else if (vdir == 'S') {
                    currentWaypointVertical.first = 'N';
                }
            } else if (el.second == 270) {
                //flip magnitude
                int temp = currentWaypointVertical.second;
                currentWaypointVertical.second = currentWaypointHorizontal.second;
                currentWaypointHorizontal.second = temp;
                //flip dir
                if (hdir == 'E') {
                    if (vdir == 'N') {
                        currentWaypointHorizontal.first = 'W';
                    } else {
                        currentWaypointVertical.first = 'N';
                    }  
                } else if (hdir == 'W') {
                    if (vdir == 'N') {
                        currentWaypointVertical.first = 'S';
                    } else {
                        currentWaypointHorizontal.first = 'E';
                    }  
                }
                
            }
            
            continue;
        } else if (el.first == 'F') {
            // this is the only place where the ship actually moves
            horizontal += (el.second * currentWaypointHorizontal.second);
            vertical += (el.second * currentWaypointVertical.second);
            continue;
        }
        // if directions equal N,S,E, or W
        if (el.first == 'N') {
            currentWaypointVertical.second += el.second;
            if (currentWaypointVertical.second > 0) {
                currentWaypointVertical.first = 'N';
            }
        } else if (el.first == 'S') {
            currentWaypointVertical.second -= el.second;
            if (currentWaypointVertical.second < 0) {
                currentWaypointVertical.first = 'S';
            }
        } else if (el.first == 'E') {
            currentWaypointHorizontal.second += el.second;
            if (currentWaypointHorizontal.second > 0) {
                currentWaypointHorizontal.first = 'E';
            }
        } else if (el.first == 'W') {
            currentWaypointHorizontal.second -= el.second;
            if (currentWaypointHorizontal.second < 0) {
                currentWaypointHorizontal.first = 'W';
            }
        }

    }
    return abs(vertical) + abs(horizontal);
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
    int manhattan = runDirections(parsed);
    return manhattan;

}

int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc12.txt");

    int manhattan_distance = manhatDist(data);
    std::cout << "the manhattan_distance is " << manhattan_distance << std::endl;
}