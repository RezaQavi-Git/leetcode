package problems
import "strings"

func Convert(s string, numRows int) string {
    if numRows == 1 || numRows >= len(s) {
        return s
    }

    rows := make([]strings.Builder, numRows)

    row := 0
    goDown := false

    for _, char := range s {
        rows[row].WriteRune(char)

        if row == 0 || row == numRows-1 {
            goDown = !goDown
        }

        if goDown {
            row++
        } else {
            row--
        }
    }

    var result strings.Builder
    for _, builder := range rows {
        result.WriteString(builder.String())
    }

    return result.String()
}
