package main

import (
    "encoding/json"
    "fmt"
)

type Request struct {
    Action string `json:"action"`
    Data   string `json:"data"`
}

func analyzeJSON(jsonData []byte, stateTrackingService *StateTrackingService) {
    var request Request
    if err := json.Unmarshal(jsonData, &request); err != nil {
        fmt.Println("Ошибка парсинга JSON:", err)
        return
    }

    // Применение стратегии в зависимости от действия
    switch request.Action {
    case "repeat":
        fmt.Println("Применена стратегия повтора для данных:", request.Data)
        // Передача данных в сервис отслеживания состояния
        stateTrackingService.TrackRepeatAction(request.Data)
    case "cancel":
        fmt.Println("Применена стратегия отмены вызова для данных:", request.Data)
        // Передача данных в сервис отслеживания состояния
        stateTrackingService.TrackCancelAction(request.Data)
    default:
        fmt.Println("Неизвестное действие:", request.Action)
        // Другие действия для неизвестного действия
    }
}

type StateTrackingService struct {
    // Логика и методы для отслеживания состояния
}

func (sts *StateTrackingService) TrackRepeatAction(data string) {
    fmt.Println("Действие повтора отслежено:", data)
    // Логика отслеживания действия повтора
}

func (sts *StateTrackingService) TrackCancelAction(data string) {
    fmt.Println("Действие отмены вызова отслежено:", data)
    // Логика отслеживания действия отмены вызова
}

func main() {
    stateTrackingService := &StateTrackingService{}
    
    jsonData := []byte(`{"action": "repeat", "data": "Some data to repeat"}`)
    analyzeJSON(jsonData, stateTrackingService)

    jsonData = []byte(`{"action": "cancel", "data": "Call cancellation data"}`)
    analyzeJSON(jsonData, stateTrackingService)
}
