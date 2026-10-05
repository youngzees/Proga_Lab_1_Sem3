#pragma once
#include <deque>
#include <string>
using namespace std;

class MyQueue {
    deque<int> data;
public:
    void push(int val);
    int pop();
    int front() const;
    bool empty() const;
    int size() const;

    string serialize() const;
    void deserialize(const string& s);
    string to_string() const;
};