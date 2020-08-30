package main

import "fmt"

func main(){
	pyramid(7)
}

func pyramid(n int){
    count:=1
	for i:=1;i<=n;i++ {
	  for x:=1;x<=n-i;x++ {
	  	fmt.Print("\t")
	  }
      k:=i-1
	  for j:=1;j<=count;j++ {
	    if j<=((count/2)+1) {
	     k+=1
         fmt.Printf("%d\t",k)
	    } else {
	      k-=1
	      fmt.Printf("%d\t",k)
	    }
	  }
	  count+=2
	  fmt.Println()
	}
}