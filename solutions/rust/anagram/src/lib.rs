use std::collections::{HashMap, HashSet};

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {

    fn counter(s: &str) -> HashMap<char, usize> {
        let mut counter: HashMap<char, usize> = HashMap::new();
        for c in s.to_lowercase().chars() {
            *counter.entry(c).or_insert(0) += 1;
        }
        counter
    }

    let mut anagrams: HashSet<&'a str> = HashSet::new();

    let word_len: usize = word.len();
    let word_count: HashMap<char, usize> = counter(&word);
    let word_lowercase: String = word.to_lowercase();

    for &a in possible_anagrams {
        if a.len() == word_len && counter(a) == word_count && a.to_lowercase() != word_lowercase {
            anagrams.insert(a);
        }
    }

    anagrams
}
