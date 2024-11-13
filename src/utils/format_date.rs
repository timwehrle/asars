use chrono::{Duration, NaiveDate};

pub fn format_date(date_str: &str, format: &str) -> String {
    let parsed_date = NaiveDate::parse_from_str(date_str, "%Y-%m-%d").ok();
    let today = chrono::Local::now().date_naive();

    match parsed_date {
        Some(date) if date == today => "Today".to_string(),
        Some(date) if date == today - Duration::days(1) => "Yesterday".to_string(),
        Some(date) if date == today + Duration::days(1) => "Tomorrow".to_string(),
        Some(date) => date.format(format).to_string(),
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
