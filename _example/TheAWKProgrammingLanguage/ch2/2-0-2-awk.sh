awk '
{ print \
    $1,		# 国名
    $2,		# 1000平方マイルを単位にした面積
    $3 }	# 100万人を単位にした人口
' countries
