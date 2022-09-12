void main() {
  print(romanToInt("MMMDCLXXII"));
}

int romanToInt(String s) {
  int result = 0;
  for (int i = 0; i < s.length; i++) {
    var value = getValue(s[i]);
    if (i < s.length - 1 && value < getValue(s[i + 1])) {
      result -= value;
    } else {
      result += value;
    }
  }
  return result;
}

int getValue(String character) {
  int ans = 0;
  switch (character) {
    case 'I':
      ans = 1;
      break;

    case 'V':
      ans = 5;
      break;

    case 'X':
      ans = 10;
      break;

    case 'L':
      ans = 50;
      break;

    case 'C':
      ans = 100;
      break;

    case 'D':
      ans = 500;
      break;

    case 'M':
      ans = 1000;
      break;

    default:
      break;
  }

  return ans;
}
