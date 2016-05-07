fn insertion_sort(items: &mut [isize]) {
    let len = items.len();
    if len < 2 {
        return;
    }

    for j as isize in 1..len {
        let value = items[j];
        let mut i = j - 1;

        while i >= 0 && items[i] > value {
            items[i + 1] = items[i];
            i -= 1;
        }

        items[i + 1] = value;
    }
}

fn main() {
    let mut items = vec![ 10, 9, 3, 3, 4, 4, 3, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1 ];
    insertion_sort(&mut items);
    println!("{:?}", items);
}
