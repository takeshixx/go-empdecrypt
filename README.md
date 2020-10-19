# EmpDecrypt

This project implemetns decrypting of EIS passwords in Matrix42 configuration files. It is basically a port of [EmpEISDecrypt](https://github.com/S3cur3Th1sSh1t/EmpEISDecrypt), all credits go to the original authors.

The reasons why it has been ported:

* Implementation in Go is more portable and can be built/used more easily (no VS and .Net required)
* A [decoy](decoy/matrix.c) lib is included so that `Matrix42.Common.AppVerificator.dll` is not required
* Implement a more generic interface that allows to pass strings/files/paths via CLI arguments to make it easier to work with many files.

This project requires the `EmpCrypt.exe` from Matrix42. Due to copyright issues this can't be included in this project.

## Usage

Provide hashs as position arguments (*Note*: hashs most likely contain single and double quotes):

```powershell
.\bin\empdecrypt A(,'-&-#+# /"*&(',.+ )*/!$%-..,/!)*")+$% X
```

When using the non-embedded binary or when `EmpDecrypt.exe` is not in the current path, provide the path with the `-c` flag:
    
```powershell
.\bin\empdecrypt -c "..\..\EmpCrypt.exe"
```

Make sure the `Matrix42.Common.AppVerificator.dll` is in the same path as `EmpCrypt.exe`. Either use the original one or the decoy lib that has previously been built with `make dll`.

In case many files have to be decrypted, it is recommended to use the following command:
    
```powershell
.\bin\empdecrypt -p x:\path\to\ini\files\ | tee-object -filepath z:\output\path\passwords.txt
```

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

## Embedded Build

In order to create an embedded build, first create a `resources` folder and copy the external dependencies into that folder:

```bash
mkdir resources
cp bin/Matrix42.Common.AppVerificator.dll resources
cp EmpCrypt.exe resources
```

Then run [pkger](https://github.com/markbates/pkger):

```bash
pkger
```

And finally build the embedded binary:

```bash
make embedded
```

The resulting binary will have no external dependencies.
