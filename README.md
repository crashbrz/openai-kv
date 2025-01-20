![License](https://img.shields.io/badge/license-sushiware-red)
![Issues open](https://img.shields.io/github/issues/crashbrz/openai-kv)
![GitHub pull requests](https://img.shields.io/github/issues-pr-raw/crashbrz/openai-kv)
![GitHub closed issues](https://img.shields.io/github/issues-closed-raw/crashbrz/openai-kv)
![GitHub last commit](https://img.shields.io/github/last-commit/crashbrz/openai-kv)

# OpenAI API Key Validator

## Overview
This Go program validates OpenAI API keys either individually or in bulk using a multithreaded approach. It supports processing multiple keys concurrently and provides counts for valid and invalid keys.

---

## Features
- **Single Key Validation:** Use the `-k` flag to validate a single API key.
- **Bulk Validation:** Use the `-f` flag to specify a file containing API keys, one per line.
- **Debug Mode:** Use the `-d` flag to display invalid keys along with valid ones.
- **Multithreaded Processing:** Use the `-t` flag to specify the number of goroutines for concurrent validation.

---

## Flags
| Flag | Description |
|------|-------------|
| `-k` | Validates a single OpenAI API key. |
| `-f` | Specifies the file path to a list of API keys. |
| `-d` | Debug mode: displays both valid and invalid keys. |
| `-t` | Number of goroutines to use for multithreaded validation. Default is `1`. |

---

## Usage

### Single Key Validation
```bash
./openai-kv -k <API_KEY>
```
Example:
```bash
./openai-kv -k sk-1234567890abcdef
```

### Bulk Validation
```bash
./openai-kv -f <FILE_PATH>
```
Example:
```bash
./openaichecker -f keys.txt
```

### Debug Mode
```bash
.//openai-kv -f <FILE_PATH> -d
```
Example:
```bash
.//openai-kv -f keys.txt -d
```

### Multithreaded Validation
```bash
./openaichecker -f <FILE_PATH> -t <NUM_THREADS>
```
Example:
```bash
.//openai-kv -f keys.txt -t 5
```

---

## Output
- **Valid Keys:** Shown in green with a count at the end.
- **Invalid Keys:** (if `-d` is used) Shown in red with a count at the end.

Example Output:
```plaintext
sk-validkey123456 is valid.
sk-invalidkey123456 is invalid.
Valid keys: 1
Invalid keys: 1
```

## Building the Program
```bash
go build -o /openai-kv
```

---

## Cloning the Repository
```bash
git clone <repository-url>
cd <repository-name>
```

---

## License
### License ###
openai-kv is licensed under the SushiWare license. For more information, check [docs/license.txt](docs/license.txt).
