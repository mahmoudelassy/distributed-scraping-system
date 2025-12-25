#!/bin/bash

# Name of the output file
OUTPUT="all_files_content.txt"

# Remove old output file if it exists
[ -f "$OUTPUT" ] && rm "$OUTPUT"

# Loop through all files recursively and append their content
find . -type f | while read -r file; do
    echo "----- Start of $file -----" >> "$OUTPUT"
    cat "$file" >> "$OUTPUT"
    echo "----- End of $file -----" >> "$OUTPUT"
    echo "" >> "$OUTPUT"  # extra newline for readability
done

echo "All file contents saved to $OUTPUT"
