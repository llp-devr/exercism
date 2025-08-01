pub fn production_rate_per_hour(speed: u8) -> f64 {
    let production_per_hour: u8 = 221;

    let mut success_rate: f64 = 1.;

    match speed {
        0 => return 0 as f64,
        1..=4 => success_rate = 1.00,
        5..=8 => success_rate = 0.90,
        9..=10 => success_rate = 0.77,
        _ => panic!("Overspeed"),
    }

    (production_per_hour as f64) * (speed as f64) * success_rate
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    (production_rate_per_hour(speed) / 60_f64).floor() as u32
}
