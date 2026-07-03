package distro

import "strings"

type CommandItem struct {
	Cmd       string
	ArgRu     string
	ArgEn     string
	DescRu    string
	DescEn    string
	ExampleRu string
	ExampleEn string
}

func GetCommonCommands() []CommandItem {
	return []CommandItem{
		{"cd", "[путь]", "[path]", "Сменить рабочую директорию", "Change the working directory", "cd ~/Загрузки", "cd ~/Downloads"},
		{"ls", "", "", "Посмотреть содержимое текущей директории", "List directory contents", "ls -la", "ls -la"},
		{"mkdir", "[имя]", "[name]", "Создать новую папку", "Create a new directory", "mkdir Проекты", "mkdir Projects"},
		{"rm", "[файл]", "[file]", "Удалить файл", "Remove a file", "rm document.txt", "rm document.txt"},
		{"pwd", "", "", "Показать текущий абсолютный путь", "Print working directory", "pwd", "pwd"},
	}
}

func GetDistroCommands(name string) []CommandItem {
	switch strings.ToLower(name) {
	case "arch", "void":
		return GetArchCommands()
	case "ubuntu", "debian":
		return GetUbuntuCommands()
	case "fedora":
		return GetFedoraCommands()
	case "nixos":
		return GetNixosCommands()
	case "gentoo":
		return GetGentooCommands()
	default:
		return []CommandItem{}
	}
}
