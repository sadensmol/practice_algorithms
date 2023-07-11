/*
 *
 * https://leetcode.com/explore/learn/card/fun-with-arrays/525/inserting-items-into-an-array/3253/
 *
 */

import java.lang.System;
import java.util.Arrays;
class merge_sorted_arrays {

	public static void main (String[] args){

		var nums1=new int[]{1,2,3,8,0,0,0};
		var nums2=new int[]{2,5,6};

		System.out.println(Arrays.toString(nums1));
		System.out.println(Arrays.toString(nums2));
		merge(nums1, 4,nums2,3);
		System.out.println(Arrays.toString(nums1));
	}

	public static void merge(int[] nums1, int m, int[] nums2, int n) {
		for (int i=0;i<n;i++){
			boolean wasProcessed = false;
			for (int j=m+i-1;j>=0;j--){
				if (nums2[i]>=nums1[j]) {
					insertAt(j+1, nums1, nums2[i]);
					wasProcessed=true;
					break;
				}

			}
			if (!wasProcessed) insertAt(0,nums1,nums2[i]);
		}

	}

	public static void insertAt(int pos,int[] arr, int val) {
		for (int i=arr.length-2;i>=pos;i--){
			arr[i+1]=arr[i];
		}
		arr[pos]=val;

	}

}
