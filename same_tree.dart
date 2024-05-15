import 'exam1.dart';

void main() {
  var list1 = [1, 1, 2];
  var s = Solution();
  var l = TreeNode(1, TreeNode(2, TreeNode(3)));
  var m = TreeNode(1, TreeNode(2, null, TreeNode(3)));
  var t = s.isSameTree(l, m);
  print(t);
}

class TreeNode {
  int val;
  TreeNode? left;
  TreeNode? right;
  TreeNode([this.val = 0, this.left, this.right]);
}

class Solution {
  bool isSameTree(TreeNode? p, TreeNode? q) {
    if (p == null && q == null) return true;
    if (q == null || p == null) return false;
    if (p.val != q.val) return false;
    return isSameTree(p.right, q.right) && isSameTree(p.left, q.left);
  }
}
