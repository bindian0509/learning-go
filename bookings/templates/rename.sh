# Rename all *.txt to *.text
for f in *.htm; do 
    mv -- "$f" "${f%.htm}.old"
done

for f in *.tmpl; do 
    mv -- "$f" "${f%.tmpl}.htm"
done
