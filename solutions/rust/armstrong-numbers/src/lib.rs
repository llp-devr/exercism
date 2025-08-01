pub fn is_armstrong_number(num: u32) -> bool {
    let num_str: String = num.to_string();
    let num_digits: u32 = num_str.len() as u32;

    let sum: u64 = num_str
        .chars()
        .map(|d| d.to_digit(10).unwrap() as u64)
        .map(|p| p.pow(num_digits) as u64)
        .sum();

    num as u64 == sum
}
