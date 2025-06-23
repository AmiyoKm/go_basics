package challenge

import (
    "bufio"
    "fmt"
    "os"
    "sync"
)

func user(name string, in <-chan string, out chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    reader := bufio.NewReader(os.Stdin)
    for {
        // Receive message
        msg, ok := <-in
        if !ok {
            return
        }
        fmt.Printf("%s received: %s\n", name, msg)

        // Prompt for reply
        fmt.Printf("%s, enter your message: ", name)
        reply, _ := reader.ReadString('\n')
        out <- reply
    }
}

func Chat() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    var wg sync.WaitGroup

    wg.Add(2)
    go user("User 1", ch1, ch2, &wg)
    go user("User 2", ch2, ch1, &wg)


    reader := bufio.NewReader(os.Stdin)
    fmt.Print("User 1, enter your message: ")
    msg, _ := reader.ReadString('\n')
    ch2 <- msg

    wg.Wait()
}