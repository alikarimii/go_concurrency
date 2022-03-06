package concurrencypatterns

func Tee(
	done <-chan interface{}, in <-chan interface{},
) (<-chan interface{}, <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})
	go func() {
		defer close(out1)
		defer close(out2)
		// The iteration over in cannot continue until
		// both out1 and out2 have been written to
		for val := range OrDone(done, in) {
			// We will want to use local versions of out1 and out2,
			//  so we shadow these variables.
			var out1, out2 = out1, out2
			// We’re going to use one select statement so that writes to
			// out1 and out2 don’t block each other.
			// To ensure both are written to, we’ll perform two
			// iterations of the select statement: one for each outbound channel.
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case out1 <- val:
					// Once we’ve written to a channel,
					// we set its shadowed copy to nil
					// so that further writes will block
					// and the other channel may continue.
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}
