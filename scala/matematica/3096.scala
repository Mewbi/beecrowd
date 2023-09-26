import scala.math._

object Main extends App {
  def kamenetsky(n: Double): Int = {
    if (n < 0) {
      0
    } else if (n <= 1) {
      1
    } else {
      val x = ((n * log10(n / E)) + (log10(2 * Pi * n) / 2.0))
      math.floor(x).toInt + 1
    }
  }

  val n = scala.io.StdIn.readDouble()
  println(kamenetsky(n))
}
