# `cbtg` Codebase Tour Guide CLI 🧭

**Codebase Tour Guide (CBTG)** is a lightweight CLI tool built in Go that wraps a powerful Python backend to analyze and generate guided tours of your codebase.

[![Go Report Card](https://goreportcard.com/badge/github.com/mrdegbe/cbtg-cli)](https://goreportcard.com/report/github.com/mrdegbe/cbtg-cli)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Build](https://img.shields.io/badge/build-passing-brightgreen)]()
[![Platform](https://img.shields.io/badge/platform-cross--platform-lightgrey)]()
[![Made with Go](https://img.shields.io/badge/made%20with-Go-1f425f.svg)]()

## 🛠 Features

- `init`: Scan and analyze your project structure
- `tour`: Generate a high-level walkthrough of your codebase (in text or markdown)
- Lightweight Go CLI calling a Python-powered AI engine
- Works offline with fallback AI models

## 🚀 Installation

Clone the repo and build:

```bash
git clone https://github.com/yourname/cbtg-cli.git
cd cbtg-cli
go build -o cbtg
````

## 📦 Usage

```bash
# Initialize and scan a codebase
./cbtg init path/to/project

# Generate a tour in markdown or text
./cbtg tour path/to/project --format md
```

## ⚙ Requirements

* Go ≥ 1.19
* Python ≥ 3.10
* [cbtg-core](https://github.com/mrdegbe/cbtg-core) must be cloned and accessible

## 📁 Project Structure

```
cbtg-cli/
├── cmd/             # CLI commands
├── utils/           # Shared helpers
├── main.go          # CLI entrypoint
├── .editorconfig
├── .gitattributes
├── .gitignore
└── README.md
```

## 💡 Tip

Want global access?

```bash
go install .
```

Then run:

```bash
cbtg init .
cbtg tour . --format text
```

---

## 📜 License

MIT © \Raymond Degbe
