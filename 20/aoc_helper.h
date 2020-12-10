#ifndef AOC_HELPER_H 
#define AOC_HELPER_H

#include <iostream>
#include <fstream>
#include <string>
#include <vector>

std::vector<std::string> getInputString(std::string filename) {
    std::ifstream myfile(filename);
    std::vector<std::string> input;
    std::string line;
    if (myfile.is_open()) {
        while ( getline (myfile, line) ){
            input.push_back(line);
        }
        myfile.close();
    } else {
        std::cout << "There was a problem reading from " << filename;
    }
    return input;
}

std::vector<int> getInputInt(std::string filename) {
    std::vector<int> input;
    std::vector<std::string> stringInput;
    
    stringInput = getInputString(filename);
    
    for(int i = 0; i < stringInput.size(); i++) {
        int el = std::stoi(stringInput[i]);
        input.push_back(el);
    }
    return input;
}

std::vector<std::string> splitString(std::string s, std::string delim) {
    std::string token;
    size_t pos = 0;
    std::vector<std::string> split;
    while ((pos = s.find(delim)) != std::string::npos) {
        token = s.substr(0, pos);
        split.push_back(token);
        s.erase(0, pos + delim.length());
    }
    split.push_back(s);
    return split;
}

#endif 