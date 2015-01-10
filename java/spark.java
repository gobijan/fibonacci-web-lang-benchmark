import static spark.Spark.get;

public class FibBench {
    public static void main(String[] args) {

        get("/:number", (req, res) -> {
                    int number = Integer.parseInt(req.params(":number"));
                    return "Sparkjava: fib(" + number +") = " + fib(number);
                }
        );

    }

    public static int fib(int n) {
        if (n == 0) {
            return 0;
        } else if (n == 1) {
            return 1;
        } else {
            return fib(n - 1) + fib(n - 2);
        }
    }

}