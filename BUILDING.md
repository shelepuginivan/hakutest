# Building Hakutest

You can compile Hakutest from source. Follow the instructions below.

## Compiling from source

### Prerequisites

Make sure the following binaries are installed on your system and available in `$PATH`:

- `git`
- `go` (`>=1.24.0`)
- `make`
- `wget`

### Compilation steps

1.  Clone the repository:

     ```shell
     git clone https://github.com/shelepuginivan/hakutest.git
     cd hakutest
     ```

2.  Install the vendor dependencies:

    ```shell
    make web-vendor
    ```

    This will download [Alpine.js](https://alpinejs.dev/) from jsDelivr CDN
    into the `web/vendor/` directory.


3.  Optimize web static assets:

    ```shell
    make web-minify
    ```

    This step is optional, but recommended.

4.  Install the binaries:

    ```shell
    sudo make install
    ```

    This will compile Hakutest binaries (`hakutest` and `hakuctl`) and install
    them into `/usr/local/bin/`.

After completing these steps, Hakutest should be built and ready to use!
If you encounter an error, please open an issue describing what went wrong.

## Cross-compiling Hakutest from Linux to Windows

Hakutest can also be cross-compiled on Linux using
[MinGW](https://sourceforge.net/projects/mingw/).

### Requirements for cross-compilation

In addition to [prerequisites for normal compilation](#Prerequisites), you need
to ensure that the MinGW gcc (`x86_64-w64-mingw32-gcc`) is available on your
system.

### Steps

1.  Clone the repository:

     ```shell
     git clone https://github.com/shelepuginivan/hakutest.git
     cd hakutest
     ```

2.  Install the vendor dependencies:

    ```shell
    make web-vendor
    ```

    This will download [Alpine.js](https://alpinejs.dev/) from jsDelivr CDN
    into the `web/vendor/` directory.


3.  Optimize web static assets:

    ```shell
    make web-minify
    ```

    This step is optional, but recommended.

4.  Compile Hakutest binaries using `x86_64-w64-mingw32-gcc`:

    ```shell
    make build-windows
    ```

    This will compile Hakutest binaries (`hakutest.exe` and `hakuctl.exe`) into
    `target/windows/` directory.
