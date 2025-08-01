pub fn square(s: u32) -> u64 {
    if s > 64 || s < 1 {
        panic!("Square must be between 1 and 64");
    }

    // Geometric Progression
    // an = a1 . q(n-1)
    //
    // a1 = 1
    // q = 2
    
    2_u64.pow(s - 1) as u64
}

pub fn total() -> u64 {
    (1..=64).map(|x| square(x)).sum()
}
