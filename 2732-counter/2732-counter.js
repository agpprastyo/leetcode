/**
 * @param {number} n
 * @return {Function} counter
 */
var createCounter = function(n) {
    // Initialize the counter with the starting value n
    let current = n;
    
    // Return a function that increments and returns the current value
    return function() {
        return current++;
    };
};

/** 
 * const counter = createCounter(10)
 * counter() // 10
 * counter() // 11
 * counter() // 12
 */