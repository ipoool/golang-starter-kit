package models

// VehicleTypeModel - Vehicle type
type VehicleTypeModel struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	Year           int               `json:"year"`
	Wheel          int               `json:"wheel"`
	EngineCapacity int               `json:"engine_capacity"`
	Photos         string            `json:"photos"`
	VehicleBrand   VehicleBrandModel `json:"vehicle_brand"`
}

// VehicleBrandModel - Vehicle brand
type VehicleBrandModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}
