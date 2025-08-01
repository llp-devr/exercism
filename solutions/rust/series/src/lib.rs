pub fn series(digits: &str, len: usize) -> Vec<String> {
    let mut series: Vec<String> = Vec::new();

    if digits.len() < len {
        return series;
    }

    for i in 0..=digits.len() - len {
        series.push(digits[i..i + len].to_string());
    }

    series
}
