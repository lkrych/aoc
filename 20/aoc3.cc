#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include "aoc_helper.h"

int treesEncountered(std::vector<std::string>data, std::vector<int> path) {
    int horizontal = path[0];
    int down = path[1];
    int tree_count = 0;
    std::vector<int> coordinates = {0,0};
    for(int i = 0; i < data.size(); i++) {
        int current_pos = coordinates[0];
        std::string row = data[i];
        if (row[current_pos] == '#') {
            tree_count++;
        }
        //update coordinates
        coordinates[0] = (current_pos + horizontal) % row.size();
        coordinates[1] += 1;
    }
    return tree_count;
}


int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc3.txt");

    // traverse the map and count the number of trees
    // this path vector means right 3, down 1
    std::vector<int> path = {3, 1};
    int trees = treesEncountered(data, path);
    std::cout << "Encountered " << trees << " trees along the toboggan path" << std::endl;
}