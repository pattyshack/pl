func method2(x >&X, arg1 [Int], varargs ...Int) Int {}

func badInferredReceiver(>_) {}

var $(
  i = 5
  (j, k) (Int, String)
  (k, v) = ("key", "value")
)

{
  let $(
    x = 1
    y = 2
    z = 3
    a Int
  )
}

type X struct (
  field struct(
    x Int
    func method1(>_) {}
    func method2(>_) {}
  )
)

func function(
  bad struct(
    x Int
    y Int
    func method1(>_) {}
    func method2(>_) {}
  ),
) {
}
