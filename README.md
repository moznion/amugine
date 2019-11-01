amugine [![CircleCI](https://circleci.com/gh/moznion/amugine.svg?style=svg)](https://circleci.com/gh/moznion/amugine)
==

A simple encrypter/decrypter CLI tool. This tool provides GPG passphrase based encryption feature.

Pre-compiled binaries are here: https://github.com/moznion/amugine/releases

Description
--

This tool provides GPG passphrase based encryption feature (i.e. it acts like `gpg -c` and `gpg -d --batch --passphrase-fd 0` commands).

It's easy to use such a feature by one-binary and to combine existing other tools which use GPG.

Usage
--

```
Usage: amugine <flags> <subcommand> <subcommand args>

Subcommands:
        commands         list all command names
        decrypt          Decrypt a given payload.
        encrypt          Encrypt s given payload.
        flags            describe all known top-level flags
        help             describe subcommands and their syntax
        version          Show the version of this tool.
```

```
encrypt <key> <payload>:
        Encrypt a given payload. It supports to give payload through an argument and STDIN (example: "cat payload.txt | amugine encrypt <key>").
        If a parameter has "@" prefix, it will be handled as the filepath and read that.
```

```
decrypt <key> <payload>:
        Decrypt a given payload. If a parameter has "@" prefix, it will be handled as the filepath and read that.
```

Use "amugine flags" for a list of top-level flags

How to build
--

```
$ make all VERSION=x.y.z
```

License
--

```
The MIT License (MIT)
Copyright © 2019 moznion, http://moznion.net/ <moznion@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the “Software”), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```
