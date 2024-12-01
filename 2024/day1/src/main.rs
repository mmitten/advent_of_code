use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() -> std::io::Result<()> {
    if let Ok(lines) = read_lines("input") {
        let mut set1 = Vec::new();
        let mut set2 = Vec::new();
        for line in lines.flatten() {
            let parts = line.split_whitespace();
            let collection = parts.collect::<Vec<&str>>();
            let parsed1: i32 = collection[0].parse().unwrap_or(0);
            let parsed2: i32 = collection[1].parse().unwrap_or(0);
            set1.push(parsed1);
            set2.push(parsed2);
        }
        set1.sort();
        set2.sort();
        part1(&set1, &set2);
        part2(&set1, &set2);
    }
    Ok(())
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn part1(set1: &Vec<i32>, set2: &Vec<i32>) {
    let mut part1: i32 = 0;
    for n in 0..set1.len() {
        let subtotal = set1[n] - set2[n];
        let abs_subtotal = subtotal.abs();
        part1 = abs_subtotal + part1;
    }
    dbg!(part1);
}

fn part2(set1: &Vec<i32>, set2: &Vec<i32>) {
    let mut part2: i32 = 0;
    for first in set1 {
        let mut count: i32 = 0;
        for second in set2 {
            if first == second {
                count += 1;
            }
        }
        part2 = part2 + (count * first);
    }
    dbg!(part2);
}
