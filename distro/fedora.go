package distro

func GetFedoraCommands() []CommandItem {
	return []CommandItem{
		// --- СИСТЕМА И ОБНОВЛЕНИЯ ---
		{"sudo dnf upgrade --refresh", "", "", "Обновить систему и обновить кэш всех репозиториев", "Upgrade system and refresh all repository caches", "sudo dnf upgrade --refresh", "sudo dnf upgrade --refresh"},
		{"sudo dnf install", "[имя]", "[name]", "Установить пакет из официальных репозиториев Fedora", "Install a package from official Fedora repositories", "sudo dnf install git", "sudo dnf install git"},
		{"sudo dnf remove", "[имя]", "[name]", "Удалить пакет из системы", "Remove a package from the system", "sudo dnf remove htop", "sudo dnf remove htop"},
		{"sudo dnf autoremove", "", "", "Удалить все неиспользуемые пакеты-сироты из системы", "Remove all unused orphan packages from the system", "sudo dnf autoremove", "sudo dnf autoremove"},
		{"sudo dnf clean all", "", "", "Полностью очистить весь кэш пакетов и метаданных DNF", "Completely clean all DNF package and metadata cache", "sudo dnf clean all", "sudo dnf clean all"},
		
		// --- УСКОРЕНИЕ DNF И RPM FUSION ---
		{"sudo sh -c 'echo \"max_parallel_downloads=10\" >> /etc/dnf/dnf.conf'", "", "", "Увеличить количество параллельных загрузок DNF до 10 (ускоряет установку)", "Set DNF max parallel downloads to 10 (speeds up installations)", "sudo sh -c 'echo \"max_parallel_downloads=10\" >> /etc/dnf/dnf.conf'", "sudo sh -c 'echo \"max_parallel_downloads=10\" >> /etc/dnf/dnf.conf'"},
		{"sudo dnf install https://rpmfusion.org(rpm -E %fedora).noarch.rpm https://rpmfusion.org(rpm -E %fedora).noarch.rpm", "", "", "Подключить репозитории RPM Fusion для кодеков и проприетарных драйверов", "Enable RPM Fusion repositories for codecs and proprietary drivers", "sudo dnf install https://rpmfusion.org...", "sudo dnf install https://rpmfusion.org..."},
		{"sudo dnf install akmod-nvidia", "", "", "Установить официальные драйверы NVIDIA (требуется RPM Fusion)", "Install official proprietary NVIDIA drivers (requires RPM Fusion)", "sudo dnf install akmod-nvidia", "sudo dnf install akmod-nvidia"},
		{"sudo dnf install lclint-libs gstreamer1-plugins-{{bad,ugly},free-nonfree} ffmpeg gstreamer1-libav --allowerasing", "", "", "Установить полный набор кодеков FFmpeg для воспроизведения любых видео", "Install full set of codecs and FFmpeg plugins for any video playback", "sudo dnf install lclint-libs...", "sudo dnf install lclint-libs..."},

		// --- FLATPAK ---
		{"flatpak search", "[имя]", "[name]", "Найти графическое приложение в репозитории Flathub", "Search for a graphical application in Flathub", "flatpak search spotify", "flatpak search spotify"},
		{"flatpak install flathub", "[имя]", "[name]", "Установить приложение из Flathub", "Install an application from Flathub", "flatpak install flathub com.spotify.Client", "flatpak install flathub com.spotify.Client"},
		{"flatpak update", "", "", "Обновить все установленные Flatpak-приложения", "Update all installed Flatpak applications", "flatpak update", "flatpak update"},

		// --- СИСТЕМНАЯ БЕЗОПАСНОСТЬ (SELinux) ---
		{"sudo setenforce 0", "", "", "Временно перевести SELinux в разрешающий режим (если блокируются приложения)", "Temporarily set SELinux to permissive mode (helps if applications are blocked)", "sudo setenforce 0", "sudo setenforce 0"},
		{"sestatus", "", "", "Проверить текущий статус и режим работы системы безопасности SELinux", "Check the current status and mode of SELinux security system", "sestatus", "sestatus"},

		// --- КАСТОМИЗАЦИЯ И ИНТЕРФЕЙС ---
		{"sudo dnf install waybar wofi fastfetch", "", "", "Установить статус-бар Waybar, меню запуска приложений Wofi и инфо-утилиту Fastfetch", "Install Waybar status bar, Wofi application launcher and Fastfetch info utility", "sudo dnf install waybar wofi fastfetch", "sudo dnf install waybar wofi fastfetch"},
		{"sudo dnf copr enable", "[автор/проект]", "[author/project]", "Включить сторонний репозиторий COPR (нужно для установки eww/quickshell)", "Enable a third-party COPR repository (needed for eww/quickshell)", "sudo dnf copr enable kylegospo/wayland", "sudo dnf copr enable kylegospo/wayland"},

		// --- РАЗВЛЕЧЕНИЯ И ЭСТЕТИКА ---
		{"cava", "", "", "Запустить эквалайзер звука под музыку (Если нет: sudo dnf install cava)", "Run audio visualizer (If missing: sudo dnf install cava)", "cava", "cava"},
		{"pipes.sh", "", "", "Запустить заставку с трубами (Если нет установки, см. метод с curl)", "Run pipe screensaver (If not installed, see curl install method)", "pipes.sh", "pipes.sh"},
		{"cmatrix", "", "", "Запустить падающий зеленый цифровой дождь (Если нет: sudo dnf install cmatrix)", "Run falling green digital rain (If missing: sudo dnf install cmatrix)", "cmatrix", "cmatrix"},
		{"sl", "", "", "Запустить поезд при опечатке в 'ls' (Если нет: sudo dnf install sl)", "Run steam locomotive if you mistype 'ls' (If missing: sudo dnf install sl)", "sl", "sl"},

		// --- ЖЕЛЕЗО, СТАТУС И ДИСКИ ---
		{"free -h", "", "", "Показать количество свободной и занятой оперативной памяти (RAM)", "Show free and used amount of RAM in human-readable format", "free -h", "free -h"},
		{"lsblk", "", "", "Показать структуру всех дисков, их разделов, размеры и точки монтирования", "List all block devices, partitions, sizes and mount points", "lsblk", "lsblk"},
		{"df -h", "", "", "Показать список дисков и сколько на них осталось свободного места", "Show disk space usage in human-readable format", "df -h", "df -h"},
		{"nvidia-smi", "", "", "Показать статус видеокарты NVIDIA, её температуру, загрузку (GPU) и видеопамять", "Show NVIDIA GPU status, temperature, usage and VRAM", "nvidia-smi", "nvidia-smi"},
		{"glxinfo | grep 'OpenGL renderer'", "", "", "Быстро узнать, какая именно видеокарта сейчас active в системе", "Quickly check which graphics card is currently active", "glxinfo | grep 'OpenGL renderer'", "glxinfo | grep 'OpenGL renderer'"},

		// --- ЗВУК, BLUETOOTH И НОУТБУКИ ---
		{"pavucontrol", "", "", "Открыть панель управления звуком (Если нет: sudo dnf install pavucontrol)", "Open volume control panel (If missing: sudo dnf install pavucontrol)", "pavucontrol", "pavucontrol"},
		{"blueman-manager", "", "", "Открыть менеджер Bluetooth устройств (Если нет: sudo dnf install blueman)", "Open Bluetooth devices manager (If missing: sudo dnf install blueman)", "blueman-manager", "blueman-manager"},
		{"brightnessctl set", "[значение]", "[value]", "Изменить яркость экрана ноутбука (Если нет: sudo dnf install brightnessctl)", "Change laptop screen brightness (If missing: sudo dnf install brightnessctl)", "brightnessctl set 50%", "brightnessctl set 50%"},
		{"nmtui", "", "", "Открыть консольный интерфейс для настройки Wi-Fi и сети", "Open text user interface for Wi-Fi and network setup", "nmtui", "nmtui"},

		// --- МОНИТОРИНГ И ЧИСТКА ЖЕЛЕЗА ---
		{"btop", "", "", "Запустить диспетчер задач btop (Если нет: sudo dnf install btop)", "Run task manager (If missing: sudo dnf install btop)", "btop", "btop"},
		{"ncdu", "", "", "Запустить консольный сканер диска (Если нет: sudo dnf install ncdu)", "Run disk usage analyzer (If missing: sudo dnf install ncdu)", "ncdu", "ncdu"},
		{"sensors", "", "", "Показать температуры процессора и кулеров (Если нет: sudo dnf install lm_sensors)", "Show current CPU temperatures (If missing: sudo dnf install lm_sensors)", "sensors", "sensors"},

		// --- СВЕЖИЕ БЫТОВЫЕ КОМАНДЫ (Новое!) ---
		{"killall -9", "[имя_программы]", "[app_name]", "Принудительно закрыть «намертво» зависшую программу или игру", "Force kill a completely frozen application or game", "killall -9 discord", "killall -9 discord"},
		{"sudo umount /dev/", "[раздел]", "[partition]", "Безопасно извлечь флешку или внешний диск перед вытыканием", "Safely unmount a USB flash drive or external disk", "sudo umount /dev/sdb1", "sudo umount /dev/sdb1"},
		{"wget", "[ссылка]", "[url]", "Быстро скачать любой файл из интернета по прямой ссылке через консоль", "Download any file from the internet via direct link", "wget https://example.com", "wget https://example.com"},
		{"sudo shutdown -h", "[минуты]", "[minutes]", "Запланировать безопасное выключение компьютера через указанное время", "Schedule a safe system shutdown after a specified time", "sudo shutdown -h +60", "sudo shutdown -h +60"},
		{"tar -xvf", "[архив.tar.gz]", "[archive.tar.gz]", "Распаковать скачанный из интернета архив формата .tar.gz", "Extract a downloaded .tar.gz archive", "tar -xvf theme.tar.gz", "tar -xvf theme.tar.gz"},
	}
}
