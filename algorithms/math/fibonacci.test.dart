import 'fibonacci.dart';

testFibonacci() {
  var num = fibonacci(1), expected = 1;
  if (num != expected) {
    print('expect fibonacci(1) to be: $expected, got instead: $num');
  }

  num = fibonacci(2); expected = 1;
  if (num != expected) {
    print('expect fibonacci(2) to be: $expected, got instead: $num');
  }

  num = fibonacci(6); expected = 8;
  if (num != expected) {
    print('expect fibonacci(6) to be: $expected, got instead: $num');
  }
}

main() {
  testFibonacci();
}
