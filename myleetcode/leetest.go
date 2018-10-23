package myleetcode

import (
	"fmt"
	"os"
	"strconv"
)

func TwoSum(nums []int, target int) {
	var numap map[int]int
	numap = make(map[int]int)
	for k, v := range nums {
		if v < target {
			a := target - v
			_, flag := numap[v]
			if flag {
				fmt.Println("[" + strconv.Itoa(numap[v]) + "," + strconv.Itoa(k) + "]")

			} else {
				numap[a] = k
				//fmt.Println(a,k,flag)
			}
		} else {
			//fmt.Println("exit")
			os.Exit(0)
		}
	}
}

//
//var a []int=make([]int,5)
//a[0]=1
//a[1]=2
//a[2]=7
//a[3]=15
//a[4]=19
//var b int=9
//myleetcode.TwoSum(a,b)

//----------------------------------------------------------------------------------------------------------------------
func reverse(x int) {

}

//----------------------------------------------------------------------------------------------------------------------
func IsPalindrome(x int) bool {
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var revertedNumber int = 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber =  ，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10

}

//----------------------------------------------------------------------------------------------------------------------

// Definition for singly-linked list.
type ListNode struct {
	Data int
	Next *ListNode
}

//两个数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int = 0
	var numFlag = 0
	var addNum int = 0
	var addNums int = 0
	var add1 *ListNode = l1
	var add2 *ListNode = l2
	var headLink *ListNode
	var newNode *ListNode
	var node ListNode
	headLink = &node
	newNode = headLink
	var flag bool = true
	for flag {

		addNum = add1.Data  //加数
		addNums = add2.Data //被加数
		addNums = addNum + addNums + carry
		carry = 0
		if addNums >= 10 {
			carry = 1
			addNums = addNums - 10
		}
		if numFlag == 0 {
			node.Data = addNums
			node.Next = nil
			numFlag = 1

		} else {
			var nodes ListNode
			nodes.Data = addNums
			newNode.Next = &nodes
			newNode = &nodes
		}
		if add1.Next == nil {
			flag = false
			var fleg bool = true
			for fleg {
				if add2.Next == nil {
					fleg = false
				}
				if fleg {
					var xnodes ListNode
					add2 = add2.Next
					xnodes.Data = add2.Data
					xnodes.Next = add2.Next
					newNode.Next = &xnodes
					newNode = &xnodes
				}

			}
		} else if add2.Next == nil {
			flag = false
			var fleg bool = true
			for fleg {
				if add1.Next == nil {
					fleg = false
				}
				if fleg {
					var xnodes ListNode
					add1 = add1.Next
					xnodes.Data = add1.Data
					xnodes.Next = add1.Next
					newNode.Next = &xnodes
					newNode = &xnodes
				}
			}
		}
		if flag {
			add1 = add1.Next
			add2 = add2.Next
		}

	}
	return headLink

}

func Test() {
	var a1 ListNode
	var a2 ListNode
	var a3 ListNode
	var a4 ListNode
	var a5 ListNode
	a1.Data = 1
	a2.Data = 2
	a3.Data = 3
	a4.Data = 9
	a5.Data = 3
	a1.Next = &a2
	a2.Next = &a3
	a4.Next = &a5
	var xx *ListNode = addTwoNumbers(&a1, &a4)
	var flag int = 0
	for flag == 0 {
		fmt.Println(xx.Data)
		if xx.Next == nil {
			flag = 1
		} else {
			xx = xx.Next
		}

	}
}

//----------------------------------------------------------------------------------------------------------------------
//func LengthOfLongestSubstring(s string) int {
//	var before int=0
//	var after int =1
//	var res string="test"
//	var resInt int =0
//	flag:=0
//	for flag=0;flag< len(s);flag++{
//		if s[before:after]==res{
//
//		}else {
//			resInt++
//		}
//	}
//
//}

