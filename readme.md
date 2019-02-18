# Golang-Javascript Template and Expression Evaluator

```go
import "github.com/oberonus/gojst"
```

Quick and dirty wrapper around [otto](https://github.com/robertkrimen/otto) javascript 
interpreter to give a power of ES5 to the goland template engine. 

## Quick Example
### Initialization
```go
//get your favourite javascript file
script := `
    function mul(arg1, arg2) {
        return arg1 * arg2;
    }
`
//initialize external variables
vars := map[string]interface{}{
    "v1": 3,
    "v2": 4,
}
//create new engine
eng, err := gojst.NewEngine(strings.NewReader(script), vars)
if err != nil {
    panic(err)
}
```

### Evaluating Expressions
```go
//execute javascript expression
res, err := eng.EvalString(`"vars: " + data.v1 + " and " + data.v2`)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%v\n", res)
```
Will print:
```
vars: 3 and 4
```

### Rendering templates
```go
//run template renderer
res, err := eng.Render(`multiplication of two variables: {{.C "mul" .D.v1 .D.v2}}`)
if err != nil {
    log.Fatal(err)
}
//will print ""
fmt.Printf("%v\n", res)
```
Will print:
```
multiplication of two variables: 12
```

## Usage in templates
Reaching javascript functions and internal/external variables from template
### Call javascript function
```go
    {{.C "function" param1 param2 ...}}
```

### Getting internal variables from javascript engine
```go
    {{.V "variable"}}
```

### Getting external variables from provided map
```go
    {{.D.variable}}
```
