# Setup Rust Env on MacOs

## [Rust Language](https://doc.rust-lang.org/book/title-page.html)

## [Install Rust](https://www.rust-lang.org/tools/install)

1. Installation: curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
2. Installation includes rustc, cargo, rustup </p>
   cargo - like maven or gradle in java and go mod in golang</p>
   rustc - rust compiler</p>
   rustup - toolchain multiplexer to enable you to switch compiler versions</p>
3. Verify: rustc --version
4. Upgrade Rust: rustup update
4. Uninstall Rust: rustup self uninstall

## Switching Rust toolchains

Show what toolchains installed on the machine

```bash
rustup --help
rustup show
```

default - Set the default toolchain

override - Modify directory toolchain overrides

```bash
rustup override set nightly/stable
```

```bash
rustup default nightly/stable/version
```
