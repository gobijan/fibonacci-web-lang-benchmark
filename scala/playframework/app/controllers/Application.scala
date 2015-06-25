package controllers

import play.api.mvc._
import play.api.libs.concurrent.Execution.Implicits.defaultContext

object Application extends Controller {

  def findFib(num:Int) = Action.async {
    val futureInt = scala.concurrent.Future { fib(num) }
    futureInt.map(i => Ok(i.toString))
  }

  def fib(n:Int):Int = {
    n match {
      case 0 => 0
      case 1 => 1
      case _ => fib(n-1) + fib(n-2)
    }
  }
}