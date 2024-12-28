---
Title: go zerolog setup with pgx
Summary: First try writing about something more technical. My own setup with zerolog and pgx.
Author: Peter Paravinja
Published: 28. Dec 2024
---

## Overview

Golang currently holds really special place in my heart ❤️

But since I'm doing mostly user facing oriented projects I usually need a database.

My weapon of choice is Postgres (like many others).

But I really want to know how queries are doing.

I am a freak and want to optimize things.
Now since I sometimes don't think ahead enough - I get "burned" and I found out that looking at my logs in terminal
usually catch this. So I want to add zerolog to pgx so that I can see traces and how long a query took to execute. This
way I can catch it easily.

```text
9:08PM TRC Query args=[] commandTag="DELETE 0" duration=4.653 pid=123 sql="-- name: CleanupExpiredTokens :exec\nDELETE\nFROM refresh_tokens\nWHERE expires_at < CURRENT_TIMESTAMP\n   OR revoked_at IS NOT NULL\n"
```

Why I like [zerolog](https://github.com/rs/zerolog)?
I started with normal golang log API and was having okay experience. But started sniffing around the web to see what
other people were using.

At the same time I was developing my first bigger hobby project and started to hit the "wall" of project structure "
smell".

I didn't like the way the project was structured.

I have a pretty okay nose for figuring out that something isn't at the best level... I have a tougher time figuring out
a solution for that problem on my own.

So I started looking on the web and found out a
very [nice public repository](https://github.com/learning-cloud-native-go/myapp) depicting a project structure that I
really liked.

I really liked it - good job [dumindu](https://github.com/dumindu) and thanks for sharing the idea!

Afterwards I checked hashicorp who I think are really next level company in golang space... dear god - things they have
open source... the amount of knowledge someone can gain from them by just reading and implementing ideas from them...
Amazing! Even
the [comment out thoughts](https://github.com/hashicorp/packer/blob/main/helper/wrappedreadline/wrappedreadline.go#L4)
are a very nice touch!

Anyway... lets talk about zerolog and pgx working somehow hand in hand!

This is more of a general idea what you can do with it and a few tips on how to use it.

It's not a general defacto way or anything... just my own implementation that I liked.

Let's start with dependencies I guess.

I am working with versions:
zerolog: 1.33.0
pgx: v5 5.7.1

```shell
go get -u github.com/rs/zerolog/log
go get github.com/jackc/pgx/v5
```

Now I personally like to split the depth of logging into 2 separate concerns. One is for local, development and staging
environments - the other is for production.

Now this is totally up to you.

I like to triple tests stuff before publishing to production, so I don't usually include traces there. Buuuuut you
theoretically could - its just... depends how much logs can your production take. If you have high traffic... you will
fill them big boy log files up.

What we'll do now is create a function that will create a new instance of zerolog with some configs that we want

```go
package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
)

func New(isDebug bool, logPath string) (*zerolog.Logger, error) {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.TraceLevel
	}

	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		return nil, fmt.Errorf("create log directory: %w", err)
	}

	runLogFile, err := os.OpenFile(
		logPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		return nil, fmt.Errorf("open log file: %w", err)
	}

	zerolog.SetGlobalLevel(logLevel)
	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, runLogFile)
	logger := zerolog.New(multi).With().Timestamp().Logger()

	return &logger, nil
}

```

now we want 2 level writers - one for the console standard out and one for appending to the log file.

We can set global level - for production we want Information level events - and for debug environments we lower that
down to trace so we see everything!

But this config depends entirely up to you!

The last line
`logger := zerolog.New(multi).With().Timestamp().Logger()`

Creates a new logger and we return a pointer to that logger.
Alternative solution would be that we create a struct and assign the pointer to a property inside the struct. We could
add the pointer to *os.File as-well so we can do operations on it (closing it for example).

After that we instantiate this new beautiful piece of abstract code:

`l, _ := logger.New(conf.Server.Debug, "logs/app.log")`

_!!deal with your errors boys and girls!!_

Ok now we'll create a new Trace Logger!

We need to satisfy this daddy interface

```go
type Logger interface {
    Log(ctx context.Context, level LogLevel, msg string, data map[string]any)
}
```

After that we'll use this new struct and method Log and append him on this pgx5 struct:
```go
// /github.com/jackc/pgx/v5@v5.7.1/tracelog/tracelog.go:135
type TraceLog struct {
    Logger   Logger
    LogLevel LogLevel

	Config           *TraceLogConfig
	ensureConfigOnce sync.Once
}
```

Let's bring it all together now!

```go
package logger

import (
	"context"

	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
)

type zerologPgxLogger struct {
	logger zerolog.Logger
}

func (l zerologPgxLogger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
	logEvent := l.logger.With().Fields(data).Logger()
	zerologLevel := l.logger.GetLevel()
	if zerologLevel == zerolog.TraceLevel {
		level = tracelog.LogLevelTrace
	}

	switch level {
	case tracelog.LogLevelTrace:
		logEvent.Trace().Msg(msg)
	case tracelog.LogLevelDebug:
		logEvent.Debug().Msg(msg)
	case tracelog.LogLevelInfo:
		logEvent.Info().Msg(msg)
	case tracelog.LogLevelWarn:
		logEvent.Warn().Msg(msg)
	case tracelog.LogLevelError:
		logEvent.Error().Msg(msg)
	default:
		logEvent.Info().Msg(msg)
	}
}

func NewTraceLogger(zeroLog zerolog.Logger, isDebug bool) *tracelog.TraceLog {
	logLevel := tracelog.LogLevelInfo
	if isDebug {
		logLevel = tracelog.LogLevelTrace
	}

	return &tracelog.TraceLog{
		Logger:   zerologPgxLogger{logger: zeroLog},
		LogLevel: logLevel,
		Config: &tracelog.TraceLogConfig{
			TimeKey: "duration",
		},
	}
}
```

so inside Log method we basically decide what we want to do with the message that we get and data that we get - for this example we're just logging message that we get and we route it to correct log level event.

The only thing that is really important is the 

`TimeKey: "duration"`

This is how we'll get duration of the query! I think this is really, really important, and I am sad that out of the box the time value gets overwritten with time when the query took place (that is important as well! But we need both!!!!)

https://github.com/jackc/pgx/pull/2098 <- say thanks ❤️

Next comes creating a multi query which means satistfying another interface daddy, this time:

```go
// QueryTracer traces Query, QueryRow, and Exec.
type QueryTracer interface {
    // TraceQueryStart is called at the beginning of Query, QueryRow, and Exec calls. The returned context is used for the
    // rest of the call and will be passed to TraceQueryEnd.
    TraceQueryStart(ctx context.Context, conn *Conn, data TraceQueryStartData) context.Context

	TraceQueryEnd(ctx context.Context, conn *Conn, data TraceQueryEndData)
}
```

Why MultiQuery?

Because mby in the future you'll want to implement OpenTelemetry Tracer (https://github.com/exaring/otelpgx) - why should I stop you with just standard logging - you can do so much more!!! I am feeling observability slowly becoming "my thing". I will dive deeper! I will go so fu*?ing deep!

Anyway...


```go
package logger

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type MultiQueryTracer struct {
	Tracers []pgx.QueryTracer
}

func (m *MultiQueryTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	for _, t := range m.Tracers {
		ctx = t.TraceQueryStart(ctx, conn, data)
	}

	return ctx
}

func (m *MultiQueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	for _, t := range m.Tracers {
		t.TraceQueryEnd(ctx, conn, data)
	}
}
```

Basically we're just ranging over the 2 traces and telling them to do their thing - nothing else... just a smoll wrapper :)

Now we have all the components!

We'll just go and combine them.

```go
func main() {
	// prepare everything
	l, _ := logger.New(conf.Server.Debug, "logs/app.log")
	traceLogger := logger.NewTraceLogger(*l, conf.Server.Debug)
	m := logger.MultiQueryTracer{
		Tracers: []pgx.QueryTracer{
			// open telemetry tracer: https://github.com/exaring/otelpgx
			otelpgx.NewTracer(),
			// logger
			traceLogger,
		},
	}

	l.Info().Msg("Initializing database connection")
	dbString := fmt.Sprintf(fmtDBString, conf.DB.Host, conf.DB.Username, conf.DB.Password, conf.DB.DBName, conf.DB.Port)
	l.Info().
		Str("host", conf.DB.Host).
		Int("port", conf.DB.Port).
		Str("database", conf.DB.DBName).
		Msg("Connecting to database")

	dbConfig, err := pgxpool.ParseConfig(dbString)
	if err != nil {
		l.Fatal().
			Err(err).
			Msg("Failed to parse database config")
		return
	}
	dbConfig.ConnConfig.Tracer = &m

	pool, err := pgxpool.NewWithConfig(
		ctx,
		dbConfig,
	)
	if err != nil {
		l.Fatal().
			Err(err).
			Msg("Failed to create database connection pool")
		return
	}
	defer pool.Close()
	l.Info().Msg("Database connection established successfully")
}
```

This is it actually... you should have a fully functional trace logging into your log file and console output!

![tracer-query.png](/public/assets/tracer-query.png)

Let's dive into it even more sometime in the future! 😍