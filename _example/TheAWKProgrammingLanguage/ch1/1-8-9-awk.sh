awk '
{ print "max=",max }
$1 > max { max = $1; maxline = $0 }
END	 { print max, maxline }
' emp.data