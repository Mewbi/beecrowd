import kotlin.math.*
import java.util.*

fun kamenetsky(n: Double): Int {
    if (n < 0) {
        return 0
    }

    if (n <= 1) {
        return 1
    }

    val x = ((n * (ln(n / E) / ln(10.0))) +
             ((ln(2.0 * PI * n) / ln(10.0)) / 2.0))

    return floor(x).toInt() + 1
}

fun main(args: Array<String>) {
    val n = readLine()?.toDoubleOrNull() ?: 0.0

    val result = kamenetsky(n)
    println(result)
}

