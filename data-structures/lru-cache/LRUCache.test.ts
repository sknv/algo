import LRUCache from "./LRUCache.ts";

function testLRUCache() {
  const cache = new LRUCache<number, number>(2);

  cache.put(1, 1);
  cache.put(2, 2);
  let value = cache.get(1);
  console.assert(value === 1, "expect cached value to be equal 1");

  cache.put(3, 3); // evicts key 2
  value = cache.get(2);
  console.assert(!value, "expect cached value to be blank");

  cache.put(4, 4); // evicts key 1
  value = cache.get(1);
  console.assert(!value, "expect cached value to be blank");

  value = cache.get(3);
  console.assert(value === 3, "expect cached value to be equal 3");

  value = cache.get(4);
  console.assert(value === 4, "expect cached value to be equal 4");
}

testLRUCache();
