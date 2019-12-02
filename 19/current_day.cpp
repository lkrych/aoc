#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

void print_array (vector<int> arr) {
    int i;
    printf("[ ");
    for (i = 0; i < arr.size(); i++) {
        printf("%d ", arr[i]);
    }
    printf("]\n");
}

// need to process each line by the following rules
// The opcode indicates what to do; for example, 99 means that the program is finished and should immediately halt. Encountering an unknown opcode means something went wrong.

// Opcode 1 adds together numbers read from two positions and stores the result in a third position. 
// The three integers immediately after the opcode tell you these three positions - 
// the first two indicate the positions from which you should read the input values, and the third indicates the position at which the output should be stored.

// For example, if your Intcode computer encounters 1,10,20,30, it should read the values at positions 10 and 20, add those values, and then overwrite the value at position 30 with their sum.

// Opcode 2 works exactly like opcode 1, except it multiplies the two inputs instead of adding them. Again, the three integers after the opcode indicate where the inputs and outputs are, not their values.

// Once you're done processing an opcode, move to the next one by stepping forward 4 positions.

void process_p1(vector<int> arr) {
    int current_idx = 0;
    int x, y;
    while (arr[current_idx] != 99) {
        int operation = arr[current_idx];
        printf("==============================\n");
        print_array(arr);
        printf("current operation is %d\n", operation);
        if (operation == 1) {
            x = arr[arr[current_idx + 1]];
            y = arr[arr[current_idx + 2]];
            arr[arr[current_idx + 3]] = x + y;
            printf("case 1 opcode: %d, x: %d, y: %d dest: %d, \n", operation, x, y, current_idx + 3);
            current_idx += 4;
        } else if (operation == 2) {
            x = arr[arr[current_idx + 1]];
            y = arr[arr[current_idx + 2]];
            arr[arr[current_idx + 3]] = x * y;
            printf("case 2 opcode: %d, x: %d, y: %d dest: %d, \n", operation, x, y, current_idx + 3);
            current_idx += 4;
        } else {
            break;
        }
        printf("==============================\n");
    }
    printf("The value left at position zero is: %d \n", arr[0]);
}

int main() {
    fstream newfile;
    newfile.open("aoc_day2.txt", ios::in);
    if (newfile.is_open()) {
        string line;
        vector<int> arr;
        //read input into array
        while (getline(newfile, line, ',')) {
            //printf("%s\n", line.c_str());
            arr.push_back(stoi(line));
        }
        newfile.close();
        process_p1(arr);
    }
}