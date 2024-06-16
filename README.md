# 🛠️ CHX

## 📘 Introduction

CHX is a Go-based project providing command-line utilities for domain research, hashing, IP analysis, and other functions. The project includes various commands and integrates with different sources like VirusTotal and CrowdSec.

## 📑 Table of Contents

- [📘 Introduction](#📘-introduction)
- [⚙️ Installation](#⚙️-installation)
- [📝 Usage](#📝-usage)
- [✨ Features](#✨-features)
- [📦 Dependencies](#📦-dependencies)
- [🛠️ Configuration](#🛠️-configuration)
- [📚 Documentation](#📚-documentation)
- [💡 Examples](#💡-examples)
- [🔧 Troubleshooting](#🔧-troubleshooting)
- [👥 Contributors](#👥-contributors)
- [📜 License](#📜-license)

## ⚙️ Installation

To install CHX, ensure you have Go installed on your machine and run the following commands:

```sh
go install github.com/EpykLab/chx@latest
```

Prebuilt binaries are also available under releases.

## 📝 Usage

After installing the project, you can use the CHX command-line tool as follows:

```sh
chx <command> [options]
```

### Example Commands

- `chx domain <domain-name>`
- `chx hash <file-path>`
- `chx ip <ip-address>`

## ✨ Features

- 🌐 Domain analysis
- 🔑 File hashing
- 📡 IP address lookup
- 🔗 Integration with VirusTotal, CrowdSec, AlientVault, and AbuseIPDB
- ⚙️ Configurable via YAML
- 🖥️ Can take input as argument or stdin

## 📦 Dependencies

- 🐹 Go (version 1.18 or higher)
- 🌐 External APIs for certain functionalities (e.g., VirusTotal, CrowdSec)

## 🛠️ Configuration

Configuration is managed through a json file. The config file is stored in `$HOME/.config/chx/conf.json`.

To configure chx, use `chx config`.

### 🌐 Domain Analysis

```sh
chx domain example.com
```

### 🔑 File Hashing

```sh
chx hash /path/to/file
```

### 📡 IP Lookup

```sh
chx ip 8.8.8.8
```

## 🔧 Troubleshooting

If you encounter any issues, please check the following:

- Ensure you have the correct API keys in your configuration file.
- Verify your internet connection for API integrations.
- Refer to the logs for any error messages.

## 👥 Contributors

- [Your Name](https://github.com/DavidHoenisch)

## 📜 License
