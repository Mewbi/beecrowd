use std::io;

fn euclidean_division(a: i32, b: i32) -> (i32, i32) {
    let limit = if b < 0 { -b } else { b };
    let r = 0;
    let mut q = 0;

    for r in 0..limit {
        if (a - r) % b == 0 {
            q = (a - r) / b;
            return (q, r);
        }
    }

    (q, r)
}

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Failed to read line");
    let tokens: Vec<i32> = input
        .split_whitespace()
        .map(|s| s.parse().expect("Failed to parse number"))
        .collect();

    let a = tokens[0];
    let b = tokens[1];

    let (q, r) = euclidean_division(a, b);

    println!("{} {}", q, r);
}