//----------------------------------------------------------------------------------------------------------------------
//strs:=make([]string,3)
//strs[0]="flower"
//strs[1]="flow"
//strs[2]="flight"
//strs[0]="aaaaaaaaaa"
//strs[1]="aaaaaaaaaaaa"
//strs[2]="aaaaaaaa"
//strs=append(strs, "aaa")
//strs=append(strs, "aab")

//fmt.Println(myleetcode.LongestCommonPrefix(strs))
//
func LongestCommonPrefix(strs []string) string {
	var str string
	var index int
	//flagMap:=make(map[string]int)
	flag := make([]int, len(strs))
	var fleg bool = true
	for k, v := range strs {
		if k == 0 {
			//flagMap[v]=0
			//flag=append(flag, 0)
			flag[k] = 0
			index = 0
		} else {
			//flagMap[v]=len(str)
			if len(strs[index]) > len(v) {
				index = k
			}
			if v != "" {
				flag[k] = len(str)
				//flag=append(flag, len(str))
			}
		}
		str = str + v
	}
	var res string
	var temp string = str[0:1]
	//var fleg bool=true
	for i := 0; i < len(strs[index]); i++ {
		for k1, v1 := range flag {
			//fmt.Println(k1)
			if fleg {
				temp = str[v1 : v1+1]
				fleg = false
			}
			if temp == str[v1:v1+1] {
				//if fleg{
				//	flagMap[k1]=v1+1
				//	fleg=false
				//}else {
				//	temp=str[v1:v1+1]
				//	flagMap[k1]=v1+1
				//}
				temp = str[v1 : v1+1]
				flag[k1] = v1 + 1
			} else {
				return res
			}
		}
		fleg = true
		res = res + temp

	}
	return res

}

//----------------------------------------------------------------------------------------------------------------------
func StrStr(haystack string, needle string) int {
	var res int = -1
	var flag []int = make([]int, len(haystack))
	var index int = 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i:i+1] == needle[0:1] {
			flag[index] = i
			index++
		}
	}
	for _, v := range flag {
		if needle == haystack[v:v+len(needle)] {
			res = v
			return res
		}

	}
	return res

}

//----------------------------------------------------------------------------------------------------------------------
type Result struct {
	First  int
	Second int
	Third  int
}

