package mqttModels

type Device struct {
	Ids          string `json:"ids"`
	Name         string `json:"name"`
	Software     string `json:"sw"`
	Model        string `json:"model"`
	Manufacturer string `json:"mf"`
}
