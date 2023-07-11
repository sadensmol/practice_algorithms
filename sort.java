import java.lang.System;
import java.util.Arrays;


class Solution{

	public static void main (String[] params) {

		var arr = new int[]{1,3,6,7,21,2,23,4,35};

		System.out.println(Arrays.toString(arr));
		var arr1 = cloneArray(arr);
		bubbleSort(arr1);

		System.out.println(Arrays.toString(arr1));
		var arr2 = cloneArray(arr);
		insertionSort(arr2);
		System.out.println(Arrays.toString(arr2));
		var arr3 = cloneArray(arr);
		selectionSort(arr3);
		System.out.println(Arrays.toString(arr3));

		
	}

	private static void selectionSort(int[] arr) {
		for (var i=0;i<arr.length;i++) {
			int minPos=i;
			int minVal=arr[i];
			for (var j=i+1;j<arr.length;j++){
				if (arr[j]<minVal) {
					minPos=j;
					minVal = arr[j];
				}
			}

			moveLeftTo(minPos,i,arr);
			
		}

	}

	private static void insertionSort(int[] arr) {
		for (var i=0;i<arr.length;i++){
			
			boolean moved = false;
			for (var j=i-1;j>=0;j--) {


				if (arr[i]>arr[j]) {
					
					moveLeftTo(i,j+1,arr);
					moved = true;
					break;
				}
			}

			if (!moved) moveLeftTo(i,0, arr);

		}
	}

	private static void moveLeftTo(int spos, int tpos, int[] arr) {
		if (spos == tpos) return;

		int val = arr[spos];
		for (var i = spos;i>tpos;i--){
			arr[i]=arr[i-1];
		}
		arr[tpos]=val;

	}
	private static int[] cloneArray(int[] arr){
		int[] res = new int[arr.length];

		for (int i=0;i<arr.length;i++) res[i]=arr[i];

		return res;
	}

	private static void bubbleSort(int[] arr){
		boolean running;
		do {
			running= false;
		for (int i=0;i<arr.length-2;i++){
			if (arr[i+1]<arr[i]) {
				var temp = arr[i];
				arr[i] = arr[i+1];
				arr[i+1]= temp;
				running = true;
			}
		}


		}while(running);
	}
}
