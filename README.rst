Go implementation of filestorage2
=================================

This project is currently an exploration and learning exercise for Go.

Go's speed is very attractive for a storage server, which is a
potential bottleneck.  So far, I really hate programming in Go due to
it's insane approach to error handling.  Maybe if I keep plugging
away, I'll hate it less at some point. It could happen.

My main beef with Go's error handling is not just that it's a giant
DRY violation, but that idiomatic error checking leads to error
reports without tracebacks.  So, you detect errors when they occur,
but not where they occur.

Go's advocates argue that one should separate unexpected from expected
errors and only use exceptions for unexpected errors.  I can buy that
errors like Python's ``KeyError`` and ``StopIteration`` are a bad
idea. Most Pyton programmers use ``get`` rather than catching
``KeyError`` these days. And the iteration protocaol could have used
some sort of isempty test rather than ``StopIteration`` (or
``IndexError`` in the earlier sequence-based iteration protocol).
However, lots of other cases are only expected or unexpected based on
context.  Is a failure to write to an open file ever expected? Go
apparently thinks so.  The only really expected read failure when
reading from a file is end-of-file and Python deals with that pretty
sanely by returning no data.

Maybe Go's speed will be worth the hassle for a few applications. I'd
love to find time to compare Go-based and, say, Erlang/Elixir-based
implementations.

BTW, a little bit of pattern matching could go a long way to
addressing Go's error handling problem.  It might be nice if you could
state explicitly that an error was unexpected by pattern matching on nil:

  n, nil = file.Read(buf)

This would panic if a non-nil error was returned.  If course, then
you'd be stuck with panic/recover.
