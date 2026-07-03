# NVRCMD Handbook 🚀

> [!NOTE]
> 🌐 **Language / Язык:**
> * 🇺🇸 **English** (Current)
> * 🇷🇺 [Русский](README_RU.md)

An interactive, multi-language terminal utility designed for beginners to quickly lookup and learn basic and advanced Linux commands. It automatically detects your Linux distribution and features a beautiful scrolling UI.

## ✨ Features

* 📦 **Distro Auto-Detection:** Automatically parses `/etc/os-release` to show commands tailored for your OS (Arch Linux, Ubuntu, Debian, Fedora, NixOS, Gentoo).
* 🔄 **On-the-fly Translation:** Press `l` to instantly switch between Russian and English interfaces.
* 💾 **Persistent Settings:** Remembers your chosen language by saving it to `~/.config/nvrcmd.json`.
* 🛡️ **Safe and Clean:** Opens in an alternative screen buffer (like `htop` or `vim`) and leaves no mess in your shell history upon exit.
* 💡 **Practical Examples:** Every command comes with a real-world usage example displayed right next to its description.

## 📦 Supported Distributions

NVRCMD automatically recognizes and provides tailored commands for the following Linux families:
* 🔵 **Arch Linux** (including Void Linux compatibility)
* 🟠 **Ubuntu / Debian / Linux Mint**
* 🟢 **Fedora Linux**
* ⚪ **NixOS** (with declarative syntax and `nix-shell` guides)
* 🟣 **Gentoo Linux** (with Portage `emerge` and USE-flags guides)


## 📥 Installation

You can easily compile and install the utility globally using the provided script.

```bash
# Clone your repository (or go to your project folder)
cd nvrcmd

# Run the interactive installation script
./install.sh
```

## ⌨️ Controls

* `↑ / ↓` or `j / k` — Scroll through the commands list
* `l` — Toggle interface language (RU / EN)
* `q` or `Ctrl + C` — Exit the handbook safely
