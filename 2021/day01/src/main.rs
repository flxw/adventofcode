use std::{io::{BufRead, BufReader}, fs::File};

fn main() {
  let file = File::open("input.txt").unwrap();
  let reader = BufReader::new(file);
  let readings: Vec<u32> = reader.lines().filter(|res| res.is_ok()).map(|res| res.unwrap().parse::<u32>().unwrap()).collect::<Vec<u32>>();
  let depth_changes_single_step = count_depth_increases_in_single_steps(&readings);
  let depth_changes_multi_step = count_depth_increases_in_multiple_steps(&readings);

  println!("There were {} depth changes (single step)", depth_changes_single_step);
  println!("There were {} depth changes (three-step average)", depth_changes_multi_step);
}

fn count_depth_increases_in_single_steps(depth_readings: &[u32]) -> usize {
  depth_readings
    .windows(2)
    .filter(|s| s[0] < s[1])
    .count()
}

fn count_depth_increases_in_multiple_steps(depth_readings: &[u32]) -> usize {
  let aggregated_readings: Vec<u32> = depth_readings.windows(3).map(|s| s.iter().sum()).collect();
  count_depth_increases_in_single_steps(&aggregated_readings)
}

#[cfg(test)]
mod tests {
  use crate::*;

  const READOUTS: [u32;10] = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263];

  #[test]
  fn it_counts_depth_increases_single_step() {
      assert_eq!(count_depth_increases_in_single_steps(&READOUTS), 7);
  }

  #[test]
  fn it_counts_depth_increases_multi_step() {
      assert_eq!(count_depth_increases_in_multiple_steps(&READOUTS), 5);
  }
}
