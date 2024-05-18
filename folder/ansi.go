package folder

import "fmt"


func ANSIColor(r, g, b int) string {
    return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}