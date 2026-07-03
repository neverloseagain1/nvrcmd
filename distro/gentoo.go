package distro

func GetGentooCommands() []CommandItem {
	return []CommandItem{
		// --- СИСТЕМА И ОБНОВЛЕНИЯ (Компиляция) ---
		{"sudo emaint sync -a", "", "", "Синхронизировать дерево пакетов Portage (обновить списки доступного софта)", "Synchronize Portage package tree (update available software lists)", "sudo emaint sync -a", "sudo emaint sync -a"},
		{"sudo emerge --ask --verbose", "[имя]", "[name]", "Скомпилировать и установить новый пакет с выводом подробностей", "Compile and install a new package with verbose output", "sudo emerge -av app-admin/htop", "sudo emerge -av app-admin/htop"},
		{"sudo emerge --ask --update --deep --newuse @world", "", "", "Полностью обновить всю систему с учётом изменённых USE-флагов", "Update the entire system considering changed USE-flags", "sudo emerge -auDN @world", "sudo emerge -auDN @world"},
		{"sudo emerge --ask --depclean", "", "", "Удалить пакеты-сироты и неиспользуемые зависимости (очистка мусора)", "Remove orphan packages and unused dependencies from the system", "sudo emerge -a --depclean", "sudo emerge -a --depclean"},
		{"sudo dispatch-conf", "", "", "Запустить утилиту обновления конфигурационных файлов после обновления пакетов", "Run configuration file update utility after package upgrades", "sudo dispatch-conf", "sudo dispatch-conf"},
		{"sudo nano /etc/portage/make.conf", "", "", "Открыть главный конфигурационный файл Gentoo (глобальные USE-флаги, зеркала)", "Open main Gentoo configuration file (global USE-flags)", "sudo nano /etc/portage/make.conf", "sudo nano /etc/portage/make.conf"},

		// --- УТИЛИТЫ ДЛЯ СИСТЕМЫ И ИНТЕРФЕЙСА ---
		{"sudo emerge --ask app-admin/btop", "", "", "Скомпилировать и установить продвинутый диспетчер задач btop", "Compile and install advanced btop task manager", "sudo emerge -a app-admin/btop", "sudo emerge -a app-admin/btop"},
		{"sudo emerge --ask sys-fs/ncdu", "", "", "Скомпилировать и установить консольный сканер занятого места ncdu", "Compile and install ncdu disk usage analyzer", "sudo emerge -a sys-fs/ncdu", "sudo emerge -a sys-fs/ncdu"},
		{"sudo emerge --ask app-misc/fastfetch", "", "", "Скомпилировать и установить утилиту вывода информации о железе fastfetch", "Compile and install fastfetch system info utility", "sudo emerge -a app-misc/fastfetch", "sudo emerge -a app-misc/fastfetch"},

		// --- КАСТОМИЗАЦИЯ, РАЗВЛЕЧЕНИЯ И ЭСТЕТИКА ---
		{"sudo emerge --ask gui-apps/waybar", "", "", "Скомпилировать и установить статус-бар Waybar для Wayland окружений", "Compile and install Waybar status bar for Wayland", "sudo emerge -a gui-apps/waybar", "sudo emerge -a gui-apps/waybar"},
		{"sudo emerge --ask gui-apps/wofi", "", "", "Скомпилировать и установить легкое меню запуска приложений Wofi", "Compile and install Wofi application launcher", "sudo emerge -a gui-apps/wofi", "sudo emerge -a gui-apps/wofi"},
		{"sudo emerge --ask media-sound/cava", "", "", "Скомпилировать анимированный эквалайзер звука cava (может потребовать флаг alsa/pulseaudio)", "Compile and install cava audio visualizer", "sudo emerge -a media-sound/cava", "sudo emerge -a media-sound/cava"},
		{"sudo emerge --ask app-misc/cmatrix", "", "", "Скомпилировать и установить падающий цифровой дождь cmatrix", "Compile and install cmatrix screensaver", "sudo emerge -a app-misc/cmatrix", "sudo emerge -a app-misc/cmatrix"},

		// --- FLATPAK (Управление готовыми приложениями) ---
		{"sudo emerge --ask sys-apps/flatpak", "", "", "Скомпилировать и установить менеджер пакетов Flatpak в Gentoo", "Compile and install Flatpak package manager into Gentoo", "sudo emerge -a sys-apps/flatpak", "sudo emerge -a sys-apps/flatpak"},
		{"flatpak remote-add --if-not-exists flathub https://flathub.org", "", "", "Подключить репозиторий Flathub для быстрой установки Discord/Spotify без компиляции", "Add the main Flathub repository", "flatpak remote-add --if-not-exists flathub...", "flatpak remote-add --if-not-exists flathub..."},
		{"flatpak install flathub", "[имя]", "[name]", "Установить графическое приложение из Flathub", "Install an application from Flathub", "flatpak install flathub com.spotify.Client", "flatpak install flathub com.spotify.Client"},

		// --- ЖЕЛЕЗО, СТАТУС И ДИСКИ ---
		{"free -h", "", "", "Показать количество свободной и занятой оперативной памяти (RAM)", "Show free and used amount of RAM in human-readable format", "free -h", "free -h"},
		{"lsblk", "", "", "Показать структуру всех дисков, их разделов, размеры и точки монтирования", "List all block devices, partitions, sizes and mount points", "lsblk", "lsblk"},
		{"df -h", "", "", "Показать список дисков и сколько на них осталось свободного места", "Show disk space usage in human-readable format", "df -h", "df -h"},
		{"nvidia-smi", "", "", "Показать статус видеокарты NVIDIA, её температуру и видеопамять (Требует x11-drivers/nvidia-drivers)", "Show NVIDIA GPU status, temperature and VRAM", "nvidia-smi", "nvidia-smi"},
		{"uptime", "", "", "Показать время аптайма компьютера и среднюю нагрузку системы", "Show how long the system has been running and load average", "uptime", "uptime"},

		// --- ЗВУК, BLUETOOTH И СЕТЬ ---
		{"sudo emerge --ask media-sound/pavucontrol", "", "", "Скомпилировать и установить панель управления звуком и микрофонами", "Compile and install GUI volume control panel", "sudo emerge -a media-sound/pavucontrol", "sudo emerge -a media-sound/pavucontrol"},
		{"nmtui", "", "", "Открыть консольный интерфейс для настройки Wi-Fi (Требует net-misc/networkmanager)", "Open text user interface for Wi-Fi and network setup", "nmtui", "nmtui"},
		{"brightnessctl set", "[значение]", "[value]", "Изменить яркость экрана ноутбука (Если нет: sudo emerge -a sys-power/brightnessctl)", "Change laptop screen brightness", "brightnessctl set 50%", "brightnessctl set 50%"},

		// --- СКОРАЯ ПОМОЩЬ И ПИТАНИЕ ---
		{"killall -9", "[имя_программы]", "[app_name]", "Принудительно закрыть «намертво» зависшую программу или игру", "Force kill a completely frozen application or game", "killall -9 discord", "killall -9 discord"},
		{"sudo shutdown -h", "[минуты]", "[minutes]", "Запланировать безопасное выключение компьютера через указанное время", "Schedule a safe system shutdown after a specified time", "sudo shutdown -h +60", "sudo shutdown -h +60"},
		{"tar -xvf", "[архив.tar.gz]", "[archive.tar.gz]", "Распаковать скачанный архив формата .tar.gz", "Extract a downloaded .tar.gz archive", "tar -xvf theme.tar.gz", "tar -xvf theme.tar.gz"},
	}
}
