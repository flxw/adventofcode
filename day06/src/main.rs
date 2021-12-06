fn main() {
    let state = vec![
        1, 1, 3, 5, 3, 1, 1, 4, 1, 1, 5, 2, 4, 3, 1, 1, 3, 1, 1, 5, 5, 1, 3, 2, 5, 4, 1, 1, 5, 1,
        4, 2, 1, 4, 2, 1, 4, 4, 1, 5, 1, 4, 4, 1, 1, 5, 1, 5, 1, 5, 1, 1, 1, 5, 1, 2, 5, 1, 1, 3,
        2, 2, 2, 1, 4, 1, 1, 2, 4, 1, 3, 1, 2, 1, 3, 5, 2, 3, 5, 1, 1, 4, 3, 3, 5, 1, 5, 3, 1, 2,
        3, 4, 1, 1, 5, 4, 1, 3, 4, 4, 1, 2, 4, 4, 1, 1, 3, 5, 3, 1, 2, 2, 5, 1, 4, 1, 3, 3, 3, 3,
        1, 1, 2, 1, 5, 3, 4, 5, 1, 5, 2, 5, 3, 2, 1, 4, 2, 1, 1, 1, 4, 1, 2, 1, 2, 2, 4, 5, 5, 5,
        4, 1, 4, 1, 4, 2, 3, 2, 3, 1, 1, 2, 3, 1, 1, 1, 5, 2, 2, 5, 3, 1, 4, 1, 2, 1, 1, 5, 3, 1,
        4, 5, 1, 4, 2, 1, 1, 5, 1, 5, 4, 1, 5, 5, 2, 3, 1, 3, 5, 1, 1, 1, 1, 3, 1, 1, 4, 1, 5, 2,
        1, 1, 3, 5, 1, 1, 4, 2, 1, 2, 5, 2, 5, 1, 1, 1, 2, 3, 5, 5, 1, 4, 3, 2, 2, 3, 2, 1, 1, 4,
        1, 3, 5, 2, 3, 1, 1, 5, 1, 3, 5, 1, 1, 5, 5, 3, 1, 3, 3, 1, 2, 3, 1, 5, 1, 3, 2, 1, 3, 1,
        1, 2, 3, 5, 3, 5, 5, 4, 3, 1, 5, 1, 1, 2, 3, 2, 2, 1, 1, 2, 1, 4, 1, 2, 3, 3, 3, 1, 3, 5,
    ];

    let population_size = calculate_count_after(80, &state);
    println!("Number of lanternfish after 80 days: {}", population_size);


    let population_size = calculate_count_after(256, &state);
    println!("Number of lanternfish after 256 days: {}", population_size);
}

fn calculate_count_after(days: u32, initial_state: &Vec<u32>) -> usize {
    let mut compressed_state = compress_state(initial_state);

    for day in 0..days {
        let pregnant_fish = compressed_state[0];
        compressed_state.rotate_left(1);
        compressed_state[6] += pregnant_fish;
        compressed_state[8] = pregnant_fish;
    }

    compressed_state.iter().sum()
}

fn compress_state(state: &Vec<u32>) -> Vec<usize>{
    vec![
        state.iter().filter(|e| **e == 0).count(),
        state.iter().filter(|e| **e == 1).count(),
        state.iter().filter(|e| **e == 2).count(),
        state.iter().filter(|e| **e == 3).count(),
        state.iter().filter(|e| **e == 4).count(),
        state.iter().filter(|e| **e == 5).count(),
        state.iter().filter(|e| **e == 6).count(),
        state.iter().filter(|e| **e == 7).count(),
        state.iter().filter(|e| **e == 8).count(),
    ]
}

#[cfg(test)]
mod tests {
    use crate::*;

    #[test]
    fn computes_population_size_after_time() {
        let initial_state = vec![3,4,3,1,2];
        assert_eq!(calculate_count_after(18, &initial_state), 26);
        assert_eq!(calculate_count_after(80, &initial_state), 5934);
        assert_eq!(calculate_count_after(256, &initial_state), 26984457539);
    }
}
