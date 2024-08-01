# neo4j-go-driver-panic

Start DB:
```
./db.sh
```

Populate DB with dummy data:
```
% go run ./main.go populate
Populating DB ca8256fb-1ac3-4095-8701-20e63b931ad6
```

Run the query:
```
% go run ./main.go -uuid ca8256fb-1ac3-4095-8701-20e63b931ad6 query
Node1 917
Node2 1016
panic: runtime error: makeslice: len out of range
        panic: runtime error: makeslice: len out of range

goroutine 1 [running]:
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*hydrator).path(0x14000196018, 0x0?)
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/hydrator.go:696 +0x6f0
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*hydrator).value(0x14000196018)
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/hydrator.go:465 +0x300
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*hydrator).record(0x14000196018, 0x52f4?)
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/hydrator.go:428 +0x264
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*hydrator).hydrate(0x14000196018, {0x14000236000?, 0x1033c8048?, 0x14000184060?})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/hydrator.go:162 +0x21c
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*incoming).next(0x14000196000, {0x102e6bd70?, 0x10302db20?}, {0x1033c8048?, 0x14000184060?})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/incoming.go:40 +0xb4
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*messageQueue).receiveMsg(0x14000190048, {0x102e6bd70, 0x10302db20})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/message_queue.go:209 +0x80
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*messageQueue).receive(0x14000190048, {0x102e6bd70, 0x10302db20})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/message_queue.go:153 +0x28
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*messageQueue).receiveAll(0x14000190048, {0x102e6bd70, 0x10302db20})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/message_queue.go:146 +0x40
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*bolt5).ForceReset(0x14000190000, {0x102e6bd70, 0x10302db20})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/bolt5.go:808 +0x68
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*bolt5).Reset(0xc1a323a342581130?, {0x102e6bd70?, 0x10302db20?})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/bolt5.go:796 +0x74
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/pool.(*Pool).Return(0x1400012c090, {0x102e6bd70, 0x10302db20}, {0x102e6e928, 0x14000190000})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/pool/pool.go:410 +0x1e4
github.com/neo4j/neo4j-go-driver/v5/neo4j.(*sessionWithContext).executeTransactionFunction.func2()
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/session_with_context.go:476 +0x3c
panic({0x102e29d20?, 0x102e69aa0?})
        /opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/panic.go:770 +0x124
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*hydrator).path(0x14000196018, 0x0?)
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/hydrator.go:696 +0x6f0
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*hydrator).value(0x14000196018)
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/hydrator.go:465 +0x300
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*hydrator).record(0x14000196018, 0x52f4?)
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/hydrator.go:428 +0x264
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*hydrator).hydrate(0x14000196018, {0x14000236000?, 0x1033c8048?, 0x14000184060?})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/hydrator.go:162 +0x21c
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*incoming).next(0x14000196000, {0x102e6bd70?, 0x10302db20?}, {0x1033c8048?, 0x14000184060?})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/incoming.go:40 +0xb4
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*messageQueue).receiveMsg(0x14000190048, {0x102e6bd70, 0x10302db20})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/message_queue.go:209 +0x80
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*messageQueue).receive(0x14000190048, {0x102e6bd70, 0x10302db20})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/message_queue.go:153 +0x28
github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/bolt.(*bolt5).Next(0x14000190000, {0x102e6bd70, 0x10302db20}, {0x102e3f580?, 0x1400021c140?})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/internal/bolt/bolt5.go:673 +0x10c
github.com/neo4j/neo4j-go-driver/v5/neo4j.(*resultWithContext).advance(0x140001b6080, {0x102e6bd70?, 0x10302db20?})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/result_with_context.go:246 +0x48
github.com/neo4j/neo4j-go-driver/v5/neo4j.(*resultWithContext).Collect(0x140001b6080, {0x102e6bd70, 0x10302db20})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/result_with_context.go:153 +0x9c
main.queryDB.func1({0x102e6b3b8, 0x140002046f0})
        /Users/x/Projects/neo4j-go-driver-panic/main.go:156 +0x1b4
github.com/neo4j/neo4j-go-driver/v5/neo4j.(*sessionWithContext).executeTransactionFunction(0x14000132160, {0x102e6bd70, 0x10302db20}, 0x1, {0x14000139988?, 0x0?}, 0x140001c6000, 0x140001a62d0, 0x1, 0x0)
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/session_with_context.go:504 +0x2e4
github.com/neo4j/neo4j-go-driver/v5/neo4j.(*sessionWithContext).runRetriable(0x14000132160, {0x102e6bd70, 0x10302db20}, 0x1, 0x140001a62d0, 0x1, 0x0, {0x0, 0x0, 0x0?})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/session_with_context.go:443 +0x2dc
github.com/neo4j/neo4j-go-driver/v5/neo4j.(*sessionWithContext).ExecuteRead(0x102e6a538?, {0x102e6bd70?, 0x10302db20?}, 0x2?, {0x0?, 0x5?, 0x0?})
        /Users/x/go/pkg/mod/github.com/neo4j/neo4j-go-driver/v5@v5.23.0/neo4j/session_with_context.go:379 +0x44
main.queryDB({0x102e6bd70, 0x10302db20}, {0x16d24b0f7, 0x24}, 0x394, 0x3f7)
        /Users/x/Projects/neo4j-go-driver-panic/main.go:130 +0x5bc
main.main()
        /Users/x/Projects/neo4j-go-driver-panic/main.go:246 +0x2c8
exit status 2
```