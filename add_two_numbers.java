
import java.lang.*;

class AddTwoNumbers {

  public static class ListNode {
      int val;
      ListNode next;
      ListNode() {}
      ListNode(int val) { this.val = val; }
      ListNode(int val, ListNode next) { this.val = val; this.next = next; }



      public void trace() {
	      trace(this);
      }

      public void trace(ListNode ln) {
	      if (ln == null) return;

	      System.out.println("->"+ln.val);

	      trace(ln.next);
      }
  }
	public static void main (String[] args) {
		ListNode l1=new ListNode(2,new ListNode(4, new ListNode(3)));
		ListNode l2=new ListNode(5,new ListNode(6, new ListNode(4)));
		

		addTwoNumbers(l1, l2).trace();



	}

    public static ListNode addTwoNumbers(ListNode l1, ListNode l2) {
		  ListNode result=null,currentResult=null;                                                                    
        int carry = 0;                                                                                              
        while (l1 != null || l2!=null) {                                                                                        

                int valt = carry;      
            if (l1 != null) valt+=l1.val;
            if (l2 != null) valt+=l2.val;
            


                carry = valt/10;


                valt=valt%10;


                if (result == null){
                        result = new ListNode(valt);
                        currentResult = result;
                }else{
                
                        currentResult.next = new ListNode(valt);
                        currentResult = currentResult.next;
                }


            if (l1!=null)
                l1 = l1.next;
            
            if (l2!=null)
                l2 = l2.next;


        }

     //check for carry
     if (carry > 0) {
         currentResult.next = new ListNode(carry);
     }

            return result;
    }

}

