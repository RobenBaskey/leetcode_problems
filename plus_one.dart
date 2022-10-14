import 'dart:convert';

void main() {
  var str = [
    7,
    2,
    8,
    5,
    0,
    9,
    1,
    2,
    9,
    5,
    3,
    6,
    6,
    7,
    3,
    2,
    8,
    4,
    3,
    7,
    9,
    5,
    7,
    7,
    4,
    7,
    4,
    9,
    4,
    7,
    0,
    1,
    1,
    1,
    7,
    4,
    0,
    0,
    6
  ];
  var t = [1, 9, 9];
  var s = Solution();
  print(s.plusOne(t));
}

class Solution {
  List<int> plusOne(List<int> digits) {
    int size = digits.length - 1;

    if (digits[size] < 9) {
      digits[size]++;
      return digits;
    } else {
      for (int i = size; i >= 0; i--) {
        print(digits[i]);
        if (digits[i] == 9) {
          digits[i] = 0;
          digits[i - 1] = digits[i - 1]++;
        }
      }
    }
    return digits;
  }
}
