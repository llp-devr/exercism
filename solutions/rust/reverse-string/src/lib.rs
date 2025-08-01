use unicode_segmentation::UnicodeSegmentation;

pub fn reverse(input: &str) -> String {
    let mut reversed = String::new();

    for grapheme in input.graphemes(true).rev() {
        reversed.push_str(grapheme);
    }

    reversed
}
