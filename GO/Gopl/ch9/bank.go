package main

//var (
//	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
//	balance int
//)
//
//// 共享变量
//func Deposit(amount int) {
//	sema <- struct{}{} // acquire token
//	balance = balance + amount
//	<-sema // release token
//}
//
//func Balance() int {
//	sema <- struct{}{} // acquire token
//	b := balance
//	<-sema // release token
//	return b
//}
