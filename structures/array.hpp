#pragma once
#include <vector>
#include <string>
using namespace std;

class MyArray {
    vector<int> data;
public:
    void push(int val);
    bool del(int index);
    int get(int index) const;
    void set(int index, int val);
    int length() const;
    void clear();
    const vector<int>& raw() const { return data; }

    string serialize() const;
    void deserialize(const string& s);
};