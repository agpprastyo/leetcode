// Last updated: 4/21/2026, 12:17:25 PM
func intToRoman(num int) string {

    type RomanNumeral struct {
        value int
        symbol string
    }

    numerals := []RomanNumeral{
        {1000, "M"},
        {900, "CM"},
        {500, "D"},
        {400, "CD"},
        {100, "C"},
        {90, "XC"},
        {50, "L"},
        {40, "XL"},
        {10, "X"},
        {9, "IX"},
        {5, "V"},
        {4, "IV"},
        {1, "I"},
    }

    result := ""

    for _, numeral := range numerals {
        for num >= numeral.value {
            result += numeral.symbol
            num -= numeral.value
        }
    }

    return result
    
}