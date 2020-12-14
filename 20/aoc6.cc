#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"

// part 2
int calculateYeses(std::map<char, int> answers, int group_count) {
    int yeses = 0;
    std::map<char, int>::iterator it;

    for ( it = answers.begin(); it != answers.end(); it++ )
    {
        if (it->second == group_count) {
            yeses++;
        }
    }

    return yeses;
}

//part1
int yesOnCustoms (std::vector<std::string> arr) {
    int count = 0;
    std::map<char, int> group_answers;
    int group_n = 0;
    for (int i = 0; i < arr.size(); i++ ){
        std::string row = arr[i];
        if (row.size() == 0) {
            // we are done parsing a group
            count += calculateYeses(group_answers, group_n);
            group_answers.clear();
            group_n = 0;
            continue;
        }
        group_n++;
        for (int j = 0; j < row.size(); j++) {
            char c = row[j];
            if ( group_answers.find(c) == group_answers.end() ) {
            // not found
                group_answers[c] = 1;
            } else {
            // found
                group_answers[c] += 1;
            }
        }
    }
    count += calculateYeses(group_answers, group_n);
    return count;
 }


int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc6.txt");

    int yes = yesOnCustoms(data);
    std::cout << "the amount of yeses " << yes << std::endl;
}