void main() {
  ///logic -> a=z, b=y , c=x
  var data = getTheWord("give");
  print(data);
}

String getTheWord(String text) {
  text = text.toLowerCase();
  String alphabet = "abcdefghijklmnopqrstuvwxyz";
  String reverseAlphabet = alphabet.split('').reversed.join('');
  String result = "";
  for (int i = 0; i < text.length; i++) {
    var index = alphabet.indexOf(text[i]);
    print(index);
    result = result + "${reverseAlphabet[index]}";
  }
  return result;
}
