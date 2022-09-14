import 'dart:convert';

import 'exam1.dart';

void main() {
  var list1 = [1, 1, 2];
  var s = Solution();
  print(s.removeDuplicates(list1));
}

class Solution {
  int removeDuplicates(List<int> nums) {
    int length = 0;
    for (int i = 1; i < nums.length; i++) {
      if (nums[i] != nums[length]) {
        length++;
        nums[length] = nums[i];
      }
    }
    return length + 1;
  }
}
