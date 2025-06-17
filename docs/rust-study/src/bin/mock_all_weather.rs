#[mockall::automock]
trait WeatherApi {
    fn get_weather(&self, city: &str) -> String;
}

struct RealWeatherApi;

impl WeatherApi for RealWeatherApi {
    fn get_weather(&self, _zip_code: &str) -> String {
        "is sunny today".to_string()
    }
}

fn create_message_weather(api: &dyn WeatherApi, zip_code: &str) -> String {
    if zip_code.trim().is_empty() {
        panic!( "Invalid zip code");
    }

    let msg = api.get_weather(&zip_code);
    format!("The weather in your area (zip code: {} {}", zip_code, msg)
}

// struct WeatherService {
//     api: Box<dyn WeatherApi + 'static>,
// }
//
// impl WeatherService {
//     fn new(api: Box<dyn WeatherApi + 'static>) -> Self {
//         WeatherService {
//             api,
//         }
//     }
//
//     fn get_weather(&self, city: &str) -> String {
//         self.api.get_weather(city)
//     }
// }

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    #[should_panic(expected = "Invalid zip code")]
    fn test_weather_message_empty_zip_code() {
        let api = RealWeatherApi;
        let zip_code = "";
        create_message_weather(&api, zip_code);
    }

    #[test]
    fn test_create_message_weather() {
        let mut mock_api = MockWeatherApi::new();
        let zip_code = "12345";

        mock_api
            .expect_get_weather()
            .once()
            .returning(move |_| "is cloudy today".to_string());
        let result = create_message_weather(&mock_api, zip_code);
        assert_eq!(result, "The weather in your area (zip code: 12345 is cloudy today");
    }
}

fn main() {
    let api = RealWeatherApi;
    let zip_code = "12345";
    let result = create_message_weather(&api, zip_code);
    println!("{}", result);
}