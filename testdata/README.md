This directory contains end-to-end integration tests for gopatch. Each test is
specified as a [txtar] file. This is the same format used by Go Playground to
support multiple files.

  [txtar]: https://godoc.org/github.com/rogpeppe/go-internal/txtar

Tests contain a one or more patch files which are executed upon one or more Go
file pairs, also specified in the patch file.

-   Patch files must be specified with the ".patch" suffix
-   Input files must be specified with the ".in.go" suffix
-   Output files must be specified with the ".out.go" suffix
-   Each input file must have an output file and vice versa

Input files will be renamed from ".in.go" to ".go" and the patches will be
executed on them in-order. Their new contents will be matched against the
associated ".out.go" files, marking the test as failed if they don't.

Test files will generally take the form,

    <comments>

    -- p1.patch --
    @@
    @@
    -foo
    +bar

    <more patches if needed>

    -- foo.in.go --
    package x

    // input source

    -- foo.out.go --
    package x

    // output source

    <more files if needed>