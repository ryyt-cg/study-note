# Collections

- Collections are data structures that can hold multiple values.
- Rust standard library provides a number of useful collection types.
- The most common collections are:
    - Sequences: Vec, VecDeque, LinkedList
      - Vec: A growable array
      - VecDeque: A double-ended queue
      - LinkedList: A doubly-linked list (rarely used)
    - Maps: HashMap, BTreeMap
      - HashMap: A collection of key-value pairs
      - BTreeMap: A collection of key-value pairs sorted by key and use for search
    - Sets: HashSet, BTreeSet
      - HashSet: A collection of unique values
      - BTreeSet: A collection of unique values sorted by key - optimized for search
    - Misc.: BinaryHeap
      - BinaryHeap: A priority queue - is very niche and not used often

## Most Useful Collections

- Vec: A growable array
- HashMap: A collection of key-value pairs
- Strings: A growable UTF-8 encoded text

```rust
fn main() {
    // Simple vectors/strings
    let mut my_vec = vec![1, 2, 3, 4, 5];
    // adding a value to the vector
    my_vec.push(6);
    
    // initializing a String - high-level wrapper around Vec<byte>
    let message = String::from("Hello, World!");
    
    // initializing a HashMap
    let mut coffee_ratings = HashMap::new();
    coffee_ratings.insert("Latte", 10);
    
    for (name, rating) in &coffee_ratings {
        println!("{}: {}", name, rating);
    };
}
```

## Primitive Types

Array is fixed-size and cannot grow or shrink.
Tuple a fixed number of elements, sequence of elements, and these elements can have different types.
Slices are references to a contiguous sequence of elements in a collection.

```rust
fn main() {
    // Array
    let my_array: [u8; 5] = [1, 2, 3, 4, 5];
    
    // Tuple
    let my_tuple = (1, "hello", 4.5, true);
    
    // Slices
    let my_slice = &my_array[1..3];
}
```

