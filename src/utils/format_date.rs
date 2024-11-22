use chrono::{Datelike, Duration, NaiveDate};

pub fn format_date(date_str: &str, format: &str) -> String {
    let parsed_date = NaiveDate::parse_from_str(date_str, "%Y-%m-%d").ok();
    let today = chrono::Local::now().date_naive();

    match parsed_date {
        Some(date) if date == today => "Today".to_string(),
        Some(date) if date == today - Duration::days(1) => "Yesterday".to_string(),
        Some(date) if date == today + Duration::days(1) => "Tomorrow".to_string(),
        Some(date) => {
            let days_diff = (date - today).num_days();
            if days_diff > 0 && days_diff < 7 {
                let weekday = date.weekday();
                match weekday {
                    chrono::Weekday::Mon => "Mon".to_string(),
                    chrono::Weekday::Tue => "Tue".to_string(),
                    chrono::Weekday::Wed => "Wed".to_string(),
                    chrono::Weekday::Thu => "Thu".to_string(),
                    chrono::Weekday::Fri => "Fri".to_string(),
                    chrono::Weekday::Sat => "Sat".to_string(),
                    chrono::Weekday::Sun => "Sun".to_string(),
                }
            } else {
                date.format(format).to_string()
            }
        }
        None => "Unknown".to_string(),
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_format_date() {
        let date_str = "2021-08-01";
        let format = "%b %d, %Y";
        let expected = "Aug 01, 2021".to_string();

        assert_eq!(format_date(date_str, format), expected);
    }

    #[test]
    fn test_format_date_today() {
        let date_str = chrono::Local::now()
            .date_naive()
            .format("%Y-%m-%d")
            .to_string();
        let format = "%b %d, %Y";
        let expected = "Today".to_string();

        assert_eq!(format_date(&date_str, format), expected);
    }

    #[test]
    fn test_format_date_yesterday() {
        let date_str = (chrono::Local::now().date_naive() - Duration::days(1))
            .format("%Y-%m-%d")
            .to_string();
        let format = "%b %d, %Y";
        let expected = "Yesterday".to_string();

        assert_eq!(format_date(&date_str, format), expected);
    }

    #[test]
    fn test_format_date_tomorrow() {
        let date_str = (chrono::Local::now().date_naive() + Duration::days(1))
            .format("%Y-%m-%d")
            .to_string();
        let format = "%b %d, %Y";
        let expected = "Tomorrow".to_string();

        assert_eq!(format_date(&date_str, format), expected);
    }
}
