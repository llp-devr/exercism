pub fn build_proverb(list: &[&str]) -> String {
    let mut proverb: String = String::new();

    if list.is_empty() {
        return proverb;
    }

    for i in 1..list.len() {
        proverb.push_str(format!("For want of a {} the {} was lost.\n", list[i - 1], list[i]).as_str());
    }

    proverb.push_str(format!("And all for the want of a {}.", list[0]).as_str());

    proverb
}
