void main() {
  var value = palindrom(x: 0);
  print(value);
}

bool palindrom({required int x}) {
  if (x < 0) return false;
  int reversed = 0, remainder, original = x;
  while (x != 0) {
    remainder = x % 10;
    reversed = reversed * 10 + remainder;
    x = x ~/ 10;
  }
  return original == reversed;
}
