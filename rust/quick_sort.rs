fn quick_sort(items: &mut [isize]) {
    let len = items.len();
    if len < 2 {
        return;
    }

    let mut median_pos = 0;

    for i in 1..len {
        if items[i] < items[median_pos] {
            if i > median_pos + 1 {
                items.swap(median_pos, median_pos + 1);
            }
            items.swap(median_pos, i);
            median_pos += 1;
        }
    }

    quick_sort(&mut items[0..median_pos]);
    quick_sort(&mut items[median_pos+1..len]);
}

fn main() {
    let mut items = vec![ 10, 9, 3, 3, 4, 4, 3, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1 ];
    quick_sort(&mut items);
    println!("{:?}", items);
}
