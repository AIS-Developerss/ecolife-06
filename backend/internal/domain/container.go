package domain

// Container представляет контейнер для мусора
type Container struct {
	ID          string `json:"id"`
	Volume      int    `json:"volume"` // в литрах
	Price       int    `json:"price"`  // в рублях
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

// ContainerVolume представляет объемы контейнеров
const (
	ContainerVolume120  = 120
	ContainerVolume240  = 240
	ContainerVolume1100 = 1100
)

