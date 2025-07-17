# ntptest

A simple cross-platform CLI tool for NTP testing.

## 📦 Installation

Prebuilt binaries are available for Linux, macOS, and Windows under the [Releases](https://github.com/Whats-A-MattR/ntptest/releases) tab.

### 🖥️ Download and Install

#### Linux/macOS

```bash
curl -LO https://github.com/Whats-A-MattR/ntptest/releases/latest/download/ntptest-linux-amd64
chmod +x ntptest-linux-amd64
sudo mv ntptest-linux-amd64 /usr/local/bin/ntptest
```

#### Windows

1. Download `ntptest-windows-amd64.exe` from the [Releases](https://github.com/Whats-A-MattR/ntptest/releases) page.
2. (Optional) Rename it to `ntptest.exe`
3. Move it to a folder like `C:\ntptest\`
4. Add that folder to your system `Path`:
   - Control Panel → System → Advanced → Environment Variables → System Variables → `Path`

Now you can run `ntptest` from any terminal window.

---

## 🚀 Usage

```bash
ntptest --help
```

### Test against a specific NTP server

```bash
ntptest -server time.google.com
```

---

## 🛠 Building from Source

If you have Go installed:

```bash
git clone https://github.com/Whats-A-MattR/ntptest.git
cd ntptest
go build -o ntptest main.go
```

---

## 📄 License

MIT
