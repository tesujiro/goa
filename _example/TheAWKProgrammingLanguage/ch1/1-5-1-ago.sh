ago -g '
$3 > 15	{ emp = emp + 1}
END	{ print emp, "employees worked more than hours" }
' emp.data