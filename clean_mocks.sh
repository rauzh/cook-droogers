#!/bin/bash

root_dir="."

find "$root_dir" -type d -name "mocks" | while read dir; do
    echo "Removing directory: $dir"
    rm -rf "$dir"
done

echo "All 'mocks' directories have been removed successfully."
