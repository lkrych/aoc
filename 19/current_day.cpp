#include <fstream>
#include <iostream>
#include <sstream>
#include <vector>

using namespace std;
//day 3 crossed wires
//What is the Manhattan distance from the central port to the closest intersection?
// ...........
// .+-----+...
// .|.....|...
// .|..+--X-+.
// .|..|..|.|.
// .|.-X--+.|.
// .|..|....|.
// .|.......|.
// .o-------+.
// ...........
void print_array (vector<string> arr) {
    int i;
    printf("arr size: %lu \n", arr.size());
    printf("[ ");
    for (i = 0; i < arr.size(); i++) {
        printf("%s ", arr[i].c_str());
    }
    printf("]\n");
}

vector<string> get_wire(int idx) {
    fstream newfile;
    vector<string> arr;

    newfile.open("aoc_day3.txt", ios::in);
    if (newfile.is_open()) {
        string line;
        //skip to correct line
        for ( int i = 1; i <= idx; i++) {
            getline(newfile, line);
        }
        //split string
        stringstream ss(line);
        string token;
        while (getline(ss, token, ',')) {
            arr.push_back(token);
        }
    }
    newfile.close();
    return arr;
}


int main() {
    vector<string> w1;
    vector<string> w2;
    w1 = get_wire(1);
    w2 = get_wire(2);
    print_array(w1);
    print_array(w2);
}