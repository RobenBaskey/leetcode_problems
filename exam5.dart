void main() {
  Solution solution = Solution();
  print(solution.isValid("([}}])"));
}

class Solution {
  bool isValid(String s) {
    List<String> list = [];
    if (s[0] == ")" || s[0] == "}" || s[0] == "]") {
      return false;
    } else if (s.length % 2 != 0) {
      return false;
    } else {
      for (int i = 0; i < s.length; i++) {
        if (s[i] == "(" || s[i] == "{" || s[i] == "[") {
          list.add(s[i]);
        } else if (s[i] == ")" &&
            list.isNotEmpty &&
            list[list.length - 1] == "(") {
          list.removeAt(list.length - 1);
        } else if (s[i] == "}" &&
            list.isNotEmpty &&
            list[list.length - 1] == "{") {
          list.removeAt(list.length - 1);
        } else if (s[i] == "]" &&
            list.isNotEmpty &&
            list[list.length - 1] == "[") {
          list.removeAt(list.length - 1);
        } else {
          return false;
        }
      }
      return list.isEmpty;
    }
  }
}

///it's successfully worked on leetcode.
///thanks god
