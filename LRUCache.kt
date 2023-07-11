import java.util.LinkedList

class LRUCache(private val capacity: Int) {
    private val keyVal = mutableMapOf<Int, Int>()
    private var lruKeyLinkedList = LinkedList<Int>()

    fun processKeyEvict(key: Int) {
        lruKeyLinkedList.remove(key)
        lruKeyLinkedList.add(key)
    }
    fun get(key: Int): Int {
        //make lru key as last accessed key
        val result = keyVal.getOrDefault(key, -1)

        if (result != -1) processKeyEvict(key)
        return result
    }

    fun put(key: Int, value: Int) {
        if (keyVal.size == capacity) {
            if (!keyVal.containsKey(key)) {
                keyVal.remove(lruKeyLinkedList.removeFirst())
            }
        }
        keyVal[key] = value
        processKeyEvict(key)
    }

    fun trace() {
        println("$keyVal -> $lruKeyLinkedList ")
    }
}


fun main() {
    test1(LRUCache(2))
    test2(LRUCache(2))

}

fun test1(obj: LRUCache) {
    println("--- test1 ---")
    obj.put(1, 1)
    obj.put(2, 2)
    obj.trace()
    println(obj.get(1)) // 1
    obj.trace()
    obj.put(3, 3)
    obj.trace()
    println(obj.get(2)) // -1
    obj.put(4, 4)
    println(obj.get(1)) // -1
    println(obj.get(3)) // 3
    println(obj.get(4)) //4

}

fun test2(obj: LRUCache) {
    println("--- test2 ---")
    obj.get(2)
    obj.trace()
    obj.put(2, 6)
    obj.trace()
    obj.get(1)
    obj.trace()
    obj.put(1, 5)
    obj.trace()
    obj.put(1, 2)
    obj.trace()
    obj.get(1)
    obj.trace()
    obj.get(2)
    obj.trace()


}