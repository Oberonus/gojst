function multiply(arg1, arg2) {
    if (typeof arg1 === 'string') {
        arg1 = value(arg1)
    }
    if (typeof arg2 === 'string') {
        arg2 = value(arg2)
    }
    return arg1 * arg2
}

function value(index) {
    return index.split('.').reduce(function index(obj, i) {
        return obj[i]
    }, data)
}
