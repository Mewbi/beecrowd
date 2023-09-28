use std::io;

fn kamenetsky(n: f64) -> i32 {
    if n < 0.0 {
        return 0;
    }

    if n <= 1.0 {
        return 1;
    }

    let x = (n * (n / std::f64::consts::E).log10()) +
             ((2.0 * std::f64::consts::PI * n).log10() / 2.0);

    return x.floor() as i32 + 1;
}

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Failed to read line");
    let n: f64 = input.trim().parse().expect("Invalid input");

    println!("{}", kamenetsky(n));
}

