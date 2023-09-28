import 'dart:io';
import 'dart:math';

int kamenetsky(double n) {
  if (n < 0) {
    return 0;
  }

  if (n <= 1) {
    return 1;
  }

  double x = ((n * (log(n / e) / ln10) ) + ( (log(2 * pi * n) / ln10) / 2.0));

  return (x.floor() + 1);
}

void main() {
  double n = double.parse(stdin.readLineSync()!);
  print(kamenetsky(n));
}

