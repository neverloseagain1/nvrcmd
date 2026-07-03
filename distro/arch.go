package distro

func GetArchCommands() []CommandItem {
	return []CommandItem{
		// --- СИСТЕМА И ОБНОВЛЕНИЯ ---
		{"sudo pacman -Syu", "", "", "Обновить всю систему и все установленные пакеты", "Synchronize and update all system packages", "sudo pacman -Syu", "sudo pacman -Syu"},
		{"sudo pacman -S", "[имя]", "[name]", "Установить новый пакет из официальных репозиториев", "Install a new package from official repositories", "sudo pacman -S firefox", "sudo pacman -S firefox"},
		{"sudo pacman -Rns", "[имя]", "[name]", "Полностью удалить пакет вместе с его зависимостями и конфигами", "Remove a package along with its dependencies and configs", "sudo pacman -Rns vlc", "sudo pacman -Rns vlc"},
		{"sudo pacman -Rns $(pacman -Qdtq)", "", "", "Удалить все пакеты-сироты (остаточный мусор) из системы", "Remove all orphan packages from the system", "sudo pacman -Rns $(pacman -Qdtq)", "sudo pacman -Rns $(pacman -Qdtq)"},
		{"sudo journalctl -p 3 -xb", "", "", "Показать критические ошибки текущей загрузки системы (при сбоях)", "Show errors for the current system boot (important on crashes)", "sudo journalctl -p 3 -xb", "sudo journalctl -p 3 -xb"},
		
		// --- РЕПОЗИТОРИЙ MULTILIB (32-bit / Steam) ---
		{"sudo sed -i '/\\[multilib\\]/,+1 s/^#//' /etc/pacman.conf && sudo pacman -Syu", "", "", "Включить репозиторий multilib для запуска Steam и 32-битных приложений", "Enable multilib repository to run Steam and 32-bit applications", "sudo sed -i '/\\[multilib\\]/,+1 s/^#//' /etc/pacman.conf && sudo pacman -Syu", "sudo sed -i '/\\[multilib\\]/,+1 s/^#//' /etc/pacman.conf && sudo pacman -Syu"},
		{"sudo pacman -S steam", "", "", "Установить Steam (требуется предварительное включение multilib)", "Install Steam (requires multilib enabled first)", "sudo pacman -S steam", "sudo pacman -S steam"},

		// --- FLATPAK (Управление приложениями) ---
		{"sudo pacman -S flatpak", "", "", "Установить менеджер пакетов Flatpak для простых графических программ", "Install Flatpak package manager for easy GUI apps", "sudo pacman -S flatpak", "sudo pacman -S flatpak"},
		{"flatpak remote-add --if-not-exists flathub https://flathub.org", "", "", "Подключить главный репозиторий графических приложений Flathub", "Add the main Flathub repository", "flatpak remote-add --if-not-exists flathub...", "flatpak remote-add --if-not-exists flathub..."},
		{"flatpak search", "[имя]", "[name]", "Найти графическое приложение в репозитории Flathub", "Search for a graphical application in Flathub", "flatpak search spotify", "flatpak search spotify"},
		{"flatpak install flathub", "[имя]", "[name]", "Установить графическое приложение из Flathub", "Install an application from Flathub", "flatpak install flathub com.spotify.Client", "flatpak install flathub com.spotify.Client"},
		{"flatpak update", "", "", "Обновить все установленные Flatpak-приложения", "Update all installed Flatpak applications", "flatpak update", "flatpak update"},
		{"flatpak uninstall", "[имя]", "[name]", "Удалить Flatpak-приложение из системы", "Uninstall a Flatpak application", "flatpak uninstall com.spotify.Client", "flatpak uninstall com.spotify.Client"},

		// --- ДРАЙВЕРЫ И ГРАФИКА ---
		{"sudo pacman -S nvidia nvidia-utils", "", "", "Установить официальные проприетарные драйверы для видеокарт NVIDIA", "Install official proprietary NVIDIA drivers", "sudo pacman -S nvidia nvidia-utils", "sudo pacman -S nvidia nvidia-utils"},
		{"sudo pacman -S mesa lib32-mesa", "", "", "Установить открытые драйверы OpenGL/Vulkan для AMD и Intel (нужен multilib)", "Install open-source OpenGL/Vulkan drivers for AMD and Intel", "sudo pacman -S mesa lib32-mesa", "sudo pacman -S mesa lib32-mesa"},

		// --- РАБОТА С AUR (ДЛЯ YAY) ---
		{"yay -Syu", "", "", "Обновить всю систему и все пакеты из AUR (через хелпер yay)", "Update system and all AUR packages (via yay helper)", "yay -Syu", "yay -Syu"},

		// --- КАСТОМИЗАЦИЯ И ИНТЕРФЕЙС ---
		{"sudo pacman -S waybar", "", "", "Установить Waybar — настраиваемую панель для Wayland", "Install Waybar — highly customizable status bar for Wayland", "sudo pacman -S waybar", "sudo pacman -S waybar"},
		{"yay -S quickshell", "", "", "Установить Quickshell — создание виджетов на QML (Если нет: yay -S quickshell)", "Install Quickshell — widgetizer via QML (If missing: yay -S quickshell)", "yay -S quickshell", "yay -S quickshell"},
		{"yay -S eww", "", "", "Установить Eww — мощный движок виджетов (Если нет: yay -S eww)", "Install Eww — standalone widgetizer (If missing: yay -S eww)", "yay -S eww", "yay -S eww"},
		{"sudo pacman -S rofi-wayland", "", "", "Установить Rofi — удобное меню запуска приложений и окон", "Install Rofi — application launcher and window switcher", "sudo pacman -S rofi-wayland", "sudo pacman -S rofi-wayland"},
		{"sudo pacman -S hyprland", "", "", "Установить Hyprland — красивый динамический оконный менеджер", "Install Hyprland — beautiful dynamic tiling window manager", "sudo pacman -S hyprland", "sudo pacman -S hyprland"},

		// --- РАЗВЛЕЧЕНИЯ И ЭСТЕТИКА ---
		{"cava", "", "", "Запустить эквалайзер звука под музыку (Если нет: sudo pacman -S cava)", "Run audio visualizer (If missing: sudo pacman -S cava)", "cava", "cava"},
		{"pipes.sh", "", "", "Запустить заставку с растущими трубами (Если нет: yay -S pipes.sh)", "Run animated pipe screensaver (If missing: yay -S pipes.sh)", "pipes.sh", "pipes.sh"},
		{"cmatrix", "", "", "Запустить падающий цифровой дождь (Если нет: sudo pacman -S cmatrix)", "Run falling green digital rain (If missing: sudo pacman -S cmatrix)", "cmatrix", "cmatrix"},
		{"sl", "", "", "Запустить паровоз на случай опечатки в команде 'ls' (Если нет: yay -S sl)", "Run steam locomotive if you mistype 'ls' (If missing: yay -S sl)", "sl", "sl"},

		// --- ЖЕЛЕЗО, СТАТУС И ДИСКИ ---
		{"free -h", "", "", "Показать количество свободной и занятой оперативной памяти (RAM)", "Show free and used amount of RAM in human-readable format", "free -h", "free -h"},
		{"lsblk", "", "", "Показать структуру всех дисков, их разделов, размеры и точки монтирования", "List all block devices, partitions, sizes and mount points", "lsblk", "lsblk"},
		{"df -h", "", "", "Показать список дисков и сколько на них осталось свободного места", "Show disk space usage in human-readable format", "df -h", "df -h"},
		{"nvidia-smi", "", "", "Показать статус видеокарты NVIDIA, её температуру и видеопамять", "Show NVIDIA GPU status, temperature, usage and VRAM", "nvidia-smi", "nvidia-smi"},
		{"glxinfo | grep 'OpenGL renderer'", "", "", "Быстро узнать, какая именно видеокарта сейчас активна в системе", "Quickly check which graphics card is currently active", "glxinfo | grep 'OpenGL renderer'", "glxinfo | grep 'OpenGL renderer'"},
		{"uptime", "", "", "Показать время аптайма компьютера и среднюю нагрузку системы", "Show how long the system has been running and load average", "uptime", "uptime"},

		// --- ЗВУК, BLUETOOTH И НОУТБУКИ ---
		{"pavucontrol", "", "", "Открыть панель управления звуком (Если нет: sudo pacman -S pavucontrol)", "Open volume control panel (If missing: sudo pacman -S pavucontrol)", "pavucontrol", "pavucontrol"},
		{"sudo systemctl enable --now bluetooth", "", "", "Включить и запустить службу Bluetooth для наушников/мышек", "Enable and start Bluetooth service for devices", "sudo systemctl enable --now bluetooth", "sudo systemctl enable --now bluetooth"},
		{"blueman-manager", "", "", "Открыть менеджер Bluetooth устройств (Если нет: sudo pacman -S blueman)", "Open Bluetooth devices manager (If missing: sudo pacman -S blueman)", "blueman-manager", "blueman-manager"},
		{"brightnessctl set", "[значение]", "[value]", "Изменить яркость экрана ноутбука (Если нет: sudo pacman -S brightnessctl)", "Change laptop screen brightness (If missing: sudo pacman -S brightnessctl)", "brightnessctl set 50%", "brightnessctl set 50%"},
		{"nmtui", "", "", "Открыть консольный интерфейс для удобной настройки Wi-Fi и сети", "Open text user interface for Wi-Fi and network setup", "nmtui", "nmtui"},

		// --- МОНИТОРИНГ И ЧИСТКА ЖЕЛЕЗА ---
		{"btop", "", "", "Запустить красивый понятный диспетчер задач (Если нет: sudo pacman -S btop)", "Run task manager (If missing: sudo pacman -S btop)", "btop", "btop"},
		{"ncdu", "", "", "Запустить консольный сканер диска (Если нет: sudo pacman -S ncdu)", "Run disk usage analyzer (If missing: sudo pacman -S ncdu)", "ncdu", "ncdu"},
		{"sensors", "", "", "Показать температуры процессора и скорость кулеров (Если нет: sudo pacman -S lm_sensors)", "Show current CPU temperatures and fan speeds (If missing: sudo pacman -S lm_sensors)", "sensors", "sensors"},

		// --- СКОРАЯ ПОМОЩЬ И СВЕЖИЕ БЫТОВЫЕ КОМАНДЫ (Новое!) ---
		{"killall -9", "[имя_программы]", "[app_name]", "Принудительно закрыть «намертво» зависшую программу или игру", "Force kill a completely frozen application or game", "killall -9 discord", "killall -9 discord"},
		{"sudo umount /dev/", "[раздел]", "[partition]", "Безопасно извлечь флешку или внешний диск перед вытыканием", "Safely unmount a USB flash drive or external disk", "sudo umount /dev/sdb1", "sudo umount /dev/sdb1"},
		{"wget", "[ссылка]", "[url]", "Быстро скачать любой файл из интернета по прямой ссылке через консоль", "Download any file from the internet via direct link", "wget https://example.com", "wget https://example.com"},
		{"sudo shutdown -h", "[минуты]", "[minutes]", "Запланировать безопасное выключение компьютера через время", "Schedule a safe system shutdown after a specified time", "sudo shutdown -h +60", "sudo shutdown -h +60"},
		{"tar -xvf", "[архив.tar.gz]", "[archive.tar.gz]", "Распаковать скачанный архив формата .tar.gz", "Extract a downloaded .tar.gz archive", "tar -xvf theme.tar.gz", "tar -xvf theme.tar.gz"},
	}
}
