## Variables & Constants
- var <변수명> <타입> = 값
- := 으로도 사용할 수 있으나 지역변수를 선언할 때만 사용할 수 있다.
```
var a int
var f float32 = 11.
var i,j,k int = 1,2,3
var d = 5
```
- const 키워드는 상수를 선언한다.
- 여러개를 묶어서 한번에 const를 선언할 수 있다.
```
const hi = "Hi"
const(
    sky = "blue"
    rose = "red"
    Gyuri = "awesome"
)
```
## Data Type
### String
- Back Quote와 Double Quote로 선언할 수 있다.
- String은 immutable type이므로, 한번 생성되면 수정할 수 없다.
```
rawLiteral := `메롱\n`
interLiteral := `hi\n`
```
### Type Conversion
- Go에서는 암묵적인 type conversion이 이뤄지지 않는다. (즉, 항상 명시적으로 지정해야 함.)
```
func main() {
    var i int = 100
    var u uint = uint(i)
    var f float32 = float32(i)
    println(f, u)

    str := "ABC"
    bytes := []byte(str)

    str2 := string(bytes)
    println(bytes, str2)
}
```

## Operator
```
var p *int
var q ^int
var k int = 10
var p = &k //k의 주소를 할당
println(*p) //p가 가리키는 주소에 있는 실제 내용을 출력
```

## 조건문
- if 다음에는 반드시 boolean식으로 표현한다.
- if문에서 for문 처럼 optional statement를 사용할 수 있다.
```
if k == 1{
    println("one")
} else if k==2 {
    println("Two")
} else {
    println("other")
}


if val := i * 2; val < max{
    println(val)
}
```

## Loop
- while은 없고, for문만 있다.
```
for i := 1; i <= 100; i++ {
    logic
}
for n < 100{
    //while문 역할
}
for {
    //무한 반복
}
```
- python 처럼 for range를 사용할 수 있다.(단, index변수까지 2개 필요)
```
names := []string{"one","two","three"}

for index, name := range names{
    println(index, name)
}
```
```
package main

func main(){
    i := 0
L1:
    for{
        if i==0{
            break L1
        }
    }
    println("Ok")
}
```

## Array
- var 변수이름 [배열크기]type
```
var a [3]int
var a1 = [3]int{1,2,3}
var a3 = [...]int{1,2,3} //배열크기 자동으로 할당

var a = [2][3]int{
    {1,2,3},
    {4,5,6}, //끝에 항상 콤마 추가
}
```

## Slice
- 선언은 배열과 똑같으며, 대신 [] 안에 크기는 지정하지 않는다.
```
func main(){
    var s[]int
    var a[]int{1,2,3}

    if s == nil{
        println("Nul slice")
    }
    println(len(s), cap(s)) // 모두 0
}
```
```
a := make([]int, 5, 10) // slice with len :5, capacity :10
```
```
s := []int{0,1,2,3,4,5}
s = s[2:5] // 2,3,4
s = s[1:] //3, 4
```
- slice와 slice는 append()를 사용하여 연결할 수 있다.
    - 주의할 점은 slice 중 두번째에 ellipsis(...)를 붙여야 한다.
    - Go에서 ellipsis는 여러 의미로 사용되는데, 여기서는 unpacking의 의미이다.
```
package main
import "fmt"

func main(){
    sliceA := []int{1,2,3}
    sliceB := []int{4,5,6}
    sliceA = append(sliceA, sliceB...)

    fmt.Println(sliceA) //[1 2 3 4 5 6] 출력
}

source := []int{0,1,2}
target := make([]int, len(source), cap(source)*2)
copy(target, source)
```

## Map
- map[key type]value type 의 형태로 선언할 수 있다.
- Map을 활용하면 dict처럼 활용할 수 있다.
- Map은 Unordered인 Hash이므로 순서가 무작위이다.
    - 즉, for range를 이용했을 때, 모든 요소가 매번 다른 순서로 반복문에서 돌려진다.
```
var idMap map[int]string
idMap = make(map[int]string)
```
```
package main
func main(){
    tickers := map[string]string{
        "goog": "Google Inc",
        "MSFT": "Microsoft",
        "FB": "FaceBook",
        "AMZN": "Amazon",
    }

    val, exists := tickers["MSFT"]
    if !exists{
        println("No MSFT ticker")
    }
}
```