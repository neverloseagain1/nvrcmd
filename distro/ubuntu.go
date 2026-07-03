package distro

func GetUbuntuCommands() []CommandItem {
	return []CommandItem{
		// --- СИСТЕМА И ОБНОВЛЕНИЯ ---
		{"sudo apt update", "", "", "Обновить списки пакетов из подключенных репозиториев", "Update package lists from repositories", "sudo apt update", "sudo apt update"},
		{"sudo apt upgrade", "", "", "Обновить все установленные пакеты до актуальных версий", "Upgrade all installed packages to current versions", "sudo apt upgrade", "sudo apt upgrade"},
		{"sudo apt install", "[имя]", "[name]", "Установить новый пакет из официальных репозиториев", "Install a new package from official repositories", "sudo apt install vlc", "sudo apt install vlc"},
		{"sudo apt autoremove", "", "", "Удалить ненужные пакеты и старые ядра, оставшиеся после удаления программ", "Remove unneeded orphan packages and old system kernels", "sudo apt autoremove", "sudo apt autoremove"},
		{"sudo apt clean", "", "", "Полностью очистить локальный кэш скачанных файлов .deb (освобождает место)", "Completely clear local cache of downloaded .deb files", "sudo apt clean", "sudo apt clean"},
		
		// --- СКОРАЯ ПОМОЩЬ (Поломки APT) ---
		{"sudo dpkg --configure -a", "", "", "Починить базу данных пакетов, если обновление было прервано", "Fix broken package database if update was interrupted", "sudo dpkg --configure -a", "sudo dpkg --configure -a"},
		{"sudo apt install -f", "", "", "Найти и автоматически исправить сломанные зависимости пакетов", "Find and automatically fix broken package dependencies", "sudo apt install -f", "sudo apt install -f"},

		// --- РАБОТА С PPA (СТОРОННИЕ РЕПОЗИТОРИИ) ---
		{"sudo add-apt-repository", "[ppa:имя]", "[ppa:name]", "Добавить сторонний PPA-репозиторий для установки свежего софта", "Add a third-party PPA repository for newer software", "sudo add-apt-repository ppa:obsproject/obs-studio", "sudo add-apt-repository ppa:obsproject/obs-studio"},

		// --- МЕНЕДЖЕР ПАКЕТОВ SNAP ---
		{"snap list", "", "", "Показать список всех установленных пакетов формата Snap", "List all installed Snap packages", "snap list", "snap list"},
		{"sudo snap install", "[имя]", "[name]", "Установить обычное графическое приложение (Discord, Telegram) из Snap", "Install a GUI application from the Snap Store", "sudo snap install discord", "sudo snap install discord"},
		{"sudo snap refresh", "", "", "Обновить все установленные Snap-пакеты до последних версий", "Update all installed Snap packages to the latest versions", "sudo snap refresh", "sudo snap refresh"},

		// --- FLATPAK ---
		{"flatpak search", "[имя]", "[name]", "Найти графическое приложение в репозитории Flathub", "Search for a graphical application in Flathub", "flatpak search spotify", "flatpak search spotify"},
		{"flatpak install flathub", "[имя]", "[name]", "Установить приложение из Flathub", "Install an application from Flathub", "flatpak install flathub com.spotify.Client", "flatpak install flathub com.spotify.Client"},
		{"flatpak update", "", "", "Обновить все установленные Flatpak-приложения", "Update all installed Flatpak applications", "flatpak update", "flatpak update"},

		// --- КАСТОМИЗАЦИЯ И ИНТЕРФЕЙС ---
		{"sudo apt install waybar", "", "", "Установить Waybar — статус-бар для современных графических окружений Wayland", "Install Waybar status bar for Wayland environments", "sudo apt install waybar", "sudo apt install waybar"},
		{"sudo apt install wofi", "", "", "Установить Wofi — легкое меню запуска приложений для Wayland (аналог Rofi)", "Install Wofi — application launcher for Wayland (Rofi analogue)", "sudo apt install wofi", "sudo apt install wofi"},

		// --- РАЗВЛЕЧЕНИЯ И ЭСТЕТИКА ---
		{"cava", "", "", "Запустить эквалайзер звука под музыку (Если нет: sudo apt install cava)", "Run audio visualizer (If missing: sudo apt install cava)", "cava", "cava"},
		{"pipes.sh", "", "", "Запустить заставку с растущими трубами (Если нет: sudo apt install pipes.sh)", "Run animated pipe screensaver (If missing: sudo apt install pipes.sh)", "pipes.sh", "pipes.sh"},
		{"cmatrix", "", "", "Запустить падающий цифровой дождь (Если нет: sudo apt install cmatrix)", "Run falling green digital rain (If missing: sudo apt install cmatrix)", "cmatrix", "cmatrix"},
		{"sl", "", "", "Запустить поезд при опечатке в 'ls' (Если нет: sudo apt install sl)", "Run steam locomotive if you mistype 'ls' (If missing: sudo apt install sl)", "sl", "sl"},

		// --- ЖЕЛЕЗО, СТАТУС И ДИСКИ ---
		{"free -h", "", "", "Показать количество свободной и занятой оперативной памяти (RAM)", "Show free and used amount of RAM in human-readable format", "free -h", "free -h"},
		{"lsblk", "", "", "Показать структуру всех дисков, их разделов, размеры и точки монтирования", "List all block devices, partitions, sizes and mount points", "lsblk", "lsblk"},
		{"df -h", "", "", "Показать список дисков и сколько на них осталось свободного места", "Show disk space usage in human-readable format", "df -h", "df -h"},
		{"nvidia-smi", "", "", "Показать статус видеокарты NVIDIA, её температуру/загрузку (GPU) и видеопамять", "Show NVIDIA GPU status, temperature, usage and VRAM", "nvidia-smi", "nvidia-smi"},
		{"glxinfo | grep 'OpenGL renderer'", "", "", "Быстро узнать, какая именно видеокарта сейчас активна в системе", "Quickly check which graphics card is currently active", "glxinfo | grep 'OpenGL renderer'", "glxinfo | grep 'OpenGL renderer'"},

		// --- ЗВУК, BLUETOOTH И НОУТБУКИ ---
		{"pavucontrol", "", "", "Открыть панель управления звуком (Если нет: sudo apt install pavucontrol)", "Open volume control panel (If missing: sudo apt install pavucontrol)", "pavucontrol", "pavucontrol"},
		{"blueman-manager", "", "", "Открыть менеджер Bluetooth устройств (Если нет: sudo apt install blueman)", "Open Bluetooth devices manager (If missing: sudo apt install blueman)", "blueman-manager", "blueman-manager"},
		{"brightnessctl set", "[значение]", "[value]", "Изменить яркость экрана ноутбука (Если нет: sudo apt install brightnessctl)", "Change laptop screen brightness (If missing: sudo apt install brightnessctl)", "brightnessctl set 50%", "brightnessctl set 50%"},
		{"nmtui", "", "", "Открыть консольный интерфейс для удобной настройки Wi-Fi и сети", "Open text user interface for Wi-Fi and network setup", "nmtui", "nmtui"},

		// --- МОНИТОРИНГ И ЧИСТКА ЖЕЛЕЗА ---
		{"btop", "", "", "Запустить диспетчер задач btop (Если нет: sudo apt install btop)", "Run task manager (If missing: sudo apt install btop)", "btop", "btop"},
		{"ncdu", "", "", "Запустить консольный сканер диска (Если нет: sudo apt install ncdu)", "Run disk usage analyzer (If missing: sudo apt install ncdu)", "ncdu", "ncdu"},
		{"sensors", "", "", "Показать температуры процессора и кулеров (Если нет: sudo apt install lm-sensors)", "Show current CPU temperatures (If missing: sudo apt install lm-sensors)", "sensors", "sensors"},

		// --- СВЕЖИЕ БЫТОВЫЕ КОМАНДЫ (Новое!) ---
		{"killall -9", "[имя_программы]", "[app_name]", "Принудительно закрыть «намертво» зависшую программу или игру", "Force kill a completely frozen application or game", "killall -9 discord", "killall -9 discord"},
		{"sudo umount /dev/", "[раздел]", "[partition]", "Безопасно извлечь флешку или внешний диск перед вытыканием", "Safely unmount a USB flash drive or external disk", "sudo umount /dev/sdb1", "sudo umount /dev/sdb1"},
		{"wget", "[ссылка]", "[url]", "Быстро скачать любой файл из интернета по прямой ссылке через консоль", "Download any file from the internet via direct link", "wget https://example.com", "wget https://example.com"},
		{"sudo shutdown -h", "[минуты]", "[minutes]", "Запланировать безопасное выключение компьютера через указанное время", "Schedule a safe system shutdown after a specified time", "sudo shutdown -h +60", "sudo shutdown -h +60"},
		{"tar -xvf", "[архив.tar.gz]", "[archive.tar.gz]", "Распаковать скачанный из интернета архив формата .tar.gz", "Extract a downloaded .tar.gz archive", "tar -xvf theme.tar.gz", "tar -xvf theme.tar.gz"},
	}
}
