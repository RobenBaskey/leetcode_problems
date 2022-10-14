import 'dart:convert';

void main() {
  var str = " Hello World ";
  var s = Solution();
  print(s.lengthOfLastWord(str));
}

class Solution {
  int lengthOfLastWord(String s) {
    int i = s.length - 1;
    if (s.length == 1) {
      return 1;
    }
    int count = 0;
    while (s[i] == ' ') {
      i--;
    }
    while (s[i] != ' ') {
      count++;
      if (i == 0) {
        break;
      }
      i--;
    }
    return count;
  }
}
