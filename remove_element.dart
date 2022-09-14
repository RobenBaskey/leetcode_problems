import 'dart:convert';

import 'exam1.dart';

void main() {
  var list1 = [0, 1, 2, 2, 3, 0, 4, 2];
  var s = Solution();
  print(s.removeDuplicates(list1, 2));
}

class Solution {
  int removeDuplicates(List<int> nums, int val) {
    nums.removeWhere((element) => element == val);
    return nums.length;
  }
}
