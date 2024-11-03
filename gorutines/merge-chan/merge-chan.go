package main

func merge(chans ...chan int) chan int {
	out := make(chan int)
	done := make(chan struct{})

	for _, ch := range chans {
		for {
		go func() {
			select {
				case v, ok := <-ch:
				if !ok {
					close(done)
					return
				}
				select {
					case out <-v:
					case <-done:
						return
					}
				case <-done:
					return
				}
			}()
		}
	}
	go func() {
		<-done

		for _, ch := range chans {
			close(ch)
		}
		close(out)
	}()
	return out
}
