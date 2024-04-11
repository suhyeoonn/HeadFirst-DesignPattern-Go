# 옵저버 패턴

독자가 신문을 보고 싶으면 신문사에 구독 신청을 하고, 더 이상 보고 싶지 않으면 구독 해지 신청을 하는 신문 구독 메커니즘을 이해하면 옵저버 패턴을 쉽게 이해할 수 있다.
신문사를 주제(subject), 구독자를 옵저버(observer)라고 부른다.

> 옵저버 패턴(Observer Pattern)은 한 객체의 상태가 바뀌면 그 객체에 의존하는 다른 객체에게 연락이 가고
> 자동으로 내용이 갱신되는 방식으로 일대다 의존성을 정의한다.

### 주제

- 데이터를 관리한다.
- 데이터가 바뀌면 옵저버에게 소식을 전달한다.

### 옵저버

- 주제를 구독하고 있으며 주제 데이터가 바뀌면 갱신 내용을 전달받는다.

## 예제

다음과 같은 기능을 구현해야 한다.

- 온도, 습도, 기압 데이터가 새로 들어올 때마다 3가지 디스플레이를 갱신해야 한다. (현재 조건 디스플레이, 기상 통계 디스플레이, 기상 예보 디스플레이)
- 확장성을 고려해야 한다.
  - 이후에 디스플레이가 더 생길 수 있다.

## Bad Practice

```go
func (w WeatherData) MeasurementsChanged() {
	temp := w.temp
	humidity := w.humidity
	pressure := w.pressure

	CurrentConditionsDisplay{}.Update(temp, humidity, pressure)
	StatisticsDisplay{}.Update(temp, humidity, pressure)
	ForecaseDisplay{}.Update(temp, humidity, pressure)
}
```

위 코드는 다음과 같은 문제가 있다.

- 인터페이스가 아닌 구체적인 구현을 바탕으로 코딩을 하고 있다.
- 새로운 디스플레이 항목이 추가될 때마다 코드를 변경해야 한다.
- 실행 중에 디스플레이 항목을 추가하거나 제거할 수 없다.
- 바뀌는 부분을 캡슐화하지 않았다.
- WeatherData 클래스를 캡슐화하지 않고 있다.

##
