import 'dart:io';


void main() {
  List<String> input = stdin.readLineSync()!.split(' ');
  int a = int.parse(input[0]);
  int b = int.parse(input[1]);

  int q, r, limit;
  q = 0;
  limit = b < 0 ? -b : b;
  r = 0;
  for (; r < limit; r++) {
    if ((a - r) % b == 0) {
      q = (a - r) ~/ b;
      break;
    }
  }

  print('$q $r');
}
