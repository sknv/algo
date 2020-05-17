import 'factorial.dart';

testFactorial() {
  var num = factorial(0), expected = 1;
  if (num != expected) {
    print('expect factorial(0) to be: $expected, got instead: $num');
  }

  num = factorial(1); expected = 1;
  if (num != expected) {
    print('expect factorial(1) to be: $expected, got instead: $num');
  }

  num = factorial(2); expected = 2;
  if (num != expected) {
    print('expect factorial(2) to be: $expected, got instead: $num');
  }

  num = factorial(6); expected = 720;
  if (num != expected) {
    print('expect factorial(6) to be: $expected, got instead: $num');
  }
}

main() {
  testFactorial();
}
