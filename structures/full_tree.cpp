#include "full_tree.hpp"
#include <queue>
#include <sstream>
#include <stdexcept>
#include <cctype>
using namespace std;

FullBinaryTree::~FullBinaryTree() { clear(); }

void FullBinaryTree::clear() {
    destroy(root);
    root = nullptr;
}

void FullBinaryTree::destroy(TreeNode* n) {
    if (!n) return;
    destroy(n->left);
    destroy(n->right);
    delete n;
}

void FullBinaryTree::insert(int val) {
    TreeNode* n = new TreeNode(val);
    if (!root) { root = n; return; }

    queue<TreeNode*> q;
    q.push(root);
    while (!q.empty()) {
        TreeNode* cur = q.front(); q.pop();
        if (!cur->left)  { cur->left  = n; return; }
        q.push(cur->left);
        if (!cur->right) { cur->right = n; return; }
        q.push(cur->right);
    }
    delete n;
}

bool FullBinaryTree::find(int val) const {
    if (!root) return false;
    queue<TreeNode*> q;
    q.push(root);
    while (!q.empty()) {
        TreeNode* cur = q.front(); q.pop();
        if (cur->val == val) return true;
        if (cur->left)  q.push(cur->left);
        if (cur->right) q.push(cur->right);
    }
    return false;
}

bool FullBinaryTree::isFull() const { return checkFull(root); }

bool FullBinaryTree::checkFull(TreeNode* n) const {
    if (!n) return true;
    if (!n->left && !n->right) return true;
    if (n->left && n->right)
        return checkFull(n->left) && checkFull(n->right);
    return false;
}

string FullBinaryTree::ser(TreeNode* n) {
    if (!n) return "-";
    if (!n->left && !n->right) return std::to_string(n->val);
    return std::to_string(n->val) + "(" + ser(n->left) + ")(" + ser(n->right) + ")";
}

string FullBinaryTree::serialize() const { return ser(root); }

size_t FullBinaryTree::parseNode(const string& s, size_t pos, TreeNode*& out) {
    if (pos >= s.size()) throw runtime_error("bad tree format");
    if (s[pos] == '-') { out = nullptr; return pos + 1; }

    int val = 0;
    bool has = false;
    while (pos < s.size() && isdigit((unsigned char)s[pos])) {
        val = val * 10 + (s[pos] - '0');
        ++pos;
        has = true;
    }
    if (!has) throw runtime_error("bad tree format: no number");

    out = new TreeNode(val);

    if (pos < s.size() && s[pos] == '(') {
        ++pos;
        pos = parseNode(s, pos, out->left);
        if (pos >= s.size() || s[pos] != ')') throw runtime_error("expected )");
        ++pos;
        if (pos >= s.size() || s[pos] != '(') throw runtime_error("expected (");
        ++pos;
        pos = parseNode(s, pos, out->right);
        if (pos >= s.size() || s[pos] != ')') throw runtime_error("expected )");
        ++pos;
    }
    return pos;
}

void FullBinaryTree::deserialize(const string& s) {
    clear();
    if (s.empty()) return;
    size_t pos = 0;
    parseNode(s, pos, root);
}

string FullBinaryTree::to_string() const { return serialize(); }