package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

// Функция для переключения входа монитора
func switchMonitorInput(monitorID string, inputCode int) error {
	controlMyMonitorPath := filepath.Join("ControlMyMonitor", "ControlMyMonitor.exe")
	// Команда для переключения входа
	cmd := exec.Command(controlMyMonitorPath, "/SetValue", monitorID, "60", fmt.Sprintf("%d", inputCode))

	// Запуск команды
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ошибка при выполнении команды: %v, вывод: %s", err, string(output))
	}

	fmt.Printf("Успешно переключено на вход %d\n", inputCode)
	return nil
}

func main() {
	// Идентификатор монитора (замените на ваш)
	monitorID := `\\.\DISPLAY1\Monitor0`

	// Коды входов (замените на актуальные для вашего монитора)
	hdmi1Code := 17 // Код для HDMI1
	hdmi2Code := 18 // Код для HDMI2

	// Переключение на HDMI1
	fmt.Println("Переключаем на HDMI1...")
	err := switchMonitorInput(monitorID, hdmi1Code)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Переключение на HDMI2
	fmt.Println("Переключаем на HDMI2...")
	err = switchMonitorInput(monitorID, hdmi2Code)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
}
