package main

import "fmt"

func Maximumlength( x string ) int {

	max :=0
	aa:=0
	bb:=0
	cc:=0
	for i:=0;i<len(x);i++{

		if x[i]=='a' {
			if bb==0&&cc==0{
				aa+=1
			}else{
				bb=0
				cc = 0
			}

		}
		if x[i]=='b'{
			if bb<aa&&cc==0{
				bb+=1
			}else{
				bb=aa
				if cc!=0{
					bb =
				}
			}
		}
		if x[i]=='c' {
			 if cc<bb{
				cc+=1
			}
		}
		if cc==bb&&bb==aa{
			if 3*cc>max{
				max =3*cc
			}
			cc =0
			bb=0
			aa =0
		}
	}

	return  max
}
func main(){
   fmt.Println(Maximumlength("aaaabbcccc"))
}