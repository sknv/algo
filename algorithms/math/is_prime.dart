import 'dart:math';

bool isPrime(int n) {
  if (n <= 1) {
    return false;
  }

  for (var i = 2; i <= sqrt(n); i++) {
    if (n % i == 0) {
      return false;
    }
  }
  return true;
}
