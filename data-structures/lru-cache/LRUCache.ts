class LRUCache<K, V> {
  cache = new Map<K, V>();

  constructor(public capacity: number) {}

  /**
     * Return the value in cache by key, or undefined if one does not exist.
     */
  get(key: K): V | undefined {
    const item = this.cache.get(key);
    if (!item) {
      return undefined;
    }

    // promote current value
    this.cache.delete(key);
    this.cache.set(key, item);
    return item;
  }

  /**
     * Put the value in cache by key.
     */
  put(key: K, value: V) {
    if (this.cache.has(key)) { // promote current value
      this.cache.delete(key);
    } else if (this.cache.size >= this.capacity) {
      this.cache.delete(this.oldestKey!);
    }
    this.cache.set(key, value);
  }

  private get oldestKey(): K | undefined {
    return this.cache.keys().next().value; // take the first key
  }
}

export default LRUCache;
