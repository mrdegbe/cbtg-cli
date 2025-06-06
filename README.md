# CBTG (Codebase Tour Guide) CLI 🧭

**Codebase Tour Guide (CBTG)** is a lightweight CLI tool built in Go that wraps a powerful Python backend to analyze and generate guided tours of your codebase.

## 🛠 Features

- `init`: Scan and analyze your project structure
- `tour`: Generate a high-level walkthrough of your codebase (in text or markdown)
- Lightweight Go CLI calling a Python-powered AI engine
- Works offline with fallback AI models

## 🚀 Installation

Clone the repo and build:

```bash
git clone https://github.com/mrdegbe/cbtg-cli.git
cd cbtg-cli
go build -o cbtg
