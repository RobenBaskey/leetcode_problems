import 'dart:convert';

import 'exam1.dart';

void main() {
  var list1 = [1, 2, 4], list2 = [1, 3, 4];
  var s = Solution();
  print(s.mergeTwoLists(ListNode(22), ListNode(3)));
}

/**
 * Definition for singly-linked list.
 * class ListNode {
 *   int val;
 *   ListNode? next;
 *   ListNode([this.val = 0, this.next]);
 * }
 */

class Solution {
  ListNode? mergeTwoLists(ListNode? list1, ListNode? list2) {
    List<int>? listMerged = <int>[];

    while (list1 != null) {
      listMerged.add(list1.val);
      list1 = list1.next;
    }

    while (list2 != null) {
      listMerged.add(list2.val);
      list2 = list2.next;
    }

    listMerged.sort();

    ListNode? linkedListMerged = ListNode();
    ListNode head = linkedListMerged;

    for (var i in listMerged) {
      linkedListMerged?.next = ListNode(i);
      linkedListMerged = linkedListMerged?.next;
    }

    return head.next;
  }
}
