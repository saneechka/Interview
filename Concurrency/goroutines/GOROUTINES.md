# Goroutines and Concurrency in Go

Goroutines are a fundamental concept for concurrency in Go. They are lightweight and managed by Go's runtime scheduler. The scheduler in Go is based on a model known as M:N scheduling, which multiplexes M goroutines on N OS threads.

## Context Switching

Context switching between goroutines is extremely fast and efficient because it is managed by the Go runtime, not the operating system. This is possible due to the small stack size of goroutines which starts with a default size that grows and shrinks as required, without the need for large preallocated stacks.

## The `runtime` Package

The `runtime` package in Go provides a set of functions that are crucial for working with the concurrency of goroutines. It includes functions to manipulate the runtime environment and scheduler.

## Cooperative Scheduling

Go employs a cooperative scheduling model, which means that the goroutines themselves yield control to the scheduler, allowing other goroutines to run. This is in contrast to preemptive scheduling, where the operating system kernel interrupts the process to switch the executing thread.

## Task Scheduler

The Go scheduler uses a work-stealing algorithm to distribute goroutines over available threads. When a goroutine becomes blocked, the scheduler assigns another goroutine to the thread, ensuring CPU time is used efficiently.


# Горутины и параллельное выполнение в Go

Горутины - основа параллельности в Go. Они легковесны и управляются планировщиком времени выполнения Go. Планировщик в Go основан на модели M:N, которая мультиплексирует M горутин на N потоков ОС.

## Переключение Контекстов

Переключение контекстов между горутинами очень быстрое и эффективное, так как оно управляется временем выполнения Go, а не операционной системой. Это возможно благодаря небольшому размеру стека горутин, который начинается с размера по умолчанию и растет или уменьшается по мере необходимости, без необходимости в больших предварительно выделенных стеках.

## Пакет `runtime`

Пакет `runtime` в Go предоставляет набор функций, критически важных для работы с параллельностью горутин. Он в&#8203;``【oaicite:0】``&#8203;


# Горутины в Go

## Обзор
Горутины — это основной концепт в Go для достижения конкурентности. Это функции или методы, выполняющиеся параллельно с другими функциями или методами. Горутины легковесны и управляются средой выполнения Go.

## Создание Горутины
Чтобы запустить горутину, используйте ключевое слово `go`, за которым следует вызов функции. Это планирует выполнение функции параллельно.

## Синхронизация
Горутины выполняются независимо, но иногда требуется синхронизировать их выполнение. Для этих целей Go предоставляет каналы и примитивы синхронизации, такие как `sync.WaitGroup`.

## Каналы
Каналы — предпочтительный способ коммуникации между горутинами. Данные могут быть отправлены одной горутиной и получены другой, что позволяет безопасно и синхронно обмениваться данными.

См. `goroutines.go`.

