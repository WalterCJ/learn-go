package iteration

const repeatCount = 5

// Repeat a character n times
func Repeat(character string) string {
    var repeated string
    for i := 0; i < repeatCount; i++ {
        repeated += character
    }
    return repeated
}