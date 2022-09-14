import 'dart:convert';

import 'exam1.dart';

void main() {
  var list1 = [1, 3, 5, 6];
  var s = Solution();
  print(s.removeDuplicates(list1, 7));
}

class Solution {
  int removeDuplicates(List<int> nums, int target) {
    int index = 0;
    if (nums.contains(target)) {
      for (int i = 0; i < nums.length; i++) {
        if (nums[i] == target) {
          index = i;
        }
      }
    } else if (nums[nums.length - 1] > target) {
      for (int i = 0; i < nums.length; i++) {
        if (nums[i] < target && nums[i + 1] > target) {
          index = i + 1;
        }
      }
    } else {
      index = nums.length;
    }

    return index;
  }
}
