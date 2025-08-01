#[derive(Debug)]
pub struct HighScores {
    scores: Vec<u32>,
}

impl HighScores {
    pub fn new(scores: &[u32]) -> Self {
        Self { scores: Vec::from(scores) }
    }

    pub fn scores(&self) -> &[u32] {
        &self.scores
    }

    pub fn latest(&self) -> Option<u32> {
        self.scores.last().cloned()
    }

    pub fn personal_best(&self) -> Option<u32> {
        self.scores.iter().max().cloned()
    }

    pub fn personal_top_three(&self) -> Vec<u32> {
        let mut top_scores = self.scores.clone();
        let n = top_scores.len();

        // Bubble Sort is ok for small Vec of scores.
        for i in 0..n {
            for j in 0..n - i - 1 {
                if top_scores[j] < top_scores[j + 1] {
                    top_scores.swap(j, j + 1);
                }
            }
        }

        top_scores.into_iter().take(3).collect()
    }
}
