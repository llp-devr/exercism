use std::collections::LinkedList;

pub fn brackets_are_balanced(string: &str) -> bool {
    let mut balance: LinkedList<char> = LinkedList::new();

    for char in string.chars() {
        match char {
            '[' => balance.push_back(char),
            '{' => balance.push_back(char),
            '(' => balance.push_back(char),
            ']' => match balance.pop_back() {
                Some('[') => (),
                _ => return false,
            }
            '}' => match balance.pop_back() {
                Some('{') => (),
                _ => return false,
            }
            ')' => match balance.pop_back() {
                Some('(') => (),
                _ => return false,
            }
            _ => ()
        }
    }

    balance.len() == 0
}
