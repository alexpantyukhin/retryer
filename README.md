# Retryer

The tiny library for retrying any action.

# Motivation

I needed in some similar behavior in the one of project. A half of a minute of googling showed nothing. That's why the decision to extract the library was made.

# Plans
   - [] Add cancellation context.

# Usages

## Retry number of times

```go
	Retry(10).
		ExecuteBool(func() bool {
            counter++

			if counter == 5 {
				return true
			}

			return false
		})
```

## Retry with error

```go
	Retry(10).
		ExecuteError(func() error {
            conn, err := MakeConnection()

			return err
		})
```

## Retry with wait

```go
	RetryAndWait([]time.Duration{1000 * time.Millisecond, 1000 * time.Millisecond}).
		ExecuteError(func() error {
            conn, err := MakeConnection()

			return err
		})
```

## retry with wait forever

```go
	RetryAndWaitForever(func(attempt int) time.Duration { return time.Duration(attempt*100) * time.Millisecond }).
		ExecuteError(func() error {
            conn, err := MakeConnection()

			return err
		})
```

# Installation
Just go get this repository with the following way:

```
go get github.com/alexpantyukhin/retryer
```