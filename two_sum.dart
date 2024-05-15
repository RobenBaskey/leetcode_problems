void main() {
  Solution s = Solution();
  var data = s.twoSum([2, 5, 5, 4], 10);
  print(data);
}

class Solution {
  List<int> twoSum(List<int> nums, int target) {
    List<int> result = [];
    if (nums.isNotEmpty) {
      var myList = nums;
      for (int i = 0; i < nums.length; i++) {
        print(i);
        for (int j = i + 1; j < myList.length; j++) {
          if (nums[i] + myList[j] == target) {
            return [i, j];
          }
        }
      }
    }
    return result;
  }
}
