goparsetime
===========

Go wrapper for parsetime.c from UNIX "at" command.

Parse date strings and return a `time.Time`

Usage
-----

Parse basic timestring

```Go
    t, err := goparsetime.Parsetime("10:00 AM UTC Oct 21 2015")
    fmt.Printf("%v\n", t.UTC()) // 2015-10-21 10:00:00 +0000 UTC
```

Relative date math

```Go
    t, err = goparsetime.Parsetime("10:00 AM UTC Oct 21 2015 + 5 days")
    fmt.Printf("%v\n", t.UTC()) // 2015-10-26 10:00:00 +0000 UTC
```

Relative math to current time

```Go
    now := time.Now() // 2015-10-21 19:45:18.035108719 +0000 UTC
    t, err = goparsetime.Parsetime("now - 2 weeks + 3 hours")
    fmt.Printf("%v\n", t.UTC()) // 2015-10-07 22:45:18 +0000 UTC
```
