package constants

const (
	WeatherClear = iota
	WeatherSunny
	WeatherSunnyHarsh
	WeatherRain
	WeatherRainHarsh
	WeatherSandstorm
	WeatherHail
	WeatherAirCurrent
)

var (
	WeatherNames map[int]string
)

func init() {
	WeatherNames = make(map[int]string)

	WeatherNames[WeatherClear] = "Clear Skies"
	WeatherNames[WeatherSunny] = "Harsh Sunlight"
	WeatherNames[WeatherSunnyHarsh] = "Extremely Harsh Sunlight"
	WeatherNames[WeatherRain] = "Rain"
	WeatherNames[WeatherRainHarsh] = "Heavy Rain"
	WeatherNames[WeatherSandstorm] = "Sandstorm"
	WeatherNames[WeatherHail] = "Hail"
	WeatherNames[WeatherAirCurrent] = "Mysterious Air Current"
}
