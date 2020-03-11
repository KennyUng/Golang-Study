// in built make function for channels
// ch := make(chan int) -- Unbuffered
// ch := make(chan int, cap int) -- buffered

// Send channel ch <- x
// Receive channel x = <- ch
// Close: Close(ch)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func{
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done

}