func findElement(m map[int]int, i int) bool {
	if _, ok := m[i]; !ok {
		return false
	}
	return true
}
func ThreeSum(nums []int) [][]int {
	index := 0
	var res_s [][]int = make([][]int, 0, 5)
	//max_flag:=0
	maps := make(map[int]int)
	for _, v := range nums {
		//if v>=0 {
		//	if max_flag>=0{
		//		if v>max_flag {
		//			max_flag=v
		//		}
		//	}else {
		//		if 0-v<max_flag{
		//			max_flag=v
		//		}
		//	}
		//}else {
		//	if max_flag>=0{
		//		if 0-v>max_flag {
		//			max_flag=v
		//		}
		//	}else {
		//		if v<max_flag{
		//			max_flag=v
		//		}
		//	}
		//}
		if findElement(maps, v) {
			maps[v] = maps[v] + 1
		} else {
			maps[v] = 1
		}
	}
	for k, _ := range maps {
		//temp_v:=v
		temp_k := k
		res := Result{}
		res.First = temp_k
		if k > 0 {
			for i := k; i > 0; i-- {
				fleg := false
				if findElement(maps, 0-i) {
					res.Second = 0 - i
					temp_k = 0 - k + i
					fleg = true
				}
				if findElement(maps, temp_k) && fleg {
					if temp_k != 0 {
						if 0-i == temp_k {
							if maps[temp_k] == 2 || maps[temp_k] == 12 {
								//maps[temp_k]=2
								//if maps[0-i]==11 || maps[0-i]==1{
								//	maps[0-i]=1
								//}else {
								//	maps[0-i]=2
								//}
								res.Third = temp_k
								res_s = append(res_s, []int{res.First, res.Second, res.Third})
								index++
								fmt.Println(res)
							}
						} else {
							if maps[temp_k] != 1 && (maps[temp_k] == 12 || maps[temp_k] == 11) {
								if maps[temp_k] == 11 {
									maps[temp_k] = 1
									if maps[0-i] == 11 {
										maps[0-i] = 1
									} else {
										if maps[0-i] != 1 {
											maps[0-i] = 2
										}
									}
								} else {
									maps[temp_k] = 2
									if maps[0-i] == 11 {
										maps[0-i] = 1
									} else {
										if maps[0-i] != 1 {
											maps[0-i] = 2
										}
									}
								}
							} else {
								res.Third = temp_k
								maps[temp_k] = maps[temp_k] + 10
								maps[0-i] = maps[0-i] + 10
								res_s = append(res_s, []int{res.First, res.Second, res.Third})
								index++
								fmt.Println(res)
							}
						}

					}

					//if 0-i==temp_k{
					//	if maps[temp_k]>1{
					//		res.Third=temp_k
					//		fmt.Println(res)
					//	}
					//}else{
					//	res.Third=temp_k
					//	fmt.Println(res)
					//}
				}
			}
		} else if k < 0 {
			for i := k; i < 0; i++ {
				fleg := false
				if findElement(maps, 0-i) {
					res.Second = 0 - i
					temp_k = 0 - k + i
					fleg = true
				}
				if findElement(maps, temp_k) && fleg {
					if 0-i == temp_k {
						if maps[temp_k] == 2 || maps[temp_k] == 12 {
							//maps[temp_k]=2
							//if maps[0-i]==11 || maps[0-i]==1{
							//	maps[0-i]=1
							//}else {
							//	maps[0-i]=2
							//}
							res.Third = temp_k
							res_s = append(res_s, []int{res.First, res.Second, res.Third})
							index++
							fmt.Println(res)
						}
					} else {
						if maps[temp_k] != 1 && (maps[temp_k] == 12 || maps[temp_k] == 11) {
							if maps[temp_k] == 11 {
								maps[temp_k] = 1
								if maps[0-i] == 11 {
									maps[0-i] = 1
								} else {
									if maps[0-i] != 1 {
										maps[0-i] = 2
									}
								}
							} else {
								maps[temp_k] = 2
								if maps[0-i] == 11 {
									maps[0-i] = 1
								} else {
									if maps[0-i] != 1 {
										maps[0-i] = 2
									}
								}
							}
						} else {
							if temp_k == 0 {
								res.Third = temp_k
								//maps[temp_k]=maps[temp_k]+10
								maps[0-i] = maps[0-i] + 10
								res_s = append(res_s, []int{res.First, res.Second, res.Third})
								index++
								fmt.Println(res)
							} else {
								res.Third = temp_k
								maps[temp_k] = maps[temp_k] + 10
								maps[0-i] = maps[0-i] + 10
								res_s = append(res_s, []int{res.First, res.Second, res.Third})
								index++
								fmt.Println(res)
							}
						}
					}

					//if 0-i==temp_k{
					//	if maps[temp_k]>1{
					//		res.Third=temp_k
					//		fmt.Println(res)
					//	}
					//}else{
					//	res.Third=temp_k
					//	fmt.Println(res)
					//}
				}
			}
		} else {
			if maps[k] >= 3 {
				res_s = append(res_s, []int{0, 0, 0})
				index++
				fmt.Println([]int{0, 0, 0})
			}
		}
		//	if max_flag>0{
		//		if k>0 {
		//			for i:=k;i>0 ;i--  {
		//				fleg:=false
		//				if findElement(maps,0-i) {
		//					res.Second=0-i
		//					temp_k=0-k+i
		//					fleg=true
		//				}
		//				if findElement(maps,temp_k) && fleg{
		//					if 0-i==temp_k{
		//						if maps[temp_k]>1{
		//							res.Third=temp_k
		//							fmt.Println(res)
		//						}
		//					}else{
		//						res.Third=temp_k
		//						fmt.Println(res)
		//					}
		//				}
		//			}
		//		}
		//	}else{
		//		if k<0 {
		//			for i:=k;i<0 ;i++  {
		//				fleg:=false
		//				if findElement(maps,0-i) {
		//					res.Second=0-i
		//					temp_k=0-k+i
		//					fleg=true
		//				}
		//				if findElement(maps,temp_k) && fleg{
		//					if 0-i==temp_k{
		//						if maps[temp_k]>1{
		//							res.Third=temp_k
		//							fmt.Println(res)
		//						}
		//					}else{
		//						res.Third=temp_k
		//						fmt.Println(res)
		//					}
		//				}
		//			}
		//		}
		//	}
	}
	return res_s

}

