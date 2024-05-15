import 'dart:convert';

import 'exam1.dart';

void main() {
  var list1 = [1, 1, 2];
  var s = Solution();
  var l = ListNode(1, ListNode(1, ListNode(2, ListNode(2, ListNode(3)))));
  var t = s.deleteDuplicates(l);
  print(t);
}

class Solution {
  ListNode? deleteDuplicates(ListNode? head) {
    if (head == null) return null;
    ListNode? p = head;
    ListNode q = p;
    p = p.next;
    while (p != null) {
      if (p.val == q.val) {
        q.next = p.next;
        p = q.next;
      } else {
        q = p;
        p = p.next;
      }
    }
    return head;
  }
}
