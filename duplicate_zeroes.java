/**
 *
 * https://leetcode.com/explore/learn/card/fun-with-arrays/525/inserting-items-into-an-array/3245/
 *
 */

import java.lang.System;
import java.util.Arrays;

class duplicate_zeroes{
public static void main(String[] args){
	var ar = new int[]{1,0,2,3,0,4,5,0};
	

	System.out.println(Arrays.toString(ar));
	duplicateZeros(ar);
	System.out.println(Arrays.toString(ar));
}
public static void duplicateZeros(int[] arr) {
        for (int i=0;i<arr.length-1;i++){
		if (arr[i] == 0) {
			insertAt(0,i,arr);
			i++;
		}
	}
    }
private static void insertAt(int value, int pos, int[] arr){
	arr[arr.length-1]=arr[arr.length-2];
	for (int i=arr.length-2;i>=pos;i--){
		arr[i+1] = arr[i];

	}
	arr[pos]=value;
}
}

