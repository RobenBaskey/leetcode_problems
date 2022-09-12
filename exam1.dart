import 'dart:convert';

void main() {
  var s = Solution();
  s.twoSum([2, 4, 3], [5, 6, 4]);
}

class ListNode {
  int val;
  ListNode? next;
  ListNode([this.val = 0, this.next]);
}

class Solution {
  ListNode? addTwoNumbers(ListNode? l1, ListNode? l2) {
    List<int> res1 = [];
    res1.add(l1!.val);
    List<int> res2 = [];
    res2.add(l2!.val);

    var li = res1.reversed.toList();
    var li2 = res2.reversed.toList();

    var result =
        int.parse(li.join().toString()) + int.parse(li2.join().toString());

    var sp = result.toString();
    List re = sp.runes.map((rune) => new String.fromCharCode(rune)).toList();
    var r = re.reversed.toList();
    var myResult = strToNodeConverter(r.join().toString());
    return myResult;
  }

  ListNode? strToNodeConverter(String input) {
    var l = List.generate(
        input.length, (index) => int.parse(input.substring(index, index + 1)));
    List<ListNode?> resultChain = [];

    for (int i = 0; i < l.length; i++) {
      ListNode? stepResult = ListNode();
      stepResult.val = l[i];
      stepResult.next = i == 0 ? null : resultChain[i - 1];
      resultChain.add(stepResult);
    }
    var result = resultChain.last;
    return result;
  }

  void twoSum(List<int> list, List<int> demo) {
    var li = list.reversed.toList();
    var li2 = demo.reversed.toList();

    var result =
        int.parse(li.join().toString()) + int.parse(li2.join().toString());

    var sp = result.toString();
    List re = sp.runes.map((rune) => new String.fromCharCode(rune)).toList();
    var r = re.reversed.toList();
    print(r);

    print(r.join().toString());
  }
}
