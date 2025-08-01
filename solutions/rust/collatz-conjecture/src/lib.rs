pub fn collatz(n: u64) -> Option<u64> {
    if n == 0 {
        return None;
    }
    let mut steps: u64 = 0;
    let mut temp: u64 = n;

    while temp != 1 {
        if temp % 2 == 0 {
            temp /= 2
        } else {
            match temp.checked_mul(3) {
                Some(x) => match x.checked_add(1) {
                    Some (y) => temp = y,
                    None => return None,
                }
                None => return None,
            }
        }
        steps += 1
    }
    Option::from(steps)
}
