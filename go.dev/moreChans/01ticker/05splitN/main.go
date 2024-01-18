package main

func splitN(from chan string, n uint) []chan string {
	// []chan string: slice of channels that can send and receive strings.
	tos := make([]chan string, int(n)) // a slice of channels (with zero values) with a capacity equal to the integer value of n.
	//  tos is capable of holding n channels of type chan string.

	for i := uint(0); i < n; i++ {
		tos[i] = make(chan string, cap(from))
	}

	go func ()  {
		defer func ()  {
			for _, to := range tos {
				close(to)
			}
		}()
		for data := range from {
			for _, to := range tos {
				to <- data
			}
		}
	}
}