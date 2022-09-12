void main() {
  Solution solution = Solution();
  List<MyName> heyList = [
    MyName(id: 1, name: "nai"),
    MyName(id: 2, name: "nai"),
    MyName(id: 3, name: "nai"),
    MyName(id: 4, name: "nai"),
  ];
  List<MyName> noList = [
    MyName(id: 2, name: "hjjh"),
    MyName(id: 3, name: "ljkjh"),
    MyName(id: 4, name: "ljkjh"),
  ];

  var list = solution.uncommonData(heyList, noList);
  print(list.length);
}

class Solution {
  List<MyName> uncommonData(List<MyName> s, List<MyName> a) {
    List<MyName> newList = [];
    s.forEach((element) {
      var contain = a.where((ele) => ele.id == element.id);
      if (contain.isEmpty) {
        newList.add(element);
      } else {
        print("not");
      }
    });
    return newList;
  }
}

class MyName {
  int id;
  String name;

  MyName({required this.id, required this.name});
}
