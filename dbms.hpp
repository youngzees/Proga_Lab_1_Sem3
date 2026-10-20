#pragma once
#include <string>
#include <map>
using namespace std;

#include "structures/array.hpp"
#include "structures/forward_list.hpp"
#include "structures/doubly_list.hpp"
#include "structures/stack.hpp"
#include "structures/queue.hpp"
#include "structures/full_tree.hpp"

class DBMS {
    string filename;

    map<string, MyArray>       arrays;
    map<string, ForwardList>   flists;
    map<string, DoublyList>    dlists;
    map<string, MyStack>       stacks;
    map<string, MyQueue>       queues;
    map<string, FullBinaryTree> trees;

    void load();
    void save() const;

    void printAll() const;

public:
    DBMS(const string& fname);
    ~DBMS();

    void execute(const string& query);
};