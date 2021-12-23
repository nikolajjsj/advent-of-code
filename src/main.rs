use anyhow::Result;
use itertools::Itertools;

fn main() {
    println!("Result for day 1 part 1 is: {:?}", day_1(2));
    println!("Result for day 1 part 2 is: {:?}", day_1(4));
}

fn day_1(size: usize) -> Result<usize> {
    let path = String::from("./data/1.input");
    Ok(aoc::read_one_per_line::<u32>(&path)?
        .windows(size)
        .filter(|win| win[0] < win[size - 1])
        .collect_vec()
        .len())
}
