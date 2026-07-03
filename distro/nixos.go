package distro

func GetNixosCommands() []CommandItem {
	return []CommandItem{
		// --- СИСТЕМА И ОБНОВЛЕНИЯ (Декларативный стиль) ---
		{"sudo nano /etc/nixos/configuration.nix", "", "", "Открыть главный конфигурационный файл системы для кастомизации и добавления программ", "Open main system configuration file to add packages", "sudo nano /etc/nixos/configuration.nix", "sudo nano /etc/nixos/configuration.nix"},
		{"sudo nixos-rebuild switch", "", "", "Применить изменения из конфигурационного файла (создать новое поколение системы)", "Apply changes from configuration file and rebuild system", "sudo nixos-rebuild switch", "sudo nixos-rebuild switch"},
		{"sudo nixos-rebuild switch --upgrade", "", "", "Обновить все каналы пакетов NixOS и пересобрать систему до актуальных версий", "Upgrade NixOS channels and rebuild the system", "sudo nixos-rebuild switch --upgrade", "sudo nixos-rebuild switch --upgrade"},
		{"nix-env -qaP", "[запрос]", "[query]", "Найти точное имя пакета в репозиториях Nix для вставки в конфиг", "Search for package name in Nix repositories for config", "nix-env -qaP btop", "nix-env -qaP btop"},
		
		// --- ВРЕМЕННЫЙ ЗАПУСК БЕЗ УСТАНОВКИ (Фишка NixOS) ---
		{"nix-shell -p btop", "", "", "Временно запустить диспетчер задач btop без прописывания в конфиг системы", "Temporarily run btop task manager without adding to configuration", "nix-shell -p btop", "nix-shell -p btop"},
		{"nix-shell -p fastfetch", "", "", "Временно запустить утилиту fastfetch для вывода информации о вашем железе", "Temporarily run fastfetch utility to display system info", "nix-shell -p fastfetch", "nix-shell -p fastfetch"},
		{"nix-shell -p sys-fs/ncdu", "", "", "Временно запустить консольный сканер диска ncdu для поиска тяжелых папок", "Temporarily run ncdu disk usage analyzer to find heavy folders", "nix-shell -p ncdu", "nix-shell -p ncdu"},

		// --- КАСТОМИЗАЦИЯ И ИНТЕРФЕЙС (Wayland/Эстетика) ---
		{"programs.hyprland.enable = true;", "", "", "Строка для включения оконного менеджера Hyprland в файле configuration.nix", "Line to enable Hyprland window manager in configuration.nix", "programs.hyprland.enable = true;", "programs.hyprland.enable = true;"},
		{"nix-shell -p waybar wofi", "", "", "Временно запустить панель Waybar и меню приложений Wofi без добавления в конфиг", "Temporarily run Waybar status bar and Wofi app launcher", "nix-shell -p waybar wofi", "nix-shell -p waybar wofi"},
		{"grep -E 'hyprland|waybar|wofi' /etc/nixos/configuration.nix", "", "", "Быстро проверить через терминал, прописаны ли уже эти утилиты в вашем конфиге", "Quickly check if these utilities are already written in your system config", "grep -E 'hyprland|waybar|wofi'...", "grep -E 'hyprland|waybar|wofi'..."},
		{"nix-shell -p cava", "", "", "Временно запустить анимированный спектральный эквалайзер звука cava под музыку", "Temporarily run cava audio visualizer under music without installation", "nix-shell -p cava", "nix-shell -p cava"},
		{"nix-shell -p pipes-sh cmatrix", "", "", "Временно запустить классические скринсейверы pipes.sh (трубы) и cmatrix (Матрица)", "Temporarily run screensavers pipes-sh and cmatrix together", "nix-shell -p pipes-sh cmatrix", "nix-shell -p pipes-sh cmatrix"},

		// --- FLATPAK (Управление готовыми приложениями) ---
		{"flatpak search", "[имя]", "[name]", "Найти графическое приложение в репозитории Flathub (Включается в конфиге через services.flatpak.enable)", "Search for a graphical application in Flathub", "flatpak search spotify", "flatpak search spotify"},
		{"flatpak install flathub", "[имя]", "[name]", "Установить готовое графическое приложение (Discord, Spotify) напрямую из Flathub", "Install a graphical application from Flathub repository", "flatpak install flathub com.spotify.Client", "flatpak install flathub com.spotify.Client"},
		{"flatpak update", "", "", "Обновить все установленные Flatpak-приложения до актуальных версий", "Update all currently installed Flatpak applications", "flatpak update", "flatpak update"},

		// --- ЖЕЛЕЗО, СТАТУС И ДИСКИ ---
		{"free -h", "", "", "Показать количество свободной и занятой оперативной памяти (RAM) в удобном виде", "Show free and used amount of RAM in human-readable format", "free -h", "free -h"},
		{"lsblk", "", "", "Показать структуру всех подключенных дисков, их разделов, размеры и точки монтирования", "List all block devices, partitions, sizes and mount points", "lsblk", "lsblk"},
		{"df -h", "", "", "Показать список дисков и сколько на них осталось свободного места", "Show disk space usage in human-readable format", "df -h", "df -h"},
		{"nvidia-smi", "", "", "Показать статус видеокарты NVIDIA, её температуру и загрузку (Драйвер включается через конфиг)", "Show NVIDIA GPU status, temperature and usage (Enabled via config)", "nvidia-smi", "nvidia-smi"},
		{"uptime", "", "", "Показать время непрерывной работы компьютера с момента включения и нагрузку", "Show how long the system has been running and load average", "uptime", "uptime"},

		// --- ЗВУК, BLUETOOTH И СЕТЬ ---
		{"nix-shell -p pavucontrol blueman", "", "", "Временно запустить графическую панель управления звуком и Bluetooth-устройствами", "Temporarily run GUI volume panel and Bluetooth manager shell", "nix-shell -p pavucontrol blueman", "nix-shell -p pavucontrol blueman"},
		{"nmtui", "", "", "Открыть удобный консольный интерфейс для настройки Wi-Fi и проводного интернета", "Open text user interface for Wi-Fi and network setup", "nmtui", "nmtui"},
		{"brightnessctl set", "[значение]", "[value]", "Изменить яркость экрана ноутбука (Пакет brightnessctl должен быть прописан в конфиге)", "Change laptop screen brightness (Requires brightnessctl in configuration)", "brightnessctl set 50%", "brightnessctl set 50%"},

		// --- СКОРАЯ ПОМОЩЬ И МУСОР (Поколения NixOS) ---
		{"nix-env --list-generations", "", "", "Показать список всех созданных поколений системы (точки отката при сбоях обновлений)", "List all created system generations (rollback points)", "nix-env --list-generations", "nix-env --list-generations"},
		{"sudo nix-collect-garbage -d", "", "", "Полностью удалить старые поколения системы и очистить кэш (освобождает гигабайты)", "Collect garbage and delete old generations (frees up gigabytes)", "sudo nix-collect-garbage -d", "sudo nix-collect-garbage -d"},
		{"killall -9", "[имя_программы]", "[app_name]", "Принудительно закрыть «намертво» зависшую программу или игру", "Force kill a completely frozen application or game", "killall -9 discord", "killall -9 discord"},
		{"sudo shutdown -h", "[минуты]", "[minutes]", "Запланировать безопасное выключение компьютера через указанное время", "Schedule a safe system shutdown after a specified time", "sudo shutdown -h +60", "sudo shutdown -h +60"},
	}
}
