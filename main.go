package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

var (
    vocab = make([]map[string][]string, 0)
)

func contains(vals []string, input string) int {
    for i, s := range vals {
        if s == input {
            return i
        }
    }
    return -1
}

func ask() {
    i := 0
    l := len(vocab)
    for {
        for k, vals := range vocab[i] {
            for {
                fmt.Print(k + " : ")
                var input string
                fmt.Scanln(&input)

                if input == "n" {
                    i++
                    break
                }
                if input == "s" {
                    continue
                }
                if input == "h" {
                    for _, v := range vals {
                        fmt.Printf("%v...\n", v[:3])
                    }
                    continue
                }
                if input == "hh" {
                    for _, v := range vals {
                        fmt.Printf("%v\n", v)
                    }
                    continue
                }

                index := contains(vals, input)
                if index != -1 {
                    fmt.Println(":)")
                    vals = append(vals[:index], vals[index+1:]...)
                    if len(vals) == 0 {
                        vocab = append(vocab[:i], vocab[i+1:]...)
                        l--
                        break
                    }
                    fmt.Printf("%d left\n", len(vals))
                } else {
                    fmt.Println(":(")
                }
            }
            if i == l {
                ask()
            }
            if l == 0 {
                fmt.Println("Yeah yoah finished")
                var input string
                fmt.Scanln(&input)
                if input == "y" {
                    load()
                    ask()
                } else {
                    return
                }
            }
        }
    }
}

func load() {
	file, err := os.Open("./vocab/en-de")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    i := 0
    for scanner.Scan() {
        vocab = append(vocab, make(map[string][]string, 1))
        q := strings.Split(scanner.Text(), "-")
        key := strings.TrimSpace(q[0])
        values := strings.Split(q[1], ",")
        for _, val := range values {
            if _, ok := vocab[i][key]; !ok {
                vocab[i][key] = make([]string, 0)
            }
            vocab[i][key] = append(vocab[i][key], strings.TrimSpace(val))
        }
        i++
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    log.Printf("%v\n", vocab)
}
func main() {
    load()
    ask()
}
