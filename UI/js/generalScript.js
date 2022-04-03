//console.log(mask("5022291052143078","**** **** **** ****"))
function mask(value, pattern) {
    let count = 0;
    return pattern.replace(/\*/g, () => value[count++] || '')
}
