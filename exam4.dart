void main() {
  var strs = ["flower", "flow", "flight"];
  // var strs = ["dog", "racecar", "car"];
  //var strs = ["cir", "car"];
  Solution sl = Solution();
  var result = sl.longestCommonPrefix(strs);
  print("The result is $result");
}

class Solution {
  String longestCommonPrefix(List<String> strs) {
    String res = strs[0];
    strs.forEach((element) {
      print(res);
      print(element.indexOf(res));
      while (element.indexOf(res) != 0) {
        print(element);
        res = res.substring(0, res.length - 1);
        print(res);
      }
      print("co");
    });
    return res.isEmpty ? "" : res;
  }
}
