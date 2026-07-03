#!/bin/bash

# Меню выбора языка / Language selection menu
echo "Select installation language / Выберите язык установки:"
echo "1) Русский"
echo "2) English"
read -p "Choose (1-2): " lang_choice

if [ "$lang_choice" == "1" ]; then
    MSG_START="🚀 Начинаем установку Справочника NVRCMD..."
    MSG_COMPILING="📦 Компиляция бинарного файла..."
    MSG_ERR_COMP="❌ Ошибка компиляции кода!"
    MSG_SUDO_WARN="⚙️ Делаем команду системной.\n⚠️ Сейчас скрипт попросит пароль Администратора (sudo).\nСимволы при вводе отображаться НЕ БУДУТ. Это нормально.\nНажмите ENTER для продолжения, затем введите пароль."
    MSG_SUCCESS="✅ Установка успешно завершена!\nПерезапустите терминал или откройте новую вкладку и просто введите: nvrcmd"
    MSG_ERR_COPY="❌ Ошибка при копировании файла в системную директорию."
else
    MSG_START="🚀 Starting NVRCMD Handbook installation..."
    MSG_COMPILING="📦 Compiling binary file..."
    MSG_ERR_COMP="❌ Code compilation error!"
    MSG_SUDO_WARN="⚙️ Making command global.\n⚠️ The script will now ask for your Administrator password (sudo).\nCharacters WILL NOT be displayed while typing. This is normal.\nPress ENTER to continue, then type your password."
    MSG_SUCCESS="✅ Installation successfully finished!\nRestart your terminal or open a new tab and just type: nvrcmd"
    MSG_ERR_COPY="❌ Error while copying file to the system directory."
fi

echo "--------------------------------------------------"
echo -e "$MSG_START"
echo "--------------------------------------------------"

# 1. Компиляция
echo -e "$MSG_COMPILING"
go build -o nvrcmd .

if [ $? -ne 0 ]; then
    echo -e "$MSG_ERR_COMP"
    exit 1
fi

# 2. Предупреждение о sudo и принудительная пауза
echo "--------------------------------------------------"
echo -e "$MSG_SUDO_WARN"
echo "--------------------------------------------------"
read -p "" # Скрипт остановится и будет ждать нажатия Enter от пользователя

# 3. Вызов sudo для перемещения бинарника
sudo mv nvrcmd /usr/local/bin/

if [ $? -eq 0 ]; then
    echo "--------------------------------------------------"
    echo -e "$MSG_SUCCESS"
else
    echo "--------------------------------------------------"
    echo -e "$MSG_ERR_COPY"
fi
