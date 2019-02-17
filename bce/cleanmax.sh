# Replaces minInt, maxInt, maxUint with prettier text representations
sed -i '' 's/smin/smin/g' *$1
sed -i '' 's/smax/smax/g' *$1
sed -i '' 's/umax/umax/g' *$1
