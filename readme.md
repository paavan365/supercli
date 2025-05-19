# supercli

![Go](https://img.shields.io/badge/Go-1.24.2-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Status](https://img.shields.io/badge/status-active-brightgreen.svg)

> **Automate Infra login and Kubernetes log fetching with ease!**

---

## 🚀 Overview

**supercli** is a command-line tool written in Go that streamlines your workflow for logging into Infra and fetching Kubernetes logs. With interactive prompts and seamless integration, it saves you time and reduces manual errors.

---

## ✨ Features

- **Infra Login Automation:** Securely log in to Infra using credentials from your `.env` file.
- **Cluster Selection:** Interactively select your target cluster.
- **Pod Discovery:** List pods and their labels in your Kubernetes cluster.
- **Log Streaming:** Fetch and stream logs for selected app labels.
- **User-Friendly Prompts:** Powered by [promptui](https://github.com/manifoldco/promptui) for a smooth CLI experience.
- **Environment Management:** Uses [godotenv](https://github.com/joho/godotenv) for easy environment variable loading.

---

## 📦 Project Structure

```
supercli/
│
├── .env                # Environment variables (INFRA_USER, INFRA_PASSWORD)
├── main.go             # Entry point
├── go.mod / go.sum     # Go modules
├── cmd/
│   ├── init.go         # Core CLI logic (login, pod listing, logs)
│   └── root.go         # Cobra root command setup
└── README.md           # You're here!
```

---

## 🛠️ Prerequisites

Before using **supercli**, you must have the following tools installed and available in your system's `PATH`:

- [Infra CLI](https://github.com/infrahq/infra)  
  Follow the [Infra installation guide](https://github.com/infrahq/infra#installation) to download and install the `infra` CLI.

- [kubectl](https://kubernetes.io/docs/reference/kubectl/)  
  Follow the [kubectl installation guide](https://kubernetes.io/docs/tasks/tools/) to download and install the Kubernetes CLI.

### Add to PATH

After downloading, move the binaries to a directory in your `PATH` (e.g., `/usr/local/bin`):

```sh
sudo mv /path/to/infra /usr/local/bin/
sudo mv /path/to/kubectl /usr/local/bin/
```

Verify installation:

```sh
infra version
kubectl version --client
```

---

## 🛠️ Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/supercli.git
   cd supercli
   ```

2. **Set up your `.env` file:**
   ```
   INFRA_USER=your_infra_username
   INFRA_PASSWORD=your_infra_password
   ```

3. **Build the CLI:**
   ```sh
   go build -o supercli
   ```

---

## ⚡ Usage

### Login to Infra

```sh
./supercli login
```
- Loads credentials from `.env`
- Logs you into Infra

### Initialize and Fetch Logs

```sh
./supercli init
```
- Select your cluster interactively
- List pods and select an app label
- Stream logs for the selected label

---

## 🧩 Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [promptui](https://github.com/manifoldco/promptui) - Interactive prompts
- [godotenv](https://github.com/joho/godotenv) - Environment variable loader

---

## 🤝 Contributing

Contributions are welcome! Please open issues or submit pull requests for improvements and bug fixes.

---

## 📄 License

This project is licensed under the MIT License.

---

## 🙌 Author

Made with ❤️ by [paavan](https://github.com/paavan365)

---

> **supercli** — Automate, simplify, and supercharge your Infra & Kubernetes workflow!