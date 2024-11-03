package main

import (
)

// Data структура, которая содержит ID и Payload
type Data struct {
	ID      int
	Payload map[string]interface{}
}

// Reader интерфейс для чтения данных
type Reader interface {
	Read() []*Data
}

// Processor интерфейс для обработки данных
type Processor interface {
	Process(data Data) ([]*Data, error)
}

// Writer интерфейс для записи данных
type Writer interface {
	Write(data []*Data)
}

// Manager интерфейс для управления процессом
type Manager interface {
	Manage() // blocking
}

// DataManager структура, которая реализует интерфейс Manager
type DataManager struct {
	reader    Reader
	processors []Processor
	writer    Writer
}

// NewDataManager создает новый экземпляр DataManager
func NewDataManager(reader Reader, processors []Processor, writer Writer) *DataManager {
	return &DataManager{
		reader:    reader,
		processors: processors,
		writer:    writer,
	}
}

// Manage реализует метод Manage интерфейса Manager
func (dm *DataManager) Manage() {
	for {
		databatch := dm.reader.Read()
		if len(databatch) == 0 {
			break
		}
		var res []*Data
		for _, data := range databatch {
			currentData := data
			for _, p := range dm.processors {
				r, err := p.Process(*data)
				if err != nil {
					break
				}
				currentData = r[0]
			}

			if res != nil {
				res = append(res, currentData)
			}
		}
		if len(res) > 0 {
			dm.writer.Write(res)
		}
	}

}