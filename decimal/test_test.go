package decimal

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"testing"

	anzbank "github.com/anz-bank/decimal"
	ericlagergren "github.com/ericlagergren/decimal"
	shopspring "github.com/shopspring/decimal"
)

type testCaseStrings struct {
	testName       string
	testFunc       string
	val1           string
	val2           string
	val3           string
	expectedResult string
}

var numloops = 1000

var testPaths = []string{
	"ddAdd.decTest",
	"ddMultiply.decTest",
	"ddAbs.decTest",
	"ddDivide.decTest",
}

type testcase struct {
	op     string
	v1, v2 interface{}
}

var prettyNames = map[string]string{
	"float64":           "float64",
	"decimal.Decimal64": "anzDecimal",
	"big.Float":         "bigFloat",
}


func BenchmarkDecimal(b *testing.B) {
	types := map[string]interface{}{
		"decimal.Decimal64": anzbank.Decimal64{},
		"float64":           float64(0),
		"big.Float":         big.Float{},
	}
	typeMap := make(map[string][]testcase)
	for _, dectest := range testPaths {
		file, _ := os.Open("dectest/" + dectest)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			testVal := getInput(scanner.Text())
			if testVal.testName != "" {
				for key, val := range types {
					typeMap[key] = append(typeMap[key],
						ParseDecimal(
							testVal.val1,
							testVal.val2,
							val,
							testVal.testFunc))
				}
			}
		}
		for key, val := range types {
			b.Run(fmt.Sprintf("%s_%s", dectest, prettyNames[key]), func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					for _, test := range typeMap[key] {
						runtests(test, val)

					}
				}
			})
		}
	}
}

// Parse the vals as type of interface v
func ParseDecimal(val1, val2 string, v interface{}, op string) (test testcase) {
	switch v.(type) {
	case float64:
		c, _ := strconv.ParseFloat(val1, 64)
		d, _ := strconv.ParseFloat(val2, 64)
		test.v1 = c
		test.v2 = d
	case big.Float:
		c := big.Float{}
		d := big.Float{}
		fmt.Sscan(val1, c)
		fmt.Sscan(val2, d)
		test.v1 = c
		test.v2 = d
	case ericlagergren.Big:
		c := ericlagergren.Big{}
		d := ericlagergren.Big{}
		test.v1, _ = c.SetString(val1)
		test.v2, _ = d.SetString(val2)
	case anzbank.Decimal64:
		c, _ := anzbank.Parse64(val1)
		d, _ := anzbank.Parse64(val2)
		test.v1 = c
		test.v2 = d
	case shopspring.Decimal:
		c, _ := shopspring.NewFromString(val1)
		d, _ := shopspring.NewFromString(val2)
		test.v1 = &c
		test.v2 = &d
	}
	test.op = op
	return
}

func runtests(test testcase, t interface{}) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	switch (t).(type) {
	case float64:
		a := test.v1.(float64)
		b := test.v2.(float64)
		execOpFloat(a, b, test.op)
	case big.Float:
		a := test.v1.(big.Float)
		b := test.v2.(big.Float)
		execOpBig(a, b, test.op)
	case ericlagergren.Big:
		a := test.v1.(*ericlagergren.Big)
		b := test.v2.(*ericlagergren.Big)
		execOpEric(a, b, test.op)
	case anzbank.Decimal64:
		a := test.v1.(anzbank.Decimal64)
		b := test.v1.(anzbank.Decimal64)
		execOp(a, b, test.op)
	case shopspring.Decimal:
		a := test.v1.(*shopspring.Decimal)
		b := test.v1.(*shopspring.Decimal)
		execOpShop(a, b, test.op)
	default:
	}

}

