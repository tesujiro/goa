awk '$2 * $3 >50 { printf("$%.2f for %s\n", $2 * $3, $1) }' emp.data
