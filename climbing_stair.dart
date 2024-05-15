void main() {
  var s = Solution();
  print(s.climbStairs(6));
}

//explaination
/** 
 * every time we can climb 1 or two step
 * so if we think from last step then it work perfectly.
 * it's one kind of dynamic programing  
 */

class Solution {
  int climbStairs(int n) {
    int x = 1, y = 1;
    for (int i = n; i > 1; i--) {
      int temp = x;
      x = x + y;
      y = temp;
    }
    return x;
  }
}