// getInput gets the test file and extracts test using regex, then returns a map object and a list of test names.
func getInput(line string) testCaseStrings {
	testRegex := regexp.MustCompile(
		`(?P<testName>dd[\w]*)` + // first capturing group: testfunc made of anything that isn't a whitespace
			`(?:\s*)` + // match any whitespace (?: non capturing group)
			`(?P<testFunc>[\S]*)` + // testfunc made of anything that isn't a whitespace
			`(?:\s*\'?)` + // after can be any number of spaces and quotations if they exist (?: non capturing group)
			`(?P<val1>\+?-?[^\t\f\v\' ]*)` + // first test val is anything that isnt a whitespace or a quoteation mark
			`(?:'?\s*'?)` + // match any quotation marks and any space (?: non capturing group)
			`(?P<val2>\+?-?[^\t\f\v\' ]*)` + // second test val is anything that isnt a whitespace or a quoteation mark
			`(?:'?\s*'?)` +
			`(?P<val3>\+?-?[^->]?[^\t\f\v\' ]*)` + //testvals3 same as 1 but specifically dont match with '->'
			`(?:'?\s*->\s*'?)` + // matches the indicator to answer and surrounding whitespaces (?: non capturing group)
			`(?P<expectedResult>\+?-?[^\r\n\t\f\v\' ]*)`) // matches the answer that's anything that is plus minus but not quotations

	// capturing gorups are testName, testFunc, val1,  val2, and expectedResult)
	ans := testRegex.FindStringSubmatch(line)
	if len(ans) < 6 {
		return testCaseStrings{}
	}
	data := testCaseStrings{
		testName:       ans[1],
		testFunc:       ans[2],
		val1:           ans[3],
		val2:           ans[4],
		val3:           ans[5],
		expectedResult: ans[6],
	}
	return data
}

func execOpEric(a, b *ericlagergren.Big, op string) {
	if a == nil || b == nil {
		return
	}
	var g ericlagergren.Big
	switch op {
	case "add":
		for i := 0; i < numloops; i++ {
			g.Add(a, b)
		}
	case "multiply":
		for i := 0; i < numloops; i++ {
			g.Mul(a, b)
		}
	case "abs":
		for i := 0; i < numloops; i++ {
			g.Abs(a)
		}
	case "divide":
		for i := 0; i < numloops; i++ {
			g.Quo(a, b)
		}
	default:

	}
}
func execOp(a, b anzbank.Decimal64, op string) {
	switch op {
	case "add":
		for i := 0; i < numloops; i++ {
			a.Add(b)
		}
	case "multiply":
		for i := 0; i < numloops; i++ {
			a.Mul(b)
		}
	case "abs":
		for i := 0; i < numloops; i++ {
			a.Abs()
		}
	case "divide":
		for i := 0; i < numloops; i++ {
			a.Quo(b)
		}
	default:
	}
}

func execOpBig(a, b big.Float, op string) {
	switch op {
	case "add":
		for i := 0; i < numloops; i++ {
			a.Add(&a, &b)
		}
	case "multiply":
		for i := 0; i < numloops; i++ {
			a.Mul(&a, &b)
		}
	case "abs":
		for i := 0; i < numloops; i++ {
			a.Abs(&a)
		}
	case "divide":
		for i := 0; i < numloops; i++ {
			a.Quo(&a, &b)
		}
	default:
	}
}

func execOpShop(a, b *shopspring.Decimal, op string) {
	switch op {
	case "add":
		for i := 0; i < numloops; i++ {
			a.Add(*b)
		}
	case "multiply":
		for i := 0; i < numloops; i++ {
			a.Mul(*b)
		}
	case "abs":
		for i := 0; i < numloops; i++ {
			a.Abs()
		}
	case "divide":
		for i := 0; i < numloops; i++ {
			a.Div(*b)
		}
	}
}

func execOpFloat(a, b float64, op string) {
	var e float64
	switch op {
	case "add":
		for i := 0; i < numloops; i++ {
			e = a + b
		}
	case "multiply":
		for i := 0; i < numloops; i++ {
			e = a * b
		}
	case "abs":
	case "divide":
		for i := 0; i < numloops; i++ {
			e = a / b
		}
	default:
	}
	if false {
		println(e)
	}
}
