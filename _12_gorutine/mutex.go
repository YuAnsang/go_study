package main

//var mutex sync.Mutex
//
//type Account struct {
//	Balance int
//}
//
//func DepositAndWithdraw(account *Account) {
//	mutex.Lock()
//	defer mutex.Unlock() // 한번 획득한 뮤택스는 반드시 Unlock() 호출을 해야함
//	if account.Balance < 0 {
//		panic(fmt.Sprintf("Balance must not be negative value : %d\n", account.Balance))
//	}
//
//	account.Balance += 1000
//	time.Sleep(time.Millisecond)
//	account.Balance -= 1000
//	fmt.Printf("Account Balance : %d\n", account.Balance)
//}

//func main() {
//	var wg sync.WaitGroup
//
//	account := &Account{0}
//	wg.Add(8)
//	for i := 0; i < 8; i++ {
//		go func() {
//			for {
//				DepositAndWithdraw(account)
//			}
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//}
