import 'exam1.dart';

void main() {
  var list1 = [1, 1, 2];
  var s = Solution();
  var l = TreeNode(1, TreeNode(1, TreeNode(2, TreeNode(2, TreeNode(3)))));
  var t = s.inorderTraversal(l);
  print(t);
}

class TreeNode {
  int val;
  TreeNode? left;
  TreeNode? right;
  TreeNode([this.val = 0, this.left, this.right]);
}

class Solution {
  List<int> inorderTraversal(TreeNode? root) {
    List<int> output = [];
    helper(root, output);
    return output;
  }

  void helper(TreeNode? data, List<int> out) {
    if (data != null) {
      helper(data.left, out);
      out.add(data.val);
      helper(data.right, out);
    }
  }
}
