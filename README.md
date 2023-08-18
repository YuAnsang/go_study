# Go 언어
- Go 언어에 대해 학습한다.
- 모든 내용을 정리하는 것이 아니라 개인적으로 생각했을 때, 적어둘 법한 것만 README.md에 작성한다.
- 나머지는 예제 코드를 통해서 진행한다.

## package
- 패키지는 Go 언어에서 코드를 묶는 가장 큰 단위이다.
- 다른 언어에서는 namespace라는 키워드로 코드영역을 분리하기도 하지만 Go는 namespace를 지원하지 않고 패키지를 사용한다.

### main package
- 프로그램 시작점(main())을 포함한 패키지
- 프로그램이 실행되면 운영체제는 프로그램을 메모리로 올린다. (Load라고도 한다.)

### 패키지 사용
```
import "fmt"
imrpot (
    "fmt"
    "math/rand"  // 경로가 있는 패키지 -> 패키지명은 rand
)

import (
    "text/template"
    htemplate "html/template" // template이 겹치므로 별칭    
)
```
- Go 언어는 패캐지를 가져오면 반드시 사용해야 하고 사용하지 않으면 에러가 발생함
- 하지만 아래와 같이 직접 사용하지 않지만 부가효과를 얻고자 하는 경우 _을 붙일수 있음
```
import (
    "database/sql"
    _ "fmt"
)
```

### 패키지 경로
1. Go 언어에서 기본 제공하는 패키지는 Go 설치 경로에서 찾음
2. 외부 저장소에 저장된 패키지의 경우 다운로드 받아서 GOPATH/pkg 폴더에 설치

### Go 모듈
- Go 모듈은 Go 패키지를 모아놓은 프로젝트 단위.
- Go 1.16부터 모듈 사용이 기본으로 됐음. 이전까지는 Go 모듈을 만들지 않는 코드는 모두 GOPATH/src 폴더에 있어야 했지만 이제는 모든 Go코드들이 Go 모듈 아래 있어야함.
- go build를 하려면 반드시 Go 모듈 루트 폴더에 go.mod 파일이 존재해야함.
- Go는 go build를 통해 실행파일을 만들 때, go.mod와 외부 패키지 버전정보를 담고 있는 go.sum 파일을 통해 실행파일을 만들게 됨

## 슬라이스
- 슬라이스는 내장 타입으로 내부 구현이 감춰져 있지만 reflect 패키지의 SliceHeader 구조체를 사용함.
```
type SliceHeader struct {
	Data uintptr // 배열을 가르키는 포인터
	Len  int // 현재 요소의 개수
	Cap  int // 실제 배열의 길이
}
```

### 배열과 슬라이스의 동작 차이
```
func chanceArray(array [5]int) {
    array[2] = 200
}

func chanceSlice(slice2 []int) {
    slice[2] = 200
}

func main() {
  array := [5]int {1, 2, 3, 4, 5}
  slice := []int {1, 2, 3, 4, 5}
  
  changeArray(array) // 동작안함
  changeSlice(slice) // 동작함
}
```
- Go 언어에서 모든 값의 대입은 복사로 일어나고, 함수의 인수로 전달될 때의 값의 이동도 복사로 일어남
- 슬라이스는 포인터, len, cap을 복사하기 때문에 같은 주소값을 가르킴

## 메서드
- Go언어에서는 클래스가 없어서 구조체 밖에서 메서드를 지정함.
- 구조체 밖에 메서드를 정의할 때, 리시버라는 특별한 기능을 사용
- 리시버 : 메서드가 어떤 구조체에 속하는지 표시하는 기법
```
func (r Rabbit) info() int {
	return r.width * r.height
}
```