//myleetcode.ThreeSum([]int{-1,0,1,2,-1,-4})
//{-1 1 0}
//{2 -2 0}
//{2 -1 -1}
//{-5 4 1}
//{-5 1 4}
//{6 -5 -1}
//{6 -1 -5}
//{-1 1 0}
//{-3 2 1}
//{-3 1 2}
//{-2 2 0}
//{4 -3 -1}
//{4 -1 -3}
//----------------------------------------------------------------------------------------------------------------------
//func threeSumClosest(nums []int, target int) int {
//	//temp_nums:=make([]int,0, len(nums))
//	//for k,v:= range  nums{
//	//	if k==0 {
//	//		temp_nums=append(temp_nums,v)
//	//	}else {
//	//
//	//	}
//	//}
//	//flag:=true
//	//before_flag:=0
//	//after_flag:=0
//	//for k,v:= range nums{
//	//	if v>target && flag {
//	//		before_flag=k-1
//	//		after_flag=k
//	//		flag=false
//	//	}
//	//}
//	//before:=nums[before_flag]+nums[after_flag]+nums[before_flag-1]
//	//after:=nums[before_flag]+nums[after_flag]+nums[after_flag+1]
//	//if before_flag {
//	//
//	//}
//
//}

//---------------------------Z字型遍历----------------------------------------------------------------------------------------
func Convert(s string, numRows int) string {
	var res string
	lines := 1
	flag := numRows % 2
	indes_first := 0
	//index_second:=0
	//if flag == 0{
	//	indes_first=(1+numRows)/2
	//	//index_second=(1+numRows)/2+1
	//}else {
	//	indes_first=(1+numRows)/2
	//}
	indes_first = (1 + numRows) / 2
	for lines = 1; lines <= numRows; lines++ {
		index := lines - 1
		index_bool := true
		flag_bool := true
		for flag_bool {
			if lines > indes_first {
				if index < len(s) {
					if index_bool {
						res = res + s[index:index+1]
						//index=2*(numRows-2*indes_first+lines)
						index = index + 2*(numRows-lines)
						index_bool = false
					} else {
						if lines != numRows {
							res = res + s[index:index+1]
						}
						index = index + 2*(lines-1)
						index_bool = true
					}
				} else {
					flag_bool = false
				}
			} else if lines < indes_first || flag == 0 {
				if index < len(s) {
					if index_bool {
						res = res + s[index:index+1]
						index = index + 2*(numRows-lines)
						index_bool = false
					} else {
						if lines-1 != 0 {
							res = res + s[index:index+1]
						}
						index = index + 2*(lines-1)
						index_bool = true
					}
				} else {
					flag_bool = false
				}
			} else {
				if index < len(s) {
					res = res + s[index:index+1]
					index = index + 2*(numRows-indes_first)
				} else {
					flag_bool = false
				}
			}
		}
		//fmt.Println(res)
	}
	return res
}

//---------------------------------------有效的数独----------------------------------------------------------------------
func findElements(m map[byte]int, i byte) bool {
	if _, ok := m[i]; !ok {
		return false
	}
	return true
}
func IsValidSudoku(board [][]byte) {
	//var res bool=true
	maps := make(map[int]byte)
	map_line := make(map[byte]int)
	for k, v := range board {
		for ks, vs := range v {
			if vs != '.' {
				if findElements(map_line, vs) {
					fmt.Println(false)
					os.Exit(0)
				}
				map_line[vs] = k*10 + 10 + ks + 1
				maps[k*10+10+ks+1] = vs
			}
		}
	}
	fmt.Println(maps)

}
