/*
 *
 * https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3247/
 *
 */

import java.lang.System;
import java.util.Arrays;

class remove_element{

	public static void main(String[] args){
		var arr = new int[]{2,3,5,5,6};

		var i =	removeElement(arr, 5);// 2,3,6,6,6 3

		System.out.println(Arrays.toString(arr));
		System.out.println(i);
	}

	public static int removeElement(int[] nums, int val) {
       		int size = 0;
		for (int i=0;i<nums.length-size;i++) {
			if (nums[i] == val) {
				removeElementAt(i,nums);
				i--;
				size++;
			}
		}
		return nums.length -size;
    	}

	public static void removeElementAt(int pos, int[] arr) {
		for (int i=pos;i<arr.length-1;i++) {

			arr[i] =arr[i+1];
		}
	}

}
