import 'is_prime.dart';

testIsPrime() {
  var prime = isPrime(1), expected = false;
  if (prime != expected) {
    print('expect isPrime(1) to be: $expected, got instead: $prime');
  }

  prime = isPrime(2); expected = true;
  if (prime != expected) {
    print('expect isPrime(2) to be: $expected, got instead: $prime');
  }

  prime = isPrime(3); expected = true;
  if (prime != expected) {
    print('expect isPrime(3) to be: $expected, got instead: $prime');
  }

  prime = isPrime(4); expected = false;
  if (prime != expected) {
    print('expect isPrime(4) to be: $expected, got instead: $prime');
  }

  prime = isPrime(6); expected = false;
  if (prime != expected) {
    print('expect isPrime(6) to be: $expected, got instead: $prime');
  }

  prime = isPrime(11); expected = true;
  if (prime != expected) {
    print('expect isPrime(11) to be: $expected, got instead: $prime');
  }

  prime = isPrime(97); expected = true;
  if (prime != expected) {
    print('expect isPrime(97) to be: $expected, got instead: $prime');
  }
}

main() {
  testIsPrime();
}
