export function decodedValue(input: Array<string>): number {
    let total = 0;
    let myMap = new Map<string, number>([
        ["black", 0],
        ["brown", 1],
        ["red", 2],
        ["orange", 3],
        ["yellow", 4],
        ["green", 5],
        ["blue", 6],
        ["violet", 7],
        ["grey", 8],
        ["white", 9],
    ]);
    input.forEach(el => {
        let amt = myMap.get(el.toLowerCase())
        if (amt) {
            total = Number('' + total + amt)
        }
    })
    return total
}
