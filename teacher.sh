interviewn=$(grep -H "license" interviews/* | grep "\" " | cut -f1 -d ":" | rev | cut -f1 -d "-" | rev)
interview="interviews/interview-$interviewn"
export interviewn
echo $interviewn
cat "$interview"
echo $MAIN_SUSPECT




