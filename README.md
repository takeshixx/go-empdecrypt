# EmpDecrypt

This project implemetns decrypting of EIS passwords in Matrix42 configuration files. It is basically a port of [EmpEISDecrypt](https://github.com/S3cur3Th1sSh1t/EmpEISDecrypt), all credits go to the original authors.

The reasons why it has been ported:

* Implementation in Go is more portable and can be built/used more easily (no VS and .Net required)
* A [decoy](decoy/matrix.c) lib is included so that `Matrix42.Common.AppVerificator.dll` is not required
* Implement a more generic interface that allows to pass strings/files/paths via CLI arguments to make it easier to work with many files.

This project requires the `EmpCrypt.exe` from Matrix42. Due to copyright issues this can't be included in this project.

## Building

Install the building dependencies on Windows with [Chocolatey](https://chocolatey.org/):

```
choco install make
choco install mingw -x86
```

Build the main binary:

```bash
make
```

Build the decoy DLL:

```bash
make dll
```

Copy `EmpCrypt.exe` into the `bin/` directory.
