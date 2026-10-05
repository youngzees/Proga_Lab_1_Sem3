#pragma once
#include <deque>
#include <string>
using namespace std;

class MyStack {
    deque<int> data;
public:
    void push(int val);
    int pop();
    int top() const;
    bool empty() const;
    int size() const;

    string serialize() const;
    void deserialize(const string& s);
    string to_string() const;
};