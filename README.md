This is just a little hack tool I made because I couldn't figure out an appropriate `sed`, `awk`, or `perl` script.

All this does is removes blank lines inside the `import ( )` block of a Go file.

Follow this up with a call to `goimports -w` or `goimports -w -local` to put
the blanks back in at the desired spots. Leverage your shell to do it to a
bunch of files at once:

```sh
for f in **/*.go; do condenseimports "$f" && goimports -w "$f"; done
```

It is very hacky and generally "bad" code.
