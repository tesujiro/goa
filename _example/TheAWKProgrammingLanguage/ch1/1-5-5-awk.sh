awk '
	{ names = names $1 " " }
END	{ print names }
' emp.data
