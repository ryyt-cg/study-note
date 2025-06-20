use std::num::ParseIntError;
#[derive(Debug)]
struct SummationError;

fn sum_of_strings(strs: Vec<&str>) -> Result<String, ParseIntError> {
    let mut sum = 0;
    for s in strs {
        let val =  to_int(&s)?;
        sum += val;

    }

    Ok(sum.to_string())
}

fn to_int(s: &str) -> Result<i32, ParseIntError> {
    s.parse()
}

fn main() {
    let strs = vec!["1", "2", "3"];
    let sum = sum_of_strings(strs);
    println!("sum: {:?}", sum);

    let v = vec!["1", "2", "abc"];
    let total = sum_of_strings(v);
    println!("total: {:?}", total);
}