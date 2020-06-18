using System;

namespace temperature
{
    public class WeatherForecast
    {
        public int TemperatureC { get; set; }

        public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);

    }
}
