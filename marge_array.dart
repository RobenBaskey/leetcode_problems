import 'dart:convert';

import 'exam1.dart';

void main() {
  var nums1 = [1, 2, 3, 0, 0, 0], m = 3, nums2 = [2, 5, 6], n = 3;
  var s = Solution();
  s.merge(nums1, m, nums2, n);
}

class Solution {
  void merge(List<int> nums1, int m, List<int> nums2, int n) {
    nums1.removeRange(m, nums1.length);
    nums2.removeRange(n, nums2.length);
    nums1.addAll(nums2);
    nums1.sort();
  }
}
