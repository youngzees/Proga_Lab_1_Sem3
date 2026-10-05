#pragma once
#include <string>
using namespace std;

struct NodeF {
    int val;
    NodeF* next;
    NodeF(int v) : val(v), next(nullptr) {}
};

class ForwardList {
    NodeF* head = nullptr;
public:
    ForwardList() = default;
    ~ForwardList();
    ForwardList(const ForwardList&) = delete;
    ForwardList& operator=(const ForwardList&) = delete;

    // 4 способа добавления
    void push_head(int val);
    void push_tail(int val);
    bool push_before(int target, int val);
    bool push_after (int target, int val);

    // 4 способа удаления
    bool pop_head();
    bool pop_tail();
    bool del(int val);
    bool del_index(int idx);

    bool find(int val) const;
    void clear();

    string serialize() const;
    void deserialize(const string& s);
    string to_string() const;
};