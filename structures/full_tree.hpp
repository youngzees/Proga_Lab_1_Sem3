#pragma once
#include <string>
using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int v) : val(v), left(nullptr), right(nullptr) {}
};

class FullBinaryTree {
    TreeNode* root = nullptr;
public:
    FullBinaryTree() = default;
    ~FullBinaryTree();
    FullBinaryTree(const FullBinaryTree&) = delete;
    FullBinaryTree& operator=(const FullBinaryTree&) = delete;

    void insert(int val);
    bool find(int val) const;
    bool isFull() const;
    void clear();

    string serialize() const;
    void deserialize(const string& s);
    string to_string() const;

private:
    bool checkFull(TreeNode* n) const;
    void destroy(TreeNode* n);
    static string ser(TreeNode* n);
    static size_t parseNode(const string& s, size_t pos, TreeNode*& out);
};