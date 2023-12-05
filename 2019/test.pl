use strict;
use warnings;

# Open the file in read mode
open(my $file, '<', 'data/1.txt') or die "Cannot open file: $!";
print("ok",<$file>);
# Read the file line by line
# while (my $line = <$file>) {
#     chomp($line);  # Remove the trailing newline character
    
#     print "$line\n";  # Example: Print the line
#     print "ok",<$file>;
# }

# Close the file
close($file);