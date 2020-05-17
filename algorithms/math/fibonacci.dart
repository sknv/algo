int fibonacci(int n) {
  var cur = 0, next = 1;
  for (var i = 0; i < n; i++) {
    final tmp = cur;
    cur = next;
    next = tmp + next;
  }
  return cur;
}
