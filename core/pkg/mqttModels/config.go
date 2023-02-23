package mqttModels

type Config struct {
	Schema string `json:"schema"`

	DevClass          string `json:"dev_cla"`
	UnitOfMeasurement string `json:"unit_of_meas"`
	StatisticsClass   string `json:"stat_cla"`

	ColorMode           bool     `json:"clrm"`
	SupportedColorModes []string `json:"supported_color_modes"`

	Name string `json:"name"`

	StateTopic        string `json:"stat_t"`
	CommandTopic      string `json:"cmd_t"`
	AvailabilityTopic string `json:"avty_y"`

	UniqueId string `json:"uniq_id"`

	Dev Device `json:"dev"`
}
