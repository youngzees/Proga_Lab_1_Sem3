#pragma once
#include <string>
using namespace std;

struct NodeL {
    int val;
    NodeL* prev;
    NodeL* next;
    NodeL(int v) : val(v), prev(nullptr), next(nullptr) {}
};

class DoublyList {
    NodeL* head = nullptr;
    NodeL* tail = nullptr;
public:
    DoublyList() = default;
    ~DoublyList();
    DoublyList(const DoublyList&) = delete;
    DoublyList& operator=(const DoublyList&) = delete;

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