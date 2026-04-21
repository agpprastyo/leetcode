// Last updated: 4/21/2026, 12:33:00 PM
func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
    sum1 := int(coordinate1[0] + coordinate1[1])
    sum2 := int(coordinate2[0] + coordinate2[1])

    return (sum1 % 2) == (sum2 % 2)
}