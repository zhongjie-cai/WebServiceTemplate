mkdir -p coverage

allTestCoverResultFile="coverage/all.test.cover.result"
# Clean up all test coverage result file
rm -f $allTestCoverResultFile

for package in $(go list ./../... | grep -v '/vendor/')
do
	echo "Running coverage for $package"

	outFileName="coverage/"$(echo $package | sed -e 's/\//-/g' )
	rm -f $outFileName.cover.json

	# Run test coverage
	./gocov test $package > $outFileName.cover.json 2>> $allTestCoverResultFile

	# Create coverage report for JResultArchiver
	cat $outFileName.cover.json | ./gocov-xml > $outFileName.cover.xml

	# Create coverage report as html for local system
	./gocov-html $outFileName.cover.json > $outFileName.cover.html
done

echo ""
echo "Coverage result:"

# Echo result for visibility
FullReportLines=$(grep -e '/WebServiceTemplate' $allTestCoverResultFile)

FailedLines=$(echo "$FullReportLines" | grep -e 'FAIL')
if [ -n "$FailedLines" ]
then
	# some tests failed during coverage call
	echo "Failed packages:"
	echo "$FailedLines"
	echo ""
	echo "Some tests are not even passing! You should fix them NOW before anything committed!"
	exit 3
fi

UncoveredLines=$(echo "$FullReportLines" | grep -e 'ok' | grep -e '100.0% of statements' -e 'no test files' -e 'FAIL' -v)
if [ -n "$UncoveredLines" ]
then
	# some not fully covered
	echo "Not fully covered packages:"
	echo "$UncoveredLines"
	echo ""
	echo "Some code lines are not covered! You should cover them ASAP before anything deployed!"
	exit 2
fi

WronglySkippedLines=$(echo "$FullReportLines" | grep -e 'no test files' | grep -e '/enum' -e '/constant' -e '/model' -v)
if [ -n "$WronglySkippedLines" ]
then
	# some wrongly skipped packages
	echo "Wrongly skipped packages:"
	echo "$WronglySkippedLines"
	echo ""
	echo "Some code packages are wrongly skipped for coverage! You should cover them ASAP before anything deployed!"
	exit 1
fi

# all fully covered, so success
echo "All code lines are covered! You are good to go!"
exit 0
