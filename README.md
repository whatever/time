# Jitter Time

A ticker that "ticks" within a range of time, instead of on a set interval.


That is, `NewJitterTicker(1 * time.Second, 10 * time.Second)` will tick every
1-10 seconds.


# Usage

```go
import (
  "time"
  jitter_time "github.com/whatever/time"
)

ticker := jitter_time.NewJitterTicker(500 * time.Millisecond, 1  *time.Second)

for _ = range ticker.C {
}
```
