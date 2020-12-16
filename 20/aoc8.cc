#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"

int findLoop(std::vector<std::string> data) {
    int acc = 0;
    int curr_ins_idx = 0;
    int back_track_idx = 0;
    int back_track_acc = 0;
    std::map<int,bool> past_instructions;
    std::map<int,bool> backtrack_past_instructions;
    std::map<int, bool> backtracked;
    bool backtrack_in_progress = false;
    while (true) {
        // program has reached the end
        std::cout << "checking ins: " << curr_ins_idx << std::endl;
        if (curr_ins_idx == data.size()) {
            std::cout << "the end has been reached! praise be!" << std::endl;
            return acc;
        }
        // check if instruction has been seen before
        if ( past_instructions.find(curr_ins_idx) == past_instructions.end() ) {
            past_instructions[curr_ins_idx] = true;
        } else {
            std::cout << "discovered a loop, returning back to the old snapshot" << std::endl;
            backtracked[back_track_idx] = true;
            backtrack_in_progress = false;
            //reset instruction values
            curr_ins_idx = back_track_idx;
            acc = back_track_acc;
            past_instructions = backtrack_past_instructions;
        }
        // perform the instruction
        std::string curr_ins = data[curr_ins_idx];
        std::cout << "ins: " << curr_ins << std::endl;
        std::vector<std::string> split_ins = splitString(curr_ins, " ");
        if (split_ins.size() != 2) {
            std::cout << "There was an error in the test input" << std::endl;
            exit(1);
        }
        std::string action = split_ins[0];
        std::string val = split_ins[1];
        char op = val.at(0);
        int amount = std::stoi(split_ins[1].substr(1, split_ins[1].size() - 1));
        if (action == "acc") {
            if (op == '+') {
                acc += amount;
            } else if (op == '-') {
                acc -= amount;
            }
            curr_ins_idx++;
        } else if (action == "jmp") {
            if ( backtracked.find(curr_ins_idx) == backtracked.end() && !backtrack_in_progress ) {
                back_track_idx = curr_ins_idx;
                back_track_acc = acc;
                backtrack_past_instructions = past_instructions;
                backtrack_in_progress = true;
                //nop action
                curr_ins_idx++;
                continue;
            } 

            if (op == '+') {
                curr_ins_idx += amount;
            } else if (op == '-') {
                curr_ins_idx -= amount;
            }
        } else if (action == "nop") {
            if ( backtracked.find(curr_ins_idx) == backtracked.end() && !backtrack_in_progress) {
                back_track_idx = curr_ins_idx;
                back_track_acc = acc;
                backtrack_past_instructions = past_instructions;
                backtrack_in_progress = true;
                //jmp action
                if (op == '+') {
                    curr_ins_idx += amount;
                } else if (op == '-') {
                    curr_ins_idx -= amount;
                }
                continue;
            } 
            curr_ins_idx++;
        }
    }
}


int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc8.txt");

    int loopAcc = findLoop(data);
    std::cout << "the acc value before the loop is " << loopAcc << std::endl;
}