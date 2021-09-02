package main

//
//import (
//	"fmt"
//	"reflect"
//	"runtime"
//	"sync"
//)
//
//func main()  {
//	hy := 10
//	fmt.Println(reflect.TypeOf(hy))
//	runtime.GOMAXPROCS(1)
//
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	fmt.Println("start Goroutines")
//
//	go func() {
//		defer wg.Done()
//
//		for count:=0;count<3;count++{
//			for char :='a'; char<'a'+26; char++{
//				fmt.Printf("%c ", char)
//			}
//		}
//	}()
//
//	go func() {
//		defer wg.Done()
//
//		for count:=0;count<3;count++{
//			for char :='A'; char<'A'+26; char++{
//				fmt.Printf("%c ", char)
//			}
//		}
//	}()
//
//	fmt.Println("wait to finish")
//	wg.Wait()
//	fmt.Println("termination")
//}
