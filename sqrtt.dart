void main() {
  var s = Solution();
  print(s.mySqrt(9));
}

class Solution {
  int mySqrt(int x) {
    int result = x;
    if (x == 0) return 0;

    while ((result * result) > x) {
      result = result ~/ 2;
      if ((result * result) < x) {
        while ((result * result) < x) {
          result = result + 1;
        }
        if ((result * result) > x) {
          result = result - 1;
        }
        break;
      } else if ((result * result) == x) {
        break;
      }
    }
    return result;
  }
}
