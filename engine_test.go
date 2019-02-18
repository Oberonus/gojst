package gojst

import (
	"os"
	"testing"

	"fmt"
	"log"
	"strings"

	"github.com/stretchr/testify/require"
)

func TestEvalSimpleTrueExpression(t *testing.T) {
	env := newTestEnv(nil)
	ok, err := env.Check("12 === multiply(3,4)")
	require.Nil(t, err)
	require.True(t, ok)
}

func TestEvalSimpleFalseExpression(t *testing.T) {
	env := newTestEnv(nil)
	ok, err := env.Check("12 === multiply(4,4)")
	require.Nil(t, err)
	require.False(t, ok)
}

func TestRenderWithFunctionCallAndData(t *testing.T) {
	env := newTestEnv(map[string]interface{}{
		"var1": 2,
		"var2": 6,
		"mp":   map[string]interface{}{"var3": 3},
	})
	ok, err := env.Check("data.mp.var4 = 21; globalVar = 78;")
	require.Nil(t, err)
	require.True(t, ok)

	res, err := env.Render(`foo {{.C "multiply" (.C "multiply" .D.var1 .D.var2) "mp.var3"}} bar {{.D.mp.var4}} {{.V "globalVar"}}`)
	require.Nil(t, err)
	require.Equal(t, "foo 36 bar 21 78", res)
}

func TestItGetsVariable(t *testing.T) {
	env := newTestEnv(nil)
	env.Eval(`global = "foo"`)
	res, err := env.Render(`{{.V "global"}}`)
	require.Nil(t, err)
	require.Equal(t, "foo", res)
}

func TestUnderscore(t *testing.T) {
	env := newTestEnv(nil)
	ok, err := env.Check("_.intersection([1, 2, 3], [101, 2, 1, 10], [2, 1])[0] === 1")
	require.Nil(t, err)
	require.True(t, ok)
}

func newTestEnv(data interface{}) *Engine {
	f, _ := os.Open("./testdata/test.js")
	defer f.Close()
	env, _ := NewEngine(f, data)
	return env
}

func TestRenderUnknownFunc(t *testing.T) {
	env := newTestEnv(nil)
	res, err := env.Render(`{{.C "unknownFunc"}}`)
	require.Equal(t, "", res)
	require.NotNil(t, err)
}

func TestRenderUnknownVariable(t *testing.T) {
	env := newTestEnv(nil)
	res, err := env.Render(`{{.V "unknownVariable"}}`)
	require.Equal(t, "", res)
	require.Nil(t, err)
}

func TestEvaluateInt(t *testing.T) {
	env := newTestEnv(nil)
	v, err := env.EvalInt("2+3")
	require.Nil(t, err)
	require.Equal(t, int64(5), v)
}

func TestEvaluateString(t *testing.T) {
	env := newTestEnv(nil)
	v, err := env.EvalString(`"one "+"three"`)
	require.Nil(t, err)
	require.Equal(t, "one three", v)
}

func TestExample(t *testing.T) {
	script := `
        function mul(arg1, arg2) {
            return arg1 * arg2;
        }
    `

	vars := map[string]interface{}{
		"v1": 3,
		"v2": 4,
	}

	env, err := NewEngine(strings.NewReader(script), vars)
	if err != nil {
		log.Fatal(err)
	}
	//execute javascript expression
	res, err := env.EvalString(`"vars: " + data.v1 + " and " + data.v2`)
	if err != nil {
		log.Fatal(err)
	}
	//will print "vars: 3 and 4"
	fmt.Printf("%v\n", res)

	//run template renderer
	res, err = env.Render(`multiplication of two variables: {{.C "mul" .D.v1 .D.v2}}`)
	if err != nil {
		log.Fatal(err)
	}
	//will print "multiplication of two variables: 12"
	fmt.Printf("%v\n", res)
}
