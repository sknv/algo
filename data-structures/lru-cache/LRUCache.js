class LRUCache {
    /** @type {Map<number,number>} */
    cache = new Map()

    /**
     * @param {number} capacity
     */
    constructor(capacity) {
        this.capacity = capacity
    }

    /**
     * @param {number} key
     * @returns {number}
     */
    get(key) {
        const item = this.cache.get(key)
        if (!item) {
            return -1
        }

        // promote current value
        this.cache.delete(key)
        this.cache.set(key, item)
        return item
    }

    /**
     * @param {number} key
     * @param {number} value
     */
    put(key, value) {
        if (this.cache.has(key)) { // promote current value
            this.cache.delete(key)
        } else if (this.cache.size >= this.capacity) {
            this.cache.delete(this._oldestKey())
        }
        this.cache.set(key, value)
    }

    _oldestKey() {
        return this.cache.keys().next().value // take the first key
    }
}

export default LRUCache
