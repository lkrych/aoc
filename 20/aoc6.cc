#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"

//part1
int yesOnCustoms (std::vector<std::string> arr) {
    int count = 0;
    std::map<char, bool> group_answers;
    for (int i = 0; i < arr.size(); i++ ){
        std::string row = arr[i];
        if (row.size() == 0) {
            // we are done parsing a group
            count += group_answers.size();
            group_answers.clear();
            continue;
        }
        for (int j = 0; j < row.size(); j++) {
            char c = row[j];
            group_answers[c] = true;
        }
    }
    count += group_answers.size();
    return count;
 }


int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc6.txt");

    int yes = yesOnCustoms(data);
    std::cout << "the amount of yeses " << yes << std::endl;
}