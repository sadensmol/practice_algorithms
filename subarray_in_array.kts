//XIX Задача о поиске элемента starstar
//
//        Дан упорядоченный массив чисел размером N Нужно реализовать алгоритм поиска вхождения упорядоченного подмассива размера M, где M << N
//
//func isInclude(array int[], subarray []int) bool
//
//assert(isInclude([1, 2, 3, 5, 7, 9, 11], []) == true)
//assert(isInclude([1, 2, 3, 5, 7, 9, 11], [3, 5, 7]) == true)
//assert(isInclude([1, 2, 3, 5, 7, 9, 11], [4, 5, 7]) == false)
//Что хочется увидеть:
//
//Алгоритм со сложность быстрее чем O(N) по времени

fun hasIncluded(array:IntArray, subarray:IntArray):Boolean {

    if (subarray.size ==0) return true

    val position = array.binarySearch(subarray[0]) //O(logN)

    if (position <0) return false

    val toPosition = position+subarray.size
    if (toPosition > array.size) return false

    return areEquals(array.copyOfRange(position,toPosition),subarray)
}

fun areEquals(array1:IntArray, array2:IntArray):Boolean {
    if (array1.size != array2.size) return false
    for (a1 in array1.withIndex())

            if (a1.value!=array2[a1.index]) return false


    return true
}

println(hasIncluded(intArrayOf(1, 2, 3, 5, 7, 9, 11), intArrayOf()) ) //true
println(hasIncluded(intArrayOf(1, 2, 3, 5, 7, 9, 11), intArrayOf(3, 5, 7))) //true
 println(hasIncluded(intArrayOf(1, 2, 3, 5, 7, 9, 11), intArrayOf(7,9))) //true
println(hasIncluded(intArrayOf(1, 2, 3, 5, 7, 9, 11), intArrayOf(9,11))) //true
println(hasIncluded(intArrayOf(1, 2, 3, 5, 7, 9, 11), intArrayOf(1,2))) //true
println(hasIncluded(intArrayOf(1, 2, 3, 5, 7, 9, 11), intArrayOf(1))) //true
println(hasIncluded(intArrayOf(1, 2, 3, 3, 3, 5, 7, 9, 11), intArrayOf(3, 3, 5, 7))) //true !!! this is the problem!
println(hasIncluded(intArrayOf(1, 2, 3, 5, 7, 9, 11), intArrayOf(4, 5, 7))) //false

/*
feedback: hasIncluded(intArrayOf(1, 2, 3, 3, 3, 5, 7, 9, 11), intArrayOf(3, 3, 5, 7)) давет false, должно быть true.
 */
