# 🪟 WEnum – Windows Enumeration with Golang

**WEnum** is a lightweight Windows enumeration tool written in Go, designed for offline or post-exploitation use during red team, OSCP, or CTF engagements.

---

## 📄 Overview

WEnum executes a customizable list of enumeration commands defined in a `file.txt` file, making it highly flexible and adaptable to different assessment scenarios.

---

## ⚙️ How It Works

- Reads commands line-by-line from a file named `file.txt`
- Executes each command and outputs the result
- Useful for rapid recon on compromised systems or lab machines

---

## 📝 Usage

### 1. Prepare Your Command File

Create or edit a `file.txt` in the same directory as `wenum.exe`.  
Add any Windows command-line enumeration commands you want to run.

**Example `file.txt`:**
```
whoami
ipconfig /all
net users
systeminfo
```

### 2. Run the Tool

To execute WEnum and view results in the terminal:

```bash
wenum.exe
```

To redirect output to a file:

```bash
wenum.exe > results.txt
```

---

## 🛠️ Compilation

If you prefer to build the binary yourself:

```bash
go build -o wenum.exe main.go
```

Requires [Go](https://golang.org/) installed on your system.

---

## 📦 Download

- Precompiled `wenum.exe` is available in the [Releases](../../releases) section.
- Place `wenum.exe` and `file.txt` in the same folder before running.

---

## ⚠️ Legal Notice

This tool is intended for **educational use and authorized testing only**.  
Do **not** run on systems without **explicit permission**.

---
