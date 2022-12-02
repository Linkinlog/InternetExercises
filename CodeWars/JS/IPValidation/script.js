// 192.168.0.1 - true
// 092.168.0.1 - false 
// 258.266.333.1 - false
// 0.0.0 - false
function isValidIP(str) {
    let strArr = str.split('.')
    // Check that there is 4 octets
    if ( strArr.length  !== 4 ) {
        return false
    }
    for (const el of strArr) {
        console.log(el);
        if (el.length < 1 || /([\D])/.test(el)) {
            return false
        }
        console.log(`After: ${el}`);
        // Check that the el 
        if (el.length < 1) {
            console.log("object1");
            return false
        }
        // check that none of the octets go over 255 or under 0
        if ( el > 255 || el < 0 ) {
            console.log("object2");
            return false
        }
        // check that there are no leading 0s
        if ( el[0] == 0 && el.length > 1 ) {
            console.log("object3");
            return false
        }
    };
	return true;
}


function tester() {
    console.log(isValidIP("192.168.0.1"), "true")
    console.log(isValidIP("192.1e8.0.1"), "false")
    console.log(isValidIP("192.168.0.1\n"), "false")
    console.log(isValidIP("192.a168.0.1"), "false")
    console.log(isValidIP("192.068.0.1"), "false")
    console.log(isValidIP("192.468.0.1"), "false")
    console.log(isValidIP("192.168..1"), "false")
    console.log(isValidIP("192.168"), "false")
}

tester()