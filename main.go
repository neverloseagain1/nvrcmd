package main
import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"nvrcmd/distro"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Структура конфигурационного файла для сохранения языка
type Config struct {
	Lang string `json:"lang"`
}

type displayItem struct {
	cmd     string
	desc    string
	example string
}

type model struct {
	rawCommands    []distro.CommandItem
	displayList    []displayItem
	cursor         int
	scrollOffset   int // Индекс первой видимой команды на экране
	terminalHeight int // Текущая высота окна терминала
	lang           string // "ru" или "en"
	distroName     string
	configPath     string
}

// Функция парсит /etc/os-release для автоматического определения дистрибутива
func detectDistro() string {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "generic"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ID=") {
			id := strings.TrimPrefix(line, "ID=")
			id = strings.Trim(id, "\"")
			return id
		}
	}
	return "generic"
}

// Функция загружает конфиг из домашней папки или создает дефолтный
func loadOrCreateConfig() (Config, string) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{Lang: "ru"}, ""
	}

	configDir := filepath.Join(home, ".config")
	// Создаем папку ~/.config, если её почему-то нет
	_ = os.MkdirAll(configDir, 0755)

	configPath := filepath.Join(configDir, "nvrcmd.json")
	
	file, err := os.ReadFile(configPath)
	if err != nil {
		// Если файла нет, возвращаем дефолт и путь для будущей записи
		return Config{Lang: "ru"}, configPath
	}

	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return Config{Lang: "ru"}, configPath
	}

	return cfg, configPath
}

// Функция сохраняет настройки языка в файл
func saveConfig(path string, lang string) {
	if path == "" {
		return
	}
	cfg := Config{Lang: lang}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err == nil {
		_ = os.WriteFile(path, data, 0644)
	}
}

func initialModel(distroName, lang, configPath string) model {
	allCmds := distro.GetCommonCommands()
	allCmds = append(allCmds, distro.GetDistroCommands(distroName)...)

	m := model{
		rawCommands:    allCmds,
		lang:           lang,
		distroName:     distroName,
		terminalHeight: 40, // Дефолтное значение, защищающее от вылетов при старте в тайлинге
		configPath:     configPath,
	}
	m.updateDisplayList()
	return m
}

func (m *model) updateDisplayList() {
	m.displayList = make([]displayItem, len(m.rawCommands))
	for i, item := range m.rawCommands {
		var fullCmd string
		var desc string
		var example string

		if m.lang == "en" {
			desc = item.DescEn
			example = item.ExampleEn
			if item.ArgEn != "" {
				fullCmd = item.Cmd + " " + item.ArgEn
			} else {
				fullCmd = item.Cmd
			}
		} else {
			desc = item.DescRu
			example = item.ExampleRu
			if item.ArgRu != "" {
				fullCmd = item.Cmd + " " + item.ArgRu
			} else {
				fullCmd = item.Cmd
			}
		}

		m.displayList[i] = displayItem{cmd: fullCmd, desc: desc, example: example}
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// Если терминал прислал корректную высоту, сохраняем её
		if msg.Height > 0 {
			m.terminalHeight = msg.Height
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				if m.cursor < m.scrollOffset {
					m.scrollOffset = m.cursor
				}
			}
		case "down", "j":
			if m.cursor < len(m.displayList)-1 {
				m.cursor++
				
				// Вычитаем 5 строк, которые железно заняты под шапку интерфейса
				maxVisibleItems := m.terminalHeight - 5
				if maxVisibleItems < 1 {
					maxVisibleItems = 1
				}

				if m.cursor >= m.scrollOffset+maxVisibleItems {
					m.scrollOffset = m.cursor - maxVisibleItems + 1
				}
			}
		case "l":
			if m.lang == "ru" {
				m.lang = "en"
			} else {
				m.lang = "ru"
			}
			m.updateDisplayList()
			// Автоматически сохраняем новый язык при переключении клавишей 'l'
			saveConfig(m.configPath, m.lang)
		}
	}
	return m, nil
}

func (m model) View() string {
	var s strings.Builder

	// 1. РЕНДЕРИМ ФИОЛЕТОВЫЙ ЗАГОЛОВОК С НАЗВАНИЕМ УТИЛИТЫ NVRCMD
	var titleText string
	if m.lang == "en" {
		titleText = fmt.Sprintf(" NVRCMD Handbook [%s] | Language: %s ", strings.ToUpper(m.distroName), strings.ToUpper(m.lang))
	} else {
		titleText = fmt.Sprintf(" Справочник NVRCMD [%s] | Язык: %s ", strings.ToUpper(m.distroName), strings.ToUpper(m.lang))
	}
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color("#7D56F4")).Padding(0, 1)
	s.WriteString(titleStyle.Render(titleText) + "\n\n")

	// 2. РЕНДЕРИМ СТРОКУ НАВИГАЦИИ
	var navText string
	if m.lang == "en" {
		navText = "Navigation: ↑/↓ or j/k. Language: l. Exit: q."
	} else {
		navText = "Навигация: ↑/↓ или j/k. Смена языка: l. Выход: q."
	}
	hintStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#A0A0A0"))
	s.WriteString(hintStyle.Render(navText) + "\n\n")

	// Вычисляем, сколько команд может поместиться на экране
	maxVisibleItems := m.terminalHeight - 5
	if maxVisibleItems < 1 {
		maxVisibleItems = 1
	}
	
	endIndex := m.scrollOffset + maxVisibleItems
	if endIndex > len(m.displayList) {
		endIndex = len(m.displayList)
	}

	// Корректируем scrollOffset, если размер экрана внезапно изменился
	if m.scrollOffset+maxVisibleItems > len(m.displayList) {
		m.scrollOffset = len(m.displayList) - maxVisibleItems
		if m.scrollOffset < 0 {
			m.scrollOffset = 0
		}
	}

	// 3. РЕНДЕРИМ ОГРАНИЧЕННЫЙ СПИСОК КОМАНД СО СКРОЛЛОМ
	for i := m.scrollOffset; i < endIndex; i++ {
		if i >= len(m.displayList) {
			break
		}
		item := m.displayList[i]
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		cmdStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFD2")).Bold(m.cursor == i)
		descStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA"))
		exampleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB52A")).Italic(true)

		exampleText := ""
		if item.example != "" {
			if m.lang == "en" {
				exampleText = exampleStyle.Render(fmt.Sprintf(" // Ex: %s", item.example))
			} else {
				exampleText = exampleStyle.Render(fmt.Sprintf(" // Пример: %s", item.example))
			}
		}

		s.WriteString(fmt.Sprintf("%s %s - %s%s\n", 
			cursor, 
			cmdStyle.Render(item.cmd), 
			descStyle.Render(item.desc), 
			exampleText,
		))
	}

	return s.String()
}

func main() {
	// 1. Проверяем сохраненный язык пользователя в файле конфигурации
	cfg, configPath := loadOrCreateConfig()

	// 2. Автоматически определяем текущий дистрибутив системы
	distroName := detectDistro()

	// 3. Принудительный ручной ввод дистрибутива имеет приоритет (для тестов утилиты)
	if len(os.Args) > 1 {
		distroName = os.Args[1]
	}

	p := tea.NewProgram(
		initialModel(distroName, cfg.Lang, configPath),
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Ошибка запуска: %v", err)
		os.Exit(1)
	}
}