package keygen

func KeyGen(limit int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 1; i < limit; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
