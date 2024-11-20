# Multitasking

## Definition
Multitasking is the capability of an operating system to manage multiple tasks (or processes) at the same time or to give the appearance of doing so by switching between tasks in a way that allows the user to interact with each task in almost real-time.

## Cooperative Multitasking
In cooperative multitasking, tasks are in charge of their own execution. They must periodically yield control back to the operating system to allow other tasks to run. The issue with this approach is that if a task does not yield control, it can lead to the entire system hanging.

## Preemptive Multitasking
Preemptive multitasking allows the operating system to control task execution, deciding when to switch the context. This provides a more reliable and efficient system since the OS can interrupt a task and resume another, preventing a system freeze due to a single task.

## Multithreading
Multithreading is the concept where a process can contain several concurrently executing threads. This enables a process to make efficient use of computing resources, especially on multi-core systems.

## Parallelism
Parallelism refers to situations where multiple tasks are genuinely executed simultaneously, possible on multi-core or multiprocessor systems. Parallelism aims to speed up the execution of tasks by running them concurrently.

## Asynchrony
Asynchronous task execution allows a program to continue processing while waiting for the completion of another task, which might take a significant amount of time. This helps maintain application responsiveness and improve the user experience.

Multitasking in Go is implemented with the help of goroutines, which are lightweight threads managed by the Go runtime. This enables the creation of high-performance concurrent applications with simplified state management and synchronization.


# Мультитаскинг (Многозадачность)

## Определение
Многозадачность — это возможность операционной системы выполнять несколько задач (процессов) одновременно или создавать иллюзию одновременности путём быстрого переключения между задачами, что позволяет пользователю взаимодействовать с каждой задачей почти в режиме реального времени.

## Кооперативная многозадачность
При кооперативной многозадачности задачи сами управляют выполнением. Они должны периодически возвращать управление операционной системе, чтобы могли выполняться другие задачи. Недостаток такого подхода заключается в том, что если задача не вернёт управление, это может вызвать сбой в системе.

## Вытесняющая многозадачность (Preemptive Multitasking)
Вытесняющая многозадачность позволяет операционной системе контролировать выполнение задач, принимая решение о смене контекста. Это создаёт более надёжную и эффективную систему, так как операционная система может прервать задачу и запустить другую, предотвращая зависание системы из-за одной задачи.

## Многопоточность
Многопоточность - это возможность процесса иметь несколько параллельных потоков выполнения. Это позволяет процессу эффективно использовать вычислительные ресурсы, особенно в многоядерных системах.

## Параллелизм
Параллелизм означает, что несколько задач фактически выполняются одновременно, что возможно на многоядерных или многопроцессорных системах. Цель параллелизма - ускорить выполнение задач за счёт их одновременности.

## Асинхронность
Асинхронное выполнение задач позволяет программе продолжать обработку во время ожидания завершения другой задачи, на которую может уйти значительное время. Это способствует поддержанию отзывчивости приложения и улучшению пользовательского опыта.

Многозадачность в Go реализуется через использование горутин, которые являются легковесными потоками, управляемыми средой выполнения Go. Это позволяет создавать высокопроизводительные конкурентные приложения с упрощённым управлением состоянием и синхронизацией.


