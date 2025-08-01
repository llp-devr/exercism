use std::fmt;

#[derive(Debug, PartialEq)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let total_minutes: i32 = (hours * 60 + minutes) % 1440;
        let (fixed_hours, fixed_minutes) = match total_minutes >= 0 {
            true => (total_minutes / 60 as i32,
                     total_minutes % 60),
            false => ((total_minutes + 1440) / 60 as i32,
                      (total_minutes + 1440) % 60),
        };
        Clock {
            hours: fixed_hours,
            minutes: fixed_minutes,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(self.hours, self.minutes + minutes)
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours % 24, self.minutes)
    }
}