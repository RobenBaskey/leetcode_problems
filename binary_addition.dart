import 'dart:convert';

void main() {
  var s = Solution();
  print(s.addBinary("1010", "1011"));
}

class Solution {
  String addBinary(String a, String b) {
    if (a.length != b.length) {
      if (a.length > b.length) {
        b = List.filled(a.length - b.length, (0)).join('') + b;
      } else {
        a = List.filled(b.length - a.length, (0)).join('') + a;
      }
    }
    String sum = '';
    String rem = '0';
    for (int index = a.length - 1; index >= 0; index--) {
      if (rem == '0') {
        if (a[index] == '1' && b[index] == '1') {
          sum = '0' + sum;
          rem = '1';
        } else if (a[index] == '0' && b[index] == '0')
          sum = '0' + sum;
        else
          sum = '1' + sum;
      } else {
        if (a[index] == '1' && b[index] == '1') {
          sum = '1' + sum;
          rem = '1';
        } else if (a[index] == '0' && b[index] == '0') {
          sum = '1' + sum;
          rem = '0';
        } else {
          sum = '0' + sum;
          rem = '1';
        }
      }
    }
    return rem == '1' ? '1' + sum : sum;
  }
}
