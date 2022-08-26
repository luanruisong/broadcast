## Broadcast

    module for notify with channel

### useage

```shell
go get -u github.com/luanruisong/broadcast
```

#### for signal without data
```go
    ss := NewSignalStation()
    for i := 0; i < 10; i++ {
        go func(idx int) {
            //for {
            //    select {
            //    case <-ss.CurrSignal():
            //        fmt.Println(idx, "receive signal")
            //    }
            //}
            ss.OnSignal(context.TODO(), func() {
                fmt.Println(idx, "receive signal")
            })
        }(i)
    }
    for i := 0; i < 10; i++ {
        ss.Send()
        time.Sleep(time.Second)
    }
```

#### for notice with data

```go
    ns := NewNoticeStation()
    for i := 0; i < 10; i++ {
        go func(idx int) {
            //for {
            //    fmt.Println(idx, ns.WaitForValue())
            //}
            ns.OnNotice(context.Background(), func(value interface{}) {
                fmt.Println(idx, value)
            })
        }(i)
    }
    for i := 0; i < 10; i++ {
        ns.Notice(time.Now().Format("2006-01-02 15:04:05"))
        time.Sleep(time.Second)
    }
```