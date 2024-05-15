void main() {
  Solution s = Solution();
  ListNode l1 = ListNode(2, ListNode(4, ListNode(3)));
  ListNode l2 = ListNode(5, ListNode(6, ListNode(4)));
  s.addTwoNumbers(l1, l2);
}

class Solution {
  ListNode? addTwoNumbers(ListNode? l1, ListNode? l2) {
    ListNode result = ListNode(0);
    int number1 = 0;

    if (l1 != null) {
      String d = "";
      while (l1.next != null) {
        print('work');
        d = d + l1.val.toString();
      }
      print(number1);
      number1 = int.parse(d);
    }
    return null;
  }
}

class ListNode {
  int val;
  ListNode? next;
  ListNode([this.val = 0, this.next]);
}
