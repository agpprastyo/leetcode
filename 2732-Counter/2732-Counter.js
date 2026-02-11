// Last updated: 2/11/2026, 8:30:07 PM
/**
 * @param {number} n
 * @return {Function} counter
 */
var createCounter = function(n) {
   
    return function() {
        return n++;
    };
};

/** 
 * const counter = createCounter(10)
 * counter() // 10
 * counter() // 11
 * counter() // 12
 */