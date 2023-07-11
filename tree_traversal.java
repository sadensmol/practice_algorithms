/*
 *
 *
 * https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/928/
 */


import java.lang.System;
import java.util.List;
import java.util.ArrayList;

public class tree_traversal{
private static class TreeNode {
     int val;
     TreeNode left;
     TreeNode right;
     TreeNode() {}
     TreeNode(int val) { this.val = val; }
     TreeNode(int val, TreeNode left, TreeNode right) {
         this.val = val;
         this.left = left;
         this.right = right;
     }
 }

	public static void main (String[] params){

		/**
		 *
		 *                 1
		 *            3         4
		 *          5   23
		 *
		 */
		TreeNode root = new TreeNode(1,new TreeNode(3,new TreeNode(5), new TreeNode(23)), new TreeNode(4));

		System.out.println(preorderTraversal(root));// 1,3,5,23,4
		System.out.println(inOrderTraversal(root));// 5,3,23,1,4
		System.out.println(postOrderTraversal(root));// 5,23,3,4,1


	}


	public static List<Integer> postOrderTraversal(TreeNode root) {
        
		List<Integer> result = new ArrayList();

		if (root == null) return result;

		result.addAll(postOrderTraversal(root.left));
		result.addAll(postOrderTraversal(root.right));
		result.add(root.val);

		return result;

	}
	public static List<Integer> inOrderTraversal(TreeNode root) {
        
		List<Integer> result = new ArrayList();

		if (root == null) return result;

		result.addAll(inOrderTraversal(root.left));
		result.add(root.val);
		result.addAll(inOrderTraversal(root.right));

		return result;

	}

	public static List<Integer> preorderTraversal(TreeNode root) {
        
		List<Integer> result = new ArrayList();

		if (root == null) return result;
		result.add(root.val);
		result.addAll(preorderTraversal(root.left));
		result.addAll(preorderTraversal(root.right));

		return result;
    }
}


