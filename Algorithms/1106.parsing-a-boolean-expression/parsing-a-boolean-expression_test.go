package problem1106

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	expression string
	ans        bool
}{

	{
		"&(&(|(f),&(&(&(&(t),&(!(t),!(|(f,f,t)),!(&(f))),|(f,f,t)),&(&(t),&(&(&(&(f),&(!(t),&(f),|(f)),&(!(&(f)),&(t),|(f,f,t))),|(t),&(!(t),!(|(f,f,t)),!(&(f)))),!(&(|(f,f,t),&(&(f),&(!(t),&(f),|(f)),&(!(&(f)),&(t),|(f,f,t))),&(t))),&(&(&(!(&(f)),|(t),&(!(t),!(|(f,f,t)),!(&(f)))),!(|(f,f,t)),&(t,t,f)),&(f),&(&(t),&(!(t),!(|(f,f,t)),!(&(f))),|(f,f,t)))),&(&(&(t,t,f),|(f,f,t),|(f)),!(&(t)),!(&(|(f,f,t),&(&(f),&(!(t),&(f),|(f)),&(!(&(f)),&(t),|(f,f,t))),&(t))))),&(&(!(&(f)),|(t),&(!(t),!(|(f,f,t)),!(&(f)))),!(|(f,f,t)),&(t,t,f))),!(&(f)),&(t)),!(&(|(f,f,t),&(&(f),&(!(t),&(f),|(f)),&(!(&(f)),&(t),|(f,f,t))),&(t)))),!(!(&(&(t,t,f),|(f,f,t),|(f)))),&(&(t),&(&(&(&(f),&(!(t),&(f),|(f)),&(!(&(f)),&(t),|(f,f,t))),|(t),&(!(t),!(|(f,f,t)),!(&(f)))),!(&(|(f,f,t),&(&(f),&(!(t),&(f),|(f)),&(!(&(f)),&(t),|(f,f,t))),&(t))),&(&(&(!(&(f)),|(t),&(!(t),!(|(f,f,t)),!(&(f)))),!(|(f,f,t)),&(t,t,f)),&(f),&(&(t),&(!(t),!(|(f,f,t)),!(&(f))),|(f,f,t)))),&(&(&(t,t,f),|(f,f,t),|(f)),!(&(t)),!(&(|(f,f,t),&(&(f),&(!(t),&(f),|(f)),&(!(&(f)),&(t),|(f,f,t))),&(t))))))",
		false,
	},

	{
		"!(f)",
		true,
	},

	{
		"|(f,t)",
		true,
	},

	{
		"&(t,f)",
		false,
	},

	{
		"|(&(t,f,t),!(t))",
		false,
	},

	// 可以有多个 testcase
}

func Test_parseBoolExpr(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, parseBoolExpr(tc.expression), "输入:%v", tc)
	}
}

func Benchmark_parseBoolExpr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			parseBoolExpr(tc.expression)
		}
	}
}